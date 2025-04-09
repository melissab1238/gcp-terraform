module "gcp_bucket_state" {
  source  = "./../../modules/gcp_bucket"
  project = "central-infrastructure"
  name    = "terraform-your-project"
}
