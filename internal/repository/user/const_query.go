package user

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	getUserByID *sqlx.Stmt
}

const (
	getUserByID = `
	SELECT
		*
	FROM
		users
	WHERE
		id=$1
	`
)
