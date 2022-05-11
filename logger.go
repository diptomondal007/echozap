package echozap

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Sugared struct {
	*zap.SugaredLogger
}

func WrapSugared(z *zap.SugaredLogger) *Sugared {
	return &Sugared{z}
}

type Logger interface {
	Log(status int, err error, fields []interface{})
}

// ZapLogger is a middleware and zap to provide an "access log" like logging for each request.
func ZapLogger(log Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			fields := []interface{}{
				"remote_ip", c.RealIP(),
				"latency", time.Since(start).String(),
				"host", req.Host,
				"request", fmt.Sprintf("%s %s", req.Method, req.RequestURI),
				"status", res.Status,
				"size", res.Size,
				"user_agent", req.UserAgent(),
			}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
				fields = append(fields, zap.String("request_id", id))
			}

			log.Log(res.Status, err, fields)

			return nil
		}
	}
}

func (sl *Sugared) Log(status int, err error, fields []interface{}) {
	n := status
	switch {
	case n >= 500:
		sl.With(fields...).With(zap.Error(err)).Error("Server error")
	case n >= 400:
		sl.With(fields...).With(zap.Error(err)).Warn("Client error")
	case n >= 300:
		sl.With(fields...).Info("Redirection")
	default:
		sl.With(fields...).Info("Success")
	}
}
