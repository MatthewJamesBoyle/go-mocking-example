package testify_mocks

import "github.com/stretchr/testify/mock"

type ResulterMock struct {
	*mock.Mock
}

func (m *ResulterMock) Save(result string) error {
	args := m.Called()
	return args.Error(0)
}
