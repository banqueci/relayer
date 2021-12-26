package relayer

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

type Settings struct {
	Host string `envconfig:"HOST" default:"0.0.0.0"`
	Port string `envconfig:"PORT" default:"7447"`
}

var s Settings
var log = zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{Out: os.Stderr})

var router = mux.NewRouter()

func Start(relay Relay) {
	Log = log.With().Str("name", relay.Name()).Logger()

	if err := relay.Init(); err != nil {
		Log.Fatal().Err(err).Msg("failed to start")
	}

	// NIP01
	router.Path("/").Methods("GET").HandlerFunc(handleWebsocket(relay))

	srv := &http.Server{
		Handler:           cors.Default().Handler(router),
		Addr:              s.Host + ":" + s.Port,
		WriteTimeout:      2 * time.Second,
		ReadTimeout:       2 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	log.Debug().Str("addr", srv.Addr).Msg("listening")
	srv.ListenAndServe()
}