package server

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

const nRetries = 6

type Server struct {
	db     *pgx.Conn
	mux    *http.ServeMux
	logger *zap.Logger
}

func New(ctx context.Context) (s *Server, err error) {
	lg, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	connection, err := pgx.Connect(ctx, os.Getenv("DB_CONN"))
	if err != nil {
		return nil, err
	}
	sMux := http.NewServeMux()
	return &Server{
		db:     connection,
		logger: lg,
		mux:    sMux,
	}, nil
}

func (s *Server) HealthCheck(ctx context.Context) {
	n := nRetries
	ticker := time.Tick(time.Second * 15)
	for range ticker {
		err := s.db.Ping(ctx)
		if err != nil {
			n--
			s.logger.Warn("Healthcheck failed", zap.Int("attempts left", n))
		} else {
			n = nRetries
		}
	}
}

func (s *Server) Run(ctx context.Context) {
	s.mux.HandleFunc("/fact", s.handleFact)
	s.mux.HandleFunc("/fact/", s.handleRoot)
	s.mux.HandleFunc("/", s.handleRoot)
	if err := http.ListenAndServe(":80", s.mux); err != nil {
		s.db.Close(ctx)
		s.logger.Fatal("Server down: ", zap.Error(err))
	}
	go s.HealthCheck(ctx)
}
