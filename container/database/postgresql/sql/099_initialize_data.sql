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
('018d96bf-8dce-7f68-a926-b5d7ed6ed883', '2024-01-10 12:00:00'),
('018d9b4d-9438-79ac-b533-1323d4ec9b9f', '2024-01-10 12:00:00'),
('018da09e-c6ca-795e-878d-32bb8c1e5cac', '2024-01-10 12:00:00');

INSERT INTO system_account_profiles (system_account_id, name, email, created_at, updated_at)
VALUES
('394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'John Doe', 'system_account@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', 'Login LogoutRetry', 'login_logout_login@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d6189-9ad0-7b72-801b-1e0de0d3c214', 'Unfinished Onboarding', 'unfinished_onboarding@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d96bf-8dce-7f68-a926-b5d7ed6ed883', 'Invite TestOne', 'invite_test_1@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d9b4d-9438-79ac-b533-1323d4ec9b9f', 'Invite TestTwo', 'invite_test_already_joined_any_workspace@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018da09e-c6ca-795e-878d-32bb8c1e5cac', 'Invite TestThree', 'invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

-- INSERT INTO system_account_phone_numbers (system_account_id, phone_number, created_at, updated_at)
-- VALUES
-- ('394e67b6-2850-4ddf-a4c9-c2a619d5bf70', '09012345678', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
-- ('018d6189-9ad0-7b72-801b-1e0de0d3c214', '09012345679', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '2024-01-10 12:00:00'),
('018d91d5-2ed0-7211-b2e6-cf26182ac4f9', '2024-01-10 12:00:00'),
('018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00'),
('018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_id, name, subdomain, created_at, updated_at)
VALUES
('c1bd2603-b9cd-4f84-8b83-3548f6ae150b', 'Example', 'example', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d91d5-2ed0-7211-b2e6-cf26182ac4f9', 'LoginLogoutRetry', 'login-logout-retry', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d96b9-c920-7434-b5c3-02e5e920ae9d', 'InviteTest 1', 'invite-test-1', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d9b4d-e340-74f7-914c-2476eff949bb', 'InviteTest 2', 'invite-test-2', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO members (member_id, system_account_id, workspace_id, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '2024-01-10 12:00:00'),
('018d91d5-c061-78ba-9263-d6ef9e7e6783', '018d91d4-5a5e-799c-9cfa-de8d0c02d7f2', '018d91d5-2ed0-7211-b2e6-cf26182ac4f9', '2024-01-10 12:00:00'),
('018d96b9-f674-7ff6-83eb-506eca6452be', '018d96bf-8dce-7f68-a926-b5d7ed6ed883', '018d96b9-c920-7434-b5c3-02e5e920ae9d', '2024-01-10 12:00:00'),
('018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '018d9b4d-9438-79ac-b533-1323d4ec9b9f', '018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00'),
('018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', '018da09e-c6ca-795e-878d-32bb8c1e5cac', '018d9b4d-e340-74f7-914c-2476eff949bb', '2024-01-10 12:00:00');


INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018d6bc0-3884-7420-a802-f857192c7e24', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-10 12:00:00'),
('018d91d6-34a8-7c2b-8d1b-37622cf2fa1d', '018d91d5-c061-78ba-9263-d6ef9e7e6783', '2024-01-10 12:00:00'),
('018d96ba-6ebd-77a6-9534-958f8fe487ce', '018d96b9-f674-7ff6-83eb-506eca6452be', '2024-01-10 12:00:00'),
('018d9b4e-0b6e-7f6e-8b7e-9f6e8d7e6f8e', '018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '2024-01-10 12:00:00'),
('018da0dc-7577-7e53-8db0-ac3d68801240', '018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', '2024-01-10 12:00:00');

INSERT INTO member_profiles (member_id, member_id_number, display_name, created_at, updated_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', 'DEV-12345', 'John Doe', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d91d5-c061-78ba-9263-d6ef9e7e6783', 'DEV-67890', 'Login LogoutRetry', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d96b9-f674-7ff6-83eb-506eca6452be', 'DEV-54321', 'Invite TestOne', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', 'DEV-09876', 'Invite TestTwo', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018da0dc-dad2-7ac8-95cc-9c3afdd0dbd5', 'DEV-54321', 'Invite TestThree', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO member_addresses (member_id, postal_code, building_component_id, street_address_component_id, city_component_id, state_component_id, country_component_id, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '150-0002', null, '44002c51-cc57-489f-bcf7-4f2abc6ddeb8', '0f40229e-dc58-4111-b709-b9a5266f587f', '90c3287b-2ff8-46b4-bfb5-332a979a199a', '10269b87-98ce-490e-aeab-2a5230a48d4f', '2024-01-10 12:00:00');

INSERT INTO membership_periods (member_id, start_date, end_date, activity, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018d91d5-c061-78ba-9263-d6ef9e7e6783', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018d96b9-f674-7ff6-83eb-506eca6452be', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00'),
('018d9b4e-d8fb-73be-95c5-46fbc7a37a7d', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00');

INSERT INTO invitation_units (invitation_unit_id, workspace_id, invited_by, created_at)
VALUES
('018db49e-b4dd-7828-8a9b-fa8f9d12b552', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2023-01-09 12:00:00'),
('018db4a4-c350-747b-8c4f-bd827e08174b', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-10 12:00:00'),
('018db49e-d5a5-7ed1-aea6-e018d2e4bd38', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '018d96b9-f674-7ff6-83eb-506eca6452be', '2024-01-10 12:00:00');

INSERT INTO invitations (invitation_id, invitation_unit_id)
VALUES
('018d96b8-0bb6-7822-b3b5-78a5d0e8790e', '018db49e-b4dd-7828-8a9b-fa8f9d12b552'),
('018d96b8-2211-7862-bcbe-e9f4d002a8fc', '018db4a4-c350-747b-8c4f-bd827e08174b'),
('018d96bb-2a6e-70ed-bd74-8565ac2960ac', '018db4a4-c350-747b-8c4f-bd827e08174b'),
('018d9fb4-641f-72f9-bdeb-a493a974dba1', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018d9b6d-01cf-7d3e-8328-76736b6db971', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38'),
('018da09b-ed0c-7688-a8e3-4573104e006f', '018db49e-d5a5-7ed1-aea6-e018d2e4bd38');

INSERT INTO invitation_tokens (invitation_id, token, expired_at, created_at)
VALUES
('018d96b8-0bb6-7822-b3b5-78a5d0e8790e', '018d96b7-587c-7614-b234-e086b1944e74', '2023-01-10 12:00:00', '2023-01-09 12:00:00'),
('018d96b8-2211-7862-bcbe-e9f4d002a8fc', '018d96b7-df68-792f-97d0-d6a044c2b4a2', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d96bb-2a6e-70ed-bd74-8565ac2960ac', '018d96bb-975d-769c-aa3d-dfe09fc9f207', '2024-01-17 12:00:00', '2024-01-10 12:00:00'),
('018d9fb4-641f-72f9-bdeb-a493a974dba1', '018d9fb5-6150-7a4b-a5c8-b5a61e51ee50', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d9b6d-01cf-7d3e-8328-76736b6db971', '018d9fb5-7e56-75ed-952f-ae8aa4fed8c6', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018da09b-ed0c-7688-a8e3-4573104e006f', '018da09e-2fa7-7d3a-ad23-2c9f5cb76b92', '2200-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO invitees (invitation_id, email)
VALUES
('018d96b8-0bb6-7822-b3b5-78a5d0e8790e', 'invite_test_expired@example.com'),
('018d96b8-2211-7862-bcbe-e9f4d002a8fc', 'invite_test_not_expired@example.com'),
('018d96bb-2a6e-70ed-bd74-8565ac2960ac', 'invite_test_already_used@example.com'),
('018d9fb4-641f-72f9-bdeb-a493a974dba1', 'invite_test_not_expired_with_display_name@example.com'),
('018d9b6d-01cf-7d3e-8328-76736b6db971', 'invite_test_already_joined_any_workspace@example.com'),
('018da09b-ed0c-7688-a8e3-4573104e006f', 'invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com');

INSERT INTO invitee_names (invitation_id, display_name)
VALUES
('018d9fb4-641f-72f9-bdeb-a493a974dba1', 'Invite Test'),
('018da09b-ed0c-7688-a8e3-4573104e006f', 'Invite TestThreeChanged');

INSERT INTO invitations_events (invitations_event_id, invitation_id, event_type, created_at)
VALUES
('018db3ff-2a1a-7445-b464-3a84071b9549', '018d96bb-2a6e-70ed-bd74-8565ac2960ac', 'verified', '2023-01-10 15:00:00');
