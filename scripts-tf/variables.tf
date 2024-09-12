variable "app_name" {
  type        = string
  nullable    = false
  description = "The app name"
}

variable "environment" {
  type        = string
  nullable    = false
  description = "The environment to deploy the infrastructure"
}

variable "aws_default_tags" {
  type        = map(string)
  description = "Default tags for all AWS resources"
}

variable "aws_region" {
  description = "The AWS region to use"
  default = "eu-west-1"
}

variable "aws_vpc_id" {
  description = "The VPC to use"
  default = "vpc-0af9ec519336f57e4"
}

variable "aws_vpc_subnets" {
  description = "The VPC subnets to use"
  type    = list(string)
  default = ["subnet-0c6dabff289245526", "subnet-0b829409571337f0e"]
}

variable "aws_profile" {
    description = "The IAM profile used for SSO"
}

variable "service_a_app_image" {
    description = "Docker image to run in the ECS cluster for service-a"
    default = "ghcr.io/majidmvulle/ps-exercise/service-a:latest"
}

variable "service_b_app_image" {
    description = "Docker image to run in the ECS cluster for service-b"
    default = "ghcr.io/majidmvulle/ps-exercise/service-b:latest"
}

variable "service_a_app_port" {
    description = "Port exposed by the service-a docker image to redirect traffic to"
    default = 3000
}

variable "service_b_app_port" {
    description = "Port exposed by the service-b docker image to redirect traffic to"
    default = 8080
}

variable "health_check_path" {
  default = "/health"
}

variable "fargate_cpu" {
    description = "Fargate instance CPU units to provision (1 vCPU = 1024 CPU units)"
    default = 1024
}

variable "fargate_memory" {
    description = "Fargate instance memory to provision (in MiB)"
    default = 2048
}

variable "service_b_api_key" {
  description = "API Key to authenticate to service B"
}

variable "service_a_api_key" {
  description = "API Key to authenticate to service A"
}

variable "service_b_dsn" {
  description = "Database DSN for service B"
}

variable "service_b_database_name" {
  description = "Database name for service B"
  default = poolside-exercise
}
