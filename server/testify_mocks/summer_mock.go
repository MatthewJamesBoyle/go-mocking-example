package testify_mocks

import "github.com/stretchr/testify/mock"

type SummerMock struct {
	mock.Mock
}

func (m SummerMock) Save(numberOne int64, numberTwo int64) error {
	args := m.Called(numberOne, numberTwo)
	return args.Error(0)
}
