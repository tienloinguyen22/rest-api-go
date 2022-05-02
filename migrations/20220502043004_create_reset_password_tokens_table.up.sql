CREATE TABLE IF NOT EXISTS reset_password_tokens (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  user_id uuid NOT NULL,
  completed bool DEFAULT false,
  expired_at timestamptz NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  created_by varchar NOT NULL,
  CONSTRAINT reset_password_tokens_pk PRIMARY KEY (id),
  CONSTRAINT reset_password_tokens_users_fk FOREIGN KEY (user_id) REFERENCES users(id)
);