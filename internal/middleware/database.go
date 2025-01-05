package middleware

import (
    "go-migration-app/internal/db"
    "github.com/gin-gonic/gin"
)

func DatabaseMiddleware(database *db.Database) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("database", database)
        c.Next()
    }
}
