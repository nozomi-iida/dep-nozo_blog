provider "aws" {
  region = "ap-northeast-1"
}

locals {
  tags = {
    Terraform = "true"
    Project = "nozo_blog"
  }
}

module "vpc" {
  source = "./modules/vpc"
  common_tags = local.tags
  app_name = var.app_name
}

# module "ec2" {
#   source = "./modules/ec2"
#   common_tags = local.tags
#   app_name = var.app_name
# }

# module "s3" {
#   source = "./modules/s3"
#   common_tags = local.tags
#   app_name = var.app_name
# }

# module "route53" {
#   source = "./modules/route53"
#   common_tags = local.tags
# }

# module "acm" {
#   source = "./modules/acm"
#   route53_zone_id = module.route53.zone_id
#   common_tags = local.tags
# }
