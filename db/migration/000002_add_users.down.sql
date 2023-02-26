ALTER TABLE IF EXISTS "account" drop constraint IF EXISTS "owner_currency_key";

ALTER TABLE IF EXISTS "account" drop constraint IF EXISTS "account_owner_fkey";

DROP TABLE IF EXISTS users;