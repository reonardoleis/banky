CREATE TABLE transactions(
  id BIGSERIAL PRIMARY KEY,
  type VARCHAR(1) NOT NULL,
  amount BIGINT NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMPTZ,
  account_id BIGINT NOT NULL,
  FOREIGN KEY (account_id) REFERENCES accounts(id)
  ON UPDATE CASCADE ON DELETE CASCADE
)
