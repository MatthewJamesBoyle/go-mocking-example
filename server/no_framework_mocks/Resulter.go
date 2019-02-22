package no_framework_mocks

type ResultStorer struct {
	SaveResult error
}

func (r ResultStorer) Save(result string) error {
	return r.SaveResult
}
