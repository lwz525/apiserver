package main

import (
	"apiserver/config"
	"apiserver/model"
	"apiserver/router"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//for {
	//	fmt.Println(viper.GetString("runmode"))
	//	time.Sleep(4*time.Second)
	//}

	//set gin mode
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()
	middlewares := []gin.HandlerFunc{}
	router.Load(
		g,
		middlewares...,
	)

	model.DB.Init()
	defer model.DB.Close()
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()
	log.Infof("Start to listening the incoming requests on http address: %s", ":8080")
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("waiting for the router,retry in 1 sec.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")

}
