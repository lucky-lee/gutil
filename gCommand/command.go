package gCommand

import (
	"os/exec"
	"io/ioutil"
	"github.com/lucky-lee/gutil/gLog"
)

func Exec(cmd *exec.Cmd) {
	out, err := cmd.StdoutPipe()

	if err != nil {
		gLog.E("ExecCommandErr", err)
	}

	defer out.Close()

	//运行
	if err := cmd.Start(); err != nil {
		gLog.E("ExecCommandErr", err)
	}

	//读取输出结果
	outBytes, err := ioutil.ReadAll(out)

	if err != nil {
		gLog.E("ExecCommandErr", err)
	}

	gLog.I("ExecCommandOut", string(outBytes))
}
