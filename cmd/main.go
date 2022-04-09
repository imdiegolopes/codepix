package main

import (
	"codepix/application/grpc"
	"codepix/infrastructure/db"
	"os"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
