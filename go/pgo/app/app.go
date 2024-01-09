package app

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"

	"log/slog"
)

var (
	ErrServerClosed = errors.New("server closed")
)

type App struct {
	srv *http.Server
	log *slog.Logger
}

func New(log *slog.Logger) *App {
	return &App{
		srv: makeServer(log),
		log: log,
	}
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		http.ListenAndServe("localhost:9091", nil)
	}()

	var (
		errChan = make(chan error, 1)
	)

	go func() {
		if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}

		errChan <- ErrServerClosed
	}()

	select {
	case <-ctx.Done():
		cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		a.Close(cctx)
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

type noise struct {
	log *slog.Logger
	m   map[string][]string
}

func (n *noise) makeNoise(wg *sync.WaitGroup, count int) {
	m := make([]string, count)
	for i := 0; i < count; i++ {
		m[i] = "some noise"
	}
	wg.Done()
}

func addNoise(ctx context.Context, log *slog.Logger, count int) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			count -= 1
			c := rand.Intn(1000000-300000) + 300000
			wg := sync.WaitGroup{}
			for i := 0; i < count; i++ {
				wg.Add(1)
				go func(wg *sync.WaitGroup, c int) {
					n := noise{log: log, m: make(map[string][]string)}
					n.makeNoise(wg, c)
				}(&wg, c)
			}
			wg.Wait()
			if count < 0 {
				return
			}
		}
	}
}

func (a *App) Close(ctx context.Context) error {
	errChan := make(chan error, 1)

	go func() {
		err := a.srv.Shutdown(ctx)
		errChan <- err
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

func makeServer(log *slog.Logger) *http.Server {
	return &http.Server{
		Addr:    ":9090",
		Handler: handlers(log),
	}
}

func handlers(log *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/noise", func(w http.ResponseWriter, r *http.Request) {
		addNoise(r.Context(), log, 1)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	return mux
}
