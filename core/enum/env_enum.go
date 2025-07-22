package enum

type ENV string

const (
	DEV  ENV = "dev"
	TEST ENV = "test"
	PROD ENV = "prod"
)

func (e ENV) String() string {
	return string(e)
}
