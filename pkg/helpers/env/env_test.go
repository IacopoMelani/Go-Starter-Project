package env

import (
	"fmt"
	"os"
	"testing"
)

const envFileTest = ".env.test"

func TestEnv(t *testing.T) {

	_, err := os.OpenFile(envFileTest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}

	// TEST 1 OK
	if err := LoadEnvFile(envFileTest, false); err != nil {
		t.Fatal(err)
	}

	// TEST 2 ERR
	err = LoadEnvFile("env.err", false)
	if err == nil {
		t.Fatal("Invalid file name but no errors")
	}
	fmt.Println(err.Error())

	// TEST 3 PANIC
	defer func() {
		if p := recover(); p == nil {
			t.Fatal("No panic occurs")
		}
	}()
	LoadEnvFile("env.panic", true)
}
