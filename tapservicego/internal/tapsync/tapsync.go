package tapsync

import (
	"ataps/internal/parsers"
	"database/sql"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func setResponse(c *gin.Context, sqlResult []map[string]interface{}, format string) error {
	switch format {
	case "votable":
		votable, err := parsers.CreateVOTable(sqlResult)
		if err != nil {
			return err
		}
		result, err := parsers.VOTableToXML(votable)
		if err != nil {
			return err
		}
		c.Header("Content-Type", "application/x-votable+xml")
		c.Header("Content-Encoding", "UTF-8")
		c.Header("Content-Length", fmt.Sprintf("%d", len(result)))
		c.String(http.StatusOK, result)
	case "csv":
		result, err := parsers.ParseCSV(sqlResult)
		if err != nil {
			return err
		}
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Encoding", "UTF-8")
		c.Header("Content-Length", fmt.Sprintf("%d", len(result)))
		c.String(http.StatusOK, result)
	case "tsv":
		result, err := parsers.ParseTSV(sqlResult)
		if err != nil {
			return err
		}
		c.Header("Content-Type", "text/tab-separated-values")
		c.Header("Content-Encoding", "UTF-8")
		c.Header("Content-Length", fmt.Sprintf("%d", len(result)))
		c.String(http.StatusOK, result)
	case "fits":
		result, err := parsers.ParseFits(sqlResult)
		if err != nil {
			return err
		}
		c.Header("Content-Type", "application/fits")
		c.Header("Content-Encoding", "UTF-8")
		c.Header("Content-Length", fmt.Sprintf("%d", result.Len()))
		c.Data(http.StatusOK, "application/fits", result.Bytes())
	case "text":
		result := parsers.ParseText(sqlResult)
		c.Header("Content-Type", "text/plain")
		c.Header("Content-Encoding", "UTF-8")
		c.Header("Content-Length", fmt.Sprintf("%d", len(result)))
		c.String(http.StatusOK, result)
	case "html":
		c.Header("Content-Type", "text/html")
		c.Header("Content-Encoding", "UTF-8")
		err := parsers.ParseHTML(sqlResult, c.Writer)
		if err != nil {
			return err
		}
		c.Status(http.StatusOK)
	default:
		return fmt.Errorf("Invalid format")
	}
	return nil
}

func caseInvalidLang(c *gin.Context) {
	code := http.StatusBadRequest
	c.XML(code, getErrorVOTable(fmt.Errorf("Invalid LANG %s", c.PostForm("LANG")), code))
}

// getFormatOrResponseFormat returns the format or response format
// from the request. If both are provided, it returns an error.
// If neither are provided, it returns the default format, "votable".
// If one is provided, it returns that format.
// If the format is invalid, it returns an error.
func getFormatOrResponseFormat(c *gin.Context) string {
	format := c.PostForm("FORMAT")
	responseFormat := c.PostForm("RESPONSEFORMAT")
	validFormats := []string{"votable", "csv", "csv", "tsv", "fits", "text", "html"}
	if format != "" && responseFormat == "" {
		if slices.Contains(validFormats, format) {
			return format
		}
		code := http.StatusBadRequest
		c.XML(code, getErrorVOTable(fmt.Errorf("Invalid format %s", format), code))
		return ""
	}
	if responseFormat != "" && format == "" {
		if slices.Contains(validFormats, responseFormat) {
			return responseFormat
		}
		code := http.StatusBadRequest
		c.XML(code, getErrorVOTable(fmt.Errorf("Invalid format %s", responseFormat), code))
		return ""
	}
	if responseFormat == "" && format == "" {
		return "votable" // default format
	}
	code := http.StatusBadRequest
	c.XML(code, getErrorVOTable(fmt.Errorf("Both FORMAT and RESPONSEFORMAT provided"), code))
	return ""
}

// SyncPostHandler handles the POST request to /sync.
// Required paraemeters:
// - LANG: the language of the query. Only "PSQL" is supported for now.
// - QUERY: the query to execute.
// Optional parameters:
// - FORMAT: the format of the response. Default is "votable".
// - RESPONSEFORMAT: the format of the response. Default is "votable".
// If both FORMAT and RESPONSEFORMAT are provided, an error is returned.
func (service *TapSyncService) SyncPostHandler(c *gin.Context) {
	lang := c.PostForm("LANG")
	switch lang {
	case "PSQL":
		query := c.PostForm("QUERY")
		if query == "" {
			code := http.StatusBadRequest
			c.XML(code, getErrorVOTable(fmt.Errorf("No query provided"), code))
			return
		}
		format := getFormatOrResponseFormat(c)
		if format == "" {
			// here the error has already been added to the response
			// so we just return
			return
		}
		sqlResult, err := HandleSQLQuery(query, service.DB)
		if err != nil {
			// consider that the default XML render does not show quotes
			// if the error message contains quotes, it will be replaced by &#34;
			code := http.StatusInternalServerError
			c.XML(code, getErrorVOTable(err, code))
			return
		}
		err = setResponse(c, sqlResult, format)
		if err != nil {
			code := http.StatusInternalServerError
			c.XML(code, getErrorVOTable(err, code))
			return
		}
	default:
		caseInvalidLang(c)
		return
	}
}

type TapSyncService struct {
	Router *gin.Engine
	DB     *sql.DB
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
		DB:     db,
		config: config,
	}
	service.Router.POST("/sync", service.SyncPostHandler)
	return service
}
