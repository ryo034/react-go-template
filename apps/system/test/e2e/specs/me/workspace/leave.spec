# ワークスペースから退出

## 1つの施設に所属しているメンバーはワークスペースから退出した後に認証すると、ワークスペース作成フローに入る
* ログイン画面を開く
* メールアドレス"once_leave_has_one_workspace_admin@example.com"で始める
* ホーム画面が表示されている
* 設定のワークスペース設定を開く
* "ワークスペースから退出"ボタンをクリック
* ダイアログ上の"退出"ボタンをクリック
* ログイン画面が表示されている
* メールアドレス"once_leave_has_one_workspace_admin@example.com"で始める
* オンボーディングのワークスペース作成画面が表示されている
* 入力欄"ドメイン"に"once-leave-workspace-has-one-create"と入力する
* "送信"ボタンをクリック
* ホーム画面が表示されている
* 選択中のワークスペース名が"once-leave-workspace-has-one-create"である
* 所属中の表示名が"LeaveWorkspace HasOneAdmin"である

## 複数施設に所属しているメンバーは一方のワークスペースから退出した後に認証すると、もう一方のワークスペースにログインされる
* ログイン画面を開く
* メールアドレス"once_leave_workspace_multiple_joined@example.com"で始める
* ホーム画面が表示されている
* 選択中のワークスペース名が"Once Leave Workspace MultipleJoined 1"である
* 所属中の表示名が"OnceLeaveWorkspace MultipleJoinedAdmin"である
* 設定のワークスペース設定を開く
* "ワークスペースから退出"ボタンをクリック
* ダイアログ上の"退出"ボタンをクリック
* ログイン画面が表示されている
* メールアドレス"once_leave_workspace_multiple_joined@example.com"で始める
* ホーム画面が表示されている
* 選択中のワークスペース名が"Once Leave Workspace MultipleJoined 2"である
* 所属中の表示名が"OnceLeaveWorkspace MultipleJoinedOwner"である

## 1度退出したメンバーは、退出したワークスペースから受け取った招待から再度参加することができる
* トークン"018e4421-aba7-7de4-9bc7-ca0f93355a28"の招待画面を開く
* "メールアドレスで始める"ボタンをクリック
* メールアドレス"once_leave_workspace_accept_receive_from_already_left_member@example.com"に送信されたワンタイムパスワードを取得
* ワンタイムパスワード確認画面にワンタイムパスワードを入力する
* オンボーディングの名前入力画面が表示されている
* 入力欄"名前"に"test name"と入力する
* "送信"ボタンをクリック
* 招待受諾画面でワークスペース"Once Leave Workspace Invite"招待者"Removed User"の参加ボタンをクリック
* 選択中のワークスペース名が"Once Leave Workspace Invite"である
