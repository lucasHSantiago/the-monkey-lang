package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	b1 := &Boolean{Value: true}
	b2 := &Boolean{Value: true}

	diff1 := &Boolean{Value: false}
	diff2 := &Boolean{Value: false}

	if b1.HashKey() != b2.HashKey() {
		t.Errorf("boolean with same value have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("boolean with same value have different hash keys")
	}

	if b1.HashKey() == diff1.HashKey() {
		t.Errorf("boolean with different value have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	int1 := &Integer{Value: 1}
	int2 := &Integer{Value: 1}

	diff1 := &Integer{Value: 2}
	diff2 := &Integer{Value: 2}

	if int1.HashKey() != int2.HashKey() {
		t.Errorf("integer with same value have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("integer with same value have different hash keys")
	}

	if int1.HashKey() == diff1.HashKey() {
		t.Errorf("integer with different value have same hash keys")
	}
}
