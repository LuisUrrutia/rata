resource "google_compute_network" "private_network" {
  provider = "google-beta"
  name     = "${var.platform_name}-network"
}

resource "google_compute_firewall" "firewall_general" {
  name    = "${var.platform_name}-general"
  network = "${google_compute_network.private_network.name}"

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["22", "80", "443"]
  }

  source_ranges = ["0.0.0.0/0"]
}

resource "google_compute_global_address" "db_private_ip_address" {
  provider = "google-beta"

  name          = "${var.platform_name}-db-private-ip"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = "${google_compute_network.private_network.self_link}"
}

resource "google_service_networking_connection" "private_vpc_connection" {
  provider = "google-beta"

  network                 = "${google_compute_network.private_network.self_link}"
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = ["${google_compute_global_address.db_private_ip_address.name}"]
}

resource "google_container_cluster" "primary" {
  name     = "${var.platform_name}-cluster"
  location = "${var.gcloud_zone}"

  depends_on = [
    "google_service_networking_connection.private_vpc_connection",
  ]

  remove_default_node_pool = true
  initial_node_count       = 1
  network                  = "${google_compute_network.private_network.self_link}"

  # Setting an empty username and password explicitly disables basic auth
  master_auth {
    username = "${var.gke_master_auth_user}"
    password = "${var.gke_master_auth_pass}"
  }
}

resource "google_container_node_pool" "primary_preemptible_nodes" {
  provider           = "google-beta"
  name               = "${var.platform_name}-pool"
  location           = "${var.gcloud_zone}"
  cluster            = "${google_container_cluster.primary.name}"
  initial_node_count = 1

  autoscaling {
    min_node_count = "${var.gke_min_nodes_count}"
    max_node_count = "${var.gke_max_nodes_count}"
  }

  node_config {
    preemptible  = true
    machine_type = "n1-standard-1"

    metadata {
      disable-legacy-endpoints = "true"
    }

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}

resource "google_sql_database_instance" "master" {
  name             = "${var.platform_name}-database"
  region           = "${var.gcloud_region}"
  database_version = "MYSQL_5_7"

  depends_on = [
    "google_service_networking_connection.private_vpc_connection",
  ]

  settings {
    tier      = "db-n1-standard-1"
    disk_size = 20

    ip_configuration {
      ipv4_enabled    = "false"
      private_network = "${google_compute_network.private_network.self_link}"
    }
  }
}

resource "google_sql_database" "scraper" {
  name      = "scraper"
  instance  = "${google_sql_database_instance.master.name}"
  charset   = "latin1"
  collation = "latin1_swedish_ci"
}

# The following outputs allow authentication and connectivity to the GKE Cluster
# by using certificate-based authentication.
output "client_certificate" {
  value = "${google_container_cluster.primary.master_auth.0.client_certificate}"
}

output "client_key" {
  value = "${google_container_cluster.primary.master_auth.0.client_key}"
}

output "cluster_ca_certificate" {
  value = "${google_container_cluster.primary.master_auth.0.cluster_ca_certificate}"
}
