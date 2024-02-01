INSERT INTO address_components (component_id, component_type, component_name)
VALUES
    ('10269b87-98ce-490e-aeab-2a5230a48d4f', 'Country', 'Japan'),
    ('0f40229e-dc58-4111-b709-b9a5266f587f', 'City', 'Tokyo'),
    ('90c3287b-2ff8-46b4-bfb5-332a979a199a', 'State', 'Kanto'),
    ('44002c51-cc57-489f-bcf7-4f2abc6ddeb8', 'Street', 'Shibuya');

INSERT INTO system_accounts (system_account_id, created_at)
VALUES
('394e67b6-2850-4ddf-a4c9-c2a619d5bf70', '2024-01-10 12:00:00'),
('018d6189-9ad0-7b72-801b-1e0de0d3c214', '2024-01-10 12:00:00');

INSERT INTO system_account_profiles (system_account_id, name, email, created_at, updated_at)
VALUES
('394e67b6-2850-4ddf-a4c9-c2a619d5bf70', 'John Doe', 'system_account@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d6189-9ad0-7b72-801b-1e0de0d3c214', 'タロウ 認証済', 'test+1@example.com', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO system_account_phone_numbers (system_account_id, phone_number, created_at, updated_at)
VALUES
('394e67b6-2850-4ddf-a4c9-c2a619d5bf70', '09012345678', '2024-01-10 12:00:00', '2024-01-10 12:00:00'),
('018d6189-9ad0-7b72-801b-1e0de0d3c214', '09012345679', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_id, name, subdomain, created_at, updated_at)
VALUES
('c1bd2603-b9cd-4f84-8b83-3548f6ae150b', 'Example', 'example', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO members (member_id, workspace_id, system_account_id, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', 'c1bd2603-b9cd-4f84-8b83-3548f6ae150b', '394e67b6-2850-4ddf-a4c9-c2a619d5bf70', '2024-01-10 12:00:00');

INSERT INTO member_profiles (member_id, member_id_number, display_name, created_at, updated_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', 'DEV-12345', 'John Doe', '2024-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO member_addresses (member_id, postal_code, building_component_id, street_address_component_id, city_component_id, state_component_id, country_component_id, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '150-0002', null, '44002c51-cc57-489f-bcf7-4f2abc6ddeb8', '0f40229e-dc58-4111-b709-b9a5266f587f', '90c3287b-2ff8-46b4-bfb5-332a979a199a', '10269b87-98ce-490e-aeab-2a5230a48d4f', '2024-01-10 12:00:00');

INSERT INTO membership_periods (member_id, start_date, end_date, activity, created_at)
VALUES
('377eba35-5560-4f48-a99d-19cbd6a82b0d', '2024-01-01', NULL, 'Active', '2024-01-10 12:00:00');
