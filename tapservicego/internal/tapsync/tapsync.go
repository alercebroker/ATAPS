package tapsync

import (
	"ataps/internal/sqlparser"
	"database/sql"
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)


func setResponse(c *gin.Context, sqlResult []map[string]interface{}, format string) error {
	switch format {
	case "votable":
		return nil 
	case "csv":
		result, err := sqlparser.ParseCSV(sqlResult)
		if err != nil {
			return err
		}
		// TODO: make actual csv response
		c.String(http.StatusOK, result)
	case "tsv":
		return nil
	case "fits":
		return nil
	case "text":
		return nil
	case "html":
		return nil
	default:
		return fmt.Errorf("Invalid format")
	}
	return nil
}


func caseInvalidLang(c *gin.Context) {
	c.XML(http.StatusBadRequest, gin.H{
		"error": fmt.Sprintf("Invalid lang"),
	})
}

func getFormatResponseFormat(c *gin.Context) string {
	format := c.PostForm("FORMAT")
	responseFormat := c.PostForm("RESPONSEFORMAT")
	validFormats := []string{"votable", "csv", "csv", "tsv", "fits", "text", "html"}
	if format != "" && responseFormat == "" {
		if slices.Contains(validFormats, format){
			return format
		}
		c.XML(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid format"),
		})
		return ""
	}
	if responseFormat != "" && format == "" {
		if slices.Contains(validFormats, responseFormat) {
			return responseFormat
		}
		c.XML(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid response format"),
		})
		return ""
	}
	if responseFormat == "" && format == "" {
		return "votable" // default format
	}
	c.XML(http.StatusBadRequest, gin.H{
		"error": fmt.Sprintf("Do not provide both FORMAT and RESPONSEFORMAT"),
	})
	return ""
}

func (service *TapSyncService) SyncPostHandler(c *gin.Context) {
	lang := c.PostForm("LANG")
	switch lang {
	case "PSQL":
		query := c.PostForm("QUERY")
		if query == "" {
			c.XML(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Query is empty"),
			})
			return
		}
		format := getFormatResponseFormat(c)
		if format == "" {
			return
		}
		sqlResult, err := HandleSQLQuery(query, service.DB)
		if err != nil {
			c.XML(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Could not execute query"),
			})
			return
		}
		err = setResponse(c, sqlResult, format)
		if err != nil {
			c.XML(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Could not set response"),
			})
			return
		}
	default:
		caseInvalidLang(c)
		return
	}
}

type TapSyncService struct {
	Router *gin.Engine
	DB *sql.DB
	config *Config
}

func NewTapSyncService() *TapSyncService {
	config := GetConfig()
	db, err := GetDB(config.DatabaseURL)
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	service := &TapSyncService{
		Router: router,
		DB: db,
		config: config,
	}
	service.Router.POST("/sync", service.SyncPostHandler)
	return service
}
