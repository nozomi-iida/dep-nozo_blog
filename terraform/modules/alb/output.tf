output "alb_sg_id" {
  value = aws_security_group.allow_http.id
}

output "zone_id" {
  value = aws_lb.app.zone_id
}

output "dns_name" {
  value = aws_lb.app.dns_name
}
