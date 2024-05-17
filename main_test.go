package xenv_test

import (
	"os"
	"testing"

	"github.com/dev3mike/go-xenv"
)

type Environment struct {
	Host       string `json:"HOST" validators:"required,minLength:3,maxLength:50"`
	AdminEmail string `json:"ADMIN_EMAIL" validators:"email"`
	Code       string `json:"CODE" transformers:"uppercase"`
}

func TestEnvironmentValidation(t *testing.T) {
	// Test valid environment setup
	t.Run("ValidEnvironment", func(t *testing.T) {
		os.Setenv("HOST", "example.com")
		os.Setenv("ADMIN_EMAIL", "admin@example.com")
		os.Setenv("CODE", "abc")
		defer os.Unsetenv("HOST")
		defer os.Unsetenv("ADMIN_EMAIL")
		defer os.Unsetenv("CODE")

		env := &Environment{}
		err := xenv.ValidateEnv(env)
		if err != nil {
			t.Errorf("ValidateEnv() error = %v, wantErr %v", err, nil)
		}
		if env.Host != "example.com" {
			t.Errorf("Expected HOST=example.com, got HOST=%v", env.Host)
		}
		if env.AdminEmail != "admin@example.com" {
			t.Errorf("Expected ADMIN_EMAIL=admin@example.com, got ADMIN_EMAIL=%v", env.AdminEmail)
		}
		if env.Code != "ABC" {
			t.Errorf("Expected CODE=ABC, got CODE=%v", env.Code)
		}
	})

	// Test host validation failures
	t.Run("HostValidationFailures", func(t *testing.T) {
		// Host too short
		os.Setenv("HOST", "ex")
		os.Setenv("ADMIN_EMAIL", "admin@example.com")
		defer os.Unsetenv("HOST")
		defer os.Unsetenv("ADMIN_EMAIL")

		env := &Environment{}
		err := xenv.ValidateEnv(env)
		if err == nil {
			t.Error("ValidateEnv() expected to fail due to HOST being too short")
		}

		// Host too long
		longHost := "a"
		for len(longHost) <= 51 {
			longHost += "a"
		}
		os.Setenv("HOST", longHost)
		err = xenv.ValidateEnv(env)
		if err == nil {
			t.Error("ValidateEnv() expected to fail due to HOST being too long")
		}
	})

	// Test email validation failures
	t.Run("EmailValidationFailures", func(t *testing.T) {
		// Invalid email
		os.Setenv("HOST", "example.com")
		os.Setenv("ADMIN_EMAIL", "admin@")
		defer os.Unsetenv("HOST")
		defer os.Unsetenv("ADMIN_EMAIL")

		env := &Environment{}
		err := xenv.ValidateEnv(env)
		if err == nil {
			t.Error("ValidateEnv() expected to fail due to invalid ADMIN_EMAIL")
		}
	})

	// Test required field missing
	t.Run("RequiredFieldMissing", func(t *testing.T) {
		// Missing required HOST
		os.Unsetenv("HOST")
		os.Setenv("ADMIN_EMAIL", "admin@example.com")
		defer os.Unsetenv("ADMIN_EMAIL")

		env := &Environment{}
		err := xenv.ValidateEnv(env)
		if err == nil {
			t.Error("ValidateEnv() expected to fail due to missing required HOST")
		}
	})
}
