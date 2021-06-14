package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// fixtures
const (
	AWS_ACCESS_KEY_ID_VAL     = "1234fkd"
	AWS_SECRET_ACCESS_KEY_VAL = "dsdfasdfadf"
	AWS_SESSION_TOKEN_VAL     = "sdf414fedszxzx4ddfs"
)

func TestGetEnv(t *testing.T) {

	t.Run("When the env vars doesn't loaded", func(t *testing.T) {

		t.Run("When AWS_ACCESS_KEY_ID doesn't found!", func(t *testing.T) {
			_, err := GetEnv("AWS_ACCESS_KEY_ID")
			assert.EqualError(t, err, errAWSAccessKeyIDNotFound.Err.Error())
		})

		t.Run("When AWS_SECRET_ACCESS_KEY doesn't found!", func(t *testing.T) {
			_, err := GetEnv("AWS_SECRET_ACCESS_KEY")
			assert.EqualError(t, err, errAWSSecretAccessKeyNotFound.Err.Error())
		})

		t.Run("When AWS_SESSION_TOKEN doesn't found!", func(t *testing.T) {
			_, err := GetEnv("AWS_SESSION_TOKEN")
			assert.EqualError(t, err, errAWSSessionTokenNotFound.Err.Error())
		})

		t.Run("When Get Unknow env var", func(t *testing.T) {
			_, err := GetEnv("HEROKU_KEY")
			assert.EqualError(t, err, unknownError)
		})

	})

	t.Run("When the env vars are loaded", func(t *testing.T) {

		t.Run("When AWS_ACCESS_KEY_ID exists!", func(t *testing.T) {
			// Put the env var into the system
			os.Setenv("AWS_ACCESS_KEY_ID", AWS_ACCESS_KEY_ID_VAL)

			Access_Key_ID, err := GetEnv("AWS_ACCESS_KEY_ID")
			if assert.NoError(t, err) {
				assert.Equal(t, Access_Key_ID, AWS_ACCESS_KEY_ID_VAL)
			}
		})

		t.Run("When AWS_SECRET_ACCESS_KEY exists!", func(t *testing.T) {
			// Put the env var into the system
			os.Setenv("AWS_SECRET_ACCESS_KEY", AWS_SECRET_ACCESS_KEY_VAL)

			Secret_Access_Key, err := GetEnv("AWS_SECRET_ACCESS_KEY")
			if assert.NoError(t, err) {
				assert.Equal(t, Secret_Access_Key, AWS_SECRET_ACCESS_KEY_VAL)
			}
		})

		t.Run("When AWS_SESSION_TOKEN exists!", func(t *testing.T) {
			// Put the env var into the system
			os.Setenv("AWS_SESSION_TOKEN", AWS_SESSION_TOKEN_VAL)

			Session_token, err := GetEnv("AWS_SESSION_TOKEN")
			if assert.NoError(t, err) {
				assert.Equal(t, Session_token, AWS_SESSION_TOKEN_VAL)
			}
		})
	})
}
