INSERT INTO accounts (account_id, created_at)
VALUES
('018e0e5c-98a5-76de-9ede-13118ba8c996', '2024-01-10 12:00:00'),
('018e1398-3d80-76ce-9623-9a6caae8378e', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018e0e5d-e8cd-7b3b-8d0b-5b4daac55cdc', '018e0e5c-98a5-76de-9ede-13118ba8c996', 'email', '', '018e0e5c-98a5-76de-9ede-13118ba8c996', 'firebase', '2024-01-10 12:00:00'),
('018e1398-3d80-796a-acee-44d10b7644ec', '018e1398-3d80-76ce-9623-9a6caae8378e', 'google', 'https://github.com/ryo034/image/assets/55078625/967e0e8c-a2be-4004-834a-d56a263b89ce', 'UPeY3R7sVON9d9i8mq7KWLkyNXXZ', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e0e5e-8486-7286-9114-f7479e18f94f', '018e0e5c-98a5-76de-9ede-13118ba8c996', 'update_me_update_profile_photo@example.com', '2024-01-10 12:00:00'),
('018e1398-3d80-7fab-8333-89b92a91eca6', '018e1398-3d80-76ce-9623-9a6caae8378e', 'test_has_photo_google_setup_photo@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e0e5e-8486-7286-9114-f7479e18f94f', '018e0e5c-98a5-76de-9ede-13118ba8c996'),
('018e1398-3d80-7fab-8333-89b92a91eca6', '018e1398-3d80-76ce-9623-9a6caae8378e');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e0e5f-16d4-7fa5-8233-022f53b54b55', '018e0e5c-98a5-76de-9ede-13118ba8c996', 'UpdateMe UpdateProfilePhoto', '2024-01-10 12:00:00'),
('018e1398-3d80-79fb-af13-e2868037907c', '018e1398-3d80-76ce-9623-9a6caae8378e', 'HasPhotoGoogle SetupPhoto', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e0e5f-16d4-7fa5-8233-022f53b54b55', '018e0e5c-98a5-76de-9ede-13118ba8c996'),
('018e1398-3d80-79fb-af13-e2868037907c', '018e1398-3d80-76ce-9623-9a6caae8378e');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('018e0e5f-9c4b-7062-aaa7-3db3fde10354', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_detail_id, workspace_id, name, subdomain, created_at)
VALUES
('018e200b-9d01-74b6-baa6-0b9e39000bfe', '018e0e5f-9c4b-7062-aaa7-3db3fde10354', 'UpdateAccountPhoto Workspace', 'update-me-account-profile', '2024-01-10 12:00:00');

INSERT INTO workspace_latest_details (workspace_detail_id, workspace_id)
VALUES
('018e200b-9d01-74b6-baa6-0b9e39000bfe', '018e0e5f-9c4b-7062-aaa7-3db3fde10354');

INSERT INTO members (member_id, account_id, workspace_id, created_at)
VALUES
('018e0e60-4f73-72bd-84f7-ae1abaa7fef1', '018e0e5c-98a5-76de-9ede-13118ba8c996', '018e0e5f-9c4b-7062-aaa7-3db3fde10354', '2024-01-10 12:00:00'),
('018e1398-3d80-7fae-a661-3c234a9c5c53', '018e1398-3d80-76ce-9623-9a6caae8378e', '018e0e5f-9c4b-7062-aaa7-3db3fde10354', '2024-01-10 12:00:00');

INSERT INTO member_roles (member_role_id, member_id, assigned_by, role, assigned_at)
VALUES
('018e0e60-ca5a-7c43-9bc3-67f7d65c4512', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1', 'owner', '2024-01-10 12:00:00'),
('018e1398-3d80-743b-a74f-3497770daaf9', '018e1398-3d80-7fae-a661-3c234a9c5c53', '018e1398-3d80-7fae-a661-3c234a9c5c53', 'admin', '2024-01-10 12:00:00');

INSERT INTO member_latest_roles (member_role_id, member_id)
VALUES
('018e0e60-ca5a-7c43-9bc3-67f7d65c4512', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1'),
('018e1398-3d80-743b-a74f-3497770daaf9', '018e1398-3d80-7fae-a661-3c234a9c5c53');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018e0e61-3bf7-726a-b12d-8644926a08fd', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1', '2024-01-10 12:00:00'),
('018e1398-3d80-7103-88f9-20c3f226f13b', '018e1398-3d80-7fae-a661-3c234a9c5c53', '2024-01-10 12:00:00');

INSERT INTO member_latest_login_histories (member_login_history_id, member_id)
VALUES
('018e0e61-3bf7-726a-b12d-8644926a08fd', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1'),
('018e1398-3d80-7103-88f9-20c3f226f13b', '018e1398-3d80-7fae-a661-3c234a9c5c53');

INSERT INTO member_profiles (member_profile_id, member_id, member_id_number, display_name, bio, created_at)
VALUES
('018e2216-64a3-7536-8c7c-327a4585300f', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1', 'DEV-54321', 'UpdateMe UpdateProfilePhoto', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-7b23-bdf9-ce4aac4fc1b2', '018e1398-3d80-7fae-a661-3c234a9c5c53', 'DEV-54321', 'HasPhotoGoogle SetupPhoto', 'bio', '2024-01-10 12:00:00');

INSERT INTO member_latest_profiles (member_profile_id, member_id)
VALUES
('018e2216-64a3-7536-8c7c-327a4585300f', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1'),
('018e2216-64a3-7b23-bdf9-ce4aac4fc1b2', '018e1398-3d80-7fae-a661-3c234a9c5c53');

INSERT INTO membership_events (membership_event_id, member_id, event_type, created_by, event_at)
VALUES
('018e2ff9-c432-70d3-8b54-8874679bbb28', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1', 'join', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1', '2024-01-10 12:00:00'),
('018e2ff9-c432-71d3-9fbe-8e9d689f55cc', '018e1398-3d80-7fae-a661-3c234a9c5c53', 'join', '018e1398-3d80-7fae-a661-3c234a9c5c53', '2024-01-10 12:00:00');

INSERT INTO latest_membership_events (membership_event_id, member_id)
VALUES
('018e2ff9-c432-70d3-8b54-8874679bbb28', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1'),
('018e2ff9-c432-71d3-9fbe-8e9d689f55cc', '018e1398-3d80-7fae-a661-3c234a9c5c53');
