package godotenv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Load reads environment variables from .env files and sets values
// only for keys that are not already present in the process environment.
func Load(filenames ...string) error {
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}

	for _, name := range filenames {
		if err := loadFile(name); err != nil {
			return err
		}
	}
	return nil
}

func loadFile(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, val, ok := strings.Cut(line, "=")
		if !ok {
			return fmt.Errorf("invalid .env line: %q", line)
		}
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		val = strings.Trim(val, `"`)
		val = strings.Trim(val, `'`)

		if key == "" {
			return fmt.Errorf("invalid empty env key")
		}
		if _, exists := os.LookupEnv(key); !exists {
			if err := os.Setenv(key, val); err != nil {
				return err
			}
		}
	}

	return s.Err()
}
