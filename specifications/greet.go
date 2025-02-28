package specifications

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("Mike")
	assert.NoError(t, err)
	assert.Equal(t, "Hello, Mike", got)
}
