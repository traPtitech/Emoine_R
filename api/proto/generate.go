//go:generate go run github.com/bufbuild/buf/cmd/buf@latest lint
//go:generate go run github.com/bufbuild/buf/cmd/buf@latest format --diff --exit-code
//go:generate go run github.com/bufbuild/buf/cmd/buf@latest breaking . --against https://github.com/traPtitech/Emoine_R.git#subdir=api/proto
//go:generate go run github.com/bufbuild/buf/cmd/buf@latest generate
package main
