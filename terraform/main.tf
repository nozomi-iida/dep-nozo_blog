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

module "ec2" {
  source = "./modules/ec2"
  common_tags = local.tags
  app_name = var.app_name
  vpc_id = module.vpc.vpc_id
  subnet_id = module.vpc.private_subnet_id
  alb_sg_id = module.alb.alb_sg_id
}

module "s3" {
  source = "./modules/s3"
  common_tags = local.tags
  app_name = "nozo-blog"
}

module "route53" {
  source = "./modules/route53"
  common_tags = local.tags
  vpc_id = module.vpc.vpc_id
  alb_zone_id = module.alb.zone_id
  alb_dns_name = module.alb.dns_name
}

module "acm" {
  source = "./modules/acm"
  route53_zone_id = module.route53.zone_id
  common_tags = local.tags
}

module "alb" {
  source = "./modules/alb"
  common_tags = local.tags
  app_name = var.app_name
  vpc_id = module.vpc.vpc_id
  public_subnet_ids= module.vpc.public_subnet_ids
  certificate_arn= module.acm.certificate_arn
  instance_id= module.ec2.instance_id
}

