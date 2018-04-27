package symbol

/**
作用域接口
 */
type Scope interface {

	GetScopeName() string

	GetEnclosingScope() Scope

	Define(symbol Symbol)

	Resolve(name string) Symbol
}