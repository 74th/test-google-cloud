terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 7.9.0"
    }
  }

  required_version = "~> 1.13.4"
}

variable "gcloud_project" {
  type = string
  description = "Google Cloud Project ID"
}

variable "user_email" {
  type = string
  description = "サービスアカウントのキーを発行できるユーザ"
}

provider "google" {
  region  = "ap-northeast-1"
  project = var.gcloud_project
}

# 短い有効期限検証用のサービスアカウント
resource "google_service_account" "shortterm_sa" {
  account_id   = "test-shortterm-sa"
  display_name = "github.com/74th/test-google-cloud/20251102-shortterm_serviceaccount_key"
}

# サービスアカウントのキーを発行できる権限を付与
resource "google_service_account_iam_member" "key_creator" {
    member = "user:${var.user_email}"
    service_account_id = google_service_account.shortterm_sa.id
    role = "roles/iam.serviceAccountTokenCreator"
}

# 権限検証用のGoogle Cloud Storageバケット
resource "google_storage_bucket" "test_bucket" {
    name     = "${var.gcloud_project}-shortterm-sa-test"
    location = "ASIA-NORTHEAST1"
}

# バケットへの読み込み権限をサービスアカウントに付与
resource "google_storage_bucket_iam_member" "sa_reader" {
    member = "serviceAccount:${google_service_account.shortterm_sa.email}"
    bucket = google_storage_bucket.test_bucket.name
    role   = "roles/storage.objectViewer"
}
