package server

import (
	"context"
	"github.com/hyonosake/Random-Facts-Server/internal/types"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const nRetries = 6

type Server struct {
	db     *pgx.Conn
	mux    *http.ServeMux
	logger *zap.Logger
	cfg    *types.Env
}

func New(ctx context.Context) (s *Server, err error) {
	lg, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	sMux := http.NewServeMux()
	cfg, err := s.parseConfig()
	if err != nil {
		return nil, err
	}

	connection, err := pgx.Connect(ctx, "postgresql://localhost:5433/postgres")
	if err != nil {
		return nil, err
	}

	return &Server{
		db:     connection,
		logger: lg,
		mux:    sMux,
		cfg:    cfg,
	}, nil
}

func (s *Server) pingWithTimeout(ctx context.Context) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	err := s.db.Ping(timeoutCtx)
	return err
}

func (s *Server) Run(ctx context.Context) {
	s.mux.HandleFunc("/fact", s.handleFact)
	s.mux.HandleFunc("/fact/", s.handleFactId)
	s.mux.HandleFunc("/", s.handleRoot)
	if err := http.ListenAndServe(":80", s.mux); err != nil {
		s.db.Close(ctx)
		s.logger.Fatal("Server down: ", zap.Error(err))
	}
	go s.HealthCheck(ctx)
}

func (s *Server) HealthCheck(ctx context.Context) {
	n := nRetries
	ticker := time.Tick(time.Second * 15)
	for range ticker {
		if s.pingWithTimeout(ctx) != nil {
			n--
			s.logger.Warn("Healthcheck failed", zap.Int("attempts left", n))
		} else {
			n = nRetries
		}
		if n == 0 {
			s.logger.Fatal("Multiple HealthChecks failed, exiting")
		}
	}
}

func (s *Server) parseConfig() (*types.Env, error) {
	var e *types.Env

	//path := os.Getenv("PWD") + "/values.yaml"
	//yamlFile, err := ioutil.ReadFile(path)
	//if err != nil {
	//	return nil, err
	//}
	//err = yaml.Unmarshal(yamlFile, e)
	return e, nil
}
