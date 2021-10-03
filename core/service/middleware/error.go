package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"octopus/core/global"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if e := strings.ToLower(se.Err.Error()); strings.Contains(e, "broken pipe") || strings.Contains(e, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpReq, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.LOGGER.Error(c.Request.URL.Path, zap.Any("err", err), zap.String("request", string(httpReq)))
					c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					global.LOGGER.Error("[Recovery from panic]", zap.Any("err", err), zap.String("request", string(httpReq)), zap.String("stack", string(debug.Stack())))
				} else {
					global.LOGGER.Error("[Recovery from panic]", zap.Any("err", err), zap.String("request", string(httpReq)))
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
