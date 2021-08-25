package ping

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hdomin/bookstore_users-api/datasources/mysql/users_db"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func PingDB(c *gin.Context) {

	if users_db.Client == nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("DB Client is null %s", users_db.Client))
		return
	}

	if err := users_db.Client.Ping(); err != nil {
		c.String(http.StatusAccepted, err.Error())
		return
	}

	c.String(http.StatusOK, "Ping DB Successfully")

}
