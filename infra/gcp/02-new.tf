variable "cf_tunnel_id" { type = string }
variable "cf_account_id" { type = string }
variable "cf_tunnel_secret" { type = string }

locals {
  cloudflared_credentials_json = jsonencode({
    AccountTag   = var.cf_account_id
    TunnelID     = var.cf_tunnel_id
    TunnelSecret = var.cf_tunnel_secret
  })
  cloudflared_config_yaml = yamlencode({
    url      = "frontend:8080"
    tunnel   = var.cf_tunnel_id
    protocol = "http2"

    # when this is finally referenced, it's from the perspective of the cloudflared container.
    # if you change this, you need to edit the docker compose under the frappe directory
    credentials_file = "/etc/cloudflared/cloudflared-creds.json"
  })

  cloudflared_config_hash      = md5(local.cloudflared_config_yaml)
  cloudflared_credentials_hash = md5(local.cloudflared_credentials_json)
  docker_service_hash          = filemd5("${path.module}/frappe/docker.service")
  docker_compose_app_hash      = filemd5("${path.module}/frappe/docker-compose.app.service")
  next_compose_yaml_hash       = filemd5("${path.module}/frappe/next-compose.yaml")
}

resource "google_compute_address" "frappe-static-next" {
  name = "frappe-static-next"
}

resource "google_compute_instance" "frappe-next" {
  name         = "frappe-next"
  machine_type = "e2-medium"

  tags = ["https-server", "http-server", "milestone-medical", "terraform"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  attached_disk {
    source      = google_compute_disk.docker-disk.id
    device_name = "docker-data"
  }

  metadata = {
    ssh-keys       = format("%s:%s", data.sops_file.gcp-secret.data["google.ssh.user"], data.sops_file.gcp-secret.data["google.ssh.public_key"])
    startup-script = <<-EOT
        #!/bin/bash
        apt-get update
        apt-get install -y \
        apt-transport-https \
        ca-certificates \
        curl \
        software-properties-common

        curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -
        add-apt-repository \
          "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable"

        apt-get update
        apt-get install -y docker-ce docker-compose
    EOT
  }


  network_interface {
    network = "default"
    access_config {
      nat_ip = google_compute_address.frappe-static-next.address
    }
  }
}

# This is done as a null resource so that diffs are calculated properly in the case of the local files here changing.
# Terraform does not track provisioners only resources so without this you need to manually taint the resource.

resource "null_resource" "upload_cloudflared_config" {
  triggers = {
    creds_hash = local.cloudflared_config_hash
  }

  provisioner "file" {
    content     = local.cloudflared_config_yaml
    destination = "/home/${data.sops_file.gcp-secret.data["google.ssh.user"]}/cloudflared-config.yaml"

    connection {
      type        = "ssh"
      host        = google_compute_instance.frappe-next.network_interface[0].access_config[0].nat_ip
      user        = data.sops_file.gcp-secret.data["google.ssh.user"]
      private_key = data.sops_file.gcp-secret.data["google.ssh.private_key"]
      timeout     = "3m"
    }
  }
  depends_on = [google_compute_instance.frappe-next]
}

resource "null_resource" "upload_cloudflared_creds" {
  triggers = {
    creds_hash = local.cloudflared_credentials_hash
  }

  provisioner "file" {
    content     = local.cloudflared_credentials_json
    destination = "/home/${data.sops_file.gcp-secret.data["google.ssh.user"]}/cloudflared-creds.json"

    connection {
      type        = "ssh"
      host        = google_compute_instance.frappe-next.network_interface[0].access_config[0].nat_ip
      user        = data.sops_file.gcp-secret.data["google.ssh.user"]
      private_key = data.sops_file.gcp-secret.data["google.ssh.private_key"]
      timeout     = "3m"
    }
  }
  depends_on = [google_compute_instance.frappe-next]
}

resource "null_resource" "upload_next_compose" {
  triggers = {
    docker_compose_hash = local.next_compose_yaml_hash
  }

  provisioner "file" {
    source      = "${path.module}/frappe/next-compose.yaml"
    destination = "docker-compose.yaml"
    connection {
      type        = "ssh"
      host        = google_compute_instance.frappe-next.network_interface[0].access_config[0].nat_ip
      user        = data.sops_file.gcp-secret.data["google.ssh.user"]
      private_key = data.sops_file.gcp-secret.data["google.ssh.private_key"]
      timeout     = "4m"
    }
  }
  depends_on = [google_compute_instance.frappe-next]
}

resource "null_resource" "upload_systemd_docker_compose_service" {
  triggers = {
    systemd_docker_compose_service_hash = local.docker_compose_app_hash
  }

  provisioner "file" {
    source      = "${path.module}/frappe/docker-compose.app.service"
    destination = "docker-compose.app.service"
    connection {
      type        = "ssh"
      host        = google_compute_instance.frappe-next.network_interface[0].access_config[0].nat_ip
      user        = data.sops_file.gcp-secret.data["google.ssh.user"]
      private_key = data.sops_file.gcp-secret.data["google.ssh.private_key"]
      timeout     = "4m"
    }
  }

  depends_on = [google_compute_instance.frappe-next]
}

resource "null_resource" "upload_systemd_docker_service" {
  triggers = {
    systemd_docker_service_hash = local.docker_service_hash
  }

  provisioner "file" {
    source      = "${path.module}/frappe/docker.service"
    destination = "docker.service"
    connection {
      type        = "ssh"
      host        = google_compute_instance.frappe-next.network_interface[0].access_config[0].nat_ip
      user        = data.sops_file.gcp-secret.data["google.ssh.user"]
      private_key = data.sops_file.gcp-secret.data["google.ssh.private_key"]
      timeout     = "4m"
    }
  }
  depends_on = [google_compute_instance.frappe-next]
}

resource "null_resource" "start_services" {
  triggers = {
    creds_hash                          = local.cloudflared_credentials_hash
    config_hash                         = local.cloudflared_config_hash
    docker_compose_hash                 = local.next_compose_yaml_hash
    systemd_docker_compose_service_hash = local.docker_compose_app_hash
    systemd_docker_service_hash         = local.docker_service_hash
  }

  depends_on = [
    google_compute_instance.frappe-next,
    null_resource.upload_cloudflared_config,
    null_resource.upload_cloudflared_creds,
    null_resource.upload_next_compose,
    null_resource.upload_systemd_docker_service,
    null_resource.upload_systemd_docker_compose_service,
  ]

  provisioner "remote-exec" {
    inline = [
      "sudo systemctl daemon-reload",
      "sudo systemctl enable --now /home/${data.sops_file.gcp-secret.data["google.ssh.user"]}/docker.service",
      "sudo systemctl enable --now /home/${data.sops_file.gcp-secret.data["google.ssh.user"]}/docker-compose.app.service",
    ]

    connection {
      type        = "ssh"
      host        = google_compute_instance.frappe-next.network_interface[0].access_config[0].nat_ip
      user        = data.sops_file.gcp-secret.data["google.ssh.user"]
      private_key = data.sops_file.gcp-secret.data["google.ssh.private_key"]
      timeout     = "3m"
    }
  }
}


resource "google_compute_disk" "docker-disk" {
  name = "frappe-docker-data-new"
  type = "pd-standard"
  size = 50
}
