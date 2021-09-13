package courses

import (
	mooc "courses-api-mysql-and-cb/internal"
	"courses-api-mysql-and-cb/internal/storage/couchbase"
	"courses-api-mysql-and-cb/internal/storage/mysql"
	"database/sql"
	"fmt"
	"net/http"
	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
)

const (
	dbUser = "user"
	dbPass = "78jsh90io9"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "coursesdb"
)

// createRequest is the adapter struct for the API HTTP that drives course data from CreateHandler to sqlCourse struct(the port between course model and mysql adapter).
type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
	Price    string `json:"price" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course := mooc.NewCourse(req.ID, req.Name, req.Duration, req.Price)

		mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		db, err := sql.Open("mysql", mysqlURI)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		courseRepository := mysql.NewCourseRepository(db)

		if err := courseRepository.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		//I should make cbURI with const's for more security i guess
		cluster, err := gocb.Connect(
			"174.138.44.19:8091",
			gocb.ClusterOptions{
				Username: "Administrator",
				Password: "gregoria123R",
			})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		cbCourseRepository := couchbase.NewCbRepository(cluster)
		if err := cbCourseRepository.Save(course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		
		ctx.Status(http.StatusCreated)
	}
}
