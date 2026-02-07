package pointer

import "testing"

func TestDoubleViaPointer(t *testing.T) {
	x := 7
	DoubleViaPointer(&x)
	if x != 14 {
		t.Fatalf("expected 14, got %d", x)
	}

	// nil darf nicht crashen
	DoubleViaPointer(nil)
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Fatalf("expected a=2 b=1, got a=%d b=%d", a, b)
	}

	// nil-Fälle: keine Änderung + kein Crash
	a, b = 3, 4
	Swap(nil, &b)
	if a != 3 || b != 4 {
		t.Fatalf("expected unchanged (nil case), got a=%d b=%d", a, b)
	}
}

func TestBirthday(t *testing.T) {
	p := Person{Name: "Max", Alter: 19}
	Birthday(&p)
	if p.Alter != 20 {
		t.Fatalf("expected 20, got %d", p.Alter)
	}

	// nil darf nicht crashen
	Birthday(nil)
}

func TestNewPerson(t *testing.T) {
	p := NewPerson("Lisa", 25)
	if p == nil {
		t.Fatalf("expected non-nil *Person")
	}
	if p.Name != "Lisa" || p.Alter != 25 {
		t.Fatalf("expected Lisa/25, got %s/%d", p.Name, p.Alter)
	}
}

func TestDoubleAll(t *testing.T) {
	s := []int{1, 2, 3, 4}
	DoubleAll(s)
	want := []int{2, 4, 6, 8}
	for i := range s {
		if s[i] != want[i] {
			t.Fatalf("at i=%d expected %d got %d", i, want[i], s[i])
		}
	}

	// nil/leer darf nicht crashen
	DoubleAll(nil)
	DoubleAll([]int{})
}

func TestSafeDeref(t *testing.T) {
	x := 99
	v, ok := SafeDeref(&x)
	if !ok || v != 99 {
		t.Fatalf("expected (99,true), got (%d,%v)", v, ok)
	}

	v, ok = SafeDeref(nil)
	if ok || v != 0 {
		t.Fatalf("expected (0,false), got (%d,%v)", v, ok)
	}
}

func TestSetThroughDoublePointer(t *testing.T) {
	x := 10
	p := &x
	pp := &p

	SetThroughDoublePointer(pp, 123)
	if x != 123 {
		t.Fatalf("expected 123, got %d", x)
	}

	// nil-Fälle: kein Crash
	SetThroughDoublePointer(nil, 5)

	var pNil *int
	pp2 := &pNil
	SetThroughDoublePointer(pp2, 7) // *pp2 == nil
}

func TestIntPtr(t *testing.T) {
	p := IntPtr(42)
	if p == nil {
		t.Fatalf("expected non-nil pointer")
	}
	if *p != 42 {
		t.Fatalf("expected 42, got %d", *p)
	}

	// Sicherstellen, dass es nicht auf eine gemeinsame globale Variable zeigt:
	p2 := IntPtr(42)
	if p == p2 {
		t.Fatalf("expected different pointers (distinct allocations), got same address")
	}
}
