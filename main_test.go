package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	out1 := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	if out1 != 5 {
		t.Errorf("got %d, want %d", out1, 5)
	}

	out := ladderLength("hot", "dog", []string{"hot", "dog", "dot"})
	if out != 3 {
		t.Errorf("got %d, want %d", out, 3)
	}

	second := ladderLength("a", "c", []string{"a", "b", "c"})
	if second != 2 {
		t.Errorf("got %d, want %d", second, 2)
	}

	third := ladderLength("hot", "dog", []string{"hot", "cog", "dog", "tot", "hog", "hop", "pot", "dot"})
	if third != 3 {
		t.Errorf("got %d, want %d", third, 3)
	}
}
