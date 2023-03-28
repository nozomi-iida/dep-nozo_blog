resource "aws_vpc" "app" {
  cidr_block = "10.1.0.0/16"
  tags       = merge(var.common_tags, { Name = "${var.app_name}_vpc" })
}

resource "aws_subnet" "subnet_public_1a" {
  vpc_id            = aws_vpc.app.id
  cidr_block        = "10.1.0.0/24"
  availability_zone = "ap-northeast-1a"
  tags              = merge(var.common_tags, { Name = "${var.app_name}_public_subnet_1a" })
}

resource "aws_subnet" "subnet_public_1c" {
  vpc_id            = aws_vpc.app.id
  cidr_block        = "10.1.1.0/24"
  availability_zone = "ap-northeast-1c"
  tags              = merge(var.common_tags, { Name = "${var.app_name}_public_subnet_1c" })
}

resource "aws_subnet" "subnet_private_1a" {
  vpc_id            = aws_vpc.app.id
  cidr_block        = "10.1.2.0/24"
  availability_zone = "ap-northeast-1a"
  tags              = merge(var.common_tags, { Name = "${var.app_name}_private_subnet_1a" })
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.app.id
  tags   = merge(var.common_tags, { Name = "${var.app_name}_igw" })
}

resource "aws_route_table" "rt" {
  vpc_id = aws_vpc.app.id
  tags   = merge(var.common_tags, { Name = "${var.app_name}_rt" })
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }
}

resource "aws_route_table_association" "rta1a" {
  route_table_id = aws_route_table.rt.id
  subnet_id      = aws_subnet.subnet_public_1a.id
}

resource "aws_route_table_association" "rta1c" {
  route_table_id = aws_route_table.rt.id
  subnet_id      = aws_subnet.subnet_public_1c.id
}
