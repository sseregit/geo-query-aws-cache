package network

import (
	"geo-query-aws-cache/module/API/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	n *Network
}

func userRouter(n *Network) {
	u := &user{n}

	n.Router(POST, "/register-user", u.RegisterUser)
	n.Router(POST, "/upload-image", u.UploadImage)
	n.Router(GET, "/around-users", u.AroundUsers)
}

func (u *user) RegisterUser(c *gin.Context) {
	var req types.RegisterUserReq

	if err := c.ShouldBind(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else if err = u.n.s.RegisterUser(req); err != nil {
		res(c, http.StatusInternalServerError, err.Error())
	} else {
		res(c, http.StatusOK, "Success")
	}
}

func (u *user) UploadImage(c *gin.Context) {
	name := c.Request.FormValue("userName")
	file, header, err := c.Request.FormFile("image")

	if err != nil || name == "" {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else {
		if err = u.n.s.UploadFile(name, header, file); err != nil {
			res(c, http.StatusInternalServerError, err.Error())
		} else {
			res(c, http.StatusOK, "Success")
		}
	}
}

func (u *user) AroundUsers(c *gin.Context) {
	var req types.AroundUsers

	if err := c.ShouldBindQuery(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else if result, err := u.n.s.FindAroundUsers(req.UserName, req.Range, req.Limit); err != nil {
		res(c, http.StatusInternalServerError, err.Error())
	} else {
		res(c, http.StatusOK, result)
	}
}
