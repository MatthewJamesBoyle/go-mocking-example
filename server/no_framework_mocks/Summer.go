package no_framework_mocks

type Summer struct {
	SaveResult error
}

func (s Summer) Save(numberOne int64, numberTwo int64) error {
	return s.SaveResult
}
