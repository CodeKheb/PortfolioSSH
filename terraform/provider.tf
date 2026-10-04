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
  region  = "us-west1"
  zone    = "us-west1-a"
}
