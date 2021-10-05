package invoker

import "github.com/Allenxuxu/gev/connection"

type AuthInvoker struct {
	password string
	c        *connection.Connection
}
