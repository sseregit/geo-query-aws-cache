package query

const (
	InsertIgnoreUser = "" +
		"INSERT IGNORE user (user_name, description, image, hobby) VALUES (?,?,'[]',?);"

	InsertIgnoreUserLocation = "" +
		"INSERT IGNORE user_location (user_name, latitude, hardness, location) VALUES (?,?,?,POINT(?, ?));"
)

const (
	GetUserByName = `
	SELECT u.user_name, u.image, u.description, u.hobby, ul.latitude, ul.hardness
	FROM user AS u JOIN user_location as ul ON u.user_name = ul.user_name WHERE u.user_name = ?;
	`

	GetAroundUsers = `
	SELECT u.user_name, u.image, u.description, ul.latitude, ul.hardness
	FROM user AS u JOIN user_location AS ul ON u.user_name = ul.user_name
	WHERE u.user_name != ? AND ST_Distance_Sphere(POINT(?, ?), POINT(ul.hardness, ul.latitude)) <= ?
	ORDER BY ST_Distance_Sphere(POINT(?, ?), POINT(ul.hardness, ul.latitude)) LIMIT ?;
	`
)
