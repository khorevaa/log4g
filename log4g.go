package log4g

import (
	"github.com/khorevaa/log4g/api"
	"github.com/khorevaa/log4g/internal"
)

// GetLogger return a instance that implements Logger interface specified by name,
// name like a/b/c or a.b.c ,
// Logger named 'a/b' is the parent of Logger named 'a/b/c'
func GetLogger(name string) api.Logger {
	return internal.GetLogger(name)
}

// GetManager return a instance that implements Manager interface
func GetManager() api.Manager {
	return internal.GetManager()
}
