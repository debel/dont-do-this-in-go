package main

func dont(f func(v any)) {}

func main() {
	defer dont(panic)
}
