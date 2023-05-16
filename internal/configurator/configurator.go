package configurator

import "os"

func GetEnv(key string, fb *string) *string {
	if value, ok := os.LookupEnv(key); ok {
		return &value
	}
	return fb
}
