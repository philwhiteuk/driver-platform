package main

import (
	"common"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
)

func main() {
	lu, err := url.Parse(fmt.Sprintf("udp://%s", os.Getenv("LOG_DAEMON")))
	if err != nil {
		log.Fatal(err)
	}
	logger, _ := common.NewUnifiedLogger(common.SyslogEstablishConnectionFn(), lu, "gateway")

	conn, err := net.Dial("udp", os.Getenv("STATSD_DAEMON"))
	if err != nil {
		log.Fatal(err)
	}
	statter, _ := common.NewDogStatsDStatter("gateway", conn)

	proxyRouter, err := startPublicGateway()
	if err != nil {
		fmt.Println("Failed to start api gateway server!")
		return
	}

	router := http.NewServeMux()
	router.Handle("/service/", common.HTTPLogger(logger, statter, common.NewSystemClock())(registerServiceHandler(proxyRouter)))

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")),
		Handler: router,
	}

	logger.Notice(fmt.Sprintf("application serving on :%s", os.Getenv("SERVICE_PORT")))

	err = server.ListenAndServe()
	if err != nil {
		logger.Crit("Failed to launch server!")
		return
	}
}

func startPublicGateway() (*http.ServeMux, error) {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PUBLIC_PORT")),
		Handler: router,
	}

	var err error
	go func() {
		err = server.ListenAndServe()
	}()

	return router, err
}
