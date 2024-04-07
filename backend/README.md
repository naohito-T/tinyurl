# Backend


## memo

- go vet
  - コンパイラでは発見できないバグを見つける
  - go testを走らせれば自動で実行される（Go 1.10から）


要件

 - go testを導入（これでgo vetもできる）
 - go lintは非推奨
   - staticcheckを導入（VSCodeではデフォルト）

SA	staticcheck	コードの正しさに関するチェック
S	simple	コードの簡潔さに関するチェック
ST	stylecheck	コーディングスタイルに関するチェック
QF	quickfix	gopls の自動リファクタリングとして表示されるチェッ
