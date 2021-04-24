package know_accumulation

import (
	"fmt"
	"os"
	"testing"
)

func Test_process_kill(t *testing.T) {

	proc,err := os.FindProcess(3909)
	if err != nil {
		panic(err)
	}

	err = proc.Kill()
	if err != nil {
		panic(err)
	}

}


func Test_process_start(t *testing.T) {
	processName := "go-create-thread"
	argv := make([]string,0)
	argv = append(argv,processName)
	proc, err := os.StartProcess("/Users/huangbocai/go/src/t-redis-like",argv,&os.ProcAttr{Files:[]*os.File{os.Stdin,os.Stdout,os.Stderr}})
	if err != nil {
		panic(err)
	}
	fmt.Println(proc.Pid)

	for true {
		fmt.Printf("\t")
	}
}