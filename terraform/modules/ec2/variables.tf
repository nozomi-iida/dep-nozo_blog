variable "common_tags" {
  type = map(string)
}

variable "app_name" {
  type = string
}

variable "instance_type" {
  default = "t4g.nano"
}

variable "key_name" {
  default = "nozo_blog_kp"
}

variable "vpc_id" {
  type = string
}

variable "subnet_id" {
  type = string
}

variable "alb_sg_id" {
  type = string
}
