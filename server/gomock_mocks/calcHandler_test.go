package server

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MatthewJamesBoyle/go-mocking-example/server"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalcHandler_AddNumbers(t *testing.T) {

	const addEndpoint = "/add/"

	t.Run("it should return 400 because first is empty", func(t *testing.T) {

		handler, err := server.NewHandler(
			nil,
			nil,
			nil,
		)

		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodGet,
			addEndpoint,
			bytes.NewBufferString(`{}`),
		)

		rr := httptest.NewRecorder()

		handler.AddNumbers(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("it should return 400 because second is empty", func(t *testing.T) {

		handler, err := server.NewHandler(
			nil,
			nil,
			nil,
		)

		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodGet,
			addEndpoint,
			bytes.NewBufferString(`{
				"first" : "10"
			}`),
		)

		rr := httptest.NewRecorder()

		handler.AddNumbers(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("it should return 400 because first is malformed", func(t *testing.T) {

		handler, err := server.NewHandler(
			nil,
			nil,
			nil,
		)

		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodGet,
			addEndpoint,
			bytes.NewBufferString(`{
				"first" : "'",
				"second" : "10"
			}`),
		)

		rr := httptest.NewRecorder()

		handler.AddNumbers(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("it should return 400 because second is malformed", func(t *testing.T) {

		handler, err := server.NewHandler(
			nil,
			nil,
			nil,
		)

		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodGet,
			addEndpoint,
			bytes.NewBufferString(`{
				"first" : "10",
				"second" : "'"
			}`),
		)

		rr := httptest.NewRecorder()

		handler.AddNumbers(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("it should return 400 because something went wrong upon calling summer save", func(t *testing.T) {

		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		testErr := errors.New("some error")

		mockSummer := NewMockSumStorer(ctrl)
		mockLogger := NewMockSuperComplexLogger(ctrl)

		handler, err := server.NewHandler(
			nil,
			mockSummer,
			mockLogger,
		)

		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodGet,
			addEndpoint,
			bytes.NewBufferString(`{
				"first" : "10",
				"second" : "10"
			}`),
		)

		// expectations
		mockSummer.EXPECT().Save(int64(10), int64(10)).Return(testErr)
		mockLogger.EXPECT().
			SuperLog("SUPER LOG").
			Return(
				"hello",
				1,
				1,
				nil,
			)

		rr := httptest.NewRecorder()

		handler.AddNumbers(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("it should return 400 because something went wrong upon calling resulter save", func(t *testing.T) {

		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		testErr := errors.New("some error")

		mockSummer := NewMockSumStorer(ctrl)
		mockResulter := NewMockResultStorer(ctrl)

		handler, err := server.NewHandler(
			mockResulter,
			mockSummer,
			nil,
		)

		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodGet,
			addEndpoint,
			bytes.NewBufferString(`{
				"first" : "10",
				"second" : "10"
			}`),
		)

		// expectations
		gomock.InOrder(
			mockSummer.EXPECT().Save(int64(10), int64(10)).Return(nil),
			mockResulter.EXPECT().Save(string(int64(20))).Return(testErr),
		)

		rr := httptest.NewRecorder()

		handler.AddNumbers(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("it should return 200", func(t *testing.T) {

		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		mockSummer := NewMockSumStorer(ctrl)
		mockResulter := NewMockResultStorer(ctrl)

		handler, err := server.NewHandler(
			mockResulter,
			mockSummer,
			nil,
		)

		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodGet,
			addEndpoint,
			bytes.NewBufferString(`{
				"first" : "10",
				"second" : "10"
			}`),
		)

		// expectations
		gomock.InOrder(
			mockSummer.EXPECT().Save(int64(10), int64(10)).Return(nil),
			mockResulter.EXPECT().Save(string(int64(20))).Return(nil),
		)

		rr := httptest.NewRecorder()

		handler.AddNumbers(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

	})

}
