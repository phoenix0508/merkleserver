package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/lib/pq"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"

	"github.com/DedicatedDev/merkleserver/internal/app"
	"github.com/DedicatedDev/merkleserver/internal/repository"
	"github.com/DedicatedDev/merkleserver/internal/service"
	desc "github.com/DedicatedDev/merkleserver/pkg"
)

const (
	DB_OWNIFY   = "database.ownify.dbname"
	DB_USER_LOG = "database.log.dbname"
)

func main() {
	// Ownify DB
	filestoreDB, err := repository.NewDB(DB_OWNIFY)
	if err != nil {
		log.Fatalf("[ERR] cannot create database %s", err)
		return
	}

	// Register all services
	//dbHandler := repository.NewDBHandler(db)
	fileDBHandler := repository.NewDBHandler(filestoreDB)

	fileService := service.NewFileService(fileDBHandler)
	verifyService := service.NewVerifyService(fileDBHandler)

	// Interceptors
	grpcOpts := app.GrpcInterceptor()
	httpOpts := app.HttpInterceptor()

	// Starting gRPC server
	go func() {
		listener, err := net.Listen("tcp", "0.0.0.0:8900")
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("start server")

		grpcServer := grpc.NewServer(grpcOpts)

		desc.RegisterMicroserviceServer(grpcServer, app.NewMicroservice(
			fileService,
			verifyService,
		))

		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// Starting HTTP server
	mux := runtime.NewServeMux(httpOpts)

	err = desc.RegisterMicroserviceHandlerServer(context.Background(), mux, app.NewMicroservice(
		fileService,
		verifyService,
	))
	if err != nil {
		log.Println("cannot register this service")
	}
	//log.Fatalln(http.ListenAndServe(":8901", setUserLog(addCORSHeaders(mux), logService)))
	log.Fatalln(http.ListenAndServe(":8901", addCORSHeaders(mux)))
}

// addCORSHeaders is a middleware function that adds the necessary CORS headers to the HTTP response.
// addCORSHeaders is a middleware function that adds the necessary CORS headers to the HTTP response.
func addCORSHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCORSHeaders(w, r)
		// Preflight request. Reply successfully:
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func setCORSHeaders(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers to the response
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
}

func rateLimitMiddleware(limiter *rate.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
