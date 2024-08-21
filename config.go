package dotenv

import (
	"fmt"
	"os"
	"regexp"
)

var pattern = regexp.MustCompile(`(?m)^(?P<key>.+?)\s*=\s*(?P<value>.+?)$`)

type Config struct {
	Names []string
}

func (c *Config) Load() {
	for _, name := range c.Names {
		if b, err := os.ReadFile(name); err != nil {
			fmt.Println(err)
		} else {
			for _, loc := range pattern.FindAllSubmatchIndex(b, -1) {
				key, value := b[loc[2]:loc[3]], b[loc[4]:loc[5]]
				if err := os.Setenv(string(key), string(value)); err != nil {
					fmt.Println()
				}
			}
		}

	}
}

func (c *Config) Get(key string, defaultValue ...string) string {
	v := os.Getenv(key)
	if len(defaultValue) > 0 && v == "" {
		return defaultValue[0]
	}
	return v
}
