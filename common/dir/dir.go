package dir

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	// project root for this project
	Root = filepath.Join(filepath.Dir(b), "../..")
)
