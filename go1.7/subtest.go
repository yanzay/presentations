package main

import "testing"

func TestFoo(t *testing.T) {
	t.Run("subtest1", func(t *testing.T) {
		// ...
	})
	t.Run("subtest2", func(t *testing.T) {
		// ...
	})
	t.Run("subtest3", func(t *testing.T) {
		// ...
	})
}
