package tapsync

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func handleSQLQuery(c *gin.Context, query string) interface {} {
	return nil
}

func parseSQLResult(sqlResult interface{}, format string) (interface{}, error) {
	switch format {
	case "votable":
		return nil, nil
	case "csv":
		return nil, nil
	case "tsv":
		return nil, nil
	case "fits":
		return nil, nil
	case "text":
		return nil, nil
	case "html":
		return nil, nil
	}
	return nil, fmt.Errorf("Could not parse result")
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

func syncPostHandler(c *gin.Context) {
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
		sqlResult := handleSQLQuery(c, query)
		parsedResult, err := parseSQLResult(sqlResult, format)
		if err != nil {
			c.XML(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Could not parse result"),
			})
			return
		}
		c.XML(http.StatusOK, parsedResult)
	default:
		caseInvalidLang(c)
		return
	}
}

func TapSyncService() *gin.Engine {
	r := gin.Default()
	r.POST("/sync", syncPostHandler)
	return r
}
