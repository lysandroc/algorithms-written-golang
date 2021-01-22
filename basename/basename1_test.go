package basename

import "testing"

func TestBasenameV1(t *testing.T) {
	compare := func(t *testing.T, want, got string) {
		t.Helper()
		if want != got {
			t.Errorf("want '%s' but got '%s'", want, got)
		}
	}

	t.Run("Expected 'a' when given 'a'", func(t *testing.T) {
		expected := "a"
		got := BasenameV1("a.go")
		compare(t, expected, got)
	})

	t.Run("Expected 'a' when given 'a.go'", func(t *testing.T) {
		expected := "a"
		got := BasenameV1("a.go")
		compare(t, expected, got)
	})

	t.Run("Expected 'a/b/c.go' when given 'c'", func(t *testing.T) {
		expected := "c"
		got := BasenameV1("a/b/c.go")
		compare(t, expected, got)
	})

	t.Run("Expected 'a/b.c.go' when given 'b.c'", func(t *testing.T) {
		expected := "b.c"
		got := BasenameV1("a/b.c.go")
		compare(t, expected, got)
	})

	t.Run("Expected 'a' when given 'a'", func(t *testing.T) {
		expected := "a"
		got := Basename("a.go")
		compare(t, expected, got)
	})

	t.Run("Expected 'a' when given 'a.go'", func(t *testing.T) {
		expected := "a"
		got := Basename("a.go")
		compare(t, expected, got)
	})

	t.Run("Expected 'a/b/c.go' when given 'c'", func(t *testing.T) {
		expected := "c"
		got := Basename("a/b/c.go")
		compare(t, expected, got)
	})

	t.Run("Expected 'a/b.c.go' when given 'b.c'", func(t *testing.T) {
		expected := "b.c"
		got := Basename("a/b.c.go")
		compare(t, expected, got)
	})
}
