# アカウントのメールアドレス変更

tags: statefulAll

* メールアドレス認証済ユーザーでログインする
* サイドバーメニューから"アカウント情報"を選択

## ユーザーは不正なメールアドレスに変更しようとした場合に適切なエラーメッセージを確認することができる
* アカウント設定画面のメールアドレス編集ボタンをクリック
* 入力欄"新しいメールアドレス"に"system_account@example.com"と入力する
* "変更する"ボタンをクリック
* エラーメッセージ"すでにそのメールアドレスは使用されています"が表示されている

## ユーザーはメールアドレスを変更することができる
* アカウント設定画面のメールアドレスが"unfinished_onboarding@example.com"
* アカウント設定画面のメールアドレス編集ボタンをクリック
* 入力欄"新しいメールアドレス"に"test+12@example.com"と入力する
* "変更する"ボタンをクリック
* 現在のURLのパスが"/login"
* 入力欄"メールアドレス"に"test+12@example.com"と入力する
* 入力欄"パスワード"に"Test123"と入力する
* ログインボタンをクリック
* 現在のURLのパスが"/confirm-email"
