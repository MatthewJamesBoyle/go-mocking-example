package server

import (
	"bytes"
	"github.com/MatthewJamesBoyle/mocking-example/server/testify_mocks"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler_AddNumbers(t *testing.T) {
	t.Run("200 if successful request", func(t *testing.T) {
		mockLogger := new(testify_mocks.LoggerMock)

		handler, err := NewHandler(mockResulter, mockSummer, mockLogger)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "", bytes.NewBufferString(`[}`))

		handler.AddNumbers(rr, req)
	})
}
