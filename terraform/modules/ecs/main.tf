resource "aws_ecr_repository" "app" {
  name = "${var.app_name}_ecr"
  tags = var.common_tags
}

resource "aws_ecs_task_definition" "app" {
  family = "${var.app_name}_ecs_task"
}

resource "aws_ecs_cluster" "app" {
  name = "${var.app_name}_ecs_cluster"
  tags = var.common_tags
}

resource "aws_ecs_service" "app" {
  launch_type = "EC2"
}
