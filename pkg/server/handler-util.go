// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"log"
	"net/http"

	"github.com/murosan/shogi-proxy-server/pkg/msg"
)

func (s *Server) Handling(meth string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Access-Log. [uri:" + r.RequestURI + "] [method:" + r.Method + "] [addr:" + r.RemoteAddr + "] [ua:" + r.Header.Get("user-agent") + "]")

		if r.Method != meth {
			log.Printf("Error: %s, URI: %s\n", msg.MethodNotAllowed, r.RequestURI)
			http.Error(w, msg.MethodNotAllowed.Error(), http.StatusBadRequest)
			return
		}

		h(w, r)
	}
}
