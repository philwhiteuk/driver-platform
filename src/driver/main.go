package main

import (
	"common"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	lu, err := url.Parse(fmt.Sprintf("udp://%s", os.Getenv("LOG_DAEMON")))
	if err != nil {
		log.Fatal(err)
	}
	logger, _ := common.NewUnifiedLogger(common.SyslogEstablishConnectionFn(), lu, "driver")

	conn, err := net.Dial("udp", os.Getenv("STATSD_DAEMON"))
	if err != nil {
		log.Fatal(err)
	}
	statter, _ := common.NewDogStatsDStatter("driver", conn)

	err = registerWithAPIGateway()
	if err != nil {
		logger.Err(err.Error())
		logger.Emerg("Registering with Gateway Failed!")
		return
	}

	drivers, err := parseDriverData()
	if err != nil {
		logger.Err(err.Error())
		logger.Emerg("Failed to parse driver data!")
		return
	}
	http.Handle("/drivers/", common.HTTPLogger(logger, statter, common.NewSystemClock())(driversHandlerFunc(drivers)))

	logger.Notice(fmt.Sprintf("application serving on :%s", os.Getenv("SERVICE_PORT")))

	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")), nil)
	if err != nil {
		logger.Err(err.Error())
		logger.Crit("Failed to launch server!")
		return
	}
}

func registerWithAPIGateway() error {
	data := url.Values{}
	data.Add("address", fmt.Sprintf("http://%s:%s", os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_PORT")))

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/service/drivers", os.Getenv("API_GATEWAY")), strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("Invalid Request %s", err)
	}
	req.PostForm = data
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("Client Request Failed %s", err)
	}

	return nil
}
