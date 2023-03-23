resource "aws_security_group" "allow_ssh" {
  name        = "${var.app_name}_sg"
  description = "Allow SSH traffic"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = var.common_tags
}

resource "aws_instance" "app" {
  ami           = "ami-0f758aaed03c79cf3" # Amazon Linux 2 LTS AMI
  instance_type = var.instance_type
  key_name      = var.key_name

  vpc_security_group_ids = [aws_security_group.allow_ssh.id]

  tags = merge(var.common_tags, { Name = "${var.app_name}_instance" })
}
