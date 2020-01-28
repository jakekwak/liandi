// LianDi - 链滴笔记，链接点滴
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

package util

import (
	"encoding/json"

	"github.com/88250/gulu"
)

type Result struct {
	Cmd string `json:"cmd"`
	*gulu.Result
}

func NewResult() *Result {
	return &Result{"", &gulu.Result{Code: 0, Msg: "", Data: nil}}
}

func NewCmdResult(cmd string) *Result {
	ret := NewResult()
	ret.Cmd = cmd

	return ret
}

func (r *Result) Bytes() []byte {
	ret, err := json.Marshal(r)
	if nil != err {
		Logger.Errorf("marshal result [%#v] failed [%s]", r, err)
	}

	return ret
}
