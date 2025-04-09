resource "google_storage_bucket" "bucket" {
  name          = var.name
  location      = var.location
  force_destroy = true
  project       = var.project

  uniform_bucket_level_access = true

}
