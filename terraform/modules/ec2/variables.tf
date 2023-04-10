variable "common_tags" {
  type = map(string)
}

variable "app_name" {
  type = string
}

variable "public_subnet_ids" {
  type = list(string)
}

variable "alb_sg_id" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "key_name" {
  type = string
}

variable "cluster_name" {
  type = string
}
