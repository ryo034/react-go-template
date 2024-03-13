INSERT INTO accounts (account_id, created_at)
VALUES
('018ddee7-3a8e-7387-a03e-2b37173b5ada', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018de2f7-0939-7cb7-a1f0-c7959bf6edd7', '018ddee7-3a8e-7387-a03e-2b37173b5ada', 'email', '', '018ddee7-3a8e-7387-a03e-2b37173b5ada', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e09c2-d3f3-73ec-bf34-bb12b5851d00', '018ddee7-3a8e-7387-a03e-2b37173b5ada', 'update_me_member_profile@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e09c2-d3f3-73ec-bf34-bb12b5851d00', '018ddee7-3a8e-7387-a03e-2b37173b5ada');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e088f-fcc3-7586-a75c-401167937632', '018ddee7-3a8e-7387-a03e-2b37173b5ada', 'UpdateMe MemberProfile', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e088f-fcc3-7586-a75c-401167937632', '018ddee7-3a8e-7387-a03e-2b37173b5ada');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('018ddee6-6446-7f9d-b750-469a7c2dfac5', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_detail_id, workspace_id, name, subdomain, created_at)
VALUES
('018e200b-9d01-7eb7-bcb6-faa2b2a24521', '018ddee6-6446-7f9d-b750-469a7c2dfac5', 'UpdateMemberMeProfile Workspace', 'update-me-member-profile', '2024-01-10 12:00:00');

INSERT INTO workspace_latest_details (workspace_detail_id, workspace_id)
VALUES
('018e200b-9d01-7eb7-bcb6-faa2b2a24521', '018ddee6-6446-7f9d-b750-469a7c2dfac5');

INSERT INTO members (member_id, account_id, workspace_id, created_at)
VALUES
('018ddee7-2419-7c62-a9be-a56a2c07916e', '018ddee7-3a8e-7387-a03e-2b37173b5ada', '018ddee6-6446-7f9d-b750-469a7c2dfac5', '2024-01-10 12:00:00');

INSERT INTO member_roles (member_role_id, member_id, assigned_by, role, assigned_at)
VALUES
('018df76b-bf83-723d-b345-0f9d6d94f0a4', '018ddee7-2419-7c62-a9be-a56a2c07916e', '018ddee7-2419-7c62-a9be-a56a2c07916e', 'owner', '2024-01-10 12:00:00');

INSERT INTO member_latest_roles (member_role_id, member_id)
VALUES
('018df76b-bf83-723d-b345-0f9d6d94f0a4', '018ddee7-2419-7c62-a9be-a56a2c07916e');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018ddee7-3a8e-7387-a03e-2b37173b5ada', '018ddee7-2419-7c62-a9be-a56a2c07916e', '2024-01-10 12:00:00');

INSERT INTO member_latest_login_histories (member_login_history_id, member_id)
VALUES
('018ddee7-3a8e-7387-a03e-2b37173b5ada', '018ddee7-2419-7c62-a9be-a56a2c07916e');

INSERT INTO member_profiles (member_profile_id, member_id, member_id_number, display_name, bio, created_at)
VALUES
('018e2216-64a3-7e62-9f98-3e7cb870fc15', '018ddee7-2419-7c62-a9be-a56a2c07916e', 'DEV-54321', 'UpdateMe MemberProfile DisplayName', 'bio', '2024-01-10 12:00:00');

INSERT INTO member_latest_profiles (member_profile_id, member_id)
VALUES
('018e2216-64a3-7e62-9f98-3e7cb870fc15', '018ddee7-2419-7c62-a9be-a56a2c07916e');

INSERT INTO membership_events (membership_event_id, member_id, event_type, created_by, event_at)
VALUES
('018e2ff9-c432-74f3-80ee-a5309e31a3d4', '018ddee7-2419-7c62-a9be-a56a2c07916e', 'join', '018ddee7-2419-7c62-a9be-a56a2c07916e', '2024-01-10 12:00:00');

INSERT INTO latest_membership_events (membership_event_id, member_id)
VALUES
('018e2ff9-c432-74f3-80ee-a5309e31a3d4', '018ddee7-2419-7c62-a9be-a56a2c07916e');
