# tinyurl

## overview

短縮URLを作成するサービス

## Feature

- SHA1ハッシュを使用して入力URLのランダムなハッシュを生成し、同じものが`dynamoDB`データベースの入力URLにマッピングされます。
- 同じURLを何度も短縮しようとしている場合は、既存の短縮URLを返します。
- 短縮URLスラッグは次のもので構成されます。[0-9A-Za-z+_]


## CONTRIBUTING

このプロジェクトに貢献したい方は以下のドキュメントを参考にしてください。
[CONTRIBUTING](./CONTRIBUTING.md)
