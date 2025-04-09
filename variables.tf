variable "project_id" {
  description = "The GCP project ID for the developer's project"
  type        = string
}

variable "project_name" {
  description = "The name of the GCP project for the developer's environment"
  type        = string
}

variable "billing_account" {
  description = "The billing account ID for the GCP project"
  type        = string
}

variable "org_id" {
  description = "The organization ID for the GCP project"
  type        = string
}
