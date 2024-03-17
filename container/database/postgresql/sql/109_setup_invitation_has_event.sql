

INSERT INTO accounts (account_id, created_at)
VALUES
('018e4922-563a-7097-bbdb-ffa9f74da283', '2024-01-10 12:00:00');

INSERT INTO auth_providers (auth_provider_id, account_id, provider, photo_url, provider_uid, provided_by, registered_at)
VALUES
('018e4922-563a-71b5-bdb4-ba06511f2590', '018e4922-563a-7097-bbdb-ffa9f74da283', 'email', '', '018e4922-563a-7097-bbdb-ffa9f74da283', 'firebase', '2024-01-10 12:00:00');

INSERT INTO account_emails (account_email_id, account_id, email, created_at)
VALUES
('018e4922-563a-7f98-b31e-e42efc811159', '018e4922-563a-7097-bbdb-ffa9f74da283', 'invite_test_has_event_inviter@example.com', '2024-01-10 12:00:00');

INSERT INTO account_latest_emails (account_email_id, account_id)
VALUES
('018e4922-563a-7f98-b31e-e42efc811159', '018e4922-563a-7097-bbdb-ffa9f74da283');

INSERT INTO account_names (account_name_id, account_id, name, created_at)
VALUES
('018e4922-563a-7cff-bc38-66e095586aa0', '018e4922-563a-7097-bbdb-ffa9f74da283', 'Invite TestHasEvent', '2024-01-10 12:00:00');

INSERT INTO account_latest_names (account_name_id, account_id)
VALUES
('018e4922-563a-7cff-bc38-66e095586aa0', '018e4922-563a-7097-bbdb-ffa9f74da283');

INSERT INTO workspaces (workspace_id, created_at)
VALUES
('018e4922-563a-7731-b389-c2a9ac0d97e9', '2024-01-10 12:00:00');

INSERT INTO workspace_details (workspace_detail_id, workspace_id, name, subdomain, created_at)
VALUES
('018e4922-563a-736e-b632-b32b7df08c67', '018e4922-563a-7731-b389-c2a9ac0d97e9', 'Invite TestHasEvent', 'invite-test-has-event', '2024-01-10 12:00:00');

INSERT INTO workspace_latest_details (workspace_detail_id, workspace_id)
VALUES
('018e4922-563a-736e-b632-b32b7df08c67', '018e4922-563a-7731-b389-c2a9ac0d97e9');

INSERT INTO members (member_id, account_id, workspace_id, created_at)
VALUES
('018e4922-563a-7807-b01f-2e630e4d22e9', '018e4922-563a-7097-bbdb-ffa9f74da283', '018e4922-563a-7731-b389-c2a9ac0d97e9', '2024-01-10 12:00:00');

INSERT INTO member_roles (member_role_id, member_id, assigned_by, role, assigned_at)
VALUES
('018e4922-563a-7abc-a66a-24ea32938e9c', '018e4922-563a-7807-b01f-2e630e4d22e9', '018e4922-563a-7807-b01f-2e630e4d22e9', 'owner', '2024-01-10 12:00:00');

INSERT INTO member_latest_roles (member_role_id, member_id)
VALUES
('018e4922-563a-7abc-a66a-24ea32938e9c', '018e4922-563a-7807-b01f-2e630e4d22e9');

INSERT INTO member_login_histories (member_login_history_id, member_id, login_at)
VALUES
('018e4922-563a-7b96-b6b2-4e4066ba1d50', '018e4922-563a-7807-b01f-2e630e4d22e9', '2024-01-10 12:00:00');

INSERT INTO member_latest_login_histories (member_login_history_id, member_id)
VALUES
('018e4922-563a-7b96-b6b2-4e4066ba1d50', '018e4922-563a-7807-b01f-2e630e4d22e9');

INSERT INTO member_profiles (member_profile_id, member_id, member_id_number, display_name, bio, created_at)
VALUES
('018e4922-563a-71f9-a7a1-25455276884f', '018e4922-563a-7807-b01f-2e630e4d22e9', 'DEV-12345', 'Invite TestHasEvent', 'bio', '2024-01-10 12:00:00');

INSERT INTO member_latest_profiles (member_profile_id, member_id)
VALUES
('018e4922-563a-71f9-a7a1-25455276884f', '018e4922-563a-7807-b01f-2e630e4d22e9');

INSERT INTO membership_events (membership_event_id, member_id, event_type, created_by, event_at)
VALUES
('018e4922-563a-7b08-8fb3-955ae030f5d1', '018e4922-563a-7807-b01f-2e630e4d22e9', 'join', '018e4922-563a-7807-b01f-2e630e4d22e9', '2024-01-10 12:00:00');

INSERT INTO latest_membership_events (membership_event_id, member_id)
VALUES
('018e4922-563a-7b08-8fb3-955ae030f5d1', '018e4922-563a-7807-b01f-2e630e4d22e9');

INSERT INTO invitation_units (invitation_unit_id, workspace_id, invited_by, created_at)
VALUES
('018e4922-563a-7c76-8ea0-815441c038cb', '018e4922-563a-7731-b389-c2a9ac0d97e9', '018e4922-563a-7807-b01f-2e630e4d22e9', '2024-01-10 12:00:00');

INSERT INTO invitations (invitation_id, invitation_unit_id)
VALUES
('018e4922-563a-7566-bc5e-65dc2f8faefe', '018e4922-563a-7c76-8ea0-815441c038cb'),
('018e4922-563a-7760-a695-d1e985fc1495', '018e4922-563a-7c76-8ea0-815441c038cb'),
('018e4922-563a-7b0c-bd1f-8f31d9d2cd27', '018e4922-563a-7c76-8ea0-815441c038cb'),
('018e4922-563a-72b9-aa38-60ddd4e4d4e4', '018e4922-563a-7c76-8ea0-815441c038cb');

INSERT INTO invitation_tokens (invitation_token_id, invitation_id, token, expired_at, created_at)
VALUES
('018e493a-1b7f-79f3-83f9-c2ee307ce23d', '018e4922-563a-7566-bc5e-65dc2f8faefe', '018e4922-563a-735c-a715-2fe940d327cf', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018e493a-1b7f-7b41-af59-aa9d53131ff5', '018e4922-563a-7760-a695-d1e985fc1495', '018e4922-563a-762b-ac52-6308978dbf70', '2200-01-10 12:00:00', '2024-01-10 12:00:00'),
('018e493a-1b7f-7857-9b13-0327f326e3bc', '018e4922-563a-7b0c-bd1f-8f31d9d2cd27', '018e4922-563a-737f-bd28-a7454b757e6e', '2024-01-17 12:00:00', '2024-01-10 12:00:00'),
('018e493a-1b7f-72bb-bb4c-d7b3f1796aba', '018e4922-563a-72b9-aa38-60ddd4e4d4e4', '018e4922-563a-75f7-9153-bd733f331541', '2200-01-10 12:00:00', '2024-01-10 12:00:00');

INSERT INTO latest_invitation_tokens (invitation_token_id, invitation_id)
VALUES
('018e493a-1b7f-79f3-83f9-c2ee307ce23d', '018e4922-563a-7566-bc5e-65dc2f8faefe'),
('018e493a-1b7f-7b41-af59-aa9d53131ff5', '018e4922-563a-7760-a695-d1e985fc1495'),
('018e493a-1b7f-7857-9b13-0327f326e3bc', '018e4922-563a-7b0c-bd1f-8f31d9d2cd27'),
('018e493a-1b7f-72bb-bb4c-d7b3f1796aba', '018e4922-563a-72b9-aa38-60ddd4e4d4e4');

INSERT INTO invitees (invitation_id, email)
VALUES
('018e4922-563a-7566-bc5e-65dc2f8faefe', 'invite_test_not_expired@example.com'),
('018e4922-563a-7760-a695-d1e985fc1495', 'invite_test_already_verified@example.com'),
('018e4922-563a-7b0c-bd1f-8f31d9d2cd27', 'invite_test_already_accepted@example.com'),
('018e4922-563a-72b9-aa38-60ddd4e4d4e4', 'invite_test_revoked@example.com');

INSERT INTO invitation_events (invitation_event_id, invitation_id, event_type, created_at)
VALUES
('018e4922-563a-7802-8d90-3ee885ece6fc', '018e4922-563a-7760-a695-d1e985fc1495', 'verified', '2023-01-10 15:00:00'),
('018e4922-563a-7317-a442-6b0d108f09d0', '018e4922-563a-7b0c-bd1f-8f31d9d2cd27', 'verified', '2023-01-10 15:00:00'),
('018e4922-563a-7886-afaf-db551d83f32e', '018e4922-563a-7b0c-bd1f-8f31d9d2cd27', 'accepted', '2023-01-11 15:00:00'),
('018e4922-563a-7866-9098-6a66fc3b1321', '018e4922-563a-72b9-aa38-60ddd4e4d4e4', 'revoked', '2024-01-10 15:00:00');

INSERT INTO latest_invitation_events (invitation_event_id, invitation_id)
VALUES
('018e4922-563a-7802-8d90-3ee885ece6fc', '018e4922-563a-7760-a695-d1e985fc1495'),
('018e4922-563a-7886-afaf-db551d83f32e', '018e4922-563a-7b0c-bd1f-8f31d9d2cd27'),
('018e4922-563a-7866-9098-6a66fc3b1321', '018e4922-563a-72b9-aa38-60ddd4e4d4e4');
