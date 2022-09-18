package main

import (
	"path/filepath"
	"runtime"
)

func getBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}
