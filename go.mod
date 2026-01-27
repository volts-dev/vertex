module github.com/volts-dev/vertex

go 1.24.0

replace github.com/volts-dev/ccacher => ../cacher

replace github.com/volts-dev/lexer => ../lexer

require (
	github.com/expr-lang/expr v1.17.7
	github.com/volts-dev/cacher v0.0.0-20260108163127-b0c79465d9d6
	github.com/volts-dev/lexer v0.0.0-20251110194144-a67bf55d1b63
	github.com/volts-dev/utils v0.0.0-20241206111447-ee54d4e2c42c
)

require (
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/volts-dev/dataset v0.0.0-20251225093026-6d6fb8739587 // indirect
)
