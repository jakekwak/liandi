// LianDi - 链滴笔记，连接点滴
// Copyright (c) 2020-present, b3log.org
//
// LianDi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package cmd

import (
	"github.com/88250/liandi/kernel/model"
	execstd "os/exec"
)

type exec struct {
	*BaseCmd
}

func (cmd *exec) Exec() {
	ret := model.NewCmdResult(cmd.Name(), cmd.id)
	execPath := cmd.param["bin"].(string)
	args := cmd.param["args"].([]interface{})
	var argsStrs []string
	for _, arg := range args {
		argsStrs = append(argsStrs, arg.(string))
	}
	go execstd.Command(execPath, argsStrs...).Run()
	cmd.Push(ret.Bytes())
}

func (cmd *exec) Name() string {
	return "exec"
}