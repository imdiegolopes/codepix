package main

import (
	"codepix/application/grpc"
	"codepix/infrastructure/db"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"log"
)

var database *gorm.DB

var (
	rootCmd = &cobra.Command{
		Use:   "Codepix",
		Short: "An example cobra program",
		Long:  `Use CodePix software to intermediate bank transaction with apache kafka and grpc`,
	}
)

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
