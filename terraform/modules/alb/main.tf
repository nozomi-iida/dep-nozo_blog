resource "aws_security_group" "allow_http" {
  name        = "${var.app_name}_alb_sg"
  description = "Allow SSH traffic"
  vpc_id = var.vpc_id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = var.common_tags
}

resource "aws_lb_target_group" "name" {
  port = 80
  protocol = "HTTP"
  vpc_id = var.vpc_id
  tags = merge(var.common_tags, { Name = "${var.app_name}_tg" })
}

resource "aws_lb" "app" {
  internal = false
  load_balancer_type = "application"
  security_groups = [aws_security_group.allow_http.id]
  subnets = var.public_subnet_ids
  tags = merge(var.common_tags, { Name = "${var.app_name}_alb" })
}
