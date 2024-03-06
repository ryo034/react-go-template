# プロフィール写真の変更

## ユーザーは自身のメールアドレスで登録したアカウントでプロフィール写真を更新することができる
* ログイン画面を開く
* メールアドレス"update_me_update_profile_photo@example.com"で始める
* サイドバーのアカウント情報にプロフィール画像が表示されていない
* ホーム画面が表示されている
* 設定のプロフィール設定を開く
* プロフィール設定画面のプロフィール画像をクリック
* 画像"setup/images/test.jpg"をアップロードする
* サイドバーのアカウント情報にプロフィール画像が表示されている
* 画面を更新する
* プロフィール設定画面のプロフィール削除ボタンをクリック
* トーストメッセージ"Photo removed"が表示されている
* 画面を更新する
* サイドバーのアカウント情報にプロフィール画像が表示されていない

## ユーザーは自身のGoogleで登録したアカウントでプロフィール写真を更新することができる
tags: stateful
ユーザーは自身のGoolgeアカウント（プロフィール写真設定済）で追加したプロフィール写真を削除してもGoogleアカウントのプロフィール写真が表示されない
* ログイン画面を開く
* "googleで始める"ボタンをクリック
* Googleアカウント選択画面でメールアドレス"test_has_photo_google_setup_photo@example.com"を選択する
* ホーム画面が表示されている
* サイドバーのアカウント情報にプロフィール画像"https://github.com/ryo034/image/assets/55078625/967e0e8c-a2be-4004-834a-d56a263b89ce"が設定されている
* 設定のプロフィール設定を開く
* プロフィール設定画面のプロフィール画像をクリック
* 画像"setup/images/test.jpg"をアップロードする
* トーストメッセージ"Photo updated"が表示されている
* 画面を更新する
* プロフィール設定画面のプロフィール削除ボタンをクリック
* トーストメッセージ"Photo removed"が表示されている
* 画面を更新する
* サイドバーのアカウント情報にプロフィール画像が表示されていない