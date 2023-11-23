# Configure lb

data "aws_acm_certificate" "api-miromie-issued-cert" {
  domain   = "api.miromie.com"
  statuses = ["ISSUED"]
}

resource "aws_lb" "ecs_alb" {
  name               = "miromie-ecs-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.security_group.id]
  subnets            = [aws_subnet.miromie-subnet1.id, aws_subnet.miromie-subnet2.id]

  access_logs {
    bucket  = aws_s3_bucket.logs_s3.id
    prefix  = "miromie-ecs-lb"
    enabled = true
  }

  tags = {
    Name = "miromie-ecs-alb"
  }
}

resource "aws_lb_target_group" "ecs_tg_tls" {
  name        = "miromie-app-ecs-tg"
  port        = 12345
  protocol    = "HTTP"
  protocol_version = "GRPC"
  target_type = "ip"
  vpc_id      = aws_vpc.miromie-vpc.id
  lifecycle {
       create_before_destroy = true
  }

  health_check {
    path = "/grpc.health.v1.Health/Check"
    matcher = "0"
    port = "12345"
  }
}



resource "aws_lb_listener" "miromie_ecs_alb_listener_https" {
  load_balancer_arn = aws_lb.ecs_alb.arn
  port              = 443
  protocol          = "HTTPS"
  certificate_arn = data.aws_acm_certificate.api-miromie-issued-cert.arn
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.ecs_tg_tls.arn
  }
}
