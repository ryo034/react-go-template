INSERT INTO accounts (account_id, created_at)
VALUES
('018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018de2f6-b536-7f9c-bd34-dcf319ee4127', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', 'email', '', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e09c2-a8d4-7eb8-966d-40069a2ad41a', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', 'login_logout_login@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e09c2-a8d4-7eb8-966d-40069a2ad41a', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e088f-9eab-78bf-9b3f-c4aacb50e666', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', 'Login LogoutRetry', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e088f-9eab-78bf-9b3f-c4aacb50e666', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('018d91d5-2ed0-7211-b2e6-cf26182ac4f9', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_detail_id, workspace_id, name, subdomain, created_at)
VALUES
('018e200b-9d01-7a70-a8ee-5b7bd09d5cf1', '018d91d5-2ed0-7211-b2e6-cf26182ac4f9', 'LoginLogoutRetry', 'login-logout-retry', '2024-01-10 12:00:00');

INSERT INTO workspace_latest_details (workspace_detail_id, workspace_id)
VALUES
('018e200b-9d01-7a70-a8ee-5b7bd09d5cf1', '018d91d5-2ed0-7211-b2e6-cf26182ac4f9');

INSERT INTO members (member_id, account_id, workspace_id, created_at)
VALUES
('018d91d5-c061-78ba-9263-d6ef9e7e6783', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', '018d91d5-2ed0-7211-b2e6-cf26182ac4f9', '2024-01-10 12:00:00');

INSERT INTO member_roles (member_role_id, member_id, assigned_by, role, assigned_at)
VALUES
('018df76b-3cbe-7e58-81cf-431eeef1bffe', '018d91d5-c061-78ba-9263-d6ef9e7e6783', '018d91d5-c061-78ba-9263-d6ef9e7e6783', 'owner', '2024-01-10 12:00:00');

INSERT INTO member_latest_roles (member_role_id, member_id)
VALUES
('018df76b-3cbe-7e58-81cf-431eeef1bffe', '018d91d5-c061-78ba-9263-d6ef9e7e6783');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018d91d6-34a8-7c2b-8d1b-37622cf2fa1d', '018d91d5-c061-78ba-9263-d6ef9e7e6783', '2024-01-10 12:00:00');

INSERT INTO member_latest_login_histories (member_login_history_id, member_id)
VALUES
('018d91d6-34a8-7c2b-8d1b-37622cf2fa1d', '018d91d5-c061-78ba-9263-d6ef9e7e6783');

INSERT INTO member_profiles (member_profile_id, member_id, member_id_number, display_name, bio, created_at)
VALUES
('018e2216-64a3-7c4c-883c-4e8eec98c901', '018d91d5-c061-78ba-9263-d6ef9e7e6783', 'DEV-67890', 'Login LogoutRetry', 'bio', '2024-01-10 12:00:00');

INSERT INTO member_latest_profiles (member_profile_id, member_id)
VALUES
('018e2216-64a3-7c4c-883c-4e8eec98c901', '018d91d5-c061-78ba-9263-d6ef9e7e6783');

INSERT INTO membership_events (membership_event_id, member_id, event_type, created_by, event_at)
VALUES
('018e2ff9-c432-7760-9548-a8c94c0340f8', '018d91d5-c061-78ba-9263-d6ef9e7e6783', 'join', '018d91d5-c061-78ba-9263-d6ef9e7e6783', '2024-01-10 12:00:00');

INSERT INTO latest_membership_events (membership_event_id, member_id)
VALUES
('018e2ff9-c432-7760-9548-a8c94c0340f8', '018d91d5-c061-78ba-9263-d6ef9e7e6783');
