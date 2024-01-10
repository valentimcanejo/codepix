package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/valentimcanejo/codepix/application/grpc"
	"github.com/valentimcanejo/codepix/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
