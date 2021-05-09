package storage

import (
	"context"
)

type ErrorInfo string

const (
	DelErrorText       = ErrorInfo("engine: %v , operator del failed , error info: %v .")
	AppendGetErrorText = ErrorInfo("engine: %v , operator append failed in get , error info: %v .")
	AppendSetErrorText = ErrorInfo("engine: %v , operator append failed in set , error info: %v .")
	GetErrorText       = ErrorInfo("engine: %v , operator get failed , error info: %v .")
	SetErrorText       = ErrorInfo("engine: %v , operator set failed , error info: %v .")
)

type Storage interface {
	Del(context context.Context, key []byte) error
	Append(context context.Context, key []byte, value []byte) error
	Get(context context.Context, key []byte) ([]byte, error)
	Set(context context.Context, key []byte, value []byte) error
	Close()
}
