# * Part 1 - Setup.
locals {
  container_name = "miromie-arceus"
  container_port = 50001 # ! Must be same port from our Dockerfile that we EXPOSE
  example = "miromie-arceus"
}

provider "aws" {
  region = "ap-southeast-1"

  default_tags {
    tags = { example = local.example }
  }
}

// Create ecr repo
resource "aws_ecr_repository" "miromie-image-repo" {
  name = "miromie-image-repo"
}

# * Give Docker permission to pusher Docker Images to AWS.
data "aws_ecr_authorization_token" "token" {}


provider "docker" {
  registry_auth {
    address = data.aws_ecr_authorization_token.token.proxy_endpoint
    username = data.aws_ecr_authorization_token.token.user_name
    password  = data.aws_ecr_authorization_token.token.password
  }
}

# build docker image
resource "docker_image" "miromie-app" {
  name = "${aws_ecr_repository.miromie-image-repo.repository_url}:latest"

  build {
    context = "${path.cwd}/app"
    dockerfile = "Dockerfile"
  }

  platform = "linux/arm64"
}

resource "docker_registry_image" "miromie-app-registry-push" {
  name = docker_image.miromie-app.name

}


