INSERT INTO accounts (account_id, created_at)
VALUES
('018df2ef-43d0-7ba3-9159-13b2b6634042', '2024-01-10 12:00:00'),
('018df2ef-d77a-784a-92d3-3f52deb284bd', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018df2f3-5922-789d-b529-9b98ab707514', '018df2ef-43d0-7ba3-9159-13b2b6634042', 'google', '', 'MuJcEqPqy9r3wJ85GWsV3SszVJ6X', 'firebase', '2024-01-10 12:00:00'),
('018df2f4-d77c-7b11-9e98-a6d03d70a27a', '018df2ef-d77a-784a-92d3-3f52deb284bd', 'google', '', 'Xk1n15UQOFbml4RoF0QdCza5n0dU', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e09c2-d3f3-7635-9973-b88c73a1d73a', '018df2ef-43d0-7ba3-9159-13b2b6634042', 'google_auth_test_no_name@example.com', '2024-01-10 12:00:00'),
('018e09c2-d3f3-7409-ac0b-00bbc07acf7f', '018df2ef-d77a-784a-92d3-3f52deb284bd', 'google_auth_test_has_name@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e09c2-d3f3-7635-9973-b88c73a1d73a', '018df2ef-43d0-7ba3-9159-13b2b6634042'),
('018e09c2-d3f3-7409-ac0b-00bbc07acf7f', '018df2ef-d77a-784a-92d3-3f52deb284bd');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e0890-2197-7ab2-a17c-b0e168cd6080', '018df2ef-d77a-784a-92d3-3f52deb284bd', 'GoogleAuthTest HasNameTest', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e0890-2197-7ab2-a17c-b0e168cd6080', '018df2ef-d77a-784a-92d3-3f52deb284bd');
