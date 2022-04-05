package main

import (
	"backend-onboarding/app"
	"backend-onboarding/config"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.InitConfig()

	mysqlConn, errMysql := config.ConnectMySQL()
	if errMysql != nil {
		log.Println("error mysql connection: ", errMysql)
	}
	//defer pgSql.Close()

	//kafkaConn, errKafka, _ := config.InitProducer()
	//if errKafka != nil {
	//	log.Println("error kafka connection: ", errKafka.Error())
	//} else {
	//	defer func() {
	//		for i := range kafkaConn {
	//			kafkaConn[i].Close()
	//		}
	//	}()
	//}

	router := app.InitRouter(mysqlConn)
	log.Println("routes Initialized")

	port := config.CONFIG["PORT"]
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Println("Server Initialized")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
