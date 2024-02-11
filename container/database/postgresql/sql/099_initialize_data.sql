INSERT INTO address_components (component_id, component_type, component_name)
VALUES
    ('10269b87-98ce-490e-aeab-2a5230a48d4f', 'Country', 'Japan'),
    ('0f40229e-dc58-4111-b709-b9a5266f587f', 'City', 'Tokyo'),
    ('90c3287b-2ff8-46b4-bfb5-332a979a199a', 'State', 'Kanto'),
    ('44002c51-cc57-489f-bcf7-4f2abc6ddeb8', 'Street', 'Shibuya');

INSERT INTO system_accounts (system_account_id, created_at)
VALUES
('394e67b6-2850-4ddf-a4c9-c2a619d5bf70', '2024-01-10 12:00:00'),
('018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', '2024-01-10 12:00:00'),
('018d6189-9ad0-7b72-801b-1e0de0d3c214', '2024-01-10 12:00:00'),
('018d96bf-8dce-7f68-a926-b5d7ed6ed883', '2024-01-10 12:00:00');

INSERT INTO system_account_profiles (system_account_id, name, email, created_at, updated_at)
VALUES
('394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'John Doe', 'system_account@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', 'Login LogoutRetry', 'login_logout_login@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d6189-9ad0-7b72-801b-1e0de0d3c214', 'Unfinished Onboarding', 'unfinished_onboarding@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d96bf-8dce-7f68-a926-b5d7ed6ed883', 'Invite Test', 'invite_test@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

-- INSERT INTO system_account_phone_numbers (system_account_id, phone_number, created_at, updated_at)
-- VALUES
-- ('394e67b6-2850-4ddf-a4c9-c2a619d5bf70', '09012345678', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
-- ('018d6189-9ad0-7b72-801b-1e0de0d3c214', '09012345679', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '2024-01-10 12:00:00'),
('018d91d5-2ed0-7211-b2e6-cf26182ac4f9', '2024-01-10 12:00:00'),
('018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_id, name, subdomain, created_at, updated_at)
VALUES
('c1bd2603-b9cd-4f84-8b83-3548f6ae150b', 'Example', 'example', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d91d5-2ed0-7211-b2e6-cf26182ac4f9', 'LoginLogoutRetry', 'login-logout-retry', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d96b9-c920-7434-b5c3-02e5e920ae9d', 'InviteTest', 'invite-test', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO members (member_id, system_account_id, workspace_id, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '2024-01-10 12:00:00'),
('018d91d5-c061-78ba-9263-d6ef9e7e6783', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', '018d91d5-2ed0-7211-b2e6-cf26182ac4f9', '2024-01-10 12:00:00'),
('018d96b9-f674-7ff6-83eb-506eca6452be', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', '018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018d6bc0-3884-7420-a802-f857192c7e24', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-10 12:00:00'),
('018d91d6-34a8-7c2b-8d1b-37622cf2fa1d', '018d91d5-c061-78ba-9263-d6ef9e7e6783', '2024-01-10 12:00:00'),
('018d96ba-6ebd-77a6-9534-958f8fe487ce', '018d96b9-f674-7ff6-83eb-506eca6452be', '2024-01-10 12:00:00');

INSERT INTO member_profiles (member_id, member_id_number, display_name, created_at, updated_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', 'DEV-12345', 'John Doe', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d91d5-c061-78ba-9263-d6ef9e7e6783', 'DEV-67890', 'Login LogoutRetry', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d96b9-f674-7ff6-83eb-506eca6452be', 'DEV-54321', 'Invite Test', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO member_addresses (member_id, postal_code, building_component_id, street_address_component_id, city_component_id, state_component_id, country_component_id, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '150-0002', null, '44002c51-cc57-489f-bcf7-4f2abc6ddeb8', '0f40229e-dc58-4111-b709-b9a5266f587f', '90c3287b-2ff8-46b4-bfb5-332a979a199a', '10269b87-98ce-490e-aeab-2a5230a48d4f', '2024-01-10 12:00:00');

INSERT INTO membership_periods (member_id, start_date, end_date, activity, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018d91d5-c061-78ba-9263-d6ef9e7e6783', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018d96b9-f674-7ff6-83eb-506eca6452be', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00');

INSERT INTO invited_members (invited_member_id, workspace_id, email, token, used, expires_at, invited_by, created_at, updated_at)
VALUES
('018d96b8-0bb6-7822-b3b5-78a5d0e8790e', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', 'invite_test_expired@example.com', '018d96b7-587c-7614-b234-e086b1944e74' , FALSE, '2023-01-10 12:00:00', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2023-01-09 12:00:00', '2023-01-09 12:00:00'),
('018d96b8-2211-7862-bcbe-e9f4d002a8fc', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', 'invite_test_not_expired@example.com', '018d96b7-df68-792f-97d0-d6a044c2b4a2' , FALSE, '2050-01-10 12:00:00', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-10 12:00:00', '2023-01-09 12:00:00'),
('018d96bb-2a6e-70ed-bd74-8565ac2960ac', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', 'invite_test_already_used@example.com', '018d96bb-975d-769c-aa3d-dfe09fc9f207' , TRUE, '2024-01-11 12:00:00', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-10 12:00:00', '2023-01-10 15:00:00');

