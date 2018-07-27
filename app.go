package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	//"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"

	"github.com/peano88/pronoStats/dataLayer"
	"github.com/peano88/pronoStats/handler"
)

var db dataLayer.DataBridge
var hb handler.HandlerBridge

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	session, err := mgo.Dial("0.0.0.0")
	defer session.Close()

	logErr(err)

	dataBase := session.DB("<---write here--->")
	if dataBase == nil {
		log.Fatal("Fatal error in instantiating the DB")
	}

	db.Coll = dataBase.C(dataLayer.PRONO_COLLECTION)

	if db.Coll == nil {
		log.Fatal("Fatal error in instantiating the collection")
	}

	hb.Db = db
}

func main() {

	r := NewRouter()

	var wait time.Duration
	wait = 13 // to be fixed later

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

}
