//go:generate go run github.com/bufbuild/buf/cmd/buf@latest lint
//go:generate go run github.com/bufbuild/buf/cmd/buf@latest format --diff --exit-code
//go:generate go run github.com/bufbuild/buf/cmd/buf@latest generate
package main
