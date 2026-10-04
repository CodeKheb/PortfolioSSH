variable "project_id" {
  type        = string
  description = "GCP project ID"
}

variable "region" {
  type        = string
  default     = "us-west1"
}

variable "zone" {
  type        = string
  default     = "us-west1-a"
}

variable "vm_name" {
  type    = string
  default = "portfolio-vm"
}

variable "machine_type" {
  type    = string
  default = "e2-micro"
}
