// package main

// import (
// 	"DM/internal/connection"
// 	"log"
// 	"net"
// )

// func main() {
// 	db := connection.ConnectDatabase()

// 	deptService, err := wire.InitializeDepartmentService(db)
// 	if err != nil {
// 		log.Fatalf("failed to initialize department service: %v", err)
// 	}

// 	//grpcServer := wire.InitializeGRPCServer(deptService)
// 	httpServer := wire.InitializeHTTPServer(deptService)

// 	lishttp, err := net.Listen("tcp", ":8000")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	log.Println("Starting gRPC server on :80000")
// 	if err := httpServer.Serve(lishttp); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}

// 	// lis, err := net.Listen("tcp", ":50051")
// 	// if err != nil {
// 	// 	log.Fatalf("failed to listen: %v", err)
// 	// }

// 	// log.Println("Starting gRPC server on :50051")
// 	// if err := grpcServer.Serve(lis); err != nil {
// 	// 	log.Fatalf("failed to serve: %v", err)
// 	// }
// }

package main

import (
	"flag"
	"os"

	"DM/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
