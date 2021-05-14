package main

import (
	"flag"
	"io/ioutil"
	"net"
	"os"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpclog "google.golang.org/grpc/grpclog"

	acc "github.com/resonatecoop/user-api/pkg/access"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/resonatecoop/user-api/gateway"
	"github.com/resonatecoop/user-api/insecure"
	"github.com/resonatecoop/user-api/pkg/config"
	jwt "github.com/resonatecoop/user-api/pkg/jwt"
	pgsql "github.com/resonatecoop/user-api/pkg/postgres"
	"github.com/resonatecoop/user-api/pkg/zerolog"
	pbIAM "github.com/resonatecoop/user-api/proto/iam"
	pbUser "github.com/resonatecoop/user-api/proto/user"
	iamserver "github.com/resonatecoop/user-api/server/iam"
	iamdb "github.com/resonatecoop/user-api/server/iam/platform/postgres"
	"github.com/resonatecoop/user-api/server/iam/secure"
	userserver "github.com/resonatecoop/user-api/server/user"

	// Static files
	_ "github.com/resonatecoop/user-api/statik"
)

func main() {

	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)

	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	checkErr(log, err)

	db, err := pgsql.New(cfg.DB.Dev.PSN, cfg.DB.Dev.LogQueries, cfg.DB.Dev.TimeoutSeconds)

	checkErr(log, err)

	zerolog := zerolog.New()

	secureSvc := secure.New(cfg.App.MinPasswordStrength)

	jwtService := jwt.New(cfg.JWT.Secret, cfg.JWT.Duration, cfg.JWT.Algorithm)

	accService := acc.New(cfg.Access.NoTokenMethods, cfg.Access.PublicUserMethods)

	iamServer := iamserver.New(db, jwtService, iamdb.NewUser(), secureSvc)

	interceptorAuth := iamserver.NewAuthInterceptor(jwtService, accService)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	opts := make([]grpc.ServerOption, 0)

	// TODO: Replace with your own certificate!
	opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)))

	opts = append(opts, grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			interceptorAuth.Unary(),
		)))

	s := grpc.NewServer(
		opts...,
	)

	pbUser.RegisterResonateUserServer(s, userserver.New(db, secureSvc))

	iamserver.NewLoggingService(iamServer, zerolog)

	pbIAM.RegisterResonateIAMServer(s, iamServer)

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatal(err)
}

func checkErr(log grpclog.LoggerV2, err error) {
	if err != nil {
		log.Fatal(err)
	}
}
