package psutil

import (
	"fmt"
	"testing"
)

func Test_Taskstate(t *testing.T) {
	taskstate, err := TaskState(int32(1))
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Printf("taskstate=%v\n", taskstate.Mem)
}
