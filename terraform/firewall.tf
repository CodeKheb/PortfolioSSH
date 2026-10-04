resource "google_compute_firewall" "iap_ssh" {
  name    = "portfolio-iap-ssh"
  network = "default"

  direction = "INGRESS"

  allow {
    protocol = "tcp"
    ports    = ["2222"]
  }

  source_ranges = [
    "35.235.240.0/20"
  ]

  target_tags = [
    "portfolio-vm"
  ]
}

resource "google_compute_firewall" "portfolio_public" {
  name    = "portfolio-public"
  network = "default"

  direction = "INGRESS"

  allow {
    protocol = "tcp"
    ports = [
      "22",
      "80",
      "42069",
    ]
  }

  source_ranges = ["0.0.0.0/0"]

  target_tags = ["portfolio-vm"]
}
