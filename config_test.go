package dotenv_test

import (
	"os"
	"testing"

	"github.com/dunstack/go-dotenv"
	"github.com/dunstack/go-jest"
)

func TestConfig(t *testing.T) {
	jest.Test(t, func(j *jest.J[jest.BuiltinMatcher]) {
		c := dotenv.Config{
			Names: []string{".env_test"},
		}
		c.Load()

		j.Expect(os.Getenv("PORT")).ToBe("8080")
		j.Expect(os.Getenv("FOO")).ToBe("foo")
		j.Expect(os.Getenv("BAR")).ToBe("")

		j.Expect(c.Get("PORT")).ToBe("8080")
		j.Expect(c.Get("BAR")).ToBe("")
		j.Expect(c.Get("BAR", "bar")).ToBe("bar")
	})

	os.Clearenv()

	jest.Test(t, func(j *jest.J[jest.BuiltinMatcher]) {
		c := dotenv.Config{
			Names:  []string{".env_test"},
			Prefix: "PREFIX_",
		}
		c.Load()

		j.Expect(os.Getenv("FOO")).ToBe("")
		j.Expect(os.Getenv("PREFIX_BAR")).ToBe("bar")

		j.Expect(c.Get("Foo")).ToBe("")
		j.Expect(c.Get("BAR")).ToBe("bar")
		j.Expect(c.Get("PREFIX_BAR")).ToBe("")
	})
}
