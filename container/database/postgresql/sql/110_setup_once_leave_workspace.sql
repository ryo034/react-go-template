INSERT INTO accounts (account_id, created_at)
VALUES
('018e3f69-4a17-7ff5-bbb6-7534f034b0a4', '2024-01-10 12:00:00'),
('018e3f69-4a17-717f-9956-61752f39ef5e', '2024-01-10 12:00:00'),
('018e3f69-4a17-7a84-b472-23d0a97aa6f7', '2024-01-10 12:00:00'),
('018e4062-05ec-73b6-9082-e836e445739e', '2024-01-10 12:00:00'),
('018e3f69-4a17-72e4-988d-ae46efe64359', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018e3f69-4a17-7c12-ba98-554f6d9367b9', '018e3f69-4a17-7ff5-bbb6-7534f034b0a4', 'email', '', '018e3f69-4a17-7ff5-bbb6-7534f034b0a4', 'firebase', '2024-01-10 12:00:00'),
('018e3f69-4a17-7a08-b0b1-ed4ea051cd7f', '018e3f69-4a17-717f-9956-61752f39ef5e', 'email', '', '018e3f69-4a17-717f-9956-61752f39ef5e', 'firebase', '2024-01-10 12:00:00'),
('018e3f69-4a17-7381-8251-31407b7ede9c', '018e3f69-4a17-7a84-b472-23d0a97aa6f7', 'email', '', '018e3f69-4a17-7a84-b472-23d0a97aa6f7', 'firebase', '2024-01-10 12:00:00'),
('018e4062-05ec-795f-985c-4401af6b59f3', '018e4062-05ec-73b6-9082-e836e445739e', 'email', '', '018e4062-05ec-73b6-9082-e836e445739e', 'firebase', '2024-01-10 12:00:00'),
('018e3f69-4a17-777d-bdbd-6648dcbee669', '018e3f69-4a17-72e4-988d-ae46efe64359', 'email', '', '018e3f69-4a17-72e4-988d-ae46efe64359', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e3f69-4a17-7bd4-a2e5-2c7e9a2ef8e0', '018e3f69-4a17-7ff5-bbb6-7534f034b0a4', 'once_leave_workspace_1_multiple_joined_owner@example.com', '2024-01-10 12:00:00'),
('018e3f69-4a17-7aec-9160-a11e05692bd7', '018e3f69-4a17-717f-9956-61752f39ef5e', 'once_leave_workspace_multiple_joined@example.com', '2024-01-10 12:00:00'),
('018e3f69-4a17-7ab6-a925-1e1e8a2dfcd1', '018e3f69-4a17-7a84-b472-23d0a97aa6f7', 'once_leave_workspace_invite_owner@example.com', '2024-01-10 12:00:00'),
('018e4062-05ec-76bc-b664-dad7ecf4f4b2', '018e4062-05ec-73b6-9082-e836e445739e', 'once_leave_workspace_already_left@example.com', '2024-01-10 12:00:00'),
('018e3f69-4a17-7603-ae69-89e3d0d8e9ff', '018e3f69-4a17-72e4-988d-ae46efe64359', 'once_leave_workspace_invited_after_left@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e3f69-4a17-7bd4-a2e5-2c7e9a2ef8e0', '018e3f69-4a17-7ff5-bbb6-7534f034b0a4'),
('018e3f69-4a17-7aec-9160-a11e05692bd7', '018e3f69-4a17-717f-9956-61752f39ef5e'),
('018e3f69-4a17-7ab6-a925-1e1e8a2dfcd1', '018e3f69-4a17-7a84-b472-23d0a97aa6f7'),
('018e4062-05ec-76bc-b664-dad7ecf4f4b2', '018e4062-05ec-73b6-9082-e836e445739e'),
('018e3f69-4a17-7603-ae69-89e3d0d8e9ff', '018e3f69-4a17-72e4-988d-ae46efe64359');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e3f69-4a17-78a6-a177-73fc0ad5d436', '018e3f69-4a17-7ff5-bbb6-7534f034b0a4', 'LeaveWorkspaceOne Owner', '2024-01-10 12:00:00'),
('018e3f69-4a17-7195-bb94-d0c5bdba2840', '018e3f69-4a17-717f-9956-61752f39ef5e', 'LeaveWorkspace MultipleJoined', '2024-01-10 12:00:00'),
('018e3f69-4a17-7947-a8ac-80fe91b7b57a', '018e3f69-4a17-7a84-b472-23d0a97aa6f7', 'LeaveWorkspace InviteOwner', '2024-01-10 12:00:00'),
('018e4062-05ec-7923-8955-7d94ea30113c', '018e4062-05ec-73b6-9082-e836e445739e', 'LeaveWorkspace AlreadyLeft', '2024-01-10 12:00:00'),
('018e3f69-4a17-7ac5-b014-64acb55ad593', '018e3f69-4a17-72e4-988d-ae46efe64359', 'LeaveWorkspace Invite', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e3f69-4a17-78a6-a177-73fc0ad5d436', '018e3f69-4a17-7ff5-bbb6-7534f034b0a4'),
('018e3f69-4a17-7195-bb94-d0c5bdba2840', '018e3f69-4a17-717f-9956-61752f39ef5e'),
('018e3f69-4a17-7947-a8ac-80fe91b7b57a', '018e3f69-4a17-7a84-b472-23d0a97aa6f7'),
('018e4062-05ec-7923-8955-7d94ea30113c', '018e4062-05ec-73b6-9082-e836e445739e'),
('018e3f69-4a17-7ac5-b014-64acb55ad593', '018e3f69-4a17-72e4-988d-ae46efe64359');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('018e3f69-4a17-7b45-b658-d6208e80d52a', '2024-01-10 12:00:00'),
('018e3f69-4a17-7141-bf9b-d2844560ec1b', '2024-01-10 12:00:00'),
('018e3f69-4a17-7af9-bdcb-ae05aadf429c', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_detail_id, workspace_id, name, subdomain, created_at)
VALUES
('018e3f69-4a17-7eab-b206-660b3060eeac', '018e3f69-4a17-7b45-b658-d6208e80d52a', 'Once Leave Workspace MultipleJoined 1', 'once-leave-workspace-1', '2024-01-10 12:00:00'),
('018e3f69-4a17-7a5a-a603-1936173c1c7d', '018e3f69-4a17-7141-bf9b-d2844560ec1b', 'Once Leave Workspace MultipleJoined 2', 'once-leave-workspace-2', '2024-01-10 12:00:00'),
('018e3f69-4a17-75af-974d-ba03e1b18478', '018e3f69-4a17-7af9-bdcb-ae05aadf429c', 'Once Leave Workspace Invite', 'once-leave-workspace-invite', '2024-01-10 12:00:00');

INSERT INTO workspace_latest_details (workspace_detail_id, workspace_id)
VALUES
('018e3f69-4a17-7eab-b206-660b3060eeac', '018e3f69-4a17-7b45-b658-d6208e80d52a'),
('018e3f69-4a17-7a5a-a603-1936173c1c7d', '018e3f69-4a17-7141-bf9b-d2844560ec1b'),
('018e3f69-4a17-75af-974d-ba03e1b18478', '018e3f69-4a17-7af9-bdcb-ae05aadf429c');

INSERT INTO members (member_id, account_id, workspace_id, created_at)
VALUES
('018e3f69-4a17-7800-82d7-67ef8cf57a3b', '018e3f69-4a17-7ff5-bbb6-7534f034b0a4', '018e3f69-4a17-7b45-b658-d6208e80d52a', '2024-01-10 12:00:00'),
('018e3f69-4a17-7757-9edc-321b783f173a', '018e3f69-4a17-717f-9956-61752f39ef5e', '018e3f69-4a17-7b45-b658-d6208e80d52a', '2024-01-10 12:00:00'),
('018e3f69-4a17-7d02-a98a-8d76b70a148a', '018e3f69-4a17-717f-9956-61752f39ef5e', '018e3f69-4a17-7141-bf9b-d2844560ec1b', '2024-01-10 12:00:00'),
('018e3f69-4a17-730a-87ee-b53eb2bc4886', '018e3f69-4a17-7a84-b472-23d0a97aa6f7', '018e3f69-4a17-7af9-bdcb-ae05aadf429c', '2024-01-10 12:00:00'),
('018e4062-05ec-78c7-8ebd-a6a693b5d246', '018e4062-05ec-73b6-9082-e836e445739e', '018e3f69-4a17-7af9-bdcb-ae05aadf429c', '2024-01-10 12:00:00'),
('018e3f69-4a17-78c6-8f60-153dd462291e', '018e3f69-4a17-72e4-988d-ae46efe64359', '018e3f69-4a17-7af9-bdcb-ae05aadf429c', '2024-01-10 12:00:00');

INSERT INTO member_roles (member_role_id, member_id, assigned_by, role, assigned_at)
VALUES
('018e3f69-4a17-76fe-befb-42e5035e1185', '018e3f69-4a17-7800-82d7-67ef8cf57a3b', '018e3f69-4a17-7800-82d7-67ef8cf57a3b', 'owner', '2024-01-10 12:00:00'),
('018e3f69-4a17-7538-86dc-909d3676d3f0', '018e3f69-4a17-7757-9edc-321b783f173a', '018e3f69-4a17-7757-9edc-321b783f173a', 'admin', '2024-01-10 12:00:00'),
('018e3f69-4a17-766c-9a59-84e8e2686a5e', '018e3f69-4a17-7d02-a98a-8d76b70a148a', '018e3f69-4a17-7d02-a98a-8d76b70a148a', 'owner', '2024-01-10 12:00:00'),
('018e3f69-4a17-7b1d-88e2-1169e028ee37', '018e3f69-4a17-730a-87ee-b53eb2bc4886', '018e3f69-4a17-730a-87ee-b53eb2bc4886', 'owner', '2024-01-10 12:00:00'),
('018e4062-05ec-7165-bda1-670c2f3f3030', '018e4062-05ec-78c7-8ebd-a6a693b5d246', '018e4062-05ec-78c7-8ebd-a6a693b5d246', 'admin', '2024-01-10 12:00:00'),
('018e3f69-4a17-75fa-b195-21f02b2a8968', '018e3f69-4a17-78c6-8f60-153dd462291e', '018e3f69-4a17-78c6-8f60-153dd462291e', 'admin', '2024-01-10 12:00:00');

INSERT INTO member_latest_roles (member_role_id, member_id)
VALUES
('018e3f69-4a17-76fe-befb-42e5035e1185', '018e3f69-4a17-7800-82d7-67ef8cf57a3b'),
('018e3f69-4a17-7538-86dc-909d3676d3f0', '018e3f69-4a17-7757-9edc-321b783f173a'),
('018e3f69-4a17-766c-9a59-84e8e2686a5e', '018e3f69-4a17-7d02-a98a-8d76b70a148a'),
('018e3f69-4a17-7b1d-88e2-1169e028ee37', '018e3f69-4a17-730a-87ee-b53eb2bc4886'),
('018e4062-05ec-7165-bda1-670c2f3f3030', '018e4062-05ec-78c7-8ebd-a6a693b5d246'),
('018e3f69-4a17-75fa-b195-21f02b2a8968', '018e3f69-4a17-78c6-8f60-153dd462291e');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018e3f69-4a17-7b9b-a8fb-26dcd8728e42', '018e3f69-4a17-7800-82d7-67ef8cf57a3b', '2024-01-10 12:00:00'),
('018e3f69-4a17-72af-9b64-b19b764c684a', '018e3f69-4a17-7757-9edc-321b783f173a', '2024-01-11 12:00:00'), -- 退出したワークスペースに最後にログインしている状態
('018e3f69-4a17-756a-a6ca-1fd6a7ccc2b4', '018e3f69-4a17-7d02-a98a-8d76b70a148a', '2024-01-10 12:00:00'),
('018e3f69-4a17-769c-8fbc-a696af6d9fc1', '018e3f69-4a17-730a-87ee-b53eb2bc4886', '2024-01-10 12:00:00'),
('018e4062-05ec-75ad-89ae-49564a982d79', '018e4062-05ec-78c7-8ebd-a6a693b5d246', '2024-01-10 12:00:00'),
('018e3f69-4a17-74cf-8cd0-5196ce11c254', '018e3f69-4a17-78c6-8f60-153dd462291e', '2024-01-10 12:00:00');

INSERT INTO member_latest_login_histories (member_login_history_id, member_id)
VALUES
('018e3f69-4a17-7b9b-a8fb-26dcd8728e42', '018e3f69-4a17-7800-82d7-67ef8cf57a3b'),
('018e3f69-4a17-72af-9b64-b19b764c684a', '018e3f69-4a17-7757-9edc-321b783f173a'),
('018e3f69-4a17-756a-a6ca-1fd6a7ccc2b4', '018e3f69-4a17-7d02-a98a-8d76b70a148a'),
('018e3f69-4a17-769c-8fbc-a696af6d9fc1', '018e3f69-4a17-730a-87ee-b53eb2bc4886'),
('018e4062-05ec-75ad-89ae-49564a982d79', '018e4062-05ec-78c7-8ebd-a6a693b5d246'),
('018e3f69-4a17-74cf-8cd0-5196ce11c254', '018e3f69-4a17-78c6-8f60-153dd462291e');

INSERT INTO member_profiles (member_profile_id, member_id, member_id_number, display_name, bio, created_at)
VALUES
('018e3f69-4a17-7b5b-a040-73cbbdb8c2e0', '018e3f69-4a17-7800-82d7-67ef8cf57a3b', 'DEV-54321', 'OnceLeaveWorkspaceOne MultipleJoinedOwner', 'bio', '2024-01-10 12:00:00'),
('018e3f69-4a17-7543-907e-401909345d5f', '018e3f69-4a17-7757-9edc-321b783f173a', 'DEV-54321', 'OnceLeaveWorkspace MultipleJoinedAdmin', 'bio', '2024-01-10 12:00:00'),
('018e3f69-4a17-7a3a-bd9e-eeac9e7c64da', '018e3f69-4a17-7d02-a98a-8d76b70a148a', 'DEV-54321', 'OnceLeaveWorkspace MultipleJoinedOwner', 'bio', '2024-01-10 12:00:00'),
('018e3f69-4a17-7aef-b4c8-acb9e3e4e0a7', '018e3f69-4a17-730a-87ee-b53eb2bc4886', 'DEV-54321', 'OnceLeaveWorkspace InviteOwner', 'bio', '2024-01-10 12:00:00'),
('018e4062-05ec-7310-a1f1-680481a44adc', '018e4062-05ec-78c7-8ebd-a6a693b5d246', 'DEV-54321', 'OnceLeaveWorkspace AlreadyLeft', 'bio', '2024-01-10 12:00:00'),
('018e3f69-4a17-7965-a9f5-3551282fcbfb', '018e3f69-4a17-78c6-8f60-153dd462291e', 'DEV-54321', 'OnceLeaveWorkspace Invite', 'bio', '2024-01-10 12:00:00');

INSERT INTO member_latest_profiles (member_profile_id, member_id)
VALUES
('018e3f69-4a17-7b5b-a040-73cbbdb8c2e0', '018e3f69-4a17-7800-82d7-67ef8cf57a3b'),
('018e3f69-4a17-7543-907e-401909345d5f', '018e3f69-4a17-7757-9edc-321b783f173a'),
('018e3f69-4a17-7a3a-bd9e-eeac9e7c64da', '018e3f69-4a17-7d02-a98a-8d76b70a148a'),
('018e3f69-4a17-7aef-b4c8-acb9e3e4e0a7', '018e3f69-4a17-730a-87ee-b53eb2bc4886'),
('018e4062-05ec-7310-a1f1-680481a44adc', '018e4062-05ec-78c7-8ebd-a6a693b5d246'),
('018e3f69-4a17-7965-a9f5-3551282fcbfb', '018e3f69-4a17-78c6-8f60-153dd462291e');

INSERT INTO membership_events (membership_event_id, member_id, event_type, created_by, event_at)
VALUES
('018e3f69-4a17-73c5-8a38-f7f57974b0ec', '018e3f69-4a17-7800-82d7-67ef8cf57a3b', 'join', '018e3f69-4a17-7800-82d7-67ef8cf57a3b', '2024-01-10 12:00:00'),
('018e3f69-4a17-7d87-8018-86d2b0360df6', '018e3f69-4a17-7757-9edc-321b783f173a', 'join', '018e3f69-4a17-7757-9edc-321b783f173a', '2024-01-10 12:00:00'),
('018e3f69-4a17-7cd5-838a-4b423c3fb3a0', '018e3f69-4a17-7d02-a98a-8d76b70a148a', 'join', '018e3f69-4a17-7d02-a98a-8d76b70a148a', '2024-01-10 12:00:00'),
('018e3f69-4a17-7bf6-9bba-decaf3bd7f65', '018e3f69-4a17-730a-87ee-b53eb2bc4886', 'join', '018e3f69-4a17-730a-87ee-b53eb2bc4886', '2024-01-10 12:00:00'),
('018e4062-05ec-7bba-aa58-6bb8a19d52c1', '018e4062-05ec-78c7-8ebd-a6a693b5d246', 'join', '018e4062-05ec-78c7-8ebd-a6a693b5d246', '2024-01-10 12:00:00'),
('018e4062-05ec-7a50-a3b8-b8baa0fd353b', '018e4062-05ec-78c7-8ebd-a6a693b5d246', 'leave', '018e4062-05ec-78c7-8ebd-a6a693b5d246', '2024-01-12 12:00:00'),
('018e3f69-4a17-7250-a25a-516606613000', '018e3f69-4a17-78c6-8f60-153dd462291e', 'join', '018e3f69-4a17-78c6-8f60-153dd462291e', '2024-01-10 12:00:00'),
('018e3f69-4a17-7c90-993b-b01f93ebbaf9', '018e3f69-4a17-78c6-8f60-153dd462291e', 'leave', '018e3f69-4a17-78c6-8f60-153dd462291e', '2024-01-12 12:00:00');

INSERT INTO latest_membership_events (membership_event_id, member_id)
VALUES
('018e3f69-4a17-73c5-8a38-f7f57974b0ec', '018e3f69-4a17-7800-82d7-67ef8cf57a3b'),
('018e3f69-4a17-7cd5-838a-4b423c3fb3a0', '018e3f69-4a17-7d02-a98a-8d76b70a148a'),
('018e3f69-4a17-7bf6-9bba-decaf3bd7f65', '018e3f69-4a17-730a-87ee-b53eb2bc4886'),
('018e4062-05ec-7a50-a3b8-b8baa0fd353b', '018e4062-05ec-78c7-8ebd-a6a693b5d246'),
('018e3f69-4a17-7c90-993b-b01f93ebbaf9', '018e3f69-4a17-78c6-8f60-153dd462291e'),
('018e3f69-4a17-7d87-8018-86d2b0360df6', '018e3f69-4a17-7757-9edc-321b783f173a');

INSERT INTO invitation_units (invitation_unit_id, workspace_id, invited_by, created_at)
VALUES
('018e4062-05ec-76a8-a7c5-c4d975bf2778', '018e3f69-4a17-7af9-bdcb-ae05aadf429c', '018e4062-05ec-78c7-8ebd-a6a693b5d246', '2024-01-13 12:00:00'),
('018e3f69-4a17-7808-9b14-c0286a7a8156', '018e3f69-4a17-7af9-bdcb-ae05aadf429c', '018e3f69-4a17-730a-87ee-b53eb2bc4886', '2024-01-13 12:00:00');

INSERT INTO invitations (invitation_id, invitation_unit_id)
VALUES
('018e4421-aba7-7a6a-8a8f-9c2ba26e69dd', '018e4062-05ec-76a8-a7c5-c4d975bf2778'),
('018e4062-05ec-7a6f-8ad9-812f3f2c33f9', '018e4062-05ec-76a8-a7c5-c4d975bf2778'),
('018e3f69-4a17-72f3-88a2-d05923a89ce4', '018e3f69-4a17-7808-9b14-c0286a7a8156');

INSERT INTO invitation_tokens (invitation_id, token, expired_at, created_at)
VALUES
('018e4421-aba7-7a6a-8a8f-9c2ba26e69dd', '018e4421-aba7-7de4-9bc7-ca0f93355a28', '2500-01-20 12:00:00', '2023-01-13 12:00:00'),
('018e4062-05ec-7a6f-8ad9-812f3f2c33f9', '018e4062-05ec-7af4-a2fa-a48e4294deeb', '2500-01-20 12:00:00', '2023-01-13 12:00:00'),
('018e3f69-4a17-72f3-88a2-d05923a89ce4', '018e3f69-4a17-7f13-a74d-0fc58ba18a32', '2500-01-20 12:00:00', '2023-01-13 12:00:00');

INSERT INTO invitees (invitation_id, email)
VALUES
('018e4421-aba7-7a6a-8a8f-9c2ba26e69dd', 'once_leave_workspace_accept_receive_from_already_left_member@example.com'),
('018e4062-05ec-7a6f-8ad9-812f3f2c33f9', 'once_leave_workspace_check_receive_from_already_left_member@example.com'),
('018e3f69-4a17-72f3-88a2-d05923a89ce4', 'once_leave_workspace_invited_after_left@example.com');
