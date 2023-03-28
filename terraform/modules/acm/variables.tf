variable "route53_zone_id" {
  type = string

}
variable "common_tags" {
  type = map(string)
}

variable "domain_name" {
  default = "nozomi-dev.net"
}
