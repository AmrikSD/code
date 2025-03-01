terraform {
  backend "s3" {
    profile = "terraform"
    bucket  = "berries-terraform"
    key     = "terraform.tfstate"
    region  = "eu-west-2"
  }
}
module "cloudflare" {
  source     = "./cloudflare/"
  ip_address = module.gcp.frappe_ip_address
}

module "gcp" {
  source = "./gcp/"
}

module "kubernetes" {
    source = "./kubernetes/"
    github_repository = "code"
    github_org = "AmrikSD"
}

module "unifi" {
  source   = "./unifi/"
  api_url  = "https://10.0.0.1"
  insecure = true
  shellys = [
    {
      mac  = "a0:dd:6c:4e:b2:34"
      name = "Shelly Plus 2PM (Entrance)"
    },
    {
      mac  = "78:ee:4c:c3:c0:48"
      name = "Shelly Plus 1PM (Downstairs Room)"
    },
    {
      mac  = "a0:dd:6c:48:14:00"
      name = "Shelly Plus 2PM (Living Room)"
    },
    {
      mac = "a0:dd:6c:4e:cc:88"
      name = "Shelly Plus 2PM (Kate Office)"
    },
    {
      mac = "a0:dd:6c:2e:63:90"
      name = "Shelly Plus 2PM (Master Bedroom)"
    },
    {
      mac  = "a0:dd:6c:47:f4:a8"
      name = "Shelly Plus 2PM (Dining Room)"
    },
    {
      mac  = "00:08:22:3f:ca:23"
      name = "Shelly WallDisplay (Upstairs)"
    },
    {
      mac  = "a0:dd:6c:4e:5b:34"
      name = "Shelly Plus 2PM (TV, Left)"
    },
    {
      mac  = "d0:ef:76:c7:5e:d0"
      name = "Shelly Plus 2PM (Hallway)"
    },
    {
      mac  = "a0:dd:6c:4f:94:10"
      name = "Shelly Plus 2PM (Bifold)"
    },
    {
      mac  = "00:08:22:52:6f:44"
      name = "Shelly WallDisplay (Kate Office)"
    },
    {
      mac  = "00:08:22:3d:0f:55"
      name = "Shelly WallDisplay (Entrance)"
    },
    {
      mac  = "d0:ef:76:c8:b2:24"
      name = "Shelly Plus 2PM (Back Door)"
    },
    {
      mac  = "78:ee:4c:c3:f3:28"
      name = "Shelly Plus 1PM (Downstairs Bathroom)"
    },
    {
      mac  = "78:ee:4c:d5:d7:b0"
      name = "Shelly Plus 1PM (Utility)"
    },
    {
      mac  = "a0:dd:6c:4f:08:14"
      name = "Shelly Plus 2PM (Amrik Office)"
    },
    {
      mac  = "78:ee:4c:c3:e9:84"
      name = "Shelly Plus 1PM (Under Stairs)"
    },
    {
      mac  = "78:ee:4c:d0:78:28"
      name = "Shelly Plus 1PM (Family Bathroom)"
    },
    {
      mac  = "a0:dd:6c:4f:03:f8"
      name = "Shelly Plus 2PM (Stairs)"
    },
    {
      mac  = "a0:dd:6c:4f:92:80"
      name = "Shelly Plus 2PM (TV, Right)"
    },
    {
      mac  = "78:ee:4c:c9:96:38"
      name = "Shelly Plus 1PM (Guest Bedroom)"
    }
  ]
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
    },
    {
        name = "media-server-config"
        description = "For the *arr apps to store config"
    },
  ]
}
