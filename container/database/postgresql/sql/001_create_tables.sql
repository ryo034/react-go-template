CREATE TABLE genders (
  code CHAR(3),
  name VARCHAR(50) NOT NULL,
  PRIMARY KEY (code)
);
COMMENT ON TABLE genders IS '性別マスタ';

CREATE TABLE weekdays (
  code CHAR(3),
  name VARCHAR(10) NOT NULL UNIQUE,
  PRIMARY KEY (code)
);
COMMENT ON TABLE weekdays IS '曜日マスタ';
