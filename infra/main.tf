terraform {
  backend "s3" {
    profile = "terraform"
    bucket  = "berries-terraform"
    key     = "terraform.tfstate"
    region  = "eu-west-2"
  }
}
module "cloudflare" {
  source = "./cloudflare/"
  ip_address = module.gcp.frappe_ip_address
}

module "gcp" {
    source = "./gcp/"
}

module "unifi" {
  source   = "./unifi/"
  api_url  = "https://10.0.0.1"
  insecure = true
}

module "truenas" {
  source = "./truenas/"
  pool   = "pool_01"
  datasets = [
    {
      name        = "Amrik"
      size_gb     = 1000
      description = "For amrik to put files and stuff"
    },
    {
      name = "Downloads"
    },
    {
      name = "games"
    },
    {
      name = "KateSMB"
    },
    {
      name = "Wormhole"
    },
    {
      name        = "s3"
      size_gb     = 100
      description = "Blob storage"
    }
  ]
}
