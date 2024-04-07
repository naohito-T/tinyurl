# Architecture

## Overview

## TODO

- 書き込み
  - frontend
    - url形式かのvalidation
      - validationに失敗したら赤文字で記載する
      - 成功したらトースターを表示する
      - コピペボタンを配置
      - クリアボタンを配置
  - backend
    - url形式かのvalidation
      - validationに失敗したら400
      - 成功したらdynamoに登録
      - レスポンスは短縮URLを返す。
- 読み取り
  - backend
    - リダイレクトを選択できるparameterがある（301 or 302）※defaultは301
    - キャッシュがなかった場合、dynamoへ確認
    - dynamoになかった場合、404
    - dynamoから取得してキャッシュがない場合はキャッシュに入れる。
