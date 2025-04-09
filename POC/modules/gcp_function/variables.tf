variable "location" {
  description = "The GCS location"
  type        = string
  default     = "US-CENTRAL1"
}

variable "project" {
  description = "the id of the project in which the resource belongs to"
  type        = string
}

variable "name" {
  description = "the name of the bucket"
  type        = string
}

variable "runtime" {
  description = "the runtime of the gcp function"
  type        = string
  default     = "nodejs20"
}

variable "description" {
  description = "the description of the bucket"
  type        = string
  default     = ""
}
