variable "common_tags" {
  type = map(string)
}

variable "app_name" {
  type = string
}


variable "alb_tg_id" {
  type = string
}

variable "autoscaling_group_arn" {
  type = string
}
