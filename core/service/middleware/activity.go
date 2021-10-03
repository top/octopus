package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"octopus/core/global"
	"octopus/core/service/model"
	"octopus/core/service/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Activity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.LOGGER.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()
		c.Next()
		record := model.Activity{
			IP:           c.ClientIP(),
			Method:       c.Request.Method,
			Path:         c.Request.URL.Path,
			Status:       c.Writer.Status(),
			Latency:      time.Since(now),
			Agent:        c.Request.UserAgent(),
			ErrorMessage: c.Errors.ByType(gin.ErrorTypePrivate).String(),
			Body:         string(body),
			Resp:         writer.body.String(),
		}
		if err := service.CreateActivity(record); err != nil {
			global.LOGGER.Error("create activity error:", zap.Error(err))
		}
	}
}
