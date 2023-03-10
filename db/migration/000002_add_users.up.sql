CREATE TABLE IF NOT EXISTS "users" (
  "username" varchar PRIMARY KEY,
  "hash_pass" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "account" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "account" ADD constraint "owner_currency_key"  UNIQUE ("owner", "currency")
