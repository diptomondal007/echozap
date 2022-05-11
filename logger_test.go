package echozap

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestZapLogger(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/something", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}

	obs, logs := observer.New(zap.InfoLevel)
	logger := zap.New(obs)

	err := ZapLogger(WrapSugared(logger.Sugar()))(h)(c)

	assert.Nil(t, err)

	logFields := logs.All()
	require.Equal(t, 1, logs.Len())

	assert.Equal(t, logFields[0].ContextMap()["user_agent"], "")
	assert.Equal(t, logFields[0].ContextMap()["status"], int64(200))
}
