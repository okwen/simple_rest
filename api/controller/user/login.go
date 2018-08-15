package user

import (
	"log"
	"net/http"
	"simple_rest/api/protocol"
	"simple_rest/database"
	"simple_rest/env"

	"github.com/gin-gonic/gin"
)

// GetLoginAccountInput : Input參數
type GetLoginAccountInput struct {
	Account string `form:"Account"`
	Password string `form:"Password"`
}

// Login API
func Login(c *gin.Context) {
	res := &protocol.Response{}
	input := &GetLoginAccountInput{}

	// 綁定Input參數至結構中
	if err := c.Bind(input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := LoginWithArguments(input.Account, input.Password)

	if !result {
		res.Code = 2
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// LoginWithArguments : 登入
func LoginWithArguments(account string, password string) (isSucceed bool) {
	dbS := database.GetConn(env.AccountDB)

	sql := " SELECT * FROM user WHERE account = ? AND password = ?"

	var params []interface{}
	params = append(params, account, password)

	rows, err := dbS.Query(sql, params...)
	_ = err

	defer rows.Close()
	isSucceed = false
	for rows.Next() {
		isSucceed = true
		break
	}

	return
}
