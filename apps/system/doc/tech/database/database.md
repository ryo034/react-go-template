# System Database

## ER図

基本的にはMiroで書いている

![ERD](./database.jpg)

### 自動生成

https://github.com/ariga/atlas

以下コマンドでER図のコード（mermaid）をデータベースから生成することができる。

※ただし全てのフィールドがオプショナルになってしまう。

```bash
> make inspect-database
```

```mermaid
erDiagram
    address_components {
      uuid component_id PK
      character_varying(50) component_type
      character_varying(256) component_name
      timestamp created_at
    }
    member_addresses {
      uuid member_id PK,FK
      character_varying(20) postal_code
      uuid building_component_id FK
      uuid street_address_component_id FK
      uuid city_component_id FK
      uuid state_component_id FK
      uuid country_component_id FK
      timestamp created_at
    }
    member_addresses }o--o| address_components : fk_member_addresses_building_component_id
    member_addresses }o--o| address_components : fk_member_addresses_city_component_id
    member_addresses }o--o| address_components : fk_member_addresses_country_component_id
    member_addresses |o--o| members : fk_member_addresses_members_member_id
    member_addresses }o--o| address_components : fk_member_addresses_state_component_id
    member_addresses }o--o| address_components : fk_member_addresses_street_address_component_id
    member_profiles {
      uuid member_id PK,FK
      character_varying(255) member_id_number
      character_varying(50) name
      timestamp created_at
      timestamp updated_at
    }
    member_profiles |o--o| members : fk_member_profiles_members_member_id
    members {
      uuid member_id PK
      uuid system_account_id FK
      uuid workspace_id FK
      timestamp created_at
    }
    members }o--o| system_accounts : fk_members_system_accounts_system_account_id
    members }o--o| workspaces : fk_members_workspaces_workspace_id
    membership_periods {
      uuid member_id PK,FK
      date start_date PK
      date end_date
      timestamp created_at
    }
    membership_periods }o--o| members : fk_membership_periods_members_member_id
    system_account_phone_numbers {
      uuid system_account_id PK,FK
      character_varying(15) phone_number
      timestamp created_at
      timestamp updated_at
    }
    system_account_phone_numbers |o--o| system_accounts : fk_system_account_phone_numbers_system_accounts_system_account_
    system_account_profiles {
      uuid system_account_id PK,FK
      character_varying(255) name
      character_varying(256) email
      boolean email_verified
      timestamp created_at
      timestamp updated_at
    }
    system_account_profiles |o--o| system_accounts : fk_system_account_profiles_system_accounts_system_account_id
    system_accounts {
      uuid system_account_id PK
      timestamp created_at
    }
    workspace_details {
      uuid workspace_id PK,FK
      character_varying(100) name
      timestamp created_at
      timestamp updated_at
    }
    workspace_details |o--o| workspaces : fk_workspace_details_workspaces_workspace_id
    workspaces {
      uuid workspace_id PK
      timestamp created_at
    }
```

### SQLからER図を確認

以下のURLでSQLファイル（001\_create\_tables.sql）からER図を確認できる

https://gh.atlasgo.cloud/explore
