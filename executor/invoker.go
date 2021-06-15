package executor

type Invoker interface {
	invoke(ctx *InvokerContext) *Result
}
