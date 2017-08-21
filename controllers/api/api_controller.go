package api

import (
	"fmt"

	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/orubin/dizzyserver/lib"
)

type APIController struct {
	beego.Controller
}

// @Summary Create User
// @Description Create a new user
// @Param	context		body
// @Success 200
// @Failure 400
// @router /create_user [post]
func (c *APIController) CreateUser() {
	c.Data["data"] = "success"
	c.ServeJSON()
}

// @Summary Get User Details
// @Description Gets all the user details
// @Param	context		body
// @Success 200
// @Failure 400
// @router /get_user [post]
func (c *APIController) GetUser() {
	users := Users{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &users)
	if err == nil {
		lib.NewUserFetcher(users.IDS).Fetch()
		c.Data["data"] = "success"
		c.ServeJSON()
	} else {
		c.AbortE("Error parsing input", err)
	}
}

func (c *APIController) AbortF(code, formatMessage string, args ...interface{}) {
	c.Ctx.Input.SetData("error", fmt.Sprintf(formatMessage, args...))
	c.Abort(code)
}

func (c *APIController) AbortE(code string, err error) {
	c.Ctx.Input.SetData("error", err.Error())
	c.Abort(code)
}

type Users struct {
	IDS []int `json:"users"`
}
