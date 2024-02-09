CREATE TABLE accounts(
  id BIGSERIAL PRIMARY KEY,
  initial_balance BIGINT NOT NULL,
  "limit" BIGINT NOT NULL
)
