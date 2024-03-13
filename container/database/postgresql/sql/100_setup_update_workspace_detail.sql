INSERT INTO accounts (account_id, created_at)
VALUES
('018e201b-67d4-771c-b67c-91433823f052', '2024-01-10 12:00:00'),
('018e21d7-6278-7bb7-bf7f-a5f9095c10dc', '2024-01-10 12:00:00'),
('018e21d7-6278-71da-ae8f-18d988f5883f', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018e201b-67d4-714a-800d-276bda4c856c', '018e201b-67d4-771c-b67c-91433823f052', 'email', '', '018e201b-67d4-771c-b67c-91433823f052', 'firebase', '2024-01-10 12:00:00'),
('018e21d7-6278-7e46-b48e-f958f37e65b6', '018e21d7-6278-7bb7-bf7f-a5f9095c10dc', 'email', '', '018e21d7-6278-7bb7-bf7f-a5f9095c10dc', 'firebase', '2024-01-10 12:00:00'),
('018e21d7-6278-79a1-bb03-bb2004307831', '018e21d7-6278-71da-ae8f-18d988f5883f', 'email', '', '018e21d7-6278-71da-ae8f-18d988f5883f', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e201b-67d4-76de-86f5-2a6b6acd72b6', '018e201b-67d4-771c-b67c-91433823f052', 'update_workspace_detail@example.com', '2024-01-10 12:00:00'),
('018e21d7-6278-74f0-bd83-a7b42f138eb2', '018e21d7-6278-7bb7-bf7f-a5f9095c10dc', 'update_workspace_detail_member_role@example.com', '2024-01-10 12:00:00'),
('018e21d7-6278-7156-9139-3d0326e1c300', '018e21d7-6278-71da-ae8f-18d988f5883f', 'update_workspace_detail_guest_role@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e201b-67d4-76de-86f5-2a6b6acd72b6', '018e201b-67d4-771c-b67c-91433823f052'),
('018e21d7-6278-74f0-bd83-a7b42f138eb2', '018e21d7-6278-7bb7-bf7f-a5f9095c10dc'),
('018e21d7-6278-7156-9139-3d0326e1c300', '018e21d7-6278-71da-ae8f-18d988f5883f');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e201b-67d4-75a4-a1ff-95f77b5bfe03', '018e201b-67d4-771c-b67c-91433823f052', 'UpdateWorkspace Detail', '2024-01-10 12:00:00'),
('018e21d7-6279-7ed5-b114-03fc1a4eca60', '018e21d7-6278-7bb7-bf7f-a5f9095c10dc', 'UpdateWorkspace DetailMemberRole', '2024-01-10 12:00:00'),
('018e21d7-6279-71de-8f55-f218b06ef69e', '018e21d7-6278-71da-ae8f-18d988f5883f', 'UpdateWorkspace DetailGuestRole', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e201b-67d4-75a4-a1ff-95f77b5bfe03', '018e201b-67d4-771c-b67c-91433823f052'),
('018e21d7-6279-7ed5-b114-03fc1a4eca60', '018e21d7-6278-7bb7-bf7f-a5f9095c10dc'),
('018e21d7-6279-71de-8f55-f218b06ef69e', '018e21d7-6278-71da-ae8f-18d988f5883f');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('018e201b-67d4-7265-a022-1b29793b2a91', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_detail_id, workspace_id, name, subdomain, created_at)
VALUES
('018e201b-67d4-75a2-b8ef-176c25cf9fe0', '018e201b-67d4-7265-a022-1b29793b2a91', 'UpdateWorkspaceDetail Workspace', 'update-workspace-detail', '2024-01-10 12:00:00');

INSERT INTO workspace_latest_details (workspace_detail_id, workspace_id)
VALUES
('018e201b-67d4-75a2-b8ef-176c25cf9fe0', '018e201b-67d4-7265-a022-1b29793b2a91');

INSERT INTO members (member_id, account_id, workspace_id, created_at)
VALUES
('018e201b-67d4-7460-bbb2-1428e7c2d949', '018e201b-67d4-771c-b67c-91433823f052', '018e201b-67d4-7265-a022-1b29793b2a91', '2024-01-10 12:00:00'),
('018e21d7-6279-70f6-8170-5d3bfdc5c378', '018e21d7-6278-7bb7-bf7f-a5f9095c10dc', '018e201b-67d4-7265-a022-1b29793b2a91', '2024-01-10 12:00:00'),
('018e21d7-6279-788c-9a65-9932d4649535', '018e21d7-6278-71da-ae8f-18d988f5883f', '018e201b-67d4-7265-a022-1b29793b2a91', '2024-01-10 12:00:00');

INSERT INTO member_roles (member_role_id, member_id, assigned_by, role, assigned_at)
VALUES
('018e201b-67d4-7006-a732-e271820a5832', '018e201b-67d4-7460-bbb2-1428e7c2d949', '018e201b-67d4-7460-bbb2-1428e7c2d949', 'owner', '2024-01-10 12:00:00'),
('018e21d7-6279-7f68-966f-269ab1df9803', '018e21d7-6279-70f6-8170-5d3bfdc5c378', '018e21d7-6279-70f6-8170-5d3bfdc5c378', 'member', '2024-01-10 12:00:00'),
('018e21d7-6279-72bb-9b61-2d8692d7c75d', '018e21d7-6279-788c-9a65-9932d4649535', '018e21d7-6279-788c-9a65-9932d4649535', 'guest', '2024-01-10 12:00:00');

INSERT INTO member_latest_roles (member_role_id, member_id)
VALUES
('018e201b-67d4-7006-a732-e271820a5832', '018e201b-67d4-7460-bbb2-1428e7c2d949'),
('018e21d7-6279-7f68-966f-269ab1df9803', '018e21d7-6279-70f6-8170-5d3bfdc5c378'),
('018e21d7-6279-72bb-9b61-2d8692d7c75d', '018e21d7-6279-788c-9a65-9932d4649535');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018e201b-67d4-7fd4-897a-88322769bd4b', '018e201b-67d4-7460-bbb2-1428e7c2d949', '2024-01-10 12:00:00'),
('018e21d7-6279-77d5-b7bf-63c61e7378eb', '018e21d7-6279-70f6-8170-5d3bfdc5c378', '2024-01-10 12:00:00'),
('018e21d7-6279-744f-b122-9be9ff63feb9', '018e21d7-6279-788c-9a65-9932d4649535', '2024-01-10 12:00:00');

INSERT INTO member_latest_login_histories (member_login_history_id, member_id)
VALUES
('018e201b-67d4-7fd4-897a-88322769bd4b', '018e201b-67d4-7460-bbb2-1428e7c2d949'),
('018e21d7-6279-77d5-b7bf-63c61e7378eb', '018e21d7-6279-70f6-8170-5d3bfdc5c378'),
('018e21d7-6279-744f-b122-9be9ff63feb9', '018e21d7-6279-788c-9a65-9932d4649535');

INSERT INTO member_profiles (member_profile_id, member_id, member_id_number, display_name, bio, created_at)
VALUES
('018e2216-64a3-71c3-b1d9-522a8d5d0d7b', '018e201b-67d4-7460-bbb2-1428e7c2d949', 'DEV-54321', 'UpdateWorkspace Detail', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-76e2-98cb-37dc6b092c49', '018e21d7-6279-70f6-8170-5d3bfdc5c378', 'DEV-54321', 'UpdateWorkspace DetailTwo', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-7dff-8c61-1c882827448a', '018e21d7-6279-788c-9a65-9932d4649535', 'DEV-54322', 'UpdateWorkspace DetailThree', 'bio', '2024-01-10 12:00:00');

INSERT INTO member_latest_profiles (member_profile_id, member_id)
VALUES
('018e2216-64a3-71c3-b1d9-522a8d5d0d7b', '018e201b-67d4-7460-bbb2-1428e7c2d949'),
('018e2216-64a3-76e2-98cb-37dc6b092c49', '018e21d7-6279-70f6-8170-5d3bfdc5c378'),
('018e2216-64a3-7dff-8c61-1c882827448a', '018e21d7-6279-788c-9a65-9932d4649535');

INSERT INTO membership_events (membership_event_id, member_id, event_type, created_by, event_at)
VALUES
('018e2ff9-c432-7cfc-acf5-1a41df8e3880', '018e201b-67d4-7460-bbb2-1428e7c2d949', 'join', '018e201b-67d4-7460-bbb2-1428e7c2d949', '2024-01-10 12:00:00'),
('018e2ff9-c432-7555-98c3-97945f1eaefc', '018e21d7-6279-70f6-8170-5d3bfdc5c378', 'join', '018e21d7-6279-70f6-8170-5d3bfdc5c378', '2024-01-10 12:00:00'),
('018e2ff9-c432-7576-8087-ce7a38bfdae5', '018e21d7-6279-788c-9a65-9932d4649535', 'join', '018e21d7-6279-788c-9a65-9932d4649535', '2024-01-10 12:00:00');

INSERT INTO latest_membership_events (membership_event_id, member_id)
VALUES
('018e2ff9-c432-7cfc-acf5-1a41df8e3880', '018e201b-67d4-7460-bbb2-1428e7c2d949'),
('018e2ff9-c432-7555-98c3-97945f1eaefc', '018e21d7-6279-70f6-8170-5d3bfdc5c378'),
('018e2ff9-c432-7576-8087-ce7a38bfdae5', '018e21d7-6279-788c-9a65-9932d4649535');
