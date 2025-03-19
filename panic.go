package main

func dont(func(v any)) {}

func main() {
	defer dont(panic)
}
