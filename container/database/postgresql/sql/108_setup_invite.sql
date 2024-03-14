INSERT INTO accounts (account_id, created_at)
VALUES
('018d96bf-8dce-7f68-a926-b5d7ed6ed883', '2024-01-10 12:00:00'),
('018d9b4d-9438-79ac-b533-1323d4ec9b9f', '2024-01-10 12:00:00'),
('018da09e-c6ca-795e-878d-32bb8c1e5cac', '2024-01-10 12:00:00'),
('018df53c-c5a6-71a2-bf90-2f751f342d4c', '2024-01-10 12:00:00'),
('018df551-4339-730c-8031-618eb8ef66b5', '2024-01-10 12:00:00'),
('018e0ea9-6b88-71d3-a887-0cf22ede3e0c', '2024-01-10 12:00:00'),
('018e0ebc-e842-7bdb-bf50-05177e07a1c7', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018de2f6-ca47-7fc9-832a-3d725120c55b', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', 'email', '', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', 'firebase', '2024-01-10 12:00:00'),
('018de2f6-de72-7b8c-92ab-b72b90d41ccd', '018d9b4d-9438-79ac-b533-1323d4ec9b9f', 'email', '', '018d9b4d-9438-79ac-b533-1323d4ec9b9f', 'firebase', '2024-01-10 12:00:00'),
('018de2f6-f23e-7a3c-ab51-3117f07c1930', '018da09e-c6ca-795e-878d-32bb8c1e5cac', 'email', '', '018da09e-c6ca-795e-878d-32bb8c1e5cac', 'firebase', '2024-01-10 12:00:00'),
('018df53c-f868-7f2a-bafd-9cda1fe15e8a', '018df53c-c5a6-71a2-bf90-2f751f342d4c', 'email', '', '018df53c-c5a6-71a2-bf90-2f751f342d4c', 'firebase', '2024-01-10 12:00:00'),
('018df551-d07a-7761-8c69-7de98d195e26', '018df551-4339-730c-8031-618eb8ef66b5', 'google', '', 'ZHjoHCDE0C1EHxLIQvNgiygTXu9A', 'firebase', '2024-01-10 12:00:00'),
('018e0ea9-aaab-7079-9046-55cfac836d3f', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c', 'google', 'https://github.com/ryo034/image/assets/55078625/af9fae15-baf3-451e-820a-99f7e246af31', 'f8fVGXfC3dmym8XHbFyCs1LvwJ7O', 'firebase', '2024-01-10 12:00:00'),
('018e0ebd-7922-7c48-b598-5562cc7fa29c', '018e0ebc-e842-7bdb-bf50-05177e07a1c7', 'google', 'https://github.com/ryo034/image/assets/55078625/ddeb3605-2291-4c19-81ec-6d890c7d0219', 'QA1ViUeGbDWJQfydWYJbRpzXNVEk', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e09c2-d3f3-71bd-bbca-5cec4e063d46', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', 'invite_test_1@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-74e3-80a7-e33c4d0ddc9c', '018d9b4d-9438-79ac-b533-1323d4ec9b9f', 'invite_test_already_joined_any_workspace@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-79eb-a643-9c160d5b998d', '018da09e-c6ca-795e-878d-32bb8c1e5cac', 'invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-7de8-8531-51a58e1e3a96', '018df53c-c5a6-71a2-bf90-2f751f342d4c', 'invite_test_already_joined_any_workspace_by_email@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-787a-94e6-adfc83e2c457', '018df551-4339-730c-8031-618eb8ef66b5', 'invite_test_already_joined_any_workspace_by_google@example.com', '2024-01-10 12:00:00'),
('018e0eaa-8b25-7a21-a221-f85f8a85d149', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c', 'invite_test_has_photo_by_google_accept_with_email@example.com', '2024-01-10 12:00:00'),
('018e0ebe-185a-75fb-8096-b6ad5b02a9fb', '018e0ebc-e842-7bdb-bf50-05177e07a1c7', 'invite_test_has_photo_by_google_accept_with_has_photo_google@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e09c2-d3f3-71bd-bbca-5cec4e063d46', '018d96bf-8dce-7f68-a926-b5d7ed6ed883'),
('018e09c2-d3f3-74e3-80a7-e33c4d0ddc9c', '018d9b4d-9438-79ac-b533-1323d4ec9b9f'),
('018e09c2-d3f3-79eb-a643-9c160d5b998d', '018da09e-c6ca-795e-878d-32bb8c1e5cac'),
('018e09c2-d3f3-7de8-8531-51a58e1e3a96', '018df53c-c5a6-71a2-bf90-2f751f342d4c'),
('018e09c2-d3f3-787a-94e6-adfc83e2c457', '018df551-4339-730c-8031-618eb8ef66b5'),
('018e0eaa-8b25-7a21-a221-f85f8a85d149', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c'),
('018e0ebe-185a-75fb-8096-b6ad5b02a9fb', '018e0ebc-e842-7bdb-bf50-05177e07a1c7');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e088f-c783-7e2c-9d8c-371168132855', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', 'Invite TestOne', '2024-01-10 12:00:00'),
('018e088f-d9f3-74ae-9c94-25c0d99ff178', '018d9b4d-9438-79ac-b533-1323d4ec9b9f', 'Invite TestTwo', '2024-01-10 12:00:00'),
('018e088f-eada-78ba-b8ea-8d27d57db944', '018da09e-c6ca-795e-878d-32bb8c1e5cac', 'Invite TestThree', '2024-01-10 12:00:00'),
('018e0890-34c0-7ed8-80d6-6626f4d37531', '018df53c-c5a6-71a2-bf90-2f751f342d4c', 'InviteGoogleAuthTest AlreadyJoined', '2024-01-10 12:00:00'),
('018e0890-4745-7655-af59-805da2375591', '018df551-4339-730c-8031-618eb8ef66b5', 'InviteGoogleAuthTest AlreadyJoinedGoogle', '2024-01-10 12:00:00'),
('018e0eab-10a4-7872-88f1-9d74c2bd04a9', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c', 'InviteTest HasPhotoAcceptWithEmail', '2024-01-10 12:00:00'),
('018e0ebe-8ca7-7ecd-92b0-45dd064dfb14', '018e0ebc-e842-7bdb-bf50-05177e07a1c7', 'InviteTest HasPhotoAcceptWithGoogle', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e088f-c783-7e2c-9d8c-371168132855', '018d96bf-8dce-7f68-a926-b5d7ed6ed883'),
('018e088f-d9f3-74ae-9c94-25c0d99ff178', '018d9b4d-9438-79ac-b533-1323d4ec9b9f'),
('018e088f-eada-78ba-b8ea-8d27d57db944', '018da09e-c6ca-795e-878d-32bb8c1e5cac'),
('018e0890-34c0-7ed8-80d6-6626f4d37531', '018df53c-c5a6-71a2-bf90-2f751f342d4c'),
('018e0890-4745-7655-af59-805da2375591', '018df551-4339-730c-8031-618eb8ef66b5'),
('018e0eab-10a4-7872-88f1-9d74c2bd04a9', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c'),
('018e0ebe-8ca7-7ecd-92b0-45dd064dfb14', '018e0ebc-e842-7bdb-bf50-05177e07a1c7');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00'),
('018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_detail_id, workspace_id, name, subdomain, created_at)
VALUES
('018e200b-9d01-7efb-b0e7-51fb5333300e', '018d96b9-c920-7434-b5c3-02e5e920ae9d', 'InviteTest 1', 'invite-test-1', '2024-01-10 12:00:00'),
('018e200b-9d01-7f75-a427-59b82d6acc97', '018d9b4d-e340-74f7-914c-2476eff949bb', 'InviteTest 2', 'invite-test-2', '2024-01-10 12:00:00');

INSERT INTO workspace_latest_details (workspace_detail_id, workspace_id)
VALUES
('018e200b-9d01-7efb-b0e7-51fb5333300e', '018d96b9-c920-7434-b5c3-02e5e920ae9d'),
('018e200b-9d01-7f75-a427-59b82d6acc97', '018d9b4d-e340-74f7-914c-2476eff949bb');

INSERT INTO members (member_id, account_id, workspace_id, created_at)
VALUES
('018d96b9-f674-7ff6-83eb-506eca6452be', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', '018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00'),
('018df53e-4c77-79de-b725-c43ebcb79450', '018df53c-c5a6-71a2-bf90-2f751f342d4c', '018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00'),
('018df552-5086-7b84-8601-d04c319d2e44', '018df551-4339-730c-8031-618eb8ef66b5', '018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00'),
('018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '018d9b4d-9438-79ac-b533-1323d4ec9b9f', '018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00'),
('018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', '018da09e-c6ca-795e-878d-32bb8c1e5cac', '018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00'),
('018e0eab-df36-7295-a917-9c81e8c6671c', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c', '018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00'),
('018e0ebf-0806-7f10-a73b-28db9f4d2349', '018e0ebc-e842-7bdb-bf50-05177e07a1c7', '018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00');

INSERT INTO member_roles (member_role_id, member_id, assigned_by, role, assigned_at)
VALUES
('018df76b-56e0-7371-a69b-1dea397a75d8', '018d96b9-f674-7ff6-83eb-506eca6452be', '018d96b9-f674-7ff6-83eb-506eca6452be', 'owner', '2024-01-10 12:00:00'),
('018df76b-6bf7-788f-bc7b-e1102924573d', '018df53e-4c77-79de-b725-c43ebcb79450', '018df53e-4c77-79de-b725-c43ebcb79450', 'admin', '2024-01-10 12:00:00'),
('018df76b-81da-7cb2-a5fe-b849b52a939e', '018df552-5086-7b84-8601-d04c319d2e44', '018df552-5086-7b84-8601-d04c319d2e44', 'admin', '2024-01-10 12:00:00'),
('018df76b-9717-788a-9b02-548a9666ac44', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', 'owner', '2024-01-10 12:00:00'),
('018df76b-ace2-7420-87d5-666e55aa18b7', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', 'admin', '2024-01-10 12:00:00'),
('018e0eac-2417-70c4-a18d-a6acc026520b', '018e0eab-df36-7295-a917-9c81e8c6671c', '018e0eab-df36-7295-a917-9c81e8c6671c', 'admin', '2024-01-10 12:00:00'),
('018e0ebf-738d-798a-b861-739f8d4810cf', '018e0ebf-0806-7f10-a73b-28db9f4d2349', '018e0ebf-0806-7f10-a73b-28db9f4d2349', 'admin', '2024-01-10 12:00:00');

INSERT INTO member_latest_roles (member_role_id, member_id)
VALUES
('018df76b-56e0-7371-a69b-1dea397a75d8', '018d96b9-f674-7ff6-83eb-506eca6452be'),
('018df76b-6bf7-788f-bc7b-e1102924573d', '018df53e-4c77-79de-b725-c43ebcb79450'),
('018df76b-81da-7cb2-a5fe-b849b52a939e', '018df552-5086-7b84-8601-d04c319d2e44'),
('018df76b-9717-788a-9b02-548a9666ac44', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d'),
('018df76b-ace2-7420-87d5-666e55aa18b7', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5'),
('018e0eac-2417-70c4-a18d-a6acc026520b', '018e0eab-df36-7295-a917-9c81e8c6671c'),
('018e0ebf-738d-798a-b861-739f8d4810cf', '018e0ebf-0806-7f10-a73b-28db9f4d2349');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018d96ba-6ebd-77a6-9534-958f8fe487ce', '018d96b9-f674-7ff6-83eb-506eca6452be', '2024-01-10 12:00:00'),
('018df53e-965e-7e7e-8842-fd0e4135caf0', '018df53e-4c77-79de-b725-c43ebcb79450', '2024-01-10 12:00:00'),
('018df552-75b8-76c7-afc0-bb51404f9359', '018df552-5086-7b84-8601-d04c319d2e44', '2024-01-10 12:00:00'),
('018d9b4e-0b6e-7f6e-8b7e-9f6e8d7e6f8e', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '2024-01-10 12:00:00'),
('018da0dc-7577-7e53-8db0-ac3d68801240', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', '2024-01-10 12:00:00'),
('018e0eac-6b49-7343-88cc-b99b00e59fbf', '018e0eab-df36-7295-a917-9c81e8c6671c', '2024-01-10 12:00:00'),
('018e0ebf-c847-7214-a31a-12f00ef1a627', '018e0ebf-0806-7f10-a73b-28db9f4d2349', '2024-01-10 12:00:00');

INSERT INTO member_latest_login_histories (member_login_history_id, member_id)
VALUES
('018d96ba-6ebd-77a6-9534-958f8fe487ce', '018d96b9-f674-7ff6-83eb-506eca6452be'),
('018df53e-965e-7e7e-8842-fd0e4135caf0', '018df53e-4c77-79de-b725-c43ebcb79450'),
('018df552-75b8-76c7-afc0-bb51404f9359', '018df552-5086-7b84-8601-d04c319d2e44'),
('018d9b4e-0b6e-7f6e-8b7e-9f6e8d7e6f8e', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d'),
('018da0dc-7577-7e53-8db0-ac3d68801240', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5'),
('018e0eac-6b49-7343-88cc-b99b00e59fbf', '018e0eab-df36-7295-a917-9c81e8c6671c'),
('018e0ebf-c847-7214-a31a-12f00ef1a627', '018e0ebf-0806-7f10-a73b-28db9f4d2349');

INSERT INTO member_profiles (member_profile_id, member_id, member_id_number, display_name, bio, created_at)
VALUES
('018e2216-64a3-7bff-b630-555c8cd00f5b', '018d96b9-f674-7ff6-83eb-506eca6452be', 'DEV-54321', 'Invite TestOne', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-7956-bae1-1d29495831b3', '018df53e-4c77-79de-b725-c43ebcb79450', 'DEV-54322', 'InviteGoogleAuthTest AlreadyJoined', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-746b-b5f4-847aef64aede', '018df552-5086-7b84-8601-d04c319d2e44', 'DEV-54323', 'InviteGoogleAuthTest AlreadyJoinedGoogle', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-761f-951d-b57715af0d47', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', 'DEV-09876', 'Invite TestTwo', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-7bd4-bf8a-b87614b92150', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', 'DEV-54321', 'Invite TestThree', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-7739-80a1-bf8048c536ea', '018e0eab-df36-7295-a917-9c81e8c6671c', 'DEV-54321', 'InviteTest HasPhotoAcceptWithEmail', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-75da-982f-52bd79c3bbba', '018e0ebf-0806-7f10-a73b-28db9f4d2349', 'DEV-54322', 'InviteTest HasPhotoAcceptWithGoogle', 'bio', '2024-01-10 12:00:00');

INSERT INTO member_latest_profiles (member_profile_id, member_id)
VALUES
('018e2216-64a3-7bff-b630-555c8cd00f5b', '018d96b9-f674-7ff6-83eb-506eca6452be'),
('018e2216-64a3-7956-bae1-1d29495831b3', '018df53e-4c77-79de-b725-c43ebcb79450'),
('018e2216-64a3-746b-b5f4-847aef64aede', '018df552-5086-7b84-8601-d04c319d2e44'),
('018e2216-64a3-761f-951d-b57715af0d47', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d'),
('018e2216-64a3-7bd4-bf8a-b87614b92150', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5'),
('018e2216-64a3-7739-80a1-bf8048c536ea', '018e0eab-df36-7295-a917-9c81e8c6671c'),
('018e2216-64a3-75da-982f-52bd79c3bbba', '018e0ebf-0806-7f10-a73b-28db9f4d2349');

INSERT INTO membership_events (membership_event_id, member_id, event_type, created_by, event_at)
VALUES
('018e2ff9-c432-7019-8282-335c45599d2d', '018d96b9-f674-7ff6-83eb-506eca6452be', 'join', '018d96b9-f674-7ff6-83eb-506eca6452be', '2024-01-10 12:00:00'),
('018e2ff9-c432-7460-ae9d-0de49f00c6b9', '018df53e-4c77-79de-b725-c43ebcb79450', 'join', '018df53e-4c77-79de-b725-c43ebcb79450', '2024-01-10 12:00:00'),
('018e2ff9-c432-761e-811e-7c1669de7648', '018df552-5086-7b84-8601-d04c319d2e44', 'join', '018df552-5086-7b84-8601-d04c319d2e44', '2024-01-10 12:00:00'),
('018e2ff9-c432-7e18-b396-7d08695aae68', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', 'join', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '2024-01-10 12:00:00'),
('018e2ff9-c432-721c-b089-5c99efc40348', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', 'join', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', '2024-01-10 12:00:00'),
('018e2ff9-c432-7c2f-a69e-f4530fe486a9', '018e0eab-df36-7295-a917-9c81e8c6671c', 'join', '018e0eab-df36-7295-a917-9c81e8c6671c', '2024-01-10 12:00:00'),
('018e2ff9-c432-74fd-91f2-d4d0ce809b3a', '018e0ebf-0806-7f10-a73b-28db9f4d2349', 'join', '018e0ebf-0806-7f10-a73b-28db9f4d2349', '2024-01-10 12:00:00');

INSERT INTO latest_membership_events (membership_event_id, member_id)
VALUES
('018e2ff9-c432-7019-8282-335c45599d2d', '018d96b9-f674-7ff6-83eb-506eca6452be'),
('018e2ff9-c432-7460-ae9d-0de49f00c6b9', '018df53e-4c77-79de-b725-c43ebcb79450'),
('018e2ff9-c432-761e-811e-7c1669de7648', '018df552-5086-7b84-8601-d04c319d2e44'),
('018e2ff9-c432-7e18-b396-7d08695aae68', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d'),
('018e2ff9-c432-721c-b089-5c99efc40348', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5'),
('018e2ff9-c432-7c2f-a69e-f4530fe486a9', '018e0eab-df36-7295-a917-9c81e8c6671c'),
('018e2ff9-c432-74fd-91f2-d4d0ce809b3a', '018e0ebf-0806-7f10-a73b-28db9f4d2349');

INSERT INTO invitation_units (invitation_unit_id, workspace_id, invited_by, created_at)
VALUES
('018db49e-d5a5-7ed1-aea6-e018d2e4bd38', '018d96b9-c920-7434-b5c3-02e5e920ae9d', '018d96b9-f674-7ff6-83eb-506eca6452be', '2024-01-10 12:00:00'),
('018e0ead-064c-7eea-b664-c414d44270f0', '018d9b4d-e340-74f7-914c-2476eff949bb', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '2024-01-10 12:00:00');

INSERT INTO invitations (invitation_id, invitation_unit_id)
VALUES
('018d96b8-0bb6-7822-b3b5-78a5d0e8790e', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018d9fb4-641f-72f9-bdeb-a493a974dba1', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018d9b6d-01cf-7d3e-8328-76736b6db971', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018da09b-ed0c-7688-a8e3-4573104e006f', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018df2f6-152e-7c71-84c3-a2b77306724f', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018df2f9-c876-700d-8bb4-8a3613bfbc71', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018df2f7-31cf-7079-899e-9e6dab375c38', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018df2f9-8eb2-78a2-accd-e44bb55dabef', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018df53b-4310-701e-b62d-bea4b4c7d667', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018df54f-9c56-7209-9082-380f1096c46e', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018e0eae-44a6-732d-ab12-54e5a1f3f27d', '018e0ead-064c-7eea-b664-c414d44270f0'),
('018e0ec0-bc20-7a22-9ffd-db6e0dccec25', '018e0ead-064c-7eea-b664-c414d44270f0'),
('018e0ec1-cb17-777b-8e4a-3149194401b4', '018e0ead-064c-7eea-b664-c414d44270f0');

INSERT INTO invitation_tokens (invitation_id, token, expired_at, created_at)
VALUES
('018d96b8-0bb6-7822-b3b5-78a5d0e8790e', '018d96b7-587c-7614-b234-e086b1944e74', '2023-01-10 12:00:00', '2023-01-09 12:00:00'),
('018d9fb4-641f-72f9-bdeb-a493a974dba1', '018d9fb5-6150-7a4b-a5c8-b5a61e51ee50', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d9b6d-01cf-7d3e-8328-76736b6db971', '018d9fb5-7e56-75ed-952f-ae8aa4fed8c6', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018da09b-ed0c-7688-a8e3-4573104e006f', '018da09e-2fa7-7d3a-ad23-2c9f5cb76b92', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018df2f6-152e-7c71-84c3-a2b77306724f', '018df2f6-50a9-7c19-94a2-575f32dd1e41', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018df2f9-c876-700d-8bb4-8a3613bfbc71', '018df2fa-2dc2-79ea-8913-e45e39379c9c', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018df2f7-31cf-7079-899e-9e6dab375c38', '018df2fa-4598-7e13-af4d-7727a9bca288', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018df2f9-8eb2-78a2-accd-e44bb55dabef', '018df2fa-5b26-78d4-ad65-7ca831064e50', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018df53b-4310-701e-b62d-bea4b4c7d667', '018df53b-82a2-7324-9b26-f17496bfcdf8', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018df54f-9c56-7209-9082-380f1096c46e', '018df54f-e057-7818-8c72-80d6393e39e6', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018e0eae-44a6-732d-ab12-54e5a1f3f27d', '018e0eae-aea6-74e5-8bd6-288b480b335a', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018e0ec0-bc20-7a22-9ffd-db6e0dccec25', '018e0ec0-e54e-7476-b65c-220bfafbf631', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018e0ec1-cb17-777b-8e4a-3149194401b4', '018e0ec2-0d64-7b7d-92af-e42be382216c', '2200-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO invitees (invitation_id, email)
VALUES
('018d96b8-0bb6-7822-b3b5-78a5d0e8790e', 'invite_test_expired@example.com'),
('018d9fb4-641f-72f9-bdeb-a493a974dba1', 'invite_test_not_expired_with_display_name@example.com'),
('018d9b6d-01cf-7d3e-8328-76736b6db971', 'invite_test_already_joined_any_workspace@example.com'),
('018da09b-ed0c-7688-a8e3-4573104e006f', 'invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com'),
('018df2f6-152e-7c71-84c3-a2b77306724f', 'invite_test_has_name_google_auth_with_display_name_when_invite@example.com'),
('018df2f9-c876-700d-8bb4-8a3613bfbc71', 'invite_test_has_name_google_auth_no_name_when_invite@example.com'),
('018df2f7-31cf-7079-899e-9e6dab375c38', 'invite_test_no_name_google_auth_with_display_name_when_invite@example.com'),
('018df2f9-8eb2-78a2-accd-e44bb55dabef', 'invite_test_no_name_google_auth_no_name_when_invite@example.com'),
('018df53b-4310-701e-b62d-bea4b4c7d667', 'invite_test_already_joined_any_workspace_by_email@example.com'),
('018df54f-9c56-7209-9082-380f1096c46e', 'invite_test_already_joined_any_workspace_by_google@example.com'),
('018e0eae-44a6-732d-ab12-54e5a1f3f27d', 'invite_test_has_photo_by_google_accept_with_email@example.com'),
('018e0ec0-bc20-7a22-9ffd-db6e0dccec25', 'invite_test_has_photo_by_google_accept_with_has_photo_google@example.com'),
('018e0ec1-cb17-777b-8e4a-3149194401b4', 'invite_test_no_account_accept_with_has_photo_google@example.com');

INSERT INTO invitee_names (invitation_id, display_name)
VALUES
('018d9fb4-641f-72f9-bdeb-a493a974dba1', 'Invite Test'),
('018da09b-ed0c-7688-a8e3-4573104e006f', 'Invite TestThreeChanged'),
('018df2f6-152e-7c71-84c3-a2b77306724f', 'InviteGoogleAuthTest HasNameTest'),
('018df2f9-c876-700d-8bb4-8a3613bfbc71', 'InviteGoogleAuthTest NoNameTest');
