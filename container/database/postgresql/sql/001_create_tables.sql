CREATE FUNCTION refresh_updated_at_step1() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at = OLD.updated_at THEN
    NEW.updated_at := NULL;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION refresh_updated_at_step2() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at IS NULL THEN
    NEW.updated_at := OLD.updated_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION refresh_updated_at_step3() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at IS NULL THEN
    NEW.updated_at := CURRENT_TIMESTAMP;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE address_components (
  component_id uuid NOT NULL,
  component_type VARCHAR(50) NOT NULL,
  component_name VARCHAR(256) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (component_id)
);

CREATE TABLE system_accounts (
  system_account_id uuid NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (system_account_id)
);

CREATE TYPE auth_provider_provider AS ENUM ('google', 'email');
CREATE TYPE auth_provider_provided_by AS ENUM ('firebase');

CREATE TABLE auth_providers (
  auth_provider_id uuid NOT NULL,
  system_account_id uuid NOT NULL,
  provider auth_provider_provider NOT NULL,
  provider_uid VARCHAR(255) NOT NULL,
  provided_by auth_provider_provided_by NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (auth_provider_id),
  CONSTRAINT fk_auth_providers_system_accounts_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_accounts(system_account_id)
);

CREATE TABLE system_account_emails (
  system_account_id uuid NOT NULL,
  email VARCHAR(256) NOT NULL UNIQUE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (system_account_id),
  CONSTRAINT fk_system_account_emails_system_accounts_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_accounts(system_account_id)
);
CREATE INDEX system_account_emails_created_at_index ON system_account_emails(created_at);

CREATE TYPE system_account_phone_number_country_codes AS ENUM ('JP', 'US');

CREATE TABLE system_account_phone_numbers (
  system_account_id uuid NOT NULL,
  phone_number VARCHAR(15) NOT NULL UNIQUE,
  country_code system_account_phone_number_country_codes NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (system_account_id),
  CONSTRAINT fk_system_account_phone_numbers_system_accounts_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_accounts(system_account_id)
);
CREATE INDEX system_account_phone_numbers_created_at_index ON system_account_phone_numbers(created_at);

CREATE TABLE system_account_profiles (
  system_account_id uuid NOT NULL,
  name VARCHAR(255),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (system_account_id),
  CONSTRAINT fk_system_account_profiles_sas_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_accounts(system_account_id)
);

CREATE TRIGGER refresh_system_account_profiles_updated_at_step1 BEFORE UPDATE ON system_account_profiles FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_system_account_profiles_updated_at_step2 BEFORE UPDATE OF updated_at ON system_account_profiles FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_system_account_profiles_updated_at_step3 BEFORE UPDATE ON system_account_profiles FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE TABLE workspaces (
  workspace_id uuid NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (workspace_id)
);

CREATE TABLE workspace_details (
  workspace_id uuid NOT NULL,
  name VARCHAR(100) NOT NULL,
  subdomain VARCHAR(63) NOT NULL UNIQUE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (workspace_id),
  CONSTRAINT fk_workspace_details_workspaces_workspace_id FOREIGN KEY (workspace_id) REFERENCES workspaces(workspace_id)
);

CREATE TRIGGER refresh_workspace_details_updated_at_step1 BEFORE UPDATE ON workspace_details FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_workspace_details_updated_at_step2 BEFORE UPDATE OF updated_at ON workspace_details FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_workspace_details_updated_at_step3 BEFORE UPDATE ON workspace_details FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE TABLE members (
  member_id uuid NOT NULL,
  system_account_id uuid NOT NULL,
  workspace_id uuid NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (member_id),
  CONSTRAINT fk_members_system_accounts_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_accounts(system_account_id),
  CONSTRAINT fk_members_workspaces_workspace_id FOREIGN KEY (workspace_id) REFERENCES workspaces(workspace_id)
);

CREATE TYPE member_role_type AS ENUM ('owner', 'admin', 'member', 'guest');

CREATE TABLE member_roles (
  member_role_id uuid NOT NULL,
  member_id uuid NOT NULL,
  role member_role_type NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (member_role_id),
  CONSTRAINT fk_member_roles_members_member_id FOREIGN KEY (member_id) REFERENCES members(member_id)
);
CREATE INDEX member_roles_created_at_index ON member_roles(created_at);

CREATE TABLE member_login_histories (
  member_login_history_id uuid NOT NULL,
  member_id uuid NOT NULL,
  login_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (member_login_history_id),
  CONSTRAINT fk_member_login_histories_members_member_id FOREIGN KEY (member_id) REFERENCES members(member_id)
);

CREATE TABLE member_profiles (
  member_id uuid NOT NULL,
  member_id_number VARCHAR(255) NOT NULL,
  display_name VARCHAR(50) NOT NULL,
  bio TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (member_id),
  CONSTRAINT fk_member_profiles_members_member_id FOREIGN KEY (member_id) REFERENCES members(member_id)
);

CREATE TRIGGER refresh_member_profiles_updated_at_step1 BEFORE UPDATE ON member_profiles FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_member_profiles_updated_at_step2 BEFORE UPDATE OF updated_at ON member_profiles FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_member_profiles_updated_at_step3 BEFORE UPDATE ON member_profiles FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE TABLE member_addresses (
  member_id uuid NOT NULL,
  postal_code VARCHAR(20),
  building_component_id uuid,
  street_address_component_id uuid NOT NULL,
  city_component_id uuid NOT NULL,
  state_component_id uuid NOT NULL,
  country_component_id uuid NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (member_id),
  CONSTRAINT fk_member_addresses_members_member_id FOREIGN KEY (member_id) REFERENCES members(member_id),
  CONSTRAINT fk_member_addresses_building_component_id FOREIGN KEY (building_component_id) REFERENCES address_components(component_id),
  CONSTRAINT fk_member_addresses_street_address_component_id FOREIGN KEY (street_address_component_id) REFERENCES address_components(component_id),
  CONSTRAINT fk_member_addresses_city_component_id FOREIGN KEY (city_component_id) REFERENCES address_components(component_id),
  CONSTRAINT fk_member_addresses_state_component_id FOREIGN KEY (state_component_id) REFERENCES address_components(component_id),
  CONSTRAINT fk_member_addresses_country_component_id FOREIGN KEY (country_component_id) REFERENCES address_components(component_id)
);

CREATE TABLE membership_periods (
  member_id uuid NOT NULL,
  start_date DATE NOT NULL,
  end_date DATE,
  activity VARCHAR(20) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (member_id, start_date),
  CONSTRAINT fk_membership_periods_members_member_id FOREIGN KEY (member_id) REFERENCES members(member_id)
);

CREATE TRIGGER refresh_membership_periods_updated_at_step1 BEFORE UPDATE ON membership_periods FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_membership_periods_updated_at_step2 BEFORE UPDATE OF updated_at ON membership_periods FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_membership_periods_updated_at_step3 BEFORE UPDATE ON membership_periods FOR EACH ROW EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE TABLE invitation_units (
  invitation_unit_id uuid NOT NULL,
  workspace_id uuid NOT NULL,
  invited_by uuid NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (invitation_unit_id),
  CONSTRAINT fk_invitation_units_workspaces_workspace_id FOREIGN KEY (workspace_id) REFERENCES workspaces(workspace_id),
  CONSTRAINT fk_invitation_units_members_invited_by FOREIGN KEY (invited_by) REFERENCES members(member_id)
);

CREATE TABLE invitations (
  invitation_id uuid NOT NULL,
  invitation_unit_id uuid NOT NULL,
  PRIMARY KEY (invitation_id),
  CONSTRAINT fk_invitations_invitation_units_invitation_unit_id FOREIGN KEY (invitation_unit_id) REFERENCES invitation_units(invitation_unit_id)
);

CREATE TABLE invitation_tokens (
  invitation_id uuid NOT NULL,
  token uuid NOT NULL UNIQUE,
  expired_at TIMESTAMP WITH TIME ZONE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (invitation_id, token),
  CONSTRAINT fk_invitations_invitation_tokens_invitation_id FOREIGN KEY (invitation_id) REFERENCES invitations(invitation_id)
);
CREATE INDEX invitation_tokens_expired_at_index ON invitation_tokens(expired_at);

CREATE TABLE invitees (
  invitation_id uuid NOT NULL,
  email VARCHAR(256) NOT NULL,
  PRIMARY KEY (invitation_id),
  CONSTRAINT fk_invitees_invitations_invitation_id FOREIGN KEY (invitation_id) REFERENCES invitations(invitation_id)
);

CREATE TABLE invitee_names (
  invitation_id uuid NOT NULL,
  display_name VARCHAR(50),
  PRIMARY KEY (invitation_id),
  CONSTRAINT fk_invitee_names_invitations_invitation_id FOREIGN KEY (invitation_id) REFERENCES invitations(invitation_id)
);

CREATE TYPE invitation_event_types AS ENUM ('verified', 'accepted', 'revoked', 'reissued');

CREATE TABLE invitation_events (
  invitation_event_id uuid NOT NULL,
  invitation_id uuid NOT NULL,
  event_type invitation_event_types NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (invitation_event_id),
  CONSTRAINT fk_invitation_events_invitations_invitation_id FOREIGN KEY (invitation_id) REFERENCES invitations(invitation_id)
);
CREATE INDEX invitation_events_created_at_index ON invitation_events(created_at);