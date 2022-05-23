package main

import "testing"

func TestGreetingSpecificJohn(t *testing.T) {
	greeting := CreateGreeting("John")
	if greeting != "Hello, John\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}

func TestGreetingSpecificDemo(t *testing.T) {
	greeting := CreateGreeting("Demo")
	if greeting != "Hello, Demo\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

func TestDummy(t *testing.T) {
	greeting := CreateGreeting("Dummy")
	if greeting != "Dummy\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
	x := 5
	if x != (10-5) {
		t.Errorf("want x: %v, got x: %v", 10-5, x)
	}
}

// func TestShowFailure(t *testing.T) {
// 	greeting := CreateGreeting("Demo1")
// 	if greeting != "Hello, Demo\n" {
// 		t.Errorf("Intentional failure. got: %s, want: %s.", greeting, "Hello, Demo\n")
// 	}
// }

func TestShowFailurePass(t *testing.T) {
	greeting := CreateGreeting("Hello, Demo\n")
	if greeting != "Hello, Demo\n" {
		t.Errorf("Intentional failure. got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

func TestGreetingDefault(t *testing.T) {
	greeting := CreateGreeting("")
	if greeting != "Hello, Guest\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Guest\n")
	}
}
 


