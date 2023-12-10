package userqueries

const (
	GET_MANY = `
		SELECT 
			id, username, fullname, created_at, updated_at, deleted_at
		FROM
			users
		ORDER BY
			username ASC;
	`
)
