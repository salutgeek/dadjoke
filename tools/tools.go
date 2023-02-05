//go:build tools
// +build tools

package tools

// tool dependencies
import (
	_ "github.com/frapposelli/wwhrd"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "mvdan.cc/gofumpt"
)

// If you want to use tools, please run the following command:
//  go generate ./tools.go
//
//go:generate go build -v -o=./bin/wwhrd github.com/frapposelli/wwhrd
//go:generate go build -v -o=./bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go build -v -o=./bin/gofumpt mvdan.cc/gofumpt
