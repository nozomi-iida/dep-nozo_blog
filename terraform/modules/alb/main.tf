resource "aws_security_group" "allow_http" {
  name        = "${var.app_name}_alb_sg"
  description = "Allow SSH traffic"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = var.common_tags
}

resource "aws_lb_target_group" "name" {
  port     = 80
  protocol = "HTTP"
  vpc_id   = var.vpc_id
  health_check {
    path = "/health"
  }
  tags = merge(var.common_tags, { Name = "${var.app_name}_tg" })
}

resource "aws_lb_target_group_attachment" "app" {
  target_group_arn = aws_lb_target_group.name.arn
  target_id        = var.instance_id
  port             = 80
}

resource "aws_lb" "app" {
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.allow_http.id]
  subnets            = var.public_subnet_ids
  tags               = merge(var.common_tags, { Name = "${var.app_name}_alb" })
}

resource "aws_lb_listener" "app" {
  load_balancer_arn = aws_lb.app.arn
  port              = "443"
  protocol          = "HTTPS"
  certificate_arn   = var.certificate_arn
  default_action {
    target_group_arn = aws_lb_target_group.name.arn
    type             = "forward"
  }
}

