package initial

import (
	"octopus/core/global"
	"time"

	"github.com/fvbock/endless"
)

func StartServer() error {
	s := endless.NewServer(global.CONFIG.System.Addr, Routers())
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	global.LOGGER.Info("starting server...")
	if global.CONFIG.System.TLS {
		return s.ListenAndServe()
	}
	return s.ListenAndServeTLS(global.CONFIG.System.CertFile, global.CONFIG.System.KeyFile)
}
