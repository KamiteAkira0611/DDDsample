package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mjibson/goon"
	"google.golang.org/appengine/datastore"
)

//  context.Background.WithValue(
// 	&http.contextKey{name:"http-server"},
// 	&http.Server{Addr:"127.0.0.1:16760",
// 	Handler:(http.HandlerFunc)(0x7e7a50),
// 	TLSConfig:(*tls.Config)(0xc00008ac00),
// 	ReadTimeout:0,
// 	ReadHeaderTimeout:0,
// 	WriteTimeout:0,
// 	IdleTimeout:0,
// 	MaxHeaderBytes:0,
// 	TLSNextProto:map[string]func(*http.Server, *tls.Conn, http.Handler){"h2":(func(*http.Server, *tls.Conn, http.Handler))(0x6888e0)},
// 	ConnState:(func(net.Conn, http.ConnState))(nil),
// 	ErrorLog:(*log.Logger)(nil),
// 	disableKeepAlives:0,
// 	inShutdown:0,
// 	nextProtoOnce:sync.Once{m:sync.Mutex{state:0, sema:0x0}, done:0x1},
// 	nextProtoErr:error(nil),
// 	mu:sync.Mutex{state:0, sema:0x0},
// 	listeners:map[*net.Listener]struct {}{(*net.Listener)(0xc0002b3400):struct {}{}},
// 	activeConn:map[*http.conn]struct {}{(*http.conn)(0xc0000bef00):struct {}{}},
// 	doneChan:(chan struct {})(nil), onShutdown:[]func(){(func())(0x6918c0)}}).WithValue(&http.contextKey{name:"local-addr"},
// 	&net.TCPAddr{IP:net.IP{0x7f, 0x0, 0x0, 0x1}, Port:16760, Zone:""}).WithCancel.WithCancel.WithValue((*string)(0xd8ca10), &internal.context{req:(*http.Request)(0xc000304000), outCode:0, outHeader:http.Header{}, outBody:[]uint8(nil), pendingLogs:struct { sync.Mutex; lines []*log.UserAppLogLine; flushes int }{Mutex:sync.Mutex{state:0, sema:0x0}, lines:[]*log.UserAppLogLine(nil), flushes:0}, apiURL:(*url.URL)(0xc0002e4080)}).WithValue(3, "twirp.task").WithValue(2, "API").WithValue(6, &internal.context{req:(*http.Request)(0xc000304000), outCode:0, outHeader:http.Header{}, outBody:[]uint8(nil), pendingLogs:struct { sync.Mutex; lines []*log.UserAppLogLine; flushes int }{Mutex:sync.Mutex{state:0, sema:0x0}, lines:[]*log.UserAppLogLine(nil), flushes:0}, apiURL:(*url.URL)(0xc0002e4080)}).WithValue(1, "ListTask")


type Task struct {
	ID int64 `datastore:"-" goon:"id"`
	Body string 
}

func TasksHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {

		case http.MethodGet:
			fmt.Fprintln(w, "Hello, GAE/go GetTask")
			n := goon.NewGoon(r)
			tasks := []Task{}
			q := datastore.NewQuery("Task")
		
			_, err := n.GetAll(q, &tasks)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "invalid request. %s", err.Error())
				return
			}
			json.NewEncoder(w).Encode(tasks)


		case http.MethodPost:
			fmt.Fprintln(w, "Hello, GAE/go CreateTask")
			n := goon.NewGoon(r)
		
			body := r.FormValue("body")
			g := &Task{Body: body}
		
			if body == "" {
				fmt.Fprintf(w, "invalid request. body is nil")
				return
			}
		
			key, err := n.Put(g)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "invalid request. %s", err.Error())
				return
			}
			json.NewEncoder(w).Encode(key)

		default:
			fmt.Fprint(w, "Method not allowed.\n")
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodDelete:
			fmt.Fprintln(w, "Hello, GAE/go DeleteTask")
			s := r.URL.Path[len("/task/"):]
			id, iderr := strconv.ParseInt(s, 10, 64)
			if iderr != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "invalid request. %s is not ID.\n", s)
				return
			}
		
			n := goon.NewGoon(r)
			task := &Task{ID: id}
			getErr := n.Get(task)
			if getErr != nil {
				fmt.Fprintf(w, "invalid request. getErr %s", getErr.Error())
				return
			}
			
			deleteErr := n.Delete(task)
			if deleteErr != nil {
				fmt.Fprintf(w, "invalid request. deleteErr %s", deleteErr.Error())
			}
		
			json.NewEncoder(w).Encode(task)


		default:
			fmt.Fprint(w, "Method not allowed.\n")
	}
}