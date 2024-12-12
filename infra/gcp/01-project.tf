resource "google_service_account" "frappe" {
  account_id   = "frappe-sa"
  display_name = "Custom SA for frappe Instance"
}

resource "google_compute_instance" "frappe" {
  name         = "frappe"
  machine_type = "e2-small"
  zone = "us-central1"

  tags = ["milestone-medical", "terraform"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  // Local SSD disk
  scratch_disk {
    interface = "NVME"
  }

  network_interface {
    network = "default"
  }

}
