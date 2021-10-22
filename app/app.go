package app

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/valp0/academy-go-q32021/handlers"
	"github.com/valp0/academy-go-q32021/repo"
	"github.com/valp0/academy-go-q32021/services"
)

const (
	port = ":8080"
	path = "./files/pokemons.csv"
)

// Will setup a server to run using the constants provided in app.go.
func RunServer() error {
	os.Setenv("PATH", path)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go logExit(c)

	lr := repo.NewLocalRepo()
	ar := repo.NewApiRepo()
	asr := repo.NewAsyncRepo()

	hSvc := services.NewHomeSvc()
	rSvc := services.NewReadSvc(lr)
	fSvc := services.NewFetchSvc(ar)
	aSvc := services.NewAsyncSvc(asr)

	hh := handlers.NewHomeHandler(hSvc)
	rh := handlers.NewReadHandler(rSvc)
	fh := handlers.NewFetchHandler(fSvc)
	ah := handlers.NewAsyncHandler(aSvc)

	http.HandleFunc("/", hh.Home)
	http.HandleFunc("/read", rh.Query)
	http.HandleFunc("/fetch", fh.ApiFetch)
	http.HandleFunc("/async", ah.Async)

	log.Println("Listening on port", port[1:])
	if err := http.ListenAndServe(port, nil); err != nil {
		return err
	}

	return nil
}

func logExit(c chan os.Signal) {
	for range c {
		log.Fatal("Process terminated")
	}
}
