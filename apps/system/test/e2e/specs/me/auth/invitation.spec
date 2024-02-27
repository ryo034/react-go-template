# 招待を受諾する
tags: invitation

観点
- 登録済みかどうか
- 認証済みかどうか
- ワークスペースに所属しているかどうか
- 途中離脱した場合はどうか
- 招待時に設定された情報でユーザーが登録されるかどうか
  - 表示名
- 招待の有効性
  - 有効期限
  - 既に招待を受諾しているかどうか
  - 既にキャンセルされているかどうか
  - 既に一度招待フローに乗っているかどうか

## 招待を受けた未登録ユーザーは招待を受諾してワークスペースに参加することができる
tags: stateful
* トークン"018d96b7-df68-792f-97d0-d6a044c2b4a2"の招待画面を開く
* "始める"ボタンをクリック
* メールアドレス"invite_test_not_expired@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
// オンボーディング
* 入力欄"名前"に"Invitation Test"と入力する
* "送信"ボタンをクリック
// 招待の受諾
* 招待受諾画面でワークスペース"Example"招待者"John Doe"の参加ボタンをクリック
// 所属情報の確認
* 所属中の表示名が"Invitation Test"である
* 選択中のワークスペース名が"Example"である

## 招待を受けた登録済ユーザーは招待を受諾してワークスペースに参加することができる
tags: stateful
* トークン"018d9fb5-7e56-75ed-952f-ae8aa4fed8c6"の招待画面を開く
* "始める"ボタンをクリック
* メールアドレス"invite_test_already_joined_any_workspace@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
* 招待受諾画面でワークスペース"Example"招待者"John Doe"の参加ボタンをクリック
* 選択中のワークスペース名が"Example"である
* 所属中の表示名が"Invite TestTwo"である

## 認証済みユーザーが招待を受諾した場合、表示中のワークスペースが受諾したワークスペースに切り替わる
tags: stateful
* ログイン画面を開く
* メールアドレス"invite_test_already_joined_any_workspace@example.com"で始める
* 所属中の表示名が"Invite TestTwo"である
* 選択中のワークスペース名が"InviteTest 2"である
* トークン"018d9fb5-7e56-75ed-952f-ae8aa4fed8c6"の招待画面を開く
* 招待画面で"invite_test_already_joined_any_workspace@example.com"としてログインしていることがわかる
* "始める"ボタンをクリック
* メールアドレス"invite_test_already_joined_any_workspace@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
* 招待受諾画面で"invite_test_already_joined_any_workspace@example.com"としてログインしていることがわかる
* 招待受諾画面でワークスペース"Example"招待者"John Doe"の参加ボタンをクリック
* 選択中のワークスペース名が"Example"である

## 招待時に表示名が設定されている場合、その表示名でユーザーが登録される
tags: stateful
* トークン"018da09e-2fa7-7d3a-ad23-2c9f5cb76b92"の招待画面を開く
* "始める"ボタンをクリック
* メールアドレス"invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
* 招待受諾画面でワークスペース"Example"招待者"Invite TestOne"の参加ボタンをクリック
* 選択中のワークスペース名が"Example"である

## ユーザーは招待中にログアウトすることができる
* トークン"018d9fb5-7e56-75ed-952f-ae8aa4fed8c6"の招待画面を開く
* "始める"ボタンをクリック
* メールアドレス"invite_test_already_joined_any_workspace@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
* 招待画面で"invite_test_already_joined_any_workspace@example.com"としてログインしていることがわかる
* 招待中画面にログアウトする
* ログイン画面が表示されている
