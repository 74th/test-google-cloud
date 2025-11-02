# 短い有効期限のサービスアカウントの利用

## 準備

オーナー権限持つアカウントで準備

```
gcloud auth application-default login
```

```
cd terraform
terraform init
terraform apply
```

```
gcloud storage cp test_gcs_file.json gs://${GCLOUD_PROJECT}-shortterm-sa-test/test_gcs_file.json
```

## 短いサービスアカウントのトークンを発効する

```
make build
```

```
./create-short-term-key/create-short-term-key \
  --service-account-email=test-shortterm-sa@${GCLOUD_PROJECT}.iam.gserviceaccount.com
```

gcloudコマンドでもできた

```
gcloud auth print-access-token \
  --project=${GCLOUD_PROJECT} \
  --lifetime=3600 \
  --impersonate-service-account=test-shortterm-sa@${GCLOUD_PROJECT}.iam.gserviceaccount.com
```

access-tokenは取得できるが、OAuth2トークンのため、Application Default Credentials (ADC) のように簡単には使えない。
というか、ADCからOAuth2トークンを作るロングタームなものなので、逆はできなさそう。

gcloudでOAuth2トークンを使う方法はなさそう。
PythonやGoのSDKでは使うことはできるが、ひと工夫いる（デフォルトがADCなので）。

https://developers.google.com/identity/protocols/oauth2?hl=ja

3600秒（1時間）有効なアクセストークンを取得できるが、それ以上に伸ばす場合は「組織の設定」で伸ばせるように登録する必要がある。
個人のGoogle Cloudではできない。

## その他

### サービスアカウントになってログインできる（有効期限なし）

```
gcloud auth application-default login \
  --impersonate-service-account=test-shortterm-sa@${GCLOUD_PROJECT}.iam.gserviceaccount.com
```
