package types

// n.Router(POST, "/register-user", u.RegisterUser)
type RegisterUserReq struct {
	UserName    string   `json:"userName" binding:"required"`
	Description string   `json:"description"`
	Hobby       []string `json:"hobby"`
	Latitude    float64  `json:"latitude" binding:"required,min=-90,max=90"`
	Hardness    float64  `json:"hardness" binding:"required,min=-180,max=180"`
}

// n.Router(GET, "/around-users", u.AroundUsers)
type AroundUsers struct {
	UserName string `form:"user" binding:"required"`
	Range    int64  `form:"range" binding:"required"`
	Limit    int64  `form:"limit" binding:"required"`
}
