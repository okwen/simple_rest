package user

import (
	"log"
	"net/http"
	"simple_rest/api/protocol"
	"simple_rest/database"
	"simple_rest/env"

	"github.com/gin-gonic/gin"
)

// GetCreateAccountInput : Input參數
type GetCreateAccountInput struct {
	Account string `form:"Account"`
	Password string `form:"Password"`
}

// GetUser API
func CreatetUser(c *gin.Context) {
	res := &protocol.Response{}
	input := &GetCreateAccountInput{}

	// 綁定Input參數至結構中
	if err := c.Bind(input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var isSuceed bool = true
	err := CreateUserWithArguments(input.Account, input.Password)
	if err != nil {
		isSuceed = false
	}

	result := UserResult {
		IsOK: isSuceed,
	}

	res.Result = result

	c.JSON(http.StatusOK, res)
	return
}

// CreateUserWithArguments : 新增使用者
func CreateUserWithArguments(account string, password string) (err error) {
	dbS := database.GetConn(env.AccountDB)

	sql := " INSERT INTO user (account, password) VALUES (?, ?);"

	var params []interface{}
	params = append(params, account, password)

	unusedRows, err := dbS.Query(sql, params...)
	_ = unusedRows
	return
}
