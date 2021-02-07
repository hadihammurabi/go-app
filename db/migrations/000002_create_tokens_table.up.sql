CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS tokens(
  id uuid DEFAULT uuid_generate_v4(),
  user_id uuid NOT NULL,
  token TEXT NOT NULL,

  expired_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,

  PRIMARY KEY (id)
);

CREATE INDEX ON tokens(
  user_id
);
