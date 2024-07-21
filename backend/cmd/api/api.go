package api

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/ajtroup1/speakeasy/service/block"
	"github.com/ajtroup1/speakeasy/service/channel"
	"github.com/ajtroup1/speakeasy/service/friend"
	"github.com/ajtroup1/speakeasy/service/message"
	"github.com/ajtroup1/speakeasy/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr   string
	db     *sql.DB
	Router *mux.Router
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	friendStore := friend.NewStore(s.db)
	blockStore := block.NewStore(s.db)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore, friendStore, blockStore)
	userHandler.RegisterRoutes(subrouter)

	channelStore := channel.NewStore(s.db)
	channelHandler := channel.NewHandler(channelStore, userStore)
	channelHandler.RegisterRoutes(subrouter)

	messageStore := message.NewStore(s.db)
	messagehandler := message.NewHandler(messageStore, userStore)
	messagehandler.RegisterRoutes(subrouter)

	s.Router = router

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	addr := listener.Addr().String()
	ip := strings.Split(addr, ":")[0]
	if ip == "::" || ip == "" {
		ip = "127.0.0.1"
	}
	log.Printf("Server listening on %s:%s", ip, s.addr)

	return http.Serve(listener, s.Router)
}
