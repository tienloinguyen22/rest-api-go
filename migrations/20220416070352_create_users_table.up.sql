CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE gender AS ENUM (
  'MALE',
  'FEMALE',
  'OTHER'
);

CREATE TYPE signup_provider AS ENUM (
  'EMAIL',
  'FACEBOOK',
  'GOOGLE',
  'APPLE'
);

CREATE TYPE owner_type AS ENUM (
  'STUDENT',
  'PARENT',
  'TEACHER',
  'OTHER'
);

CREATE TABLE IF NOT EXISTS users (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  full_name varchar NOT NULL,
  email varchar NOT NULL,
  phone_no varchar NULL,
  avatar_url varchar NULL,
  dob timestamptz NULL,
  address varchar NULL,
  grade int NULL,
  school varchar NULL,
  gender gender NULL,
  owner_type owner_type NULL,
  signup_provider signup_provider NOT NULL,
  bank_transfer_code varchar NOT NULL,
  firebase_id varchar NOT NULL,
  is_active bool DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  created_by varchar NOT NULL,
  updated_at timestamptz NOT NULL DEFAULT now(),
  updated_by varchar NOT NULL,
  CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS users_unique_email ON users(email);