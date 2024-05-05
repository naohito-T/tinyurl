#!/bin/bash

# -----------------------------------------------------------
# @desc localStack init scripts
# ホストからマウントすることで初期化処理が実行可能
# @see https://docs.localstack.cloud/references/init-hooks/#usage-example
# @note LOCALSTACK_SERVICESに対象のサービスを記載しないと動作しない。
# -----------------------------------------------------------

echo "init localstack"

echo "init dynamoDB"
# 指定のURLをhashしたやつをprimary keyにする
awslocal dynamodb create-table --table-name offline-tinyurls \
  --attribute-definitions \
        AttributeName=id,AttributeType=S \
  --key-schema \
        AttributeName=id,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST

awslocal dynamodb create-table --table-name offline-rate-limit \
  --attribute-definitions AttributeName=key,AttributeType=S \
  --key-schema AttributeName=key,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST

awslocal dynamodb update-time-to-live \
  --table-name offline-rate-limit \
  --time-to-live-specification Enabled=true,AttributeName=expired_at_sec
echo 'dynamodb initialized.'

echo "Localstack setup completed!"
