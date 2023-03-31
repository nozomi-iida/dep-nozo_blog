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
  }
}

resource "aws_ecs_cluster_capacity_providers" "app" {
  cluster_name       = aws_ecs_cluster.app.name
  capacity_providers = [aws_ecs_capacity_provider.app.name]
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
    }
  ])
  volume {
    name      = "app"
    host_path = "./app"
  }
  memory                   = 512
  requires_compatibilities = ["EC2"]
}

resource "aws_ecs_service" "app" {
  name            = "${var.app_name}_ecs_service"
  launch_type     = "EC2"
  cluster         = aws_ecs_cluster.app.id
  task_definition = aws_ecs_task_definition.app.arn
  desired_count   = 1
  load_balancer {
    target_group_arn = var.alb_tg_id
    container_name   = "app"
    container_port   = 8080
  }
}
