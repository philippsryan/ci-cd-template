###
# ECS Cluster for running app on Fargate
###

resource "aws_iam_policy" "task_execution_role_policy" {
  name        = "${local.prefix}-task-exec-role-policy"
  path        = "/"
  description = "allow ecs to retrieve images and add to logs."
  policy      = file("./templates/ecs/task-execution-role-policy.json")
}

resource "aws_iam_role" "task_assume_role" {
  name               = "${local.prefix}-task-execution-role"
  assume_role_policy = file("./templates/ecs/task-assume-role-policy.json")

}

resource "aws_iam_role_policy_attachment" "task_exectuion_role" {
  role       = aws_iam_role.task_assume_role.name
  policy_arn = aws_iam_policy.task_execution_role_policy.arn
}


resource "aws_iam_role" "app_task" {
  name               = "${local.prefix}-app-task"
  assume_role_policy = file("./templates/ecs/task-assume-role-policy.json")

}

resource "aws_cloudwatch_log_group" "ecs_task_logs" {
  name = "${local.prefix}-frontend"

}



resource "aws_ecs_cluster" "main" {
  name = "ci-cd-cluster"
}

resource "aws_ecs_task_definition" "frontend" {
  family                   = "${local.prefix}-frontend"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 256
  memory                   = 512

  execution_role_arn = aws_iam_role.task_assume_role.arn

  task_role_arn = aws_iam_role.app_task.arn
  container_definitions = jsonencode([
    {
      name               = "frontend"
      image              = var.ecr_frontend_app_image
      essential          = true
      memeoryReservation = 256
      user               = "root"
      mountPoints = [
        {
          readOnly      = false
          containerPath = "/vol/web/static"
          sourceVolume  = "static"
        }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-group         = aws_cloudwatch_log_group.ecs_task_logs.name
          awslogs-region        = data.aws_region.current.name
          awslogs-stream-prefix = "frontend"
        }
      }
    }
  ])

  volume {
    name = "static"
  }

  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }
}

resource "aws_security_group" "ecs_service" {
  description = "access rules for the ecs service"
  name        = "${local.prefix}-ecs-service"

  # outbound access to the endpoints
  egress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # inbound access via HTTP
  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

}

resource "aws_ecs_service" "frontend" {
  name                   = "${local.prefix}-frontend"
  cluster                = aws_ecs_cluster.main.name
  task_definition        = aws_ecs_task_definition.frontend.family
  desired_count          = 1
  launch_type            = "FARGATE"
  platform_version       = "1.4.0"
  enable_execute_command = true

  network_configuration {
    assign_public_ip = true

    subnets         = ["subnet-0181255b490a4ad97"]
    security_groups = [aws_security_group.ecs_service.id]
  }

}


