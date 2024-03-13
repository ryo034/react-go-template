INSERT INTO accounts (account_id, created_at)
VALUES
('018e15d1-0ba6-78d9-b255-63bb854e9817', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018e15d1-0ba6-7f70-8f44-bb7d4a6eb72c', '018e15d1-0ba6-78d9-b255-63bb854e9817', 'email', '', '018e15d1-0ba6-78d9-b255-63bb854e9817', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e15d1-0ba6-7809-9b78-6c797f821405', '018e15d1-0ba6-78d9-b255-63bb854e9817', 'long_bio@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e15d1-0ba6-7809-9b78-6c797f821405', '018e15d1-0ba6-78d9-b255-63bb854e9817');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e15d1-0ba6-7c34-8f02-460c24c95c23', '018e15d1-0ba6-78d9-b255-63bb854e9817', 'Long Bio', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e15d1-0ba6-7c34-8f02-460c24c95c23', '018e15d1-0ba6-78d9-b255-63bb854e9817');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('018e39d3-54bf-72b7-a89d-4952010c627b', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_detail_id, workspace_id, name, subdomain, created_at)
VALUES
('018e39d3-54bf-7fd4-bf39-019e496c07c5', '018e39d3-54bf-72b7-a89d-4952010c627b', 'LongBio Workspace', 'long-bio', '2024-01-10 12:00:00');

INSERT INTO workspace_latest_details (workspace_detail_id, workspace_id)
VALUES
('018e39d3-54bf-7359-909b-64a40d681ecd', '018e39d3-54bf-72b7-a89d-4952010c627b');

INSERT INTO members (member_id, account_id, workspace_id, created_at)
VALUES
('018e1398-3d80-79dc-9459-c7a3f1609124', '018e15d1-0ba6-78d9-b255-63bb854e9817', '018e39d3-54bf-72b7-a89d-4952010c627b', '2024-01-10 12:00:00');

INSERT INTO member_roles (member_role_id, member_id, assigned_by, role, assigned_at)
VALUES
('018e15d1-0ba6-7e42-bdd3-721da67bdd70', '018e1398-3d80-79dc-9459-c7a3f1609124', '018e1398-3d80-79dc-9459-c7a3f1609124', 'admin', '2024-01-10 12:00:00');

INSERT INTO member_latest_roles (member_role_id, member_id)
VALUES
('018e15d1-0ba6-7e42-bdd3-721da67bdd70', '018e1398-3d80-79dc-9459-c7a3f1609124');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018e1398-3d80-709e-9f07-1360425107ff', '018e1398-3d80-79dc-9459-c7a3f1609124', '2024-01-10 12:00:00');

INSERT INTO member_latest_login_histories (member_login_history_id, member_id)
VALUES
('018e1398-3d80-709e-9f07-1360425107ff', '018e1398-3d80-79dc-9459-c7a3f1609124');

INSERT INTO member_profiles (member_profile_id, member_id, member_id_number, display_name, bio, created_at)
VALUES
('018e2216-64a3-7fb7-93ab-d63cbfb93a16', '018e1398-3d80-79dc-9459-c7a3f1609124', 'DEV-12346', 'Login Bio', 'Long Bio is a passionate software engineer with 8 years of experience specializing in web development, particularly with React and Node.js. A graduate from MIT with a strong focus on clean architecture and Agile methodologies, John has successfully led multiple projects, from innovative startups to established tech giants. He''s a firm believer in continuous learning, contributing regularly to open-source projects, and sharing insights through tech blogs and meetups. Outside of work, John enjoys hiking 🚶‍♂️, drone photography 📸, and playing the guitar 🎸. He''s committed to using technology to drive positive social change.', '2024-01-10 12:00:00');

INSERT INTO member_latest_profiles (member_profile_id, member_id)
VALUES
('018e2216-64a3-7fb7-93ab-d63cbfb93a16', '018e1398-3d80-79dc-9459-c7a3f1609124');

INSERT INTO membership_events (membership_event_id, member_id, event_type, created_by, event_at)
VALUES
('018e2ff9-c432-7429-a92d-22a5b2d20dbc', '018e1398-3d80-79dc-9459-c7a3f1609124', 'join', '018e1398-3d80-79dc-9459-c7a3f1609124', '2024-01-10 12:00:00');

INSERT INTO latest_membership_events (membership_event_id, member_id)
VALUES
('018e2ff9-c432-7429-a92d-22a5b2d20dbc', '018e1398-3d80-79dc-9459-c7a3f1609124');
