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

data "aws_ssm_parameter" "ecs_ami" {
  name = "/aws/service/ecs/optimized-ami/amazon-linux-2/recommended"
}

resource "aws_launch_template" "app" {
  name                   = "${var.app_name}_launch_template"
  image_id               = jsondecode(data.aws_ssm_parameter.ecs_ami.value)["image_id"]
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.instance.id]
  key_name               = var.key_name
  tags                   = var.common_tags
}

resource "aws_autoscaling_group" "app" {
  name                = "${var.app_name}_autoscaling_group"
  desired_capacity    = 1
  max_size            = 3
  min_size            = 1
  vpc_zone_identifier = var.public_subnet_ids
  tag {
    key                 = "Name"
    value               = "${var.app_name}_instance"
    propagate_at_launch = true
  }

  launch_template {
    id      = aws_launch_template.app.id
    version = "$Latest"
  }
}
