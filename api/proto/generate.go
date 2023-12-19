//go:generate go run github.com/bufbuild/buf/cmd/buf lint
//go:generate go run github.com/bufbuild/buf/cmd/buf format --diff --exit-code
//go:generate go run github.com/bufbuild/buf/cmd/buf breaking . --against https://github.com/traPtitech/Emoine_R.git#subdir=api/proto
//go:generate go run github.com/bufbuild/buf/cmd/buf generate
package proto

import (
  // NOTE: go.modでbufのバージョンを管理する
  _ "github.com/bufbuild/buf/cmd/buf"
)
