package dotenv_test

import (
	"testing"

	"github.com/dunstack/go-dotenv"
	"github.com/dunstack/go-jest"
)

func TestConfigLoad(t *testing.T) {
	c := dotenv.Config{
		Names: []string{".env_test"},
	}
	c.Load()

	jest.Test(t, func(j *jest.J[jest.BuiltinMatcher]) {
		j.Expect(c.Get("PORT")).ToBe("8080")
		j.Expect(c.Get("FOO")).ToBe("foo")
		j.Expect(c.Get("BAR")).ToBe("")
		j.Expect(c.Get("BAR", "bar")).ToBe("bar")
	})
}
