variable "platform_name" {
  default = "scraper"
}

variable "gcloud_region" {
  default = "southamerica-east1"
}

variable "gcloud_zone" {
  default = "southamerica-east1-b"
}

variable "gcloud_project" {
  default = "scrapercl-01"
}

variable "gke_master_auth_user" {
  default = ""
}

variable "gke_master_auth_pass" {
  default = ""
}

variable "gke_min_nodes_count" {
  default = 1
}

variable "gke_max_nodes_count" {
  default = 5
}
