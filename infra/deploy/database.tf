
variable "db_username" {
  description = "username for rds database"
  default     = "ci-cd"
}

variable "db_password" {
  description = "password for rds database"
  sensitive   = true
  default     = "ci-cd-password"
}

resource "aws_db_subnet_group" "db_subnet" {
  name       = "${local.prefix}-main"
  subnet_ids = [aws_subnet.private_a.id, aws_subnet.private_b.id]

}

resource "aws_security_group" "rds" {
  description = "Allow access to RDS instance"
  name        = "${local.prefix}-rds-inbound-access"
  vpc_id      = aws_vpc.main.id


  ingress {
    protocol  = "tcp"
    from_port = "3306"
    to_port   = "3306"
    security_groups = [
      aws_security_group.ecs_service.id
    ]
  }

  tags = {
    Name = "${local.prefix}-db-security-group"
  }

}

resource "aws_db_instance" "main" {
  identifier                 = "${local.prefix}-db"
  db_name                    = "todos"
  engine                     = "mysql"
  storage_type               = "gp2"
  allocated_storage          = 20
  instance_class             = "db.t4g.micro"
  engine_version             = "8.0.39"
  auto_minor_version_upgrade = true
  username                   = var.db_username
  password                   = var.db_password
  skip_final_snapshot        = true

  db_subnet_group_name = aws_db_subnet_group.db_subnet.name
  multi_az             = false

  backup_retention_period = 0
  vpc_security_group_ids  = [aws_security_group.rds.id]

  tags = {
    Name = "${local.prefix}-main"
  }
}
