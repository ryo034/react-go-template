INSERT INTO accounts (account_id, created_at)
VALUES
('018d6189-9ad0-7b72-801b-1e0de0d3c214', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018de2ff-7d69-7f8d-9d19-57bb4106f594', '018d6189-9ad0-7b72-801b-1e0de0d3c214', 'email', '', '018d6189-9ad0-7b72-801b-1e0de0d3c214', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e09c2-d3f3-75d3-9b72-91a4c2b2d8aa', '018d6189-9ad0-7b72-801b-1e0de0d3c214', 'unfinished_onboarding@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e09c2-d3f3-75d3-9b72-91a4c2b2d8aa', '018d6189-9ad0-7b72-801b-1e0de0d3c214');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e088f-af9c-7524-b0bd-f3d4dfb24f26', '018d6189-9ad0-7b72-801b-1e0de0d3c214', 'Unfinished Onboarding', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e088f-af9c-7524-b0bd-f3d4dfb24f26', '018d6189-9ad0-7b72-801b-1e0de0d3c214');
