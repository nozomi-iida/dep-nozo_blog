resource "aws_route53_zone" "app" {
  name = "${var.domain_name}"

  tags = var.common_tags
}

