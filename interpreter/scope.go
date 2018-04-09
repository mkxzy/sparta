package interpreter

/**
作用域
 */
type Scope interface {

	GetScopeName() string

	GetEnclosingScope() Scope

	Define(symbol Symbol)

	Resolve(name string) Symbol
}