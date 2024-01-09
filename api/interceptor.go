package api

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type toolBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r toolBodyWriter) Write(b []byte) (int, error) {
	return r.body.Write(b)
}

func responseInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		wb := &toolBodyWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}
		c.Writer = wb
		c.Next()

		statusCode := c.Writer.Status()
		originalResponse := wb.body.String()

		// Clear the original buffer
		wb.body = &bytes.Buffer{}

		if statusCode >= 400 {
			response := fmt.Sprintf(
				`{"status": "%s", "data": %s}`,
				http.StatusText(statusCode),
				originalResponse,
			)
			wb.Write([]byte(response))
		} else {
			response := fmt.Sprintf(
				`{"status": "%s", "data": %s}`,
				http.StatusText(statusCode),
				originalResponse,
			)
			wb.Write([]byte(response))
		}
		wb.ResponseWriter.Write(wb.body.Bytes())
	}
}
