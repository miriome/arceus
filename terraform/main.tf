# * Part 1 - Setup.

provider "aws" {
  region = "ap-southeast-1"

  default_tags {
    tags = { example = "miromie-arceus" }
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
resource "docker_image" "miromie-app-auth" {
  name = "${aws_ecr_repository.miromie-image-repo.repository_url}:latest"

  build {
    context = "${path.cwd}/app"
    dockerfile = "./auth/Dockerfile"
    platform = "linux/amd64" # currently image id is ecs-optimized x86_64 linux. To change to arm.
  }

  triggers = {
    dir_sha1 = sha1(join("", [for f in fileset(path.cwd, "auth/*") : filesha1(f)]))
  }

}

# Upload to ecr
resource "docker_registry_image" "miromie-app-auth-registry-push" {
  name = docker_image.miromie-app-auth.name

}

# Create ecs cluster
resource "aws_ecs_cluster" "miromie-app-ecs-cluster" {
  name = "miromie-app-ecs"

  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}
# Create ecs task
resource "aws_ecs_task_definition" "app_task" {
  family                   = "app-first-task" # Name your task
  container_definitions    = jsonencode(
  [
    {
      name : "app-first-task",
      image : aws_ecr_repository.miromie-image-repo.repository_url,
      essential : true,
      portMappings : [
        {
          "containerPort": 12345,
          "hostPort": 12345
        }
      ],
      memory : 512,
      cpu : 256
    }
  ]
  )
#  requires_compatibilities = ["EC2"] # Not required if not using fargate
  network_mode             = "awsvpc"    # add the AWS VPN network mode as this is required for Fargate
  memory                   = 512         # Specify the memory the container requires
  cpu                      = 256         # Specify the CPU the container requires
}


data "aws_availability_zones" "miromie-AZ" {

  filter {
    name = "region-name"
    values = ["ap-southeast-1"]
  }

}

## Creating 2 subnets with public ip
resource "aws_subnet" "miromie-subnet1" {
  vpc_id = aws_vpc.miromie-vpc.id
  cidr_block = cidrsubnet(aws_vpc.miromie-vpc.cidr_block, 8, 1)
  map_public_ip_on_launch = true
  availability_zone = data.aws_availability_zones.miromie-AZ.names[0]
}

resource "aws_subnet" "miromie-subnet2" {
  vpc_id = aws_vpc.miromie-vpc.id
  cidr_block = cidrsubnet(aws_vpc.miromie-vpc.cidr_block, 8, 2)
  map_public_ip_on_launch = true
  availability_zone = data.aws_availability_zones.miromie-AZ.names[1]
}

resource "aws_internet_gateway" "miromie-internet_gateway" {
  vpc_id = aws_vpc.miromie-vpc.id
  tags = {
    Name = "miromie-vpc-internet_gateway"
  }
}

# Create routing table to route all requests to our internet gateway.
resource "aws_route_table" "miromie-route_table" {
  vpc_id = aws_vpc.miromie-vpc.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.miromie-internet_gateway.id
  }
}


resource "aws_route_table_association" "subnet_route" {
  subnet_id      = aws_subnet.miromie-subnet1.id
  route_table_id = aws_route_table.miromie-route_table.id
}

resource "aws_route_table_association" "subnet2_route" {
  subnet_id      = aws_subnet.miromie-subnet2.id
  route_table_id = aws_route_table.miromie-route_table.id
}




resource "aws_ecs_service" "ecs_service" {
  name            = "miromie-ecs-service"
  cluster         = aws_ecs_cluster.miromie-app-ecs-cluster.id
  task_definition = aws_ecs_task_definition.app_task.arn
  desired_count   = 2

  network_configuration {
    subnets         = [aws_subnet.miromie-subnet1.id, aws_subnet.miromie-subnet2.id]
    security_groups = [aws_security_group.security_group.id]
  }

  force_new_deployment = true
  placement_constraints {
    type = "distinctInstance"
  }
  
  load_balancer {
    target_group_arn = aws_lb_target_group.ecs_tg_tls.arn
    container_name   = "app-first-task"
    container_port   = 12345
  }
  depends_on = [aws_autoscaling_group.ecs_asg]
}

# Binding capacity provider

resource "aws_ecs_capacity_provider" "ecs_capacity_provider" {
  name = "Miromie-capacity-provider"

  auto_scaling_group_provider {
    auto_scaling_group_arn = aws_autoscaling_group.ecs_asg.arn

    managed_scaling {
      maximum_scaling_step_size = 1000
      minimum_scaling_step_size = 1
      status                    = "ENABLED"
      target_capacity           = 3
    }
  }
}

resource "aws_ecs_cluster_capacity_providers" "example" {
  cluster_name = aws_ecs_cluster.miromie-app-ecs-cluster.name

  capacity_providers = [aws_ecs_capacity_provider.ecs_capacity_provider.name]

  default_capacity_provider_strategy {
    base              = 1
    weight            = 100
    capacity_provider = aws_ecs_capacity_provider.ecs_capacity_provider.name
  }
}



# Launch ec2 instance
# Auto scaling
resource "aws_autoscaling_group" "ecs_asg" {
  vpc_zone_identifier = [aws_subnet.miromie-subnet1.id, aws_subnet.miromie-subnet2.id]
  desired_capacity    = 2
  max_size            = 2
  min_size            = 1

  launch_template {
    id      = aws_launch_template.ecs_lt.id
    version = "$Latest"
  }

  tag {
    key                 = "AmazonECSManaged"
    value               = true
    propagate_at_launch = true
  }
}
# EC2 launch template. image id is ecs-optimized x86_64 linux 2. To change to arm.
resource "aws_launch_template" "ecs_lt" {
  name_prefix   = "ecs-miromie"
  image_id      = "ami-080cdc1184ac6b4fa"
  instance_type = "t3.micro"

#  key_name               = "ec2ecsglog"
  vpc_security_group_ids = [aws_security_group.security_group.id]
  iam_instance_profile {
    name = "ecsInstanceRole"
  }

  block_device_mappings {
    device_name = "/dev/xvda"
    ebs {
      volume_size = 30
      volume_type = "gp2"
    }
  }

  tag_specifications {
    resource_type = "instance"
    tags = {
      Name = "ecs-instance"
    }
  }

  user_data     = base64encode("#!/bin/bash\necho ECS_CLUSTER=${aws_ecs_cluster.miromie-app-ecs-cluster.name} >> /etc/ecs/ecs.config")
  key_name = "winson-key"

}


resource "aws_key_pair" "winson-key" {
  key_name   = "winson-key"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDsnLBoUkc5kmDiMImD37SXHbQseFxwGzZQRCibYtISjqH6A9GRa2dgPan+ipWLhR97Z3TeZ5zQVUawEtl/4fmsTu/K0VkNvvFTbaQQsIFcYILc1IgjEQmMxImPV5bPH6TkR4QOPvCv//flLcv3HqRkzW1bnZeLdC3RO9ao9cpIyaU9qUxV0/EBsNttPljwmGjVEAyiHVl1yGaltUHnSkMKTKxqJPIPn1vHWxJXXJmPM8d+bNQciPHQCaZ3xFuqmpQipYTC45Uw+fjXuNpgE/pbWZgLWwVWVAfpjtI5edeWLIpQSAmrFLBVKSBrR+ZSnvvKKfLQi3do7Cp8gAJFnz5RSZkKO1EwmkLXYKzVadRS3Vfb340I5Yhpx1LwOsh/fKCjTQWbXxSvKUlxcJQJAWBOz3xIajWkCoAOGFtnu72qHE3n3LnL+7AK9fMQ37vMsfpj3HxG7am3D2Gp8UYwa+XDM/RV5iUKwO+cqeJOp9DhCVLye1UrGJOxg3pYLTzKL28= winson.heng@Winsons-MacBook-Air.local"
}




