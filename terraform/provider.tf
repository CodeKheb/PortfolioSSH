terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 7.0"
    }
  }
}

provider "google" {
  project = "project-87c218ac-87da-47e7-8c7" // this is a temp project, I hit the limit lmao
  region  = "us-central1"
  zone    = "us-central1-a"
}

resource "google_compute_instance" "portfolio" {
  name         = "portfolio-vm"
  machine_type = "e2-micro"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2204-lts"
    }
  }

  network_interface {
    network = "default"

    access_config {}
  }
}
