package logging

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type StatusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *StatusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *StatusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

func NewStatusWriter(w http.ResponseWriter) *StatusWriter {
	return &StatusWriter{
		ResponseWriter: w,
	}
}

func (w *StatusWriter) Status() int {
	return w.status
}

func LogHandler(w *StatusWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"method" : r.Method,
		"responseCode" : w.Status(),
		"host" : r.Host,
		"uri" : r.RequestURI,
		"proto" : r.Proto,
		"remoteAddr" : r.RemoteAddr,
	}).Info("Request: ")
}

