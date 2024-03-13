package generics

import (
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integer", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
}
