INSERT INTO address_components (component_id, component_type, component_name)
VALUES
('10269b87-98ce-490e-aeab-2a5230a48d4f', 'Country', 'Japan'),
('0f40229e-dc58-4111-b709-b9a5266f587f', 'City', 'Tokyo'),
('90c3287b-2ff8-46b4-bfb5-332a979a199a', 'State', 'Kanto'),
('44002c51-cc57-489f-bcf7-4f2abc6ddeb8', 'Street', 'Shibuya');

INSERT INTO accounts (account_id, created_at)
VALUES
('394e67b6-2850-4ddf-a4c9-c2a619d5bf70', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018de2f6-968d-7458-9c67-69ae5698a143', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'email', '', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e09c2-9924-7048-9f08-afa2f3ea5b53', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'account@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e09c2-9924-7048-9f08-afa2f3ea5b53', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e088e-fd36-722d-a927-8cfd34a642bd', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'John Doe', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e088e-fd36-722d-a927-8cfd34a642bd', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_detail_id, workspace_id, name, subdomain, created_at)
VALUES
('018e200b-9d01-70ed-8c5a-5a5df2a98f11', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', 'Example', 'example', '2024-01-10 12:00:00');

INSERT INTO workspace_latest_details (workspace_detail_id, workspace_id)
VALUES
('018e200b-9d01-70ed-8c5a-5a5df2a98f11', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b');

INSERT INTO members (member_id, account_id, workspace_id, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '2024-01-10 12:00:00');

INSERT INTO member_roles (member_role_id, member_id, assigned_by, role, assigned_at)
VALUES
('018df76b-260d-759f-9b47-fb5f611f5da6', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '377eba35-5560-4f48-a99d-19cbd6a82b0d', 'owner', '2024-01-10 12:00:00');

INSERT INTO member_latest_roles (member_role_id, member_id)
VALUES
('018df76b-260d-759f-9b47-fb5f611f5da6', '377eba35-5560-4f48-a99d-19cbd6a82b0d');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018d6bc0-3884-7420-a802-f857192c7e24', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-10 12:00:00');

INSERT INTO member_latest_login_histories (member_login_history_id, member_id)
VALUES
('018d6bc0-3884-7420-a802-f857192c7e24', '377eba35-5560-4f48-a99d-19cbd6a82b0d');

INSERT INTO member_profiles (member_profile_id, member_id, member_id_number, display_name, bio, created_at)
VALUES
('018e2216-64a3-7438-9300-1cdc4354d1de', '377eba35-5560-4f48-a99d-19cbd6a82b0d', 'DEV-12345', 'John Doe', 'bio', '2024-01-10 12:00:00');

INSERT INTO member_latest_profiles (member_profile_id, member_id)
VALUES
('018e2216-64a3-7438-9300-1cdc4354d1de', '377eba35-5560-4f48-a99d-19cbd6a82b0d');

INSERT INTO membership_events (membership_event_id, member_id, event_type, created_by, event_at)
VALUES
('018e2ff9-c432-7093-b091-943915c59284', '377eba35-5560-4f48-a99d-19cbd6a82b0d', 'join', '377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-10 12:00:00');

INSERT INTO latest_membership_events (membership_event_id, member_id)
VALUES
('018e2ff9-c432-7093-b091-943915c59284', '377eba35-5560-4f48-a99d-19cbd6a82b0d');

INSERT INTO member_addresses (member_id, postal_code, building_component_id, street_address_component_id, city_component_id, state_component_id, country_component_id, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '150-0002', null, '44002c51-cc57-489f-bcf7-4f2abc6ddeb8', '0f40229e-dc58-4111-b709-b9a5266f587f', '90c3287b-2ff8-46b4-bfb5-332a979a199a', '10269b87-98ce-490e-aeab-2a5230a48d4f', '2024-01-10 12:00:00');
