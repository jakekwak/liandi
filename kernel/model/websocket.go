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
	"sync"

	"gopkg.in/olahol/melody.v1"
)

var (
	sessions     = map[string]*melody.Session{}
	sessionsLock = &sync.Mutex{}
)

func AddPushChan(session *melody.Session) {
	sessionsLock.Lock()
	defer sessionsLock.Unlock()

	id, _ := session.Get("id")
	sessions[id.(string)] = session
}

func RemovePushChan(session *melody.Session) {
	sessionsLock.Lock()
	defer sessionsLock.Unlock()

	id, _ := session.Get("id")
	delete(sessions, id.(string))
}

func ClosePushChan(id string) {
	for _, session := range sessions {
		sid, _ := session.Get("id")
		if sid == id {
			session.CloseWithMsg([]byte("  close websocket"))
			RemovePushChan(session)
			return
		}
	}
}

func PushEvent(event *Result) {
	msg := event.Bytes()
	mode := event.PushMode
	if "reload" == event.Cmd {
		mode = event.ReloadPushMode
	}
	switch mode {
	case PushModeBroadcast:
		Broadcast(msg)
	case PushModeSingleSelf:
		Single(msg, event.SessionId)
	case PushModeBroadcastExcludeSelf:
		BroadcastOthers(msg, event.SessionId)
	}
}

func Single(msg []byte, sid string) {
	for _, session := range sessions {
		id, _ := session.Get("id")
		if id == sid {
			session.Write(msg)
			return
		}
	}
}

func Broadcast(msg []byte) {
	for _, session := range sessions {
		session.Write(msg)
	}
}

func BroadcastOthers(msg []byte, excludeSID string) {
	for _, session := range sessions {
		id, _ := session.Get("id")
		if id == excludeSID {
			continue
		}
		session.Write(msg)
	}
}
