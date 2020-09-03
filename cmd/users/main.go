package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/n7down/kuiper/internal/logger"
	"github.com/n7down/kuiper/internal/logger/logruslogger"
	"github.com/n7down/kuiper/internal/users/persistence/mysql"
	"google.golang.org/grpc"

	users_pb "github.com/n7down/kuiper/internal/pb/users"
	users "github.com/n7down/kuiper/internal/users/servers"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

var (
	Version     string
	Build       string
	showVersion *bool
	port        string
	log         logger.Logger
	server      *users.UsersServer
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		port = os.Getenv("PORT")
		dbConn := os.Getenv("DB_CONN")

		log = logruslogger.NewLogrusLogger(true)
		persistence, err := mysql.NewMysqlPersistence(dbConn)
		if err != nil {
			log.Fatal(err)
		}

		server = users.NewUsersServer(persistence)
	}
}

func main() {
	if *showVersion {
		fmt.Printf("settings server: version %s build %s", Version, Build)
	} else {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("Listening on port: %s\n", port)
		grpcServer := grpc.NewServer()
		users_pb.RegisterUsersServiceServer(grpcServer, server)
		grpcServer.Serve(lis)
	}
}
