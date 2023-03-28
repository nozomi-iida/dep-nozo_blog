resource "aws_ecr_repository" "app" {
  name = "${var.app_name}_ecr"
  tags = var.common_tags
}

resource "aws_ecs_task_definition" "app" {
  family = "${var.app_name}_ecs_task"
  container_definitions = jsonencode([
    {
      name  = "app"
      image = "${aws_ecr_repository.app.repository_url}:latest"
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
    name = "app"
    host_path = "./app"
  }
  requires_compatibilities = ["EC2"]
}

resource "aws_ecs_cluster" "app" {
  name = "${var.app_name}_ecs_cluster"
  tags = var.common_tags
}

resource "aws_ecs_service" "app" {
  name = "${var.app_name}_ecs_service"
  launch_type = "EC2"
  cluster = aws_ecs_cluster.app.id
  task_definition = aws_ecs_task_definition.app.arn
  desired_count = 1
}
