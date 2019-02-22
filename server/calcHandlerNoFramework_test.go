package server

import (
	"bytes"
	"github.com/MatthewJamesBoyle/mocking-example/server/no_framework_mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler_AddNumbers_No_Mock(t *testing.T) {
	t.Run("200 if successful request", func(t *testing.T) {
		mockResulter := no_framework_mocks.ResultStorer{
			SaveResult: nil,
		}
		mockSummer := no_framework_mocks.Summer{
			SaveResult: nil,
		}
		mockLogger := no_framework_mocks.SuperComplexLogger{
			SuperLogResult: no_framework_mocks.SuperLogResult{
				Result1: "",
				Result2: 0,
				Result3: 0,
				Result4: nil,
			},
		}

		handler, err := NewHandler(mockResulter, mockSummer, mockLogger)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString(`{"first": "3", "second":"40"}`))

		handler.AddNumbers(rr, req)
		assert.Equal(t, rr.Code, http.StatusOK)
	})
	t.Run("400 and log if sql db fails", func(t *testing.T) {
		DbFailureError := errors.New("some_db_failure_error")

		mockResulter := no_framework_mocks.ResultStorer{
			SaveResult: DbFailureError,
		}
		mockSummer := no_framework_mocks.Summer{
			SaveResult: nil,
		}
		mockLogger := no_framework_mocks.SuperComplexLogger{
			SuperLogResult: no_framework_mocks.SuperLogResult{
				Result1: "",
				Result2: 0,
				Result3: 0,
				Result4: nil,
			},
		}

		handler, err := NewHandler(mockResulter, mockSummer, mockLogger)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString(`{"first": "3", "second":"40"}`))

		handler.AddNumbers(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
	})

	t.Run("400 and log if mongo db fails", func(t *testing.T) {
		DbFailureError := errors.New("some_db_failure_error")

		mockResulter := no_framework_mocks.ResultStorer{
			SaveResult: nil,
		}
		mockSummer := no_framework_mocks.Summer{
			SaveResult: DbFailureError,
		}
		mockLogger := no_framework_mocks.SuperComplexLogger{
			SuperLogResult: no_framework_mocks.SuperLogResult{
				Result1: "",
				Result2: 0,
				Result3: 0,
				Result4: nil,
			},
		}

		handler, err := NewHandler(mockResulter, mockSummer, mockLogger)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString(`{"first": "3", "second":"40"}`))

		handler.AddNumbers(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
	})
}
