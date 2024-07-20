package testCases

var EnvKeyStringTestCase = []struct {
	StructKey string
	EnvKey    string
}{
	{"MyKey", "MY_KEY"},
	{"BrokerHost", "BROKER_HOST"},
}
