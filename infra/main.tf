terraform {
    backend "s3" {
        profile = "terraform"
            bucket  = "berries-terraform"
            key     = "terraform.tfstate"
            region  = "eu-west-2"
    }
}

variable "pool_name" {
    type = string
}

module "cloudflare" {
    source = "./cloudflare/"
}

module "truenas" {
    source = "./truenas/"
    pool_name = var.pool_name
}

