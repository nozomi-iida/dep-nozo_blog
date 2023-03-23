resource "aws_s3_bucket" "app" {
  bucket = "${var.app_name}-bucket"

  tags = var.common_tags
}

output "bucket_name" {
  value = aws_s3_bucket.app.bucket
}
