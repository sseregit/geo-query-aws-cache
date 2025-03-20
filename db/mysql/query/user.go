package query

const (
	InsertIgnoreUser = "" +
		"INSERT IGNORE user (user_name, description, image, hobby) VALUES (?,?,'[]',?);"

	InsertIgnoreUserLocation = "" +
		"INSERT IGNORE user_location (user_name, latitude, hardness, location) VALUES (?,?,?,POINT(?, ?));"
)
