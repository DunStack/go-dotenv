package dotenv_test

import (
	"testing"

	"github.com/dunstack/go-dotenv"
	"github.com/dunstack/go-jest"
)

func TestDefaultConfig(t *testing.T) {
	dotenv.Load(".env_test")

	jest.Test(t, func(j *jest.J[jest.BuiltinMatcher]) {
		j.Expect(dotenv.Get("PORT")).ToBe("8080")
		j.Expect(dotenv.Get("FOO")).ToBe("foo")
		j.Expect(dotenv.Get("BAR")).ToBe("")
	})
}
