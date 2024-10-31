variable "prefix" {
  description = "prefix for resources in aws"
  default     = "ci-cd"
}

variable "project" {
  description = "Project name for tagging resources"
  default     = "ci-cd-template"
}

variable "ecr_frontend_app_image" {
  description = "docker container image for the frontend application within ecr"
}
