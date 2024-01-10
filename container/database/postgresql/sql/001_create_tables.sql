CREATE TABLE address_components (
  component_id uuid NOT NULL,
  component_type VARCHAR(50) NOT NULL,
  component_name VARCHAR(256) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (component_id)
);
COMMENT ON TABLE address_components IS '住所の各構成要素';

CREATE TABLE system_accounts (
  system_account_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (system_account_id)
);
COMMENT ON TABLE system_accounts IS 'システム利用者';

CREATE TABLE system_account_profiles (
  system_account_id uuid NOT NULL,
  email VARCHAR(256) NOT NULL,
  email_verified BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (system_account_id),
  CONSTRAINT fk_system_account_profiles_system_accounts_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_accounts(system_account_id)
);
COMMENT ON TABLE system_account_profiles IS 'システム利用者のプロフィール';

CREATE TABLE system_account_phone_numbers (
  system_account_id uuid NOT NULL,
  phone_number VARCHAR(15) NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (system_account_id),
  CONSTRAINT fk_system_account_phone_numbers_system_accounts_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_accounts(system_account_id)
);
COMMENT ON TABLE system_account_phone_numbers IS 'システム利用者の電話番号';

CREATE TABLE organizations (
  organization_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (organization_id)
);
COMMENT ON TABLE organizations IS '組織';

CREATE TABLE organization_details (
  organization_id uuid NOT NULL,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (organization_id),
  CONSTRAINT fk_organization_details_organizations_organization_id FOREIGN KEY (organization_id) REFERENCES organizations(organization_id)
);
COMMENT ON TABLE organization_details IS '組織の基本情報';

CREATE TABLE employees (
  employee_id uuid NOT NULL,
  system_account_id uuid NOT NULL,
  organization_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (employee_id),
  CONSTRAINT fk_employees_system_accounts_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_accounts(system_account_id),
  CONSTRAINT fk_employees_organizations_organization_id FOREIGN KEY (organization_id) REFERENCES organizations(organization_id)
);
COMMENT ON TABLE employees IS '従業員';

CREATE TABLE employee_profiles (
  employee_id uuid NOT NULL,
  employee_id_number VARCHAR(255) NOT NULL,
  name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (employee_id),
  CONSTRAINT fk_employee_profiles_employees_employee_id FOREIGN KEY (employee_id) REFERENCES employees(employee_id)
);
COMMENT ON TABLE employee_profiles IS '従業員のプロフィール';

CREATE TABLE employee_addresses (
  employee_id uuid NOT NULL,
  postal_code VARCHAR(20),
  building_component_id uuid,
  street_address_component_id uuid NOT NULL,
  city_component_id uuid NOT NULL,
  state_component_id uuid NOT NULL,
  country_component_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (employee_id),
  CONSTRAINT fk_employee_addresses_employees_employee_id FOREIGN KEY (employee_id) REFERENCES employees(employee_id),
  CONSTRAINT fk_employee_addresses_building_component_id FOREIGN KEY (building_component_id) REFERENCES address_components(component_id),
  CONSTRAINT fk_employee_addresses_street_address_component_id FOREIGN KEY (street_address_component_id) REFERENCES address_components(component_id),
  CONSTRAINT fk_employee_addresses_city_component_id FOREIGN KEY (city_component_id) REFERENCES address_components(component_id),
  CONSTRAINT fk_employee_addresses_state_component_id FOREIGN KEY (state_component_id) REFERENCES address_components(component_id),
  CONSTRAINT fk_employee_addresses_country_component_id FOREIGN KEY (country_component_id) REFERENCES address_components(component_id)
);
COMMENT ON TABLE employee_addresses IS '従業員の住所';

CREATE TABLE employee_hires (
  hire_id uuid NOT NULL,
  employee_id uuid NOT NULL,
  organization_id uuid NOT NULL,
  hire_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_employee_hires_employees_employee_id FOREIGN KEY (employee_id) REFERENCES employees(employee_id),
  CONSTRAINT fk_employee_hires_organizations_organization_id FOREIGN KEY (organization_id) REFERENCES organizations(organization_id)
);
COMMENT ON TABLE employee_hires IS '従業員の入社履歴';

CREATE TABLE employee_separations (
  separation_id uuid NOT NULL,
  employee_id uuid NOT NULL,
  organization_id uuid NOT NULL,
  separation_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_employee_separations_employees_employee_id FOREIGN KEY (employee_id) REFERENCES employees(employee_id),
  CONSTRAINT fk_employee_separations_organizations_organization_id FOREIGN KEY (organization_id) REFERENCES organizations(organization_id)
);
COMMENT ON TABLE employee_separations IS '従業員の退社履歴';
