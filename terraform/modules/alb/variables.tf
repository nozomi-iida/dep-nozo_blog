variable "common_tags" {
  type = map(string)
}

variable "app_name" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "certificate_arn" {
  type = string
}

variable "public_subnet_ids" {
  type = list(string)
}
