package utils

import "os"

func GetEnv(name string) (string, error) {
	env := os.Getenv(name)
	if len(env) == 0 {
		return "", EnvError(name)
	}
	return env, nil
}
