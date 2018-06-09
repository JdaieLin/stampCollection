package util

import (
	"net"
	"net/http"
	"strconv"
)

func StartHttp(l net.Listener, handle http.Handler) (err error) {
	var server = http.Server{
		Handler: handle,
	}
	err = server.Serve(l)
	return
}

func TryDial(host string, port int) (err error) {
	var conn net.Conn
	if conn, err = net.Dial("tcp", net.JoinHostPort(host, strconv.Itoa(port))); err != nil {
		return
	}
	defer conn.Close()
	return
}
