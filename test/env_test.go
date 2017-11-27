package test

import (
	"env"
	"testing"
)

func TestEnvWithDefaultFile(t *testing.T) {
	if !(env.Env("PORT") == "8080") {
		t.Error("PORT != 8080 in .env")
	}
}

func TestEnvWithSpecifiedFile(t *testing.T) {
	env.SetConfigFilePath("./.anotherEnv")
	if !(env.Env("HOST") == "http://127.0.0.1") {
		t.Error("HOST != http://127.0.0.1 in .anotherEnv")
	}
}
