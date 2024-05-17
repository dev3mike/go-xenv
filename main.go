package xenv

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/dev3mike/go-xmapper"
)

func LoadEnvFile(envFilePath string) error {
	file, err := os.Open(envFilePath)
	if err != nil {
		return fmt.Errorf("error opening .env file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("bad line format: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Inject into environment
		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("error setting environment variable: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading .env file: %w", err)
	}

	return nil
}

func ValidateEnv(envStruct interface{}) error {
	err := mapEnvToStruct(envStruct)
	if err != nil {
		return fmt.Errorf("error mapping env to struct: %w", err)
	}

	if err := xmapper.ValidateStruct(envStruct); err != nil {
		return fmt.Errorf("error validating env: %w", err)
	}

	return nil
}

func mapEnvToStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct")
	}

	v = v.Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		envValue := os.Getenv(jsonTag)
		if envValue == "" {
			continue
		}

		if v.Field(i).CanSet() {
			v.Field(i).SetString(envValue)
		}
	}

	return nil
}
