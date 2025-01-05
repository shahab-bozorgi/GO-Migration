package api

import (
	"go-migration-app/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Query struct {
	SQL string `json:"sql"`
}


func GetTablesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Fetch tables"})
}

func GetTables(c *gin.Context) {
    dbInterface, exists := c.Get("database")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
        return
    }

    database, ok := dbInterface.(*db.Database)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid database connection"})
        return
    }

    tables, err := db.FetchTables(database)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tables"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"tables": tables})
}



func ApplyQuery(c *gin.Context) {
    var query Query

    if err := c.ShouldBindJSON(&query); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
        return
    }

    dbInterface, exists := c.Get("database")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
        return
    }

    database, ok := dbInterface.(*db.Database)

    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid database connection"})
        return
    }

    rows, err := database.Query(query.SQL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to execute SQL", "error": err.Error()})
        return
    }
    defer rows.Close()

    var result []map[string]interface{}
    columns, err := rows.Columns()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get columns", "error": err.Error()})
        return
    }
    for rows.Next() {
        values := make([]interface{}, len(columns))
        for i := range values {
            values[i] = new(interface{})
        }
        if err := rows.Scan(values...); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to scan rows", "error": err.Error()})
            return
        }

        rowMap := make(map[string]interface{})
        for i, col := range columns {
            rowMap[col] = *(values[i].(*interface{}))
        }
        result = append(result, rowMap)
    }

    c.JSON(http.StatusOK, gin.H{"result": result})
}
