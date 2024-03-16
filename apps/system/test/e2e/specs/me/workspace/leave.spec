# ワークスペースから退出

## 1つの施設に所属しているメンバーはワークスペースから退出した後に認証すると、ワークスペース作成フローに入る
tags: unimplemented, テストアカウント(admin role)を作成する
* ログイン画面を開く
* メールアドレス"once_leave_workspace_has_one_workspace_e2e@example.com"で始める
* 設定のワークスペース設定を開く
* "ワークスペースから退出"ボタンをクリック
* ログイン画面が表示されている
* メールアドレス"once_leave_workspace_has_one_workspace_e2e@example.com"で始める
* オンボーディングのワークスペース作成画面が表示されている
* 入力欄"ドメイン"に"once-leave-workspace-has-one"と入力する
* "送信"ボタンをクリック
* ホーム画面が表示されている
* 選択中のワークスペース名が"OnceLeaveWorkspaceHasOne WorkspaceEndToEnd"である
* 所属中の表示名が"OnceLeaveWorkspaceHasOne DisplayNameEndToEnd"である

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

// ## 1度退出したメンバーは、退出したワークスペースから受け取った招待から再度参加することができる
