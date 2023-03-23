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
