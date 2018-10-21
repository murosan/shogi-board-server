// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/murosan/shogi-proxy-server/pkg/msg"
)

func (s *Server) Handling(meth string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("AccessLog. [uri:%s] [method:%s] [addr:%s] [ua:%s]", r.RequestURI, r.Method, r.RemoteAddr, r.Header.Get("user-agent")))

		if r.Method != meth {
			log.Printf("Error: %s, URI: %s\n", msg.MethodNotAllowed, r.RequestURI)
			http.Error(w, msg.MethodNotAllowed.Error(), http.StatusBadRequest)
			return
		}

		h(w, r)
	}
}

func (s *Server) ContentTypeCheck(tpe string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")
		if ct != tpe {
			m := fmt.Sprintf("Content-Type must be %s, but got %s", tpe, ct)
			http.Error(w, m, http.StatusBadRequest)
			log.Println(m)
			return
		}

		h(w, r)
	}
}
