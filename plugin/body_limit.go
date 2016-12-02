package plugin

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	BodyLimit struct {
		Base                       `json:",squash"`
		middleware.BodyLimitConfig `json:",squash"`
	}
)

func (l *BodyLimit) Init() (err error) {
	l.Middleware = middleware.BodyLimitWithConfig(l.BodyLimitConfig)
	return
}

func (*BodyLimit) Priority() int {
	return 1
}

func (l *BodyLimit) Execute(next echo.HandlerFunc) echo.HandlerFunc {
	return l.Middleware(next)
}
