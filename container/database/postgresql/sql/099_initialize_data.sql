INSERT INTO address_components (component_id, component_type, component_name)
VALUES
('10269b87-98ce-490e-aeab-2a5230a48d4f', 'Country', 'Japan'),
('0f40229e-dc58-4111-b709-b9a5266f587f', 'City', 'Tokyo'),
('90c3287b-2ff8-46b4-bfb5-332a979a199a', 'State', 'Kanto'),
('44002c51-cc57-489f-bcf7-4f2abc6ddeb8', 'Street', 'Shibuya');

INSERT INTO accounts (account_id, created_at)
VALUES
('394e67b6-2850-4ddf-a4c9-c2a619d5bf70', '2024-01-10 12:00:00'),
('018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', '2024-01-10 12:00:00'),
('018d6189-9ad0-7b72-801b-1e0de0d3c214', '2024-01-10 12:00:00'),
('018d96bf-8dce-7f68-a926-b5d7ed6ed883', '2024-01-10 12:00:00'),
('018d9b4d-9438-79ac-b533-1323d4ec9b9f', '2024-01-10 12:00:00'),
('018da09e-c6ca-795e-878d-32bb8c1e5cac', '2024-01-10 12:00:00'),
('018ddee7-3a8e-7387-a03e-2b37173b5ada', '2024-01-10 12:00:00'),
('018df2ef-43d0-7ba3-9159-13b2b6634042', '2024-01-10 12:00:00'),
('018df2ef-d77a-784a-92d3-3f52deb284bd', '2024-01-10 12:00:00'),
('018df53c-c5a6-71a2-bf90-2f751f342d4c', '2024-01-10 12:00:00'),
('018df551-4339-730c-8031-618eb8ef66b5', '2024-01-10 12:00:00'),
('018e0e5c-98a5-76de-9ede-13118ba8c996', '2024-01-10 12:00:00'),
('018e0ea9-6b88-71d3-a887-0cf22ede3e0c', '2024-01-10 12:00:00'),
('018e0ebc-e842-7bdb-bf50-05177e07a1c7', '2024-01-10 12:00:00'),
('018e1398-3d80-76ce-9623-9a6caae8378e', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018de2f6-968d-7458-9c67-69ae5698a143', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'email', '', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'firebase', '2024-01-10 12:00:00'),
('018de2f6-b536-7f9c-bd34-dcf319ee4127', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', 'email', '', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', 'firebase', '2024-01-10 12:00:00'),
('018de2ff-7d69-7f8d-9d19-57bb4106f594', '018d6189-9ad0-7b72-801b-1e0de0d3c214', 'email', '', '018d6189-9ad0-7b72-801b-1e0de0d3c214', 'firebase', '2024-01-10 12:00:00'),
('018de2f6-ca47-7fc9-832a-3d725120c55b', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', 'email', '', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', 'firebase', '2024-01-10 12:00:00'),
('018de2f6-de72-7b8c-92ab-b72b90d41ccd', '018d9b4d-9438-79ac-b533-1323d4ec9b9f', 'email', '', '018d9b4d-9438-79ac-b533-1323d4ec9b9f', 'firebase', '2024-01-10 12:00:00'),
('018de2f6-f23e-7a3c-ab51-3117f07c1930', '018da09e-c6ca-795e-878d-32bb8c1e5cac', 'email', '', '018da09e-c6ca-795e-878d-32bb8c1e5cac', 'firebase', '2024-01-10 12:00:00'),
('018de2f7-0939-7cb7-a1f0-c7959bf6edd7', '018ddee7-3a8e-7387-a03e-2b37173b5ada', 'email', '', '018ddee7-3a8e-7387-a03e-2b37173b5ada', 'firebase', '2024-01-10 12:00:00'),
('018df2f3-5922-789d-b529-9b98ab707514', '018df2ef-43d0-7ba3-9159-13b2b6634042', 'google', '', 'MuJcEqPqy9r3wJ85GWsV3SszVJ6X', 'firebase', '2024-01-10 12:00:00'),
('018df2f4-d77c-7b11-9e98-a6d03d70a27a', '018df2ef-d77a-784a-92d3-3f52deb284bd', 'google', '', 'Xk1n15UQOFbml4RoF0QdCza5n0dU', 'firebase', '2024-01-10 12:00:00'),
('018df53c-f868-7f2a-bafd-9cda1fe15e8a', '018df53c-c5a6-71a2-bf90-2f751f342d4c', 'email', '', '018df53c-c5a6-71a2-bf90-2f751f342d4c', 'firebase', '2024-01-10 12:00:00'),
('018df551-d07a-7761-8c69-7de98d195e26', '018df551-4339-730c-8031-618eb8ef66b5', 'google', '', 'ZHjoHCDE0C1EHxLIQvNgiygTXu9A', 'firebase', '2024-01-10 12:00:00'),
('018e0e5d-e8cd-7b3b-8d0b-5b4daac55cdc', '018e0e5c-98a5-76de-9ede-13118ba8c996', 'email', '', '018e0e5c-98a5-76de-9ede-13118ba8c996', 'firebase', '2024-01-10 12:00:00'),
('018e0ea9-aaab-7079-9046-55cfac836d3f', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c', 'google', 'https://github.com/ryo034/image/assets/55078625/af9fae15-baf3-451e-820a-99f7e246af31', 'f8fVGXfC3dmym8XHbFyCs1LvwJ7O', 'firebase', '2024-01-10 12:00:00'),
('018e0ebd-7922-7c48-b598-5562cc7fa29c', '018e0ebc-e842-7bdb-bf50-05177e07a1c7', 'google', 'https://github.com/ryo034/image/assets/55078625/ddeb3605-2291-4c19-81ec-6d890c7d0219', 'QA1ViUeGbDWJQfydWYJbRpzXNVEk', 'firebase', '2024-01-10 12:00:00'),
('018e1398-3d80-796a-acee-44d10b7644ec', '018e1398-3d80-76ce-9623-9a6caae8378e', 'google', 'https://github.com/ryo034/image/assets/55078625/967e0e8c-a2be-4004-834a-d56a263b89ce', 'UPeY3R7sVON9d9i8mq7KWLkyNXXZ', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e09c2-9924-7048-9f08-afa2f3ea5b53', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'account@example.com', '2024-01-10 12:00:00'),
('018e09c2-a8d4-7eb8-966d-40069a2ad41a', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', 'login_logout_login@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-75d3-9b72-91a4c2b2d8aa', '018d6189-9ad0-7b72-801b-1e0de0d3c214', 'unfinished_onboarding@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-71bd-bbca-5cec4e063d46', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', 'invite_test_1@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-74e3-80a7-e33c4d0ddc9c', '018d9b4d-9438-79ac-b533-1323d4ec9b9f', 'invite_test_already_joined_any_workspace@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-79eb-a643-9c160d5b998d', '018da09e-c6ca-795e-878d-32bb8c1e5cac', 'invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-73ec-bf34-bb12b5851d00', '018ddee7-3a8e-7387-a03e-2b37173b5ada', 'update_me_member_profile@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-7635-9973-b88c73a1d73a', '018df2ef-43d0-7ba3-9159-13b2b6634042', 'google_auth_test_no_name@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-7409-ac0b-00bbc07acf7f', '018df2ef-d77a-784a-92d3-3f52deb284bd', 'google_auth_test_has_name@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-7de8-8531-51a58e1e3a96', '018df53c-c5a6-71a2-bf90-2f751f342d4c', 'invite_test_already_joined_any_workspace_by_email@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-787a-94e6-adfc83e2c457', '018df551-4339-730c-8031-618eb8ef66b5', 'invite_test_already_joined_any_workspace_by_google@example.com', '2024-01-10 12:00:00'),
('018e0e5e-8486-7286-9114-f7479e18f94f', '018e0e5c-98a5-76de-9ede-13118ba8c996', 'update_me_update_profile_photo@example.com', '2024-01-10 12:00:00'),
('018e0eaa-8b25-7a21-a221-f85f8a85d149', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c', 'invite_test_has_photo_by_google_accept_with_email@example.com', '2024-01-10 12:00:00'),
('018e0ebe-185a-75fb-8096-b6ad5b02a9fb', '018e0ebc-e842-7bdb-bf50-05177e07a1c7', 'invite_test_has_photo_by_google_accept_with_has_photo_google@example.com', '2024-01-10 12:00:00'),
('018e1398-3d80-7fab-8333-89b92a91eca6', '018e1398-3d80-76ce-9623-9a6caae8378e', 'test_has_photo_google_setup_photo@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e09c2-9924-7048-9f08-afa2f3ea5b53', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70'),
('018e09c2-a8d4-7eb8-966d-40069a2ad41a', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2'),
('018e09c2-d3f3-75d3-9b72-91a4c2b2d8aa', '018d6189-9ad0-7b72-801b-1e0de0d3c214'),
('018e09c2-d3f3-71bd-bbca-5cec4e063d46', '018d96bf-8dce-7f68-a926-b5d7ed6ed883'),
('018e09c2-d3f3-74e3-80a7-e33c4d0ddc9c', '018d9b4d-9438-79ac-b533-1323d4ec9b9f'),
('018e09c2-d3f3-79eb-a643-9c160d5b998d', '018da09e-c6ca-795e-878d-32bb8c1e5cac'),
('018e09c2-d3f3-73ec-bf34-bb12b5851d00', '018ddee7-3a8e-7387-a03e-2b37173b5ada'),
('018e09c2-d3f3-7635-9973-b88c73a1d73a', '018df2ef-43d0-7ba3-9159-13b2b6634042'),
('018e09c2-d3f3-7409-ac0b-00bbc07acf7f', '018df2ef-d77a-784a-92d3-3f52deb284bd'),
('018e09c2-d3f3-7de8-8531-51a58e1e3a96', '018df53c-c5a6-71a2-bf90-2f751f342d4c'),
('018e09c2-d3f3-787a-94e6-adfc83e2c457', '018df551-4339-730c-8031-618eb8ef66b5'),
('018e0e5e-8486-7286-9114-f7479e18f94f', '018e0e5c-98a5-76de-9ede-13118ba8c996'),
('018e0eaa-8b25-7a21-a221-f85f8a85d149', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c'),
('018e0ebe-185a-75fb-8096-b6ad5b02a9fb', '018e0ebc-e842-7bdb-bf50-05177e07a1c7'),
('018e1398-3d80-7fab-8333-89b92a91eca6', '018e1398-3d80-76ce-9623-9a6caae8378e');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e088e-fd36-722d-a927-8cfd34a642bd', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'John Doe', '2024-01-10 12:00:00'),
('018e088f-9eab-78bf-9b3f-c4aacb50e666', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', 'Login LogoutRetry', '2024-01-10 12:00:00'),
('018e088f-af9c-7524-b0bd-f3d4dfb24f26', '018d6189-9ad0-7b72-801b-1e0de0d3c214', 'Unfinished Onboarding', '2024-01-10 12:00:00'),
('018e088f-c783-7e2c-9d8c-371168132855', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', 'Invite TestOne', '2024-01-10 12:00:00'),
('018e088f-d9f3-74ae-9c94-25c0d99ff178', '018d9b4d-9438-79ac-b533-1323d4ec9b9f', 'Invite TestTwo', '2024-01-10 12:00:00'),
('018e088f-eada-78ba-b8ea-8d27d57db944', '018da09e-c6ca-795e-878d-32bb8c1e5cac', 'Invite TestThree', '2024-01-10 12:00:00'),
('018e088f-fcc3-7586-a75c-401167937632', '018ddee7-3a8e-7387-a03e-2b37173b5ada', 'UpdateMe MemberProfile', '2024-01-10 12:00:00'),
('018e0890-2197-7ab2-a17c-b0e168cd6080', '018df2ef-d77a-784a-92d3-3f52deb284bd', 'GoogleAuthTest HasNameTest', '2024-01-10 12:00:00'),
('018e0890-34c0-7ed8-80d6-6626f4d37531', '018df53c-c5a6-71a2-bf90-2f751f342d4c', 'InviteGoogleAuthTest AlreadyJoined', '2024-01-10 12:00:00'),
('018e0890-4745-7655-af59-805da2375591', '018df551-4339-730c-8031-618eb8ef66b5', 'InviteGoogleAuthTest AlreadyJoinedGoogle', '2024-01-10 12:00:00'),
('018e0e5f-16d4-7fa5-8233-022f53b54b55', '018e0e5c-98a5-76de-9ede-13118ba8c996', 'UpdateMe UpdateProfilePhoto', '2024-01-10 12:00:00'),
('018e0eab-10a4-7872-88f1-9d74c2bd04a9', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c', 'InviteTest HasPhotoAcceptWithEmail', '2024-01-10 12:00:00'),
('018e0ebe-8ca7-7ecd-92b0-45dd064dfb14', '018e0ebc-e842-7bdb-bf50-05177e07a1c7', 'InviteTest HasPhotoAcceptWithGoogle', '2024-01-10 12:00:00'),
('018e1398-3d80-79fb-af13-e2868037907c', '018e1398-3d80-76ce-9623-9a6caae8378e', 'HasPhotoGoogle SetupPhoto', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e088e-fd36-722d-a927-8cfd34a642bd', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70'),
('018e088f-9eab-78bf-9b3f-c4aacb50e666', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2'),
('018e088f-af9c-7524-b0bd-f3d4dfb24f26', '018d6189-9ad0-7b72-801b-1e0de0d3c214'),
('018e088f-c783-7e2c-9d8c-371168132855', '018d96bf-8dce-7f68-a926-b5d7ed6ed883'),
('018e088f-d9f3-74ae-9c94-25c0d99ff178', '018d9b4d-9438-79ac-b533-1323d4ec9b9f'),
('018e088f-eada-78ba-b8ea-8d27d57db944', '018da09e-c6ca-795e-878d-32bb8c1e5cac'),
('018e088f-fcc3-7586-a75c-401167937632', '018ddee7-3a8e-7387-a03e-2b37173b5ada'),
('018e0890-2197-7ab2-a17c-b0e168cd6080', '018df2ef-d77a-784a-92d3-3f52deb284bd'),
('018e0890-34c0-7ed8-80d6-6626f4d37531', '018df53c-c5a6-71a2-bf90-2f751f342d4c'),
('018e0890-4745-7655-af59-805da2375591', '018df551-4339-730c-8031-618eb8ef66b5'),
('018e0e5f-16d4-7fa5-8233-022f53b54b55', '018e0e5c-98a5-76de-9ede-13118ba8c996'),
('018e0eab-10a4-7872-88f1-9d74c2bd04a9', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c'),
('018e0ebe-8ca7-7ecd-92b0-45dd064dfb14', '018e0ebc-e842-7bdb-bf50-05177e07a1c7'),
('018e1398-3d80-79fb-af13-e2868037907c', '018e1398-3d80-76ce-9623-9a6caae8378e');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '2024-01-10 12:00:00'),
('018d91d5-2ed0-7211-b2e6-cf26182ac4f9', '2024-01-10 12:00:00'),
('018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00'),
('018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00'),
('018ddee6-6446-7f9d-b750-469a7c2dfac5', '2024-01-10 12:00:00'),
('018e0e5f-9c4b-7062-aaa7-3db3fde10354', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_id, name, subdomain, created_at, updated_at)
VALUES
('c1bd2603-b9cd-4f84-8b83-3548f6ae150b', 'Example', 'example', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d91d5-2ed0-7211-b2e6-cf26182ac4f9', 'LoginLogoutRetry', 'login-logout-retry', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d96b9-c920-7434-b5c3-02e5e920ae9d', 'InviteTest 1', 'invite-test-1', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d9b4d-e340-74f7-914c-2476eff949bb', 'InviteTest 2', 'invite-test-2', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018ddee6-6446-7f9d-b750-469a7c2dfac5', 'UpdateMemberMeProfile Workspace', 'update-me-member-profile', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018e0e5f-9c4b-7062-aaa7-3db3fde10354', 'UpdateAccountPhoto Workspace', 'update-me-account-profile', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO members (member_id, account_id, workspace_id, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '2024-01-10 12:00:00'),
('018d91d5-c061-78ba-9263-d6ef9e7e6783', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', '018d91d5-2ed0-7211-b2e6-cf26182ac4f9', '2024-01-10 12:00:00'),
('018d96b9-f674-7ff6-83eb-506eca6452be', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', '018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00'),
('018df53e-4c77-79de-b725-c43ebcb79450', '018df53c-c5a6-71a2-bf90-2f751f342d4c', '018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00'),
('018df552-5086-7b84-8601-d04c319d2e44', '018df551-4339-730c-8031-618eb8ef66b5', '018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00'),
('018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '018d9b4d-9438-79ac-b533-1323d4ec9b9f', '018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00'),
('018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', '018da09e-c6ca-795e-878d-32bb8c1e5cac', '018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00'),
('018ddee7-2419-7c62-a9be-a56a2c07916e', '018ddee7-3a8e-7387-a03e-2b37173b5ada', '018ddee6-6446-7f9d-b750-469a7c2dfac5', '2024-01-10 12:00:00'),
('018e0e60-4f73-72bd-84f7-ae1abaa7fef1', '018e0e5c-98a5-76de-9ede-13118ba8c996', '018e0e5f-9c4b-7062-aaa7-3db3fde10354', '2024-01-10 12:00:00'),
('018e0eab-df36-7295-a917-9c81e8c6671c', '018e0ea9-6b88-71d3-a887-0cf22ede3e0c', '018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00'),
('018e0ebf-0806-7f10-a73b-28db9f4d2349', '018e0ebc-e842-7bdb-bf50-05177e07a1c7', '018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00'),
('018e1398-3d80-7fae-a661-3c234a9c5c53', '018e1398-3d80-76ce-9623-9a6caae8378e', '018e0e5f-9c4b-7062-aaa7-3db3fde10354', '2024-01-10 12:00:00');

INSERT INTO member_roles (member_role_id, member_id, role, created_at)
VALUES
('018df76b-260d-759f-9b47-fb5f611f5da6', '377eba35-5560-4f48-a99d-19cbd6a82b0d', 'owner', '2024-01-10 12:00:00'),
('018df76b-3cbe-7e58-81cf-431eeef1bffe', '018d91d5-c061-78ba-9263-d6ef9e7e6783', 'owner', '2024-01-10 12:00:00'),
('018df76b-56e0-7371-a69b-1dea397a75d8', '018d96b9-f674-7ff6-83eb-506eca6452be', 'owner', '2024-01-10 12:00:00'),
('018df76b-6bf7-788f-bc7b-e1102924573d', '018df53e-4c77-79de-b725-c43ebcb79450', 'admin', '2024-01-10 12:00:00'),
('018df76b-81da-7cb2-a5fe-b849b52a939e', '018df552-5086-7b84-8601-d04c319d2e44', 'admin', '2024-01-10 12:00:00'),
('018df76b-9717-788a-9b02-548a9666ac44', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', 'owner', '2024-01-10 12:00:00'),
('018df76b-ace2-7420-87d5-666e55aa18b7', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', 'admin', '2024-01-10 12:00:00'),
('018df76b-bf83-723d-b345-0f9d6d94f0a4', '018ddee7-2419-7c62-a9be-a56a2c07916e', 'owner', '2024-01-10 12:00:00'),
('018e0e60-ca5a-7c43-9bc3-67f7d65c4512', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1', 'owner', '2024-01-10 12:00:00'),
('018e0eac-2417-70c4-a18d-a6acc026520b', '018e0eab-df36-7295-a917-9c81e8c6671c', 'admin', '2024-01-10 12:00:00'),
('018e0ebf-738d-798a-b861-739f8d4810cf', '018e0ebf-0806-7f10-a73b-28db9f4d2349', 'admin', '2024-01-10 12:00:00'),
('018e1398-3d80-743b-a74f-3497770daaf9', '018e1398-3d80-7fae-a661-3c234a9c5c53', 'admin', '2024-01-10 12:00:00');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018d6bc0-3884-7420-a802-f857192c7e24', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-10 12:00:00'),
('018d91d6-34a8-7c2b-8d1b-37622cf2fa1d', '018d91d5-c061-78ba-9263-d6ef9e7e6783', '2024-01-10 12:00:00'),
('018d96ba-6ebd-77a6-9534-958f8fe487ce', '018d96b9-f674-7ff6-83eb-506eca6452be', '2024-01-10 12:00:00'),
('018df53e-965e-7e7e-8842-fd0e4135caf0', '018df53e-4c77-79de-b725-c43ebcb79450', '2024-01-10 12:00:00'),
('018df552-75b8-76c7-afc0-bb51404f9359', '018df552-5086-7b84-8601-d04c319d2e44', '2024-01-10 12:00:00'),
('018d9b4e-0b6e-7f6e-8b7e-9f6e8d7e6f8e', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '2024-01-10 12:00:00'),
('018da0dc-7577-7e53-8db0-ac3d68801240', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', '2024-01-10 12:00:00'),
('018ddee7-3a8e-7387-a03e-2b37173b5ada', '018ddee7-2419-7c62-a9be-a56a2c07916e', '2024-01-10 12:00:00'),
('018e0e61-3bf7-726a-b12d-8644926a08fd', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1', '2024-01-10 12:00:00'),
('018e0eac-6b49-7343-88cc-b99b00e59fbf', '018e0eab-df36-7295-a917-9c81e8c6671c', '2024-01-10 12:00:00'),
('018e0ebf-c847-7214-a31a-12f00ef1a627', '018e0ebf-0806-7f10-a73b-28db9f4d2349', '2024-01-10 12:00:00'),
('018e1398-3d80-7103-88f9-20c3f226f13b', '018e1398-3d80-7fae-a661-3c234a9c5c53', '2024-01-10 12:00:00');

INSERT INTO member_latest_login_histories (member_login_history_id, member_id)
VALUES
('018d6bc0-3884-7420-a802-f857192c7e24', '377eba35-5560-4f48-a99d-19cbd6a82b0d'),
('018d91d6-34a8-7c2b-8d1b-37622cf2fa1d', '018d91d5-c061-78ba-9263-d6ef9e7e6783'),
('018d96ba-6ebd-77a6-9534-958f8fe487ce', '018d96b9-f674-7ff6-83eb-506eca6452be'),
('018df53e-965e-7e7e-8842-fd0e4135caf0', '018df53e-4c77-79de-b725-c43ebcb79450'),
('018df552-75b8-76c7-afc0-bb51404f9359', '018df552-5086-7b84-8601-d04c319d2e44'),
('018d9b4e-0b6e-7f6e-8b7e-9f6e8d7e6f8e', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d'),
('018da0dc-7577-7e53-8db0-ac3d68801240', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5'),
('018ddee7-3a8e-7387-a03e-2b37173b5ada', '018ddee7-2419-7c62-a9be-a56a2c07916e'),
('018e0e61-3bf7-726a-b12d-8644926a08fd', '018e0e60-4f73-72bd-84f7-ae1abaa7fef1'),
('018e0eac-6b49-7343-88cc-b99b00e59fbf', '018e0eab-df36-7295-a917-9c81e8c6671c'),
('018e0ebf-c847-7214-a31a-12f00ef1a627', '018e0ebf-0806-7f10-a73b-28db9f4d2349'),
('018e1398-3d80-7103-88f9-20c3f226f13b', '018e1398-3d80-7fae-a661-3c234a9c5c53');

INSERT INTO member_profiles (member_id, member_id_number, display_name, bio, created_at, updated_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', 'DEV-12345', 'John Doe', 'John Doe is a passionate software engineer with 8 years of experience specializing in web development, particularly with React and Node.js. A graduate from MIT with a strong focus on clean architecture and Agile methodologies, John has successfully led multiple projects, from innovative startups to established tech giants. He''s a firm believer in continuous learning, contributing regularly to open-source projects, and sharing insights through tech blogs and meetups. Outside of work, John enjoys hiking 🚶‍♂️, drone photography 📸, and playing the guitar 🎸. He''s committed to using technology to drive positive social change.', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d91d5-c061-78ba-9263-d6ef9e7e6783', 'DEV-67890', 'Login LogoutRetry', 'bio', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d96b9-f674-7ff6-83eb-506eca6452be', 'DEV-54321', 'Invite TestOne', 'bio', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018df53e-4c77-79de-b725-c43ebcb79450', 'DEV-54322', 'InviteGoogleAuthTest AlreadyJoined', 'bio', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018df552-5086-7b84-8601-d04c319d2e44', 'DEV-54323', 'InviteGoogleAuthTest AlreadyJoinedGoogle', 'bio', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', 'DEV-09876', 'Invite TestTwo', 'bio', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', 'DEV-54321', 'Invite TestThree', 'bio', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018ddee7-2419-7c62-a9be-a56a2c07916e', 'DEV-54321', 'UpdateMe MemberProfile DisplayName', 'bio', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018e0e60-4f73-72bd-84f7-ae1abaa7fef1', 'DEV-54321', 'UpdateMe UpdateProfilePhoto', 'bio', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018e0eab-df36-7295-a917-9c81e8c6671c', 'DEV-54321', 'InviteTest HasPhotoAcceptWithEmail', 'bio', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018e0ebf-0806-7f10-a73b-28db9f4d2349', 'DEV-54322', 'InviteTest HasPhotoAcceptWithGoogle', 'bio', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018e1398-3d80-7fae-a661-3c234a9c5c53', 'DEV-54321', 'HasPhotoGoogle SetupPhoto', 'bio', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO member_addresses (member_id, postal_code, building_component_id, street_address_component_id, city_component_id, state_component_id, country_component_id, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '150-0002', null, '44002c51-cc57-489f-bcf7-4f2abc6ddeb8', '0f40229e-dc58-4111-b709-b9a5266f587f', '90c3287b-2ff8-46b4-bfb5-332a979a199a', '10269b87-98ce-490e-aeab-2a5230a48d4f', '2024-01-10 12:00:00');

INSERT INTO membership_periods (member_id, start_date, end_date, activity, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018d91d5-c061-78ba-9263-d6ef9e7e6783', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018d96b9-f674-7ff6-83eb-506eca6452be', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018df53e-4c77-79de-b725-c43ebcb79450', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018df552-5086-7b84-8601-d04c319d2e44', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018ddee7-2419-7c62-a9be-a56a2c07916e', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018e0e60-4f73-72bd-84f7-ae1abaa7fef1', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018e0eab-df36-7295-a917-9c81e8c6671c', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018e0ebf-0806-7f10-a73b-28db9f4d2349', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018e1398-3d80-7fae-a661-3c234a9c5c53', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00');

INSERT INTO invitation_units (invitation_unit_id, workspace_id, invited_by, created_at)
VALUES
('018db49e-b4dd-7828-8a9b-fa8f9d12b552', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2023-01-09 12:00:00'),
('018db4a4-c350-747b-8c4f-bd827e08174b', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-10 12:00:00'),
('018db49e-d5a5-7ed1-aea6-e018d2e4bd38', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '018d96b9-f674-7ff6-83eb-506eca6452be', '2024-01-10 12:00:00'),
('018e0ead-064c-7eea-b664-c414d44270f0', '018d9b4d-e340-74f7-914c-2476eff949bb', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '2024-01-10 12:00:00');

INSERT INTO invitations (invitation_id, invitation_unit_id)
VALUES
('018d96b8-0bb6-7822-b3b5-78a5d0e8790e', '018db49e-b4dd-7828-8a9b-fa8f9d12b552'),
('018d96b8-2211-7862-bcbe-e9f4d002a8fc', '018db4a4-c350-747b-8c4f-bd827e08174b'),
('018dcee8-97b2-7ffe-a6cb-94093483fa12', '018db4a4-c350-747b-8c4f-bd827e08174b'),
('018d96bb-2a6e-70ed-bd74-8565ac2960ac', '018db4a4-c350-747b-8c4f-bd827e08174b'),
('018dbe26-f91d-7b0b-9a18-ff5181136ffb', '018db4a4-c350-747b-8c4f-bd827e08174b'),
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
('018d96b8-2211-7862-bcbe-e9f4d002a8fc', '018d96b7-df68-792f-97d0-d6a044c2b4a2', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018dcee8-97b2-7ffe-a6cb-94093483fa12', '018dcee9-4ec8-7f93-9a9c-f9ad7ae3d592', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d96bb-2a6e-70ed-bd74-8565ac2960ac', '018d96bb-975d-769c-aa3d-dfe09fc9f207', '2024-01-17 12:00:00', '2024-01-10 12:00:00'),
('018dbe26-f91d-7b0b-9a18-ff5181136ffb', '018dbe28-a7c6-7b51-885a-7c4647e4aff4', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
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
('018d96b8-2211-7862-bcbe-e9f4d002a8fc', 'invite_test_not_expired@example.com'),
('018dcee8-97b2-7ffe-a6cb-94093483fa12', 'invite_test_already_verified@example.com'),
('018d96bb-2a6e-70ed-bd74-8565ac2960ac', 'invite_test_already_accepted@example.com'),
('018dbe26-f91d-7b0b-9a18-ff5181136ffb', 'invite_test_revoked@example.com'),
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

INSERT INTO invitation_events (invitation_event_id, invitation_id, event_type, created_at)
VALUES
('018dcee9-0494-7f71-97ec-919c04becb62', '018dcee8-97b2-7ffe-a6cb-94093483fa12', 'verified', '2023-01-10 15:00:00'),
('018db3ff-2a1a-7445-b464-3a84071b9549', '018d96bb-2a6e-70ed-bd74-8565ac2960ac', 'verified', '2023-01-10 15:00:00'),
('018dca3b-8713-71ff-b176-bfaedd0cf766', '018d96bb-2a6e-70ed-bd74-8565ac2960ac', 'accepted', '2023-01-11 15:00:00'),
('018dbe29-c687-7a1d-b93d-bb3502c32988', '018dbe26-f91d-7b0b-9a18-ff5181136ffb', 'revoked', '2024-01-10 15:00:00');
