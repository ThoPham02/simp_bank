-- name: CreateTranfer :one
INSERT INTO tranfers (
  from_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTranfer :one
SELECT * FROM tranfers
WHERE id = $1;

-- name: ListTranfers :many
SELECT * FROM tranfers
ORDER BY id;

-- name: UpdateTranfer :exec
UPDATE tranfers
SET from_account_id = $2, to_account_id = $3, amount = $4
WHERE id = $1;

-- name: DeleteTranfer :exec
DELETE FROM tranfers WHERE id = $1;