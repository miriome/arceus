resource "aws_db_instance" "default" {
  identifier = "miromie-db"
  allocated_storage           = 20
  db_name                     = "miromie"
  engine                      = "postgres"
  engine_version              = "15.3"
  instance_class              = "db.t3.micro"
  password = "winsonheng"
  username                    = "root"
  parameter_group_name        = "default.postgres15"
  vpc_security_group_ids = [aws_security_group.security_group.id, aws_security_group.bastion-sg.id]
  db_subnet_group_name = aws_db_subnet_group.rds-subnet-group.id
  skip_final_snapshot = false
#  lifecycle {
#    prevent_destroy = true
#  }
}


## Creating 2 subnets with private ip
resource "aws_subnet" "rds-subnet1" {
  vpc_id = aws_vpc.miromie-vpc.id
  cidr_block = cidrsubnet(aws_vpc.miromie-vpc.cidr_block, 8, 3)
  availability_zone = data.aws_availability_zones.miromie-AZ.names[0]
}

resource "aws_subnet" "rds-subnet2" {
  vpc_id = aws_vpc.miromie-vpc.id
  cidr_block = cidrsubnet(aws_vpc.miromie-vpc.cidr_block, 8, 4)
  availability_zone = data.aws_availability_zones.miromie-AZ.names[1]
}

resource "aws_db_subnet_group" "rds-subnet-group" {
  name = "rds_subnet_group"
  subnet_ids = [aws_subnet.rds-subnet1.id, aws_subnet.rds-subnet2.id]

  tags = {
    Name = "My RDS Subnet Group"
  }
}

resource "aws_security_group" "bastion-sg" {
  description = "Allow SSH to the jumpbox"
  vpc_id      = aws_vpc.miromie-vpc.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "TCP"
    cidr_blocks = ["0.0.0.0/0"]
  }


  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_eip" "bastion-eip" {
  instance = "${aws_instance.bastion.id}"
  domain   = "vpc"

}

resource "aws_instance" "bastion" {
  ami                         = "ami-02453f5468b897e31"
  instance_type               = "t3.micro"
  key_name                    = aws_key_pair.winson-key.key_name
  vpc_security_group_ids      = [aws_security_group.bastion-sg.id]
  subnet_id                   = aws_subnet.miromie-subnet1.id // public subnet
  associate_public_ip_address = true
  tags = {
    Name = "rds-bastion"
  }


}