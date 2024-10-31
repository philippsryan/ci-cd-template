# Create ECR Repos for storing Docker images


resource "aws_ecr_repository" "frontend" {
  name                 = "ci-cd-app-fr"
  image_tag_mutability = "MUTABLE"
  force_delete         = true

  image_scanning_configuration {
    # Update to true for real projects
    scan_on_push = false
  }
}


resource "aws_ecr_repository" "backend" {
  name                 = "ci-cd-app-be"
  image_tag_mutability = "MUTABLE"
  force_delete         = true

  image_scanning_configuration {
    # Update to true for real projects
    scan_on_push = false
  }
}



output "ecr_repo_app_fr" {
  description = "ECR repo Url for the front app image"
  value       = aws_ecr_repository.frontend.repository_url
}


output "ecr_repo_app_be" {
  description = "ECR repo Url for the back end app image"
  value       = aws_ecr_repository.backend.repository_url
}
