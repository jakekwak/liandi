// LianDi - 链滴笔记，连接点滴
// Copyright (c) 2020-present, b3log.org
//
// Lute is licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.

package main

type settheme struct {
	*BaseCmd
}

func (cmd *settheme) Exec() {
	ret := NewCmdResult(cmd.Name(), cmd.id)
	theme := cmd.param["theme"].(string)
	Conf.Theme = theme
	Conf.Save()
	ret.Data = theme
	Push(ret.Bytes())
}

func (cmd *settheme) Name() string {
	return "settheme"
}