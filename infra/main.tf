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
}

module "truenas" {
    source = "./truenas/"
    pool   = "pool_01"
    datasets = [
        {
            name = "Amrik"
            size_gb = 1000
            description = "For amrik to put files and stuff"
        },
        {
            name = "AmrikSMB"
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
    ]
}

