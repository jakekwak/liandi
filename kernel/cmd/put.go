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
)

type put struct {
	*BaseCmd
}

func (cmd *put) Exec() {
	ret := model.NewCmdResult(cmd.Name(), cmd.id)
	url := cmd.param["url"].(string)
	url = model.NormalizeURL(url)
	p := cmd.param["path"].(string)
	content := cmd.param["content"].(string)
	err := model.Put(url, p, content)
	if nil != err {
		ret.Code = -1
		ret.Msg = err.Error()
	}
	cmd.Push(ret.Bytes())

	// 反向链接
	backlinks, err := model.Backlinks(url, p)
	if nil != err {
		backlinksEvent := model.NewCmdResult("backlinks", 0)
		backlinksEvent.Data = map[string]interface{}{
			"backlinks": backlinks,
		}
		model.BroadcastEvent(backlinksEvent)
	}

	// 关系图
	nodes, links := model.Graph("")
	graph := model.NewCmdResult("graph", 0)
	graph.Data = map[string]interface{}{
		"nodes": nodes,
		"links": links,
	}
	model.BroadcastEvent(graph)
}

func (cmd *put) Name() string {
	return "put"
}
