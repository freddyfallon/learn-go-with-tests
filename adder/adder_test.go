package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Adds two numbers together", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		if sum != expected {
			t.Errorf("Expected: %d, but received %d", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
