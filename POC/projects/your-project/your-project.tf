module "gcp_bucket" {
  source   = "./../../modules/resource"
  name     = "my-bucket"
  project  = "your-project"
  location = "us-central1"
}
module "gcp_function" {
  source      = "./../../modules/gcp_function"
  location    = "us-central1"
  name        = "my-function"
  runtime     = "nodejs20"
  description = "my function can do things!"
  project     = "your-project"
}
