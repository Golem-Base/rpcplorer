package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Golem-Base/rpcplorer/templates"
	"github.com/urfave/cli/v2"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &cli.App{
		Name:  "myapp",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			mux := http.NewServeMux()
			mux.HandleFunc(
				"/",
				func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path != "/" {
						http.NotFound(w, r)
						return
					}

					component := templates.Index()
					err := component.Render(r.Context(), w)
					if err != nil {
						log.Error("failed to render template", "error", err)
						http.Error(w, "Internal Server Error", http.StatusInternalServerError)
						return
					}
				},
			)

			log.Info("Starting server on :8080")
			return http.ListenAndServe(":8080", mux)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error("error running app", "error", err)
		os.Exit(1)
	}
}
