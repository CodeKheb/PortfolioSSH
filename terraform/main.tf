resource "google_compute_address" "portfolio" {
  name   = "portfolio-static-ip"
  region = var.region
}

resource "google_compute_instance" "portfolio" {
  name         = var.vm_name
  machine_type = var.machine_type
  zone         = var.zone

  tags = [
    "portfolio-vm"
  ]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
      size  = 30
      type  = "pd-standard"
    }
  }

  network_interface {
    network = "default"

    access_config {
      nat_ip = google_compute_address.portfolio.address
    }
  }
}
