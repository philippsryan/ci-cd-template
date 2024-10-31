variable "tf_state_bucket" {
  description = "Name of S3 Bucket in AWS for storing TF State"
  default     = "ci-cd-template-tf-state"
}

variable "project" {
  description = "Name of the project to tag all aws resources"
  default     = "ci-cd-template"
}
