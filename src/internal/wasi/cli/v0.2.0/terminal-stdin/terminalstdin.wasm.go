// Code generated by wit-bindgen-go. DO NOT EDIT.

package terminalstdin

import (
	"internal/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:cli@0.2.0".

//go:wasmimport wasi:cli/terminal-stdin@0.2.0 get-terminal-stdin
//go:noescape
func wasmimport_GetTerminalStdin(result *cm.Option[TerminalInput])