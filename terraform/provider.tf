provider "google-beta" {
  credentials = "${file("google-credentials.json")}"
  project     = "${var.gcloud_project}"
  region      = "${var.gcloud_region}"
  zone        = "${var.gcloud_zone}"
}

provider "google" {
  credentials = "${file("google-credentials.json")}"
  project     = "${var.gcloud_project}"
  region      = "${var.gcloud_region}"
  zone        = "${var.gcloud_zone}"
}
