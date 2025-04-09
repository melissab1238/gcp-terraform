resource "google_cloudfunctions2_function" "function" {
  name        = "function-v2"
  location    = var.location
  description = var.description

  build_config {
    runtime = var.runtime
  }
}
