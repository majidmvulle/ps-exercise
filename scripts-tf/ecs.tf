resource "aws_ecs_cluster" "poolside-exercise" {
  name = "poolside-exercise"
}

resource "aws_ecs_task_definition" "service-b" {
  family                   = "service-b"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.fargate_cpu
  memory                   = var.fargate_memory
  container_definitions = jsonencode([
    {
      name      = "service-b"
      image     = var.service_b_app_image
      cpu       = var.fargate_cpu
      memory    = var.fargate_memory
      essential = true
      portMappings = [
        {
          containerPort = var.service_b_app_port
          hostPort      = var.service_b_app_port
        }
      ],
      environment = {
        GIN_PORT: 8080,
        APP_NAME: "service-b",
        API_KEY: var.service_b_api_key,
        APP_ENV: release,
        GIN_MODE: release,
        DATABASE_DSN=var.service_b_dsn
        DATABASE_DRIVER=postgres
        DATABASE_NAME=var.service_b_database_name
      }
    }
  ])
}

resource "aws_ecs_task_definition" "service-a" {
  family                   = "service-a"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.fargate_cpu
  memory                   = var.fargate_memory
  container_definitions = jsonencode([
    {
      name      = "service-b"
      image     = var.service_a_app_image
      cpu       = var.fargate_cpu
      memory    = var.fargate_memory
      essential = true
      portMappings = [
        {
          containerPort = var.service_a_app_port
          hostPort      = var.service_a_app_port
        }
      ],
      environment = {
        APP_PORT: 3000,
        APP_NAME: "service-a",
        API_KEY: var.service_a_api_key,
        SERVICE_B_URL: "http://localhost:8080"
        SERVICE_B_API_KEY=var.service_b_api_key
      }
    }
  ])
}


resource "aws_security_group" "service-a" {
  name = "service-a-public"

  egress {
    from_port        = 0
    to_port          = 0
    protocol       = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  ingress {
    protocol       = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
    from_port         = 3000
    to_port           = 3000
  }
}

resource "aws_security_group" "service-b" {
  name = "service-b-private"

  egress {
    from_port        = 0
    to_port          = 0
    protocol       = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  ingress {
    security_groups   = [aws_security_group.service-a.id]
    protocol          = "tcp"
    from_port         = 8080
    to_port           = 8080
  }
}

data "aws_ecs_task_execution" "service-b" {
  cluster         = aws_ecs_cluster.poolside-exercise.id
  task_definition = aws_ecs_task_definition.service-b.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = var.aws_vpc_subnets
    security_groups  = [aws_security_group.service-b.id]
    assign_public_ip = true
  }
}

data "aws_ecs_task_execution" "service-a" {
  cluster         = aws_ecs_cluster.poolside-exercise.id
  task_definition = aws_ecs_task_definition.service-a.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = var.aws_vpc_subnets
    security_groups  = [aws_security_group.service-b.id]
    assign_public_ip = false
  }
}
