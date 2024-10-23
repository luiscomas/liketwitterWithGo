terraform {
    required_providers {
        aws = {
        source  = "hashicorp/aws"
        version = "~> 3.0"
        }
        mongodbatlas = {
            source = "mongodb/mongodbatlas"
            version = "1.21.2"
        }
        
        random = {
            source = "hashicorp/random"
            version = "3.4.3"
        }
        
    }

}

provider "mongodbatlas" {
        public_key = var.atlas_public_key
        private_key = var.atlas_private_key
}