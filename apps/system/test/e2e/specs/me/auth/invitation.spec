# 招待を受諾する
tags: invitation

同時に招待をもらうことはあまりないと思われるため、可能性はあるが考慮しない。

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

## 招待を受けた未登録ユーザーはメールアドレスで招待を受諾してワークスペースに参加することができる
tags: stateful
* トークン"018d96b7-df68-792f-97d0-d6a044c2b4a2"の招待画面を開く
* "メールアドレスで始める"ボタンをクリック
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

## 招待を受けた登録済ユーザーはメールアドレスで招待を受諾してワークスペースに参加することができる
tags: stateful
* トークン"018d9fb5-7e56-75ed-952f-ae8aa4fed8c6"の招待画面を開く
* "メールアドレスで始める"ボタンをクリック
* メールアドレス"invite_test_already_joined_any_workspace@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
* 招待受諾画面でワークスペース"InviteTest 1"招待者"Invite TestOne"の参加ボタンをクリック
* 選択中のワークスペース名が"InviteTest 1"である
* 所属中の表示名が"Invite TestTwo"である

## 認証済みユーザーが招待を受諾した場合、表示中のワークスペースが受諾したワークスペースに切り替わる
tags: stateful
* ログイン画面を開く
* メールアドレス"invite_test_already_joined_any_workspace@example.com"で始める
* 所属中の表示名が"Invite TestTwo"である
* 選択中のワークスペース名が"InviteTest 2"である
* トークン"018d9fb5-7e56-75ed-952f-ae8aa4fed8c6"の招待画面を開く
* "メールアドレスで始める"ボタンをクリック
* メールアドレス"invite_test_already_joined_any_workspace@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
* 招待受諾画面で"invite_test_already_joined_any_workspace@example.com"としてログインしていることがわかる
* 招待受諾画面でワークスペース"InviteTest 1"招待者"Invite TestOne"の参加ボタンをクリック
* 選択中のワークスペース名が"InviteTest 1"である

## 招待時に表示名が設定されている場合、その表示名でユーザーが登録される
tags: stateful
* トークン"018da09e-2fa7-7d3a-ad23-2c9f5cb76b92"の招待画面を開く
* "メールアドレスで始める"ボタンをクリック
* メールアドレス"invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
* 招待受諾画面でワークスペース"InviteTest 1"招待者"Invite TestOne"の参加ボタンをクリック
* 選択中のワークスペース名が"InviteTest 1"である

## ユーザーは招待中にログアウトすることができる
* トークン"018d9fb5-7e56-75ed-952f-ae8aa4fed8c6"の招待画面を開く
* "メールアドレスで始める"ボタンをクリック
* メールアドレス"invite_test_already_joined_any_workspace@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
* 招待画面で"invite_test_already_joined_any_workspace@example.com"としてログインしていることがわかる
* 招待中画面にログアウトする
* ログイン画面が表示されている

=========== Google認証に関するテスト ===========

## 招待を受けたユーザー（メールアドレスでアカウント作成済）は招待を受けていないgoogleアカウントで招待を受諾できない
tags: stateful
* トークン"018d9fb5-7e56-75ed-952f-ae8aa4fed8c6"の招待画面を開く
* "Googleで始める"ボタンをクリック
* Googleアカウント選択画面でメールアドレス"invite_test_no_name_google_auth_with_display_name_when_invite@example.com"を選択する
* トーストメッセージ"招待の受諾に失敗しました。お手数ですが、しばらくしてから再度お試しください"が表示されている

## 招待を受けた未登録ユーザーは名前の設定されていないgoogleアカウントで招待を受諾することができる
* トークン"018df2fa-4598-7e13-af4d-7727a9bca288"の招待画面を開く
* "Googleで始める"ボタンをクリック
* Googleアカウント選択画面でメールアドレス"invite_test_no_name_google_auth_with_display_name_when_invite@example.com"を選択する
* 入力欄"名前"に"NoNameGoogleAuth HasDisplayName"と入力する
* "送信"ボタンをクリック
* 招待受諾画面でワークスペース"InviteTest 1"招待者"Invite TestOne"の参加ボタンをクリック
* 選択中のワークスペース名が"InviteTest 1"である
* 所属中の表示名が"NoNameGoogleAuth HasDisplayName"である

## 招待を受けた未登録ユーザーは名前の設定されているgoogleアカウントで招待を受諾することができる
* トークン"018df2fa-2dc2-79ea-8913-e45e39379c9c"の招待画面を開く
* "Googleで始める"ボタンをクリック
* Googleアカウント選択画面でメールアドレス"invite_test_has_name_google_auth_no_name_when_invite@example.com"を選択する
* 招待受諾画面でワークスペース"InviteTest 1"招待者"Invite TestOne"の参加ボタンをクリック
* 選択中のワークスペース名が"InviteTest 1"である
* 所属中の表示名が"InviteGoogleAuthTest NoNameTest"である

## 招待を受けたユーザー（メールアドレスでアカウント作成済）はgoogleアカウントで招待を受諾することができる
* トークン"018df53b-82a2-7324-9b26-f17496bfcdf8"の招待画面を開く
* "Googleで始める"ボタンをクリック
* Googleアカウント選択画面でメールアドレス"invite_test_already_joined_any_workspace_by_email@example.com"を選択する
* 招待受諾画面でワークスペース"InviteTest 1"招待者"Invite TestOne"の参加ボタンをクリック
* 選択中のワークスペース名が"InviteTest 1"である
* 所属中の表示名が"InviteGoogleAuthTest AlreadyJoined"である

## 招待を受けたユーザー（googleアカウントでアカウント作成済）はgoogleアカウントで招待を受諾することができる
* トークン"018df54f-e057-7818-8c72-80d6393e39e6"の招待画面を開く
* "Googleで始める"ボタンをクリック
* Googleアカウント選択画面でメールアドレス"invite_test_already_joined_any_workspace_by_google@example.com"を選択する
* 招待受諾画面でワークスペース"InviteTest 1"招待者"Invite TestOne"の参加ボタンをクリック
* 選択中のワークスペース名が"InviteTest 1"である
* 所属中の表示名が"InviteGoogleAuthTest AlreadyJoinedGoogle"である

## 招待を受けたユーザー（プロフィール写真の設定されているgoogleアカウントでアカウント作成済）はメールアドレスで招待を受諾した場合、元のプロフィール画像が保持される
* トークン"018e0eae-aea6-74e5-8bd6-288b480b335a"の招待画面を開く
* "Googleで始める"ボタンをクリック
* Googleアカウント選択画面でメールアドレス"invite_test_has_photo_by_google_accept_with_email@example.com"を選択する
* 招待受諾画面でワークスペース"InviteTest 2"招待者"Invite TestTwo"の参加ボタンをクリック
* 選択中のワークスペース名が"InviteTest 2"である
* 所属中の表示名が"InviteTest HasPhotoAcceptWithEmail"である
* サイドバーのアカウント情報にプロフィール画像"https://github.com/ryo034/image/assets/55078625/af9fae15-baf3-451e-820a-99f7e246af31"が設定されている

## 招待を受けたユーザー（新たにプロフィール写真を設定しているgoogleアカウントでアカウント作成済）はメールアドレスで招待を受諾した場合、元のプロフィール画像が保持される
* トークン"018e0ec0-e54e-7476-b65c-220bfafbf631"の招待画面を開く
* "メールアドレスで始める"ボタンをクリック
* メールアドレス"invite_test_has_photo_by_google_accept_with_has_photo_google@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
* 招待受諾画面でワークスペース"InviteTest 2"招待者"Invite TestTwo"の参加ボタンをクリック
* 選択中のワークスペース名が"InviteTest 2"である
* 所属中の表示名が"InviteTest HasPhotoAcceptWithGoogle"である
* ホーム画面が表示されている
* サイドバーのアカウント情報にプロフィール画像"https://github.com/ryo034/image/assets/55078625/ddeb3605-2291-4c19-81ec-6d890c7d0219"が設定されている

## 招待を受けた未登録ユーザーはプロフィール写真が設定されているgoogleアカウントで招待を受諾した場合、googleアカウントのプロフィール画像が設定される
* トークン"018e0ec2-0d64-7b7d-92af-e42be382216c"の招待画面を開く
* "Googleで始める"ボタンをクリック
* Googleアカウント選択画面でメールアドレス"invite_test_no_account_accept_with_has_photo_google@example.com"を選択する
* 招待受諾画面でワークスペース"InviteTest 2"招待者"Invite TestTwo"の参加ボタンをクリック
* 選択中のワークスペース名が"InviteTest 2"である
* 所属中の表示名が"InviteTest HasNotPhotoAcceptWithGoogle"である
* サイドバーのアカウント情報にプロフィール画像"https://github.com/ryo034/image/assets/55078625/8db77a30-17d4-4c78-8a3a-a88ba016dcfb"が設定されている

## すでに退出済みのメンバーから招待を受けた未登録ユーザーはメールアドレスでアカウントを登録することができる
* トークン"018e4421-aba7-7de4-9bc7-ca0f93355a28"の招待画面を開く
* "メールアドレスで始める"ボタンをクリック
* メールアドレス"once_leave_workspace_accept_receive_from_already_left_member@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
* オンボーディングの名前入力画面が表示されている
* 入力欄"名前"に"test"と入力する
* "送信"ボタンをクリック
* 招待受諾画面でワークスペース"Once Leave Workspace Invite"招待者"Removed User"の参加ボタンをクリック
* 選択中のワークスペース名が"Once Leave Workspace Invite"である
* 所属中の表示名が"test"である
