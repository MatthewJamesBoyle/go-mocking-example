package server

import (
	"bytes"
	"github.com/MatthewJamesBoyle/mocking-example/server/testify_mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		mockSummer.On("Save", int64(3), int64(40)).Return(nil)
		mockResulter.On("Save", string(43)).Return(nil)
		handler, err := NewHandler(mockResulter, mockSummer, mockLogger)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString(`{"first": "3", "second":"40"}`))

		handler.AddNumbers(rr, req)
		assert.Equal(t, rr.Code, http.StatusOK)
	})
	t.Run("400 and log if sql db fails", func(t *testing.T) {
		mockLogger := new(testify_mocks.LoggerMock)
		mockSummer := new(testify_mocks.SummerMock)
		mockResulter := new(testify_mocks.ResulterMock)
		mockSummer.On("Save", int64(3), int64(40)).Return(errors.New("some-error"))
		mockResulter.On("Save", string(43)).Return(nil)
		mockLogger.On("SuperLog", mock.Anything).Return("", 1, 1, nil)

		handler, err := NewHandler(mockResulter, mockSummer, mockLogger)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString(`{"first": "3", "second":"40"}`))

		handler.AddNumbers(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
	})
	t.Run("400 and log if mongo db fails", func(t *testing.T) {
		mockLogger := new(testify_mocks.LoggerMock)
		mockSummer := new(testify_mocks.SummerMock)
		mockResulter := new(testify_mocks.ResulterMock)
		mockSummer.On("Save", int64(3), int64(40)).Return(nil)
		mockResulter.On("Save", string(43)).Return(errors.New("some-error"))
		mockLogger.On("SuperLog", mock.Anything).Return("", 1, 1, nil)

		handler, err := NewHandler(mockResulter, mockSummer, mockLogger)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString(`{"first": "3", "second":"40"}`))

		handler.AddNumbers(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
	})
}
