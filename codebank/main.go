package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/srpulga/CodeBank/infrastructure/grpc/server"
	"github.com/srpulga/CodeBank/infrastructure/kafka"
	"github.com/srpulga/CodeBank/infrastructure/repository"
	"github.com/srpulga/CodeBank/usecase"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db := setupDb()
	defer db.Close()
	producer := setupKafkaProducer()
	processTransactionUseCase := setupTransactionUseCase(db, producer)
	serveGrpc(processTransactionUseCase)
}

func setupTransactionUseCase(db *sql.DB, producer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDB(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	useCase.KafkaProducer = producer
	return useCase
}

func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer(os.Getenv("KafkaBootstrapServers"))
	return producer
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"),
	)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Error Connection to Database")
	}
	return db
}

func serveGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	fmt.Println("Starting GRPC Server")
	grpcServer.Serve()
}

//Create credit card manually
//cc := domain.NewCreditCard()
//cc.Number = "1234567890123456"
//cc.Name = "Pulga"
//cc.ExpirationMonth = 12
//cc.ExpirationYear = 29
//cc.CVV = 123
//cc.Limit = 10000
//cc.Balance = 0
//
//repo := repository.NewTransactionRepositoryDB(db)
//err := repo.CreateCreditCard(*cc)
//if err != nil {
//fmt.Println(err)
//}
