package main

import "testing"

func TestFoo(t *testing.T) {
	result := Foo()
	if result != "foo" {

		t.Error("intentional Error 1")
	}
}
func TestBar(t *testing.T) {
	result := Bar()
	if result != "bar" {
		t.Errorf("excepting bar, got %s", result)
	}
}
func TestQuiz(t *testing.T) {
	result := Quiz()
	if result != "quiz" {
		t.Error("intentional Error 2")
	}
}
