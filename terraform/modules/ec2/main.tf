resource "aws_security_group" "instance" {
  name        = "${var.app_name}_instance_sg"
  description = "Allow SSH traffic"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port       = 80
    to_port         = 80
    protocol        = "tcp"
    security_groups = [var.alb_sg_id]
  }

  tags = var.common_tags
}

resource "aws_launch_template" "app" {
  name                   = "${var.app_name}_launch_template"
  image_id               = "ami-0ec1b47781bc9d6d1"
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.instance.id]
  key_name               = var.key_name
}

resource "aws_autoscaling_group" "app" {
  name                = "${var.app_name}_autoscaling_group"
  desired_capacity    = 1
  max_size            = 1
  min_size            = 1
  vpc_zone_identifier = var.public_subnet_ids

  launch_template {
    id      = aws_launch_template.app.id
    version = "$Latest"
  }
}
