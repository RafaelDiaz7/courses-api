package courses

import (
	"courses-api-mysql-and-cb/internal/storage/mysql"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHandler returns an HTTP handler to get courses.
func GetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		db, err := sql.Open("mysql", mysqlURI)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		defer db.Close()

		courseRepository := mysql.NewCourseRepository(db)

		resultRows,err := courseRepository.ScanTableRows()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, resultRows)
	}
}
