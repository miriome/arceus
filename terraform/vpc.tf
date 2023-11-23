resource "aws_vpc" "miromie-vpc" {

  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "miromie-vpc"
  }
}

data "aws_vpc" "rds-vpc" {
  id = "vpc-00f8d0f9524f4ed1a"
}




resource "aws_security_group" "security_group" {
  name   = "ecs-security-group"
  vpc_id = aws_vpc.miromie-vpc.id

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    self        = "false"
    cidr_blocks = ["0.0.0.0/0"]
    description = "any"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
