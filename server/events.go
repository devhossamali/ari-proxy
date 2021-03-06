package server

import "github.com/devhossamali/ari"

func (s *Server) dialogsForEvent(e ari.Event) (ret []string) {
	for _, k := range e.Keys() {
		if k == nil {
			s.Log.Warn("received nil key for event", "event", e)
			continue
		}
		ret = append(ret, s.Dialog.List(k.Kind, k.ID)...)
	}
	return
}
