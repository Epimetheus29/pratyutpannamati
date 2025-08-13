package routes

import (
	"database/sql"

	"pratyutpannamati/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/users", handlers.GetUsers(db))
}
