resource "aws_security_group" "instance" {
  name        = "${var.app_name}_instance_sg"
  description = "Allow SSH & ALB traffic"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    security_groups = [var.alb_sg_id]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  tags = var.common_tags
}

data "aws_ssm_parameter" "ecs_ami" {
  name = "/aws/service/ecs/optimized-ami/amazon-linux-2/recommended"
}

resource "aws_iam_role" "app" {
  name = "${var.app_name}-ecs-instance-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_instance_policy_attachment" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
  role       = aws_iam_role.app.name
}

resource "aws_iam_instance_profile" "app" {
  name = "${var.app_name}-ecs-instance-profile"
  role = aws_iam_role.app.name
}

resource "aws_launch_template" "app" {
  name          = "${var.app_name}_launch_template"
  image_id      = jsondecode(data.aws_ssm_parameter.ecs_ami.value)["image_id"]
  instance_type = "t2.micro"
  key_name      = var.key_name
  tags          = var.common_tags
  iam_instance_profile {
    name = aws_iam_instance_profile.app.name
  }


  network_interfaces {
    associate_public_ip_address = true
    security_groups             = [aws_security_group.instance.id]
  }

  user_data = base64encode(<<-EOF
    #!/bin/bash
    echo ECS_CLUSTER=${var.cluster_name} >> /etc/ecs/ecs.config
    EOF
  )
}

resource "aws_autoscaling_group" "app" {
  name                = "${var.app_name}_autoscaling_group"
  desired_capacity    = 1
  max_size            = 1
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
