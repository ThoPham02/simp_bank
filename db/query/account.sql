-- name: CreateAccount :one
INSERT INTO account (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY id;


-- name: UpdateAccount :one
UPDATE account SET owner = $2, balance = $3, currency = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account WHERE id = $1;