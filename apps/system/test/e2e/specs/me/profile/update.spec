# プロフィールの更新

## ユーザーは自身のアカウントプロフィール情報を更新することができる
tags: stateful
* ログイン画面を開く
* メールアドレス"update_me_member_profile@example.com"で始める
TODO: assertを考える
* "1"秒待機
* 設定のアカウント設定を開く
* 入力欄"Name"に"UpdateMe MemberProfile"と入力されている
* 入力欄"Name"に"Updated UserName"と入力する
* "Update account"ボタンをクリック
* トーストメッセージ"Profile updated"が表示されている
* 画面を更新する
* 入力欄"Name"に"Updated UserName"と入力されている

## ユーザーは自身のワークスペースでのプロフィール情報を更新することができる
tags: stateful
* ログイン画面を開く
* メールアドレス"update_me_member_profile@example.com"で始める
* 所属中の表示名が"UpdateMe MemberProfile DisplayName"である
* 設定のプロフィール設定を開く
* 入力欄"DisplayName"に"Updated DisplayName"と入力する
* 入力欄"Bio"に"bio hogehoge"と入力する
* "Update profile"ボタンをクリック
* トーストメッセージ"Profile updated"が表示されている
* 所属中の表示名が"Updated DisplayName"である
