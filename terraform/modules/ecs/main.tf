resource "aws_ecr_repository" "app" {
  name = "${var.app_name}_go_api"
  tags = var.common_tags
}

resource "aws_ecs_cluster" "app" {
  name = "${var.app_name}_ecs_cluster"
  tags = var.common_tags
}

resource "aws_ecs_capacity_provider" "app" {
  name = "${var.app_name}_ecs_capacity_provider"
  auto_scaling_group_provider {
    auto_scaling_group_arn = var.autoscaling_group_arn

    managed_scaling {
      status          = "ENABLED"
      target_capacity = 80
    }
  }
}

resource "aws_ecs_cluster_capacity_providers" "app" {
  cluster_name       = aws_ecs_cluster.app.name
  capacity_providers = [aws_ecs_capacity_provider.app.name]
}

resource "aws_cloudwatch_log_group" "app" {
  name = "${var.app_name}_ecs_log_group"
}

resource "aws_ecs_task_definition" "app" {
  family = "${var.app_name}_ecs_task"
  container_definitions = jsonencode([
    {
      name      = "app"
      image     = "${aws_ecr_repository.app.repository_url}:latest"
      essential = true
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-group"         = aws_cloudwatch_log_group.app.name
          "awslogs-region"        = "ap-northeast-1"
          "awslogs-stream-prefix" = "ecs_task"
        }
      }
    }
  ])
  volume {
    name      = "app"
    host_path = "./app"
  }
  memory                   = 128 
  cpu                      = 256
  requires_compatibilities = ["EC2"]
}

resource "aws_ecs_service" "app" {
  name            = "${var.app_name}_ecs_service"
  cluster         = aws_ecs_cluster.app.id
  task_definition = aws_ecs_task_definition.app.arn
  desired_count   = 1
  capacity_provider_strategy {
    capacity_provider = aws_ecs_capacity_provider.app.name
    weight            = 1
  }
  load_balancer {
    target_group_arn = var.alb_tg_id
    container_name   = "app"
    container_port   = 8080
  }
}
