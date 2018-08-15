package user

import (
	"log"
	"net/http"
	"simple_rest/api/protocol"
	"simple_rest/database"
	"simple_rest/env"

	"github.com/gin-gonic/gin"
)

// GetAccountInput : Input參數
type GetDeleteAccountInput struct {
	Account string `form:"Account"`
}

// DeletetUser API
func DeletetUser(c *gin.Context) {
	res := &protocol.Response{}
	input := &GetAccountInput{}

	// 綁定Input參數至結構中
	if err := c.Bind(input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := DeletetUserWithAccount(input.Account)
	_ = err

	res.Result = result

	c.JSON(http.StatusOK, res)
	return
}

// CreateUserWithArguments : 新增使用者
func DeletetUserWithAccount(account string) (result *DeletetUserResult, err error) {
	dbS := database.GetConn(env.AccountDB)

	sql := " DELETE FROM user WHERE account = ?"

	var params []interface{}
	params = append(params, account)

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
