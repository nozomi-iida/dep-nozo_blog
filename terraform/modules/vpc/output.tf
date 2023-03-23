output "vpc_id" {
  value = aws_vpc.app.id 
}

output "private_subnet_id" {
  value = aws_subnet.subnet_private_1a.id
}

output "public_subnet_ids" {
  value = [
    aws_subnet.subnet_public_1a.id,
    aws_subnet.subnet_public_1c.id
  ]
}
