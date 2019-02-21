package testify_mocks

import "github.com/stretchr/testify/mock"

type LoggerMock struct {
	mock.Mock
}

func (m *LoggerMock) Log(message string) {
	m.Called(message)
}

func (m *LoggerMock) SuperLog(message string) (string, int, int, error) {
	args := m.Called(message)
	return args.String(0), args.Int(1), args.Int(2), args.Error(3)
}
