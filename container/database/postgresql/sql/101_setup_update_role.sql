INSERT INTO accounts (account_id, created_at)
VALUES
('018e18ba-dc87-705a-9ba9-2db0f8ead09f', '2024-01-10 12:00:00'),
('018e18ba-dc87-7e16-83a6-e8ffccf96552', '2024-01-10 12:00:00'),
('018e1952-009b-7a7b-b1a5-3938a11784f9', '2024-01-10 12:00:00'),
('018e18ba-dc87-7758-929e-cf1d52320f0c', '2024-01-10 12:00:00'),
('018e18ba-dc87-739e-9206-47a7b99de453', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018e18ba-dc87-7aa4-9d14-619cfed2a967', '018e18ba-dc87-705a-9ba9-2db0f8ead09f', 'email', '', '018e18ba-dc87-705a-9ba9-2db0f8ead09f', 'firebase', '2024-01-10 12:00:00'),
('018e18ba-dc87-7b40-9ea7-e956cfe732e0', '018e18ba-dc87-7e16-83a6-e8ffccf96552', 'email', '', '018e18ba-dc87-7e16-83a6-e8ffccf96552', 'firebase', '2024-01-10 12:00:00'),
('018e1952-009b-72ed-95b8-48b869c2b1d3', '018e1952-009b-7a7b-b1a5-3938a11784f9', 'email', '', '018e1952-009b-7a7b-b1a5-3938a11784f9', 'firebase', '2024-01-10 12:00:00'),
('018e18ba-dc87-7195-8b7c-3e1f4cd27f13', '018e18ba-dc87-7758-929e-cf1d52320f0c', 'email', '', '018e18ba-dc87-7758-929e-cf1d52320f0c', 'firebase', '2024-01-10 12:00:00'),
('018e18ba-dc87-7927-be6e-eb5bb6998b0c', '018e18ba-dc87-739e-9206-47a7b99de453', 'email', '', '018e18ba-dc87-739e-9206-47a7b99de453', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e18ba-dc87-78cc-bfe1-ea2fffa4b6da', '018e18ba-dc87-705a-9ba9-2db0f8ead09f', 'update_role_owner@example.com', '2024-01-10 12:00:00'),
('018e18ba-dc87-71e0-abda-948eb79352ae', '018e18ba-dc87-7e16-83a6-e8ffccf96552', 'update_role_admin@example.com', '2024-01-10 12:00:00'),
('018e1952-009b-769d-a948-9677102ef6b8', '018e1952-009b-7a7b-b1a5-3938a11784f9', 'update_role_admin_2@example.com', '2024-01-10 12:00:00'),
('018e18ba-dc87-7d52-a12f-95c843831cc6', '018e18ba-dc87-7758-929e-cf1d52320f0c', 'update_role_member@example.com', '2024-01-10 12:00:00'),
('018e18ba-dc87-7868-82b7-a44909ce5d5a', '018e18ba-dc87-739e-9206-47a7b99de453', 'update_role_guest@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e18ba-dc87-78cc-bfe1-ea2fffa4b6da', '018e18ba-dc87-705a-9ba9-2db0f8ead09f'),
('018e18ba-dc87-71e0-abda-948eb79352ae', '018e18ba-dc87-7e16-83a6-e8ffccf96552'),
('018e1952-009b-769d-a948-9677102ef6b8', '018e1952-009b-7a7b-b1a5-3938a11784f9'),
('018e18ba-dc87-7d52-a12f-95c843831cc6', '018e18ba-dc87-7758-929e-cf1d52320f0c'),
('018e18ba-dc87-7868-82b7-a44909ce5d5a', '018e18ba-dc87-739e-9206-47a7b99de453');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e18ba-dc87-7a57-a5e2-a0db4a26e2de', '018e18ba-dc87-705a-9ba9-2db0f8ead09f', 'UpdateRole Owner', '2024-01-10 12:00:00'),
('018e18ba-dc87-77ac-bd1e-c922384dd7e8', '018e18ba-dc87-7e16-83a6-e8ffccf96552', 'UpdateRole Admin', '2024-01-10 12:00:00'),
('018e1952-009b-7383-8ae8-d5d32bb9ddcb', '018e1952-009b-7a7b-b1a5-3938a11784f9', 'UpdateRole AdminTwo', '2024-01-10 12:00:00'),
('018e18ba-dc87-77a6-bbd6-2c84421cad82', '018e18ba-dc87-7758-929e-cf1d52320f0c', 'UpdateRole Member', '2024-01-10 12:00:00'),
('018e18ba-dc87-7425-a2ed-bd1c7f551f4e', '018e18ba-dc87-739e-9206-47a7b99de453', 'UpdateRole Guest', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e18ba-dc87-7a57-a5e2-a0db4a26e2de', '018e18ba-dc87-705a-9ba9-2db0f8ead09f'),
('018e18ba-dc87-77ac-bd1e-c922384dd7e8', '018e18ba-dc87-7e16-83a6-e8ffccf96552'),
('018e1952-009b-7383-8ae8-d5d32bb9ddcb', '018e1952-009b-7a7b-b1a5-3938a11784f9'),
('018e18ba-dc87-77a6-bbd6-2c84421cad82', '018e18ba-dc87-7758-929e-cf1d52320f0c'),
('018e18ba-dc87-7425-a2ed-bd1c7f551f4e', '018e18ba-dc87-739e-9206-47a7b99de453');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('018e18ba-dc87-7658-9591-672daaddb95b', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_detail_id, workspace_id, name, subdomain, created_at)
VALUES
('018e200b-9d01-7220-88b5-1ed32c7b4808', '018e18ba-dc87-7658-9591-672daaddb95b', 'UpdateRole Workspace', 'update-role', '2024-01-10 12:00:00');

INSERT INTO workspace_latest_details (workspace_detail_id, workspace_id)
VALUES
('018e200b-9d01-7220-88b5-1ed32c7b4808', '018e18ba-dc87-7658-9591-672daaddb95b');

INSERT INTO members (member_id, account_id, workspace_id, created_at)
VALUES
('018e18ba-dc87-72e2-bb4b-c43252f51492', '018e18ba-dc87-705a-9ba9-2db0f8ead09f', '018e18ba-dc87-7658-9591-672daaddb95b', '2024-01-10 12:00:00'),
('018e18ba-dc87-740c-9aeb-ba7f8f7d490e', '018e18ba-dc87-7e16-83a6-e8ffccf96552', '018e18ba-dc87-7658-9591-672daaddb95b', '2024-01-10 12:00:00'),
('018e1952-009b-7138-aea6-24b2f9596ad7', '018e1952-009b-7a7b-b1a5-3938a11784f9', '018e18ba-dc87-7658-9591-672daaddb95b', '2024-01-10 12:00:00'),
('018e18ba-dc87-7a3e-8181-7186458e84b6', '018e18ba-dc87-7758-929e-cf1d52320f0c', '018e18ba-dc87-7658-9591-672daaddb95b', '2024-01-10 12:00:00'),
('018e18ba-dc87-7c1a-81c6-6f6415c53966', '018e18ba-dc87-739e-9206-47a7b99de453', '018e18ba-dc87-7658-9591-672daaddb95b', '2024-01-10 12:00:00');

INSERT INTO member_roles (member_role_id, member_id, assigned_by, role, assigned_at)
VALUES
('018e18ba-dc87-7fcc-95f1-6d60e426a98a', '018e18ba-dc87-72e2-bb4b-c43252f51492', '018e18ba-dc87-72e2-bb4b-c43252f51492', 'owner', '2024-01-10 12:00:00'),
('018e18ba-dc87-79ff-9fa9-b898d3aac9ee', '018e18ba-dc87-740c-9aeb-ba7f8f7d490e', '018e18ba-dc87-740c-9aeb-ba7f8f7d490e', 'admin', '2024-01-10 12:00:00'),
('018e1952-009b-71ea-9dbf-9a34b8974ac3', '018e1952-009b-7138-aea6-24b2f9596ad7', '018e1952-009b-7138-aea6-24b2f9596ad7', 'admin', '2024-01-10 12:00:00'),
('018e18ba-dc87-7515-bf20-937fcdcee67b', '018e18ba-dc87-7a3e-8181-7186458e84b6', '018e18ba-dc87-7a3e-8181-7186458e84b6', 'member', '2024-01-10 12:00:00'),
('018e18ba-dc87-7c3b-a67b-1ecf60ff8130', '018e18ba-dc87-7c1a-81c6-6f6415c53966', '018e18ba-dc87-7c1a-81c6-6f6415c53966', 'guest', '2024-01-10 12:00:00');

INSERT INTO member_latest_roles (member_role_id, member_id)
VALUES
('018e18ba-dc87-7fcc-95f1-6d60e426a98a', '018e18ba-dc87-72e2-bb4b-c43252f51492'),
('018e18ba-dc87-79ff-9fa9-b898d3aac9ee', '018e18ba-dc87-740c-9aeb-ba7f8f7d490e'),
('018e1952-009b-71ea-9dbf-9a34b8974ac3', '018e1952-009b-7138-aea6-24b2f9596ad7'),
('018e18ba-dc87-7515-bf20-937fcdcee67b', '018e18ba-dc87-7a3e-8181-7186458e84b6'),
('018e18ba-dc87-7c3b-a67b-1ecf60ff8130', '018e18ba-dc87-7c1a-81c6-6f6415c53966');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018e18ba-dc87-7780-b928-cf7712970dad', '018e18ba-dc87-72e2-bb4b-c43252f51492', '2024-01-10 12:00:00'),
('018e18ba-dc87-769c-9448-4e06d6fa9d63', '018e18ba-dc87-740c-9aeb-ba7f8f7d490e', '2024-01-10 12:00:00'),
('018e1952-009b-7457-9ba7-1b002fbb0415', '018e1952-009b-7138-aea6-24b2f9596ad7', '2024-01-10 12:00:00'),
('018e18ba-dc87-7897-8585-808a56e2c48e', '018e18ba-dc87-7a3e-8181-7186458e84b6', '2024-01-10 12:00:00'),
('018e18ba-dc87-7be6-9dba-05be861174a2', '018e18ba-dc87-7c1a-81c6-6f6415c53966', '2024-01-10 12:00:00');

INSERT INTO member_latest_login_histories (member_login_history_id, member_id)
VALUES
('018e18ba-dc87-7780-b928-cf7712970dad', '018e18ba-dc87-72e2-bb4b-c43252f51492'),
('018e18ba-dc87-769c-9448-4e06d6fa9d63', '018e18ba-dc87-740c-9aeb-ba7f8f7d490e'),
('018e1952-009b-7457-9ba7-1b002fbb0415', '018e1952-009b-7138-aea6-24b2f9596ad7'),
('018e18ba-dc87-7897-8585-808a56e2c48e', '018e18ba-dc87-7a3e-8181-7186458e84b6'),
('018e18ba-dc87-7be6-9dba-05be861174a2', '018e18ba-dc87-7c1a-81c6-6f6415c53966');

INSERT INTO member_profiles (member_profile_id, member_id, member_id_number, display_name, bio, created_at)
VALUES
('018e2216-64a3-7536-85ee-1aef8764ab43', '018e18ba-dc87-72e2-bb4b-c43252f51492', 'DEV-54321', 'UpdateRole Owner', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-7155-9422-57888faecafd', '018e18ba-dc87-740c-9aeb-ba7f8f7d490e', 'DEV-54322', 'UpdateRole Admin', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-7ae9-b464-4e1d78103c23', '018e1952-009b-7138-aea6-24b2f9596ad7', 'DEV-54323', 'UpdateRole AdminTwo', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-7e49-9c65-577f12495115', '018e18ba-dc87-7a3e-8181-7186458e84b6', 'DEV-54323', 'UpdateRole Member', 'bio', '2024-01-10 12:00:00'),
('018e2216-64a3-7d09-8649-e13bd8261c89', '018e18ba-dc87-7c1a-81c6-6f6415c53966', 'DEV-54324', 'UpdateRole Guest', 'bio', '2024-01-10 12:00:00');

INSERT INTO member_latest_profiles (member_profile_id, member_id)
VALUES
('018e2216-64a3-7536-85ee-1aef8764ab43', '018e18ba-dc87-72e2-bb4b-c43252f51492'),
('018e2216-64a3-7155-9422-57888faecafd', '018e18ba-dc87-740c-9aeb-ba7f8f7d490e'),
('018e2216-64a3-7ae9-b464-4e1d78103c23', '018e1952-009b-7138-aea6-24b2f9596ad7'),
('018e2216-64a3-7e49-9c65-577f12495115', '018e18ba-dc87-7a3e-8181-7186458e84b6'),
('018e2216-64a3-7d09-8649-e13bd8261c89', '018e18ba-dc87-7c1a-81c6-6f6415c53966');

INSERT INTO membership_events (membership_event_id, member_id, event_type, created_by, event_at)
VALUES
('018e2ff9-c432-7d02-8dbb-eaac78fb0cf7', '018e18ba-dc87-72e2-bb4b-c43252f51492', 'join', '018e18ba-dc87-72e2-bb4b-c43252f51492', '2024-01-10 12:00:00'),
('018e2ff9-c432-7c67-aea9-4174c64d0e9f', '018e18ba-dc87-740c-9aeb-ba7f8f7d490e', 'join', '018e18ba-dc87-740c-9aeb-ba7f8f7d490e', '2024-01-10 12:00:00'),
('018e2ff9-c432-7687-8b49-eeccc59ea938', '018e1952-009b-7138-aea6-24b2f9596ad7', 'join', '018e1952-009b-7138-aea6-24b2f9596ad7', '2024-01-10 12:00:00'),
('018e2ff9-c432-7306-b6ce-45d6a5a3d73a', '018e18ba-dc87-7a3e-8181-7186458e84b6', 'join', '018e18ba-dc87-7a3e-8181-7186458e84b6', '2024-01-10 12:00:00'),
('018e2ff9-c432-7ca4-9d2a-923b0428211b', '018e18ba-dc87-7c1a-81c6-6f6415c53966', 'join', '018e18ba-dc87-7c1a-81c6-6f6415c53966', '2024-01-10 12:00:00');

INSERT INTO latest_membership_events (membership_event_id, member_id)
VALUES
('018e2ff9-c432-7d02-8dbb-eaac78fb0cf7', '018e18ba-dc87-72e2-bb4b-c43252f51492'),
('018e2ff9-c432-7c67-aea9-4174c64d0e9f', '018e18ba-dc87-740c-9aeb-ba7f8f7d490e'),
('018e2ff9-c432-7687-8b49-eeccc59ea938', '018e1952-009b-7138-aea6-24b2f9596ad7'),
('018e2ff9-c432-7306-b6ce-45d6a5a3d73a', '018e18ba-dc87-7a3e-8181-7186458e84b6'),
('018e2ff9-c432-7ca4-9d2a-923b0428211b', '018e18ba-dc87-7c1a-81c6-6f6415c53966');
