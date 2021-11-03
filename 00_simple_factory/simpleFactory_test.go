package simpleFactory

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Reverie")
	if s != "Hi, Reverie" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Reverie")
	if s != "Hello, Reverie" {
		t.Fatal("Type1 test fail")
	}
}
