package factory

type DataTest struct {
	Name          string
	Variables     map[string]interface{}
	ExpectedError error
}
