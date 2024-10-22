package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/uptrace/bun"
	page "github.com/vinibgoulart/gitbook-llm/packages/page/handler"
)

func ServerInit(db *bun.DB) func(ctx context.Context, waitGroup *sync.WaitGroup) {
	return func(ctx context.Context, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()

		router := chi.NewRouter()

		router.Use(middleware.RequestID)
		router.Use(middleware.RealIP)
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		router.Use(JsonContentTypeMiddleware)
		router.Use(middleware.Timeout(30 * time.Second))

		router.Get("/status", func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte("OK"))
		})
		router.Route("/ai", func(r chi.Router) {
			r.Post("/page", page.AiPromptPost(&ctx, db))
		})

		server := &http.Server{
			Addr:    ":8080",
			Handler: router,
		}

		go func() {
			fmt.Println("HTTP server listening on :8080")
			server.ListenAndServe()
		}()

		<-ctx.Done()

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatalf("HTTP server shutdown error: %s", err)
		}
	}
}