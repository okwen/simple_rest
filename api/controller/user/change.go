package user

import (
	"log"
	"net/http"
	"simple_rest/api/protocol"
	"simple_rest/database"
	"simple_rest/env"

	"github.com/gin-gonic/gin"
)

// GetChangeAccountInput : Input參數
type GetChangeAccountInput struct {
	Account string `form:"Account"`
	Password string `form:"Password"`
}

// ChangeUser API
func ChangeUser(c *gin.Context) {
	res := &protocol.Response{}
	input := &GetChangeAccountInput{}

	// 綁定Input參數至結構中
	if err := c.Bind(input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := ChangeUserWithAccount(input.Account, input.Password)
	_ = err

	res.Result = result

	c.JSON(http.StatusOK, res)
	return
}

// ChangeUserWithAccount : 修改會員密碼
func ChangeUserWithAccount(account string, password string) (result *UserResult, err error) {
	dbS := database.GetConn(env.AccountDB)

	sql := " UPDATE user SET password = ? WHERE account = ?"

	var params []interface{}
	params = append(params, password, account)

	res, err := dbS.Exec(sql, params...)

	var isSuceed bool = false

    if err == nil {
		count, errRows := res.RowsAffected()  
		_ = errRows
	    
		if count > 0 {
			isSuceed = true
		}
	}  
	
	result = &UserResult {
		IsOK: isSuceed,
	}

	return
}
