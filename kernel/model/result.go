// LianDi - 链滴笔记，连接点滴
// Copyright (c) 2020-present, b3log.org
//
// LianDi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package model

import (
	"encoding/json"
)

type PushMode int

const (
	PushModeBroadcast            PushMode = 0 // 广播
	PushModeSingleSelf           PushMode = 1 // 自我单播
	PushModeBroadcastExcludeSelf PushMode = 2 // 非自我广播
)

type Result struct {
	Cmd            string      `json:"cmd"`
	ReqId          float64     `json:"reqId"`
	SessionId      string      `json:"sid"`
	PushMode       PushMode    `json:"pushMode"`
	ReloadPushMode PushMode    `json:"reloadPushMode"`
	Callback       interface{} `json:"callback"`
	Code           int         `json:"code"`
	Msg            string      `json:"msg"`
	Data           interface{} `json:"data"`
}

func NewResult() *Result {
	return &Result{Cmd: "",
		ReqId:    0,
		PushMode: 0,
		ReloadPushMode: 0,
		Callback: "",
		Code:     0,
		Msg:      "",
		Data:     nil}
}

func NewCmdResult(cmdName string, cmdId float64, pushMode, reloadPushMode PushMode) *Result {
	ret := NewResult()
	ret.Cmd = cmdName
	ret.ReqId = cmdId
	ret.PushMode = pushMode
	ret.ReloadPushMode = reloadPushMode
	return ret
}

func (r *Result) Bytes() []byte {
	ret, err := json.Marshal(r)
	if nil != err {
		Logger.Errorf("marshal result [%+v] failed [%s]", r, err)
	}
	return ret
}
