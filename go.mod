module github.com/mkxzy/sparta

go 1.18

require (
	github.com/antlr/antlr4/runtime/Go/antlr v1.0.0
	github.com/op/go-logging v1.0.0
)

replace github.com/antlr/antlr4/runtime/Go/antlr => ./vendor/github.com/antlr/antlr4/runtime/Go/antlr
replace github.com/op/go-logging => ./vendor/github.com/op/go-logging