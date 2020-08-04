[![Run on Repl.it](https://repl.it/badge/github/matt4biz/go-class-closure)](https://repl.it/github/matt4biz/go-class-closure)


# Go class: closure examples
These programs demonstrate closures and possible issues when capturing a mutating variable by reference.

This issue is often illustrated with a goroutine example, but it actually has nothing to do with goroutines (or even Go); it's a potential issue in any language which has closures.

## Fib
The `fib/fib.go` example shows returning an anonymous function as a closure. It's called a "closure" because it closes over two variables declared outside itself, *a* and *b*.

Note that each closure returned from `fib()` has its own copy of those two variables, whose lifetime is the lifetime of the closure, even though their scope is limited to `fib()`.

	$ go run ./fib
	1 2 3 5 8 13 21
	1 2 3 5 8 13 21

## Bad
The example in `bad/bad.go` shows what happens when we store a closure to run later, where the closure has captured a mutating variable *i* (the loop index) by reference.

Each copy of the closure ends up reading and printing the same *i* with its final value.

	$ go run ./bad
	4 0xc000016088
	4 0xc000016088
	4 0xc000016088
	4 0xc000016088

## Good
In `good/good.go` we've modified the "bad" example above to create a variable local to the closure which captures the *value* of *i* each time through the loop, so each copy of the closure has its own variable and prints the expected result.

	$ go run ./good
	4 0xc0000b4010
	4 0xc0000b4028
	4 0xc0000b4048
	4 0xc0000b4058

