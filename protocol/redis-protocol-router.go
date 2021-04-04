package protocol

var routeInfo map[string]initFunc

func GetCmdInitFunc(operatorType string) initFunc {
	f := routeInfo[operatorType]
	if f == nil {
		return defaultCmdInit
	}
	return f
}

type initFunc func(cmd *Cmd) CmdFunc

func init() {
	routeInfo = make(map[string]initFunc)
	routeInfo["ping"] = pingCmdInit
	routeInfo["get"] = getCmdInit
	routeInfo["set"] = setCmdInit
	routeInfo["append"] = appendCmdInit
	routeInfo["del"] = delCmdInit
}

func appendCmdInit(cmd *Cmd) CmdFunc {
	return &AppendCmd{Cmd: cmd}
}

func setCmdInit(cmd *Cmd) CmdFunc {
	return &SetCmd{Cmd: cmd}
}

func getCmdInit(cmd *Cmd) CmdFunc {
	return &GetCmd{Cmd: cmd}
}

func pingCmdInit(cmd *Cmd) CmdFunc {
	return &PingCmd{Cmd: cmd}
}

func delCmdInit(cmd *Cmd) CmdFunc {
	return &DelCmd{Cmd: cmd}
}

func defaultCmdInit(cmd *Cmd) CmdFunc {
	return cmd
}
