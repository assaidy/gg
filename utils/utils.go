package ggu

import "github.com/assaidy/gg"

// IfElse returns the appropriate value based on a boolean condition.
//
// This generic function is useful for inline conditional expressions in
// builder-style code where you need to choose between two values without
// breaking the chain of method calls.
//
// Example:
//
//	div := Div(KV{"class": IfElse(isActive, "active", "inactive")})
//
//	Body(
//		IfElse(isAdmin,
//			Div("Admin content"),
//			P("Regular user content"),
//		),
//	)
func IfElse[T any](condition bool, result, alternative T) T {
	if condition {
		return result
	}
	return alternative
}

// Conditionally returns a Node based on a boolean condition.
//
// This function returns an empty Node (not nil) when the
// condition is false, which prevents nil pointer issues when building
// DOM trees.
//
// Example:
//
//	Body(
//		If(showHeader, Header(...)),
//		Main(...),
//	)
func If(condition bool, result gg.Node) gg.Node {
	if condition {
		return result
	}
	return gg.Empty()
}

// Repeat generates multiple Nodes by calling a function n times.
//
// The provided function is called exactly n times, and each resulting Node
// is aggregated into a single container Node. Using a function ensures each
// Node instance is unique (important for elements with mutable state).
//
// Example:
//
//	Ul(
//		Repeat(5, func() gg.Node {
//			return Li("List item")
//		}),
//	)
func Repeat(n int, f func() gg.Node) gg.Node {
	result := gg.Empty()
	for range n {
		result.Children = append(result.Children, f())
	}
	return result
}

// Map transforms a slice of items into Nodes by applying a function to each element.
//
// Each element in the input slice is transformed using the provided function, and
// all resulting Nodes are aggregated into a single container Node.
//
// Example:
//
//	items := []string{"Apple", "Banana", "Cherry"}
//	Ul(
//		Map(items, func(item string) gg.Node {
//			return Li(item)
//		}),
//	)
func Map[T any](input []T, f func(T) gg.Node) gg.Node {
	result := gg.Empty()
	for _, item := range input {
		result.Children = append(result.Children, f(item))
	}
	return result
}
