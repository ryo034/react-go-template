CREATE TABLE address_components (
  component_id uuid NOT NULL,
  component_type VARCHAR(50) NOT NULL,
  component_name VARCHAR(256) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (component_id)
);

CREATE TABLE system_accounts (
  system_account_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (system_account_id)
);

CREATE TABLE system_account_profiles (
  system_account_id uuid NOT NULL,
  name VARCHAR(255),
  email VARCHAR(256) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (system_account_id),
  CONSTRAINT fk_system_account_profiles_sas_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_accounts(system_account_id)
);

CREATE TABLE system_account_phone_numbers (
  system_account_id uuid NOT NULL,
  phone_number VARCHAR(15) NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (system_account_id),
  CONSTRAINT fk_system_account_phone_numbers_sas_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_accounts(system_account_id)
);

CREATE TABLE workspaces (
  workspace_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (workspace_id)
);

CREATE TABLE workspace_details (
  workspace_id uuid NOT NULL,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (workspace_id),
  CONSTRAINT fk_workspace_details_workspaces_workspace_id FOREIGN KEY (workspace_id) REFERENCES workspaces(workspace_id)
);

CREATE TABLE members (
  member_id uuid NOT NULL,
  system_account_id uuid NOT NULL,
  workspace_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (member_id),
  CONSTRAINT fk_members_system_accounts_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_accounts(system_account_id),
  CONSTRAINT fk_members_workspaces_workspace_id FOREIGN KEY (workspace_id) REFERENCES workspaces(workspace_id)
);

CREATE TABLE member_profiles (
  member_id uuid NOT NULL,
  member_id_number VARCHAR(255) NOT NULL,
  display_name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (member_id),
  CONSTRAINT fk_member_profiles_members_member_id FOREIGN KEY (member_id) REFERENCES members(member_id)
);

CREATE TABLE member_addresses (
  member_id uuid NOT NULL,
  postal_code VARCHAR(20),
  building_component_id uuid,
  street_address_component_id uuid NOT NULL,
  city_component_id uuid NOT NULL,
  state_component_id uuid NOT NULL,
  country_component_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (member_id, start_date),
  CONSTRAINT fk_membership_periods_members_member_id FOREIGN KEY (member_id) REFERENCES members(member_id)
);
