package user

var (
	readUserByPasscodeQuery = `
		SELECT 
			*
		FROM 
			"%s".user
		WHERE
			passcode = ?
			AND is_active = true
		LIMIT 1
	`

	readUsersQuery = `
		SELECT 
			*
		FROM 
			"%s".user
		WHERE
			is_active = true
	`

	addUserImageQuery = `
		UPDATE
			"%s".user
		SET
			image_name = ?
		WHERE
			passcode = ?
	`

	readImageNameByIdQuery = `
		SELECT 
			image_name
		FROM 
			"%s".user
		WHERE
			id = ?
			AND is_active = true
		LIMIT 1
	`
)
