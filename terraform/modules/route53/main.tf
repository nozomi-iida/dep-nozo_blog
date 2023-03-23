resource "aws_route53_zone" "app" {
  name = "${var.domain_name}"
  vpc {
    vpc_id = var.vpc_id
  }

  tags = var.common_tags
}

resource "aws_route53_record" "api" {
  zone_id = aws_route53_zone.app.zone_id
  name = "api-${var.domain_name}"
  type = "A"
  alias {
    name = var.dns_name 
    zone_id = var.zone_id
    evaluate_target_health = false
  }
}
