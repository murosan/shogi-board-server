// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import "net/http"

var (
	ConnectPath     = "/connect"
	QuitPath        = "/quit"
	SetPositionPath = "/position/set"
	StudyInitPath   = "/study/init"
	StudyStartPath  = "/study/start"
	StudyStopPath   = "/study/stop"
)

type ShogiProxyServer interface {
	ServeHome(http.ResponseWriter, *http.Request)
	Connect(http.ResponseWriter, *http.Request)
	Close(http.ResponseWriter, *http.Request)
	OptionList(http.ResponseWriter, *http.Request)
	OptionSet(http.ResponseWriter, *http.Request)
	PositionSet(http.ResponseWriter, *http.Request)
	StudyStart(http.ResponseWriter, *http.Request)
	StudyStop(http.ResponseWriter, *http.Request)
	StudyCandidateList(http.ResponseWriter, *http.Request)
	AnalyzeInit(http.ResponseWriter, *http.Request)
	AnalyzeStart(http.ResponseWriter, *http.Request)
}
