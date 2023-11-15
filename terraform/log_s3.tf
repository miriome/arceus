resource "aws_s3_bucket" "logs_s3" {
  bucket = "miromie-logs-s3"
  force_destroy = true

  tags = {
    Name        = "Logs"
    Environment = "Prod"
  }
}


data "aws_iam_policy_document" "allow_access_from_another_account" {
  statement {
    principals {
      type        = "AWS"
      identifiers = ["114774131450"]
    }

    actions = [
      "s3:PutObject",
    ]

    resources = [
      aws_s3_bucket.logs_s3.arn,
      "${aws_s3_bucket.logs_s3.arn}/*",
    ]
  }
}

resource "aws_s3_bucket_policy" "allow_access_from_another_account" {
  bucket = aws_s3_bucket.logs_s3.id
  policy = data.aws_iam_policy_document.allow_access_from_another_account.json
}