package server

import (
	"bytes"
	"github.com/MatthewJamesBoyle/mocking-example/server/testify_mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler_AddNumbers(t *testing.T) {
	t.Run("200 if successful request", func(t *testing.T) {
		mockLogger := new(testify_mocks.LoggerMock)
		mockSummer := new(testify_mocks.SummerMock)
		mockResulter := new(testify_mocks.ResulterMock)

		handler, err := NewHandler(mockResulter, mockSummer, mockLogger)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/someUrl", bytes.NewBufferString(
			`{
				
					}`))

		handler.AddNumbers(rr, req)
		assert.Equal(t, rr.Code, http.StatusOK)
	})
}
