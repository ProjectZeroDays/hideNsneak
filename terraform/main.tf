terraform {
  backend "s3" {
    bucket         = "hidensneak-terraform"
    key            = "filename.tfstate"
    dynamodb_table = "terraform-state-lock-dynamo"
    region         = "us-east-1"
    encrypt        = true
  }
}

module "aws-us-east-1" {
  source               = "modules/ec2-deployment"
  default_sg_name      = "allow_ssh"
  aws_sg_id            = "sg-1dea3956"
  region_count         = 3
  custom_ami           = ""
  aws_instance_type    = ""
  ec2_default_user     = "ubuntu"
  aws_access_key       = "${var.aws_access_key}"
  aws_secret_key       = "${var.aws_secret_key}"
  aws_region           = "us-east-1"
  aws_new_keypair      = false
  aws_keypair_name     = "do_rsa"
  aws_private_key_file = "/Users/mike.hodges/.ssh/do_rsa"
  aws_public_key_file  = "/Users/mike.hodges/.ssh/do_rsa.pub"
  ansible_groups       = []
}