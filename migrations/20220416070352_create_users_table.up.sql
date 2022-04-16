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
  full_name char(255) NOT NULL,
  email char(255) NOT NULL,
  phone_no char(12) NULL,
  avatar_url varchar NULL,
  dob timestamp NULL,
  address varchar NULL,
  grade int NULL,
  school varchar NULL,
  gender gender NULL,
  owner_type owner_type NULL,
  signup_provider signup_provider NOT NULL,
  bank_transfer_code char(255) NOT NULL,
  firebase_id char(255) NOT NULL,
  is_active bool DEFAULT true,
  created_at timestamp NOT NULL DEFAULT now(),
  created_by varchar NOT NULL,
  updated_at timestamp NOT NULL DEFAULT now(),
  updated_by varchar NOT NULL,
  CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS users_unique_email ON users(email);