module "milestone-medical" {
    source = "./modules/aws-account"
    name = "milestone-medical"
    email = "milestone-medical-admin@amrik.co.uk"
    region = "ap-southeast-1"
    secret_key = data.sops_file.aws-secret.data["aws.secret_access_key"]
    access_key = data.sops_file.aws-secret.data["aws.access_key"]
}

output "mm-access-key" {
    value = module.milestone-medical.terraform_access_key
    sensitive = true
}

output "mm-secret-key" {
    value = module.milestone-medical.terraform_secret_key
    sensitive = true
}
