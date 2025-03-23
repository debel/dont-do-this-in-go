package main

type (
	x struct{ s string }
	y struct{ s string }   // try adding another field
	z = struct{ s string } // type alias // HL
)

func f(p struct{ s string }) {}
func g(p x)                  {}
func d(p z)                  {}

func do() {
	f(x{})
	f(y{})

	g(x{})
	g(y{}) // error: cannot use y as x // HL

	d(x{})
	d(y{})
}
