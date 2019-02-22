package no_framework_mocks

type SuperComplexLogger struct {
	SuperLogResult SuperLogResult
}

type SuperLogResult struct {
	Result1 string
	Result2 int
	Result3 int
	Result4 error
}

func (s SuperComplexLogger) Log(message string) {}

func (s SuperComplexLogger) SuperLog(message string) (string, int, int, error) {
	return s.SuperLogResult.Result1, s.SuperLogResult.Result2, s.SuperLogResult.Result3, s.SuperLogResult.Result4
}
