terraform {
  backend "s3" {
    bucket         = "hidensneak-terraform"
    key            = "filename.tfstate"
    dynamodb_table = "terraform-state-lock-dynamo"
    region         = "us-east-1"
    encrypt        = true
  }
}

module "doDropletDeploy1" {
  source             = "modules/droplet-deployment"
  do_region_count    = "${map("NYC1","2")}"
  do_token           = "${var.do_token}"
  do_image           = "ubuntu-16-04-x64"
  do_private_key     = "/Users/mike.hodges/.ssh/do_rsa"
  do_ssh_fingerprint = "b3:b2:c7:b1:73:9e:28:c6:61:8d:15:e1:0e:61:7e:35"
  do_size            = "512MB"
  do_default_user    = "root"
}
