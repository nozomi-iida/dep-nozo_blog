variable "common_tags" {
  type = map(string)
}

variable "domain_name" {
  default = "nozomi-dev.net"
}

variable "vpc_id" {
  type = string
}

variable "alb_zone_id" {
  type = string
}

variable "alb_dns_name" {
  type = string
}
