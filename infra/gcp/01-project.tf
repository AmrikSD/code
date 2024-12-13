resource "google_service_account" "frappe" {
  account_id   = "frappe-sa"
  display_name = "Custom SA for frappe Instance"
}

resource "google_compute_address" "frappe-static" {
  name = "frappe-static"
}

resource "google_compute_instance" "frappe" {
  name         = "frappe"
  machine_type = "e2-small"

  tags = ["https-server", "http-server", "milestone-medical", "terraform"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  metadata = {
    ssh-keys = format("%s:%s", data.sops_file.gcp-secret.data["google.ssh.user"], data.sops_file.gcp-secret.data["google.ssh.public_key"])
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

        systemctl enable --now /home/asbotehg/docker.service
        systemctl enable --now /home/asbotehg/docker-compose.app.service
    EOT
  }



  network_interface {
    network = "default"
    access_config {
      nat_ip = google_compute_address.frappe-static.address
    }
  }

    provisioner "file" {
      source      = "${path.module}/frappe/docker.service"
      destination = "docker.service"
      connection {
          type  = "ssh"
          host = self.network_interface[0].access_config[0].nat_ip
          user = data.sops_file.gcp-secret.data["google.ssh.user"]
          private_key = data.sops_file.gcp-secret.data["google.ssh.private_key"]
          timeout = "4m"
      }
    }

    provisioner "file" {
      source      = "${path.module}/frappe/nginx.conf"
      destination = "nginx.conf"
      connection {
          type  = "ssh"
          host = self.network_interface[0].access_config[0].nat_ip
          user = data.sops_file.gcp-secret.data["google.ssh.user"]
          private_key = data.sops_file.gcp-secret.data["google.ssh.private_key"]
          timeout = "4m"
      }
    }

    provisioner "file" {
      source      = "${path.module}/frappe/docker-compose.app.service"
      destination = "docker-compose.app.service"
      connection {
          type  = "ssh"
          host = self.network_interface[0].access_config[0].nat_ip
          user = data.sops_file.gcp-secret.data["google.ssh.user"]
          private_key = data.sops_file.gcp-secret.data["google.ssh.private_key"]
          timeout = "4m"
      }
    }

  provisioner "file" {
      source = "${path.module}/frappe/docker-compose.yaml"
      destination = "docker-compose.yaml"
      connection {
          type  = "ssh"
          host = self.network_interface[0].access_config[0].nat_ip
          user = data.sops_file.gcp-secret.data["google.ssh.user"]
          private_key = data.sops_file.gcp-secret.data["google.ssh.private_key"]
          timeout = "4m"
      }
  }

}

resource "google_compute_firewall" "default-ssh" {
    name = "allow-ssh"
    network = "default"
    allow {
        protocol = "tcp"
        ports = ["22"]
    }
    source_ranges = [
        "213.78.238.32/32"
    ]
}

output "frappe_ip_address" {
  value = google_compute_address.frappe-static.address
}
