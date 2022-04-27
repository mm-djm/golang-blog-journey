package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"golang-blog-journey/controller"
	"golang-blog-journey/util"
	"golang-blog-journey/util/cipher"
	"golang-blog-journey/util/db"
	log "golang-blog-journey/util/log"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var (
	config = flag.String("c", "./conf/config.json", "configuration file, json format")
)

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadFile(*config)
	if err != nil {
		fmt.Printf("Read ulb_api_go config file %v failed: %v\n", config, err)
		panic(err)
	}

	var Conf util.Config
	if err = json.Unmarshal(bytes, &Conf); err != nil {
		fmt.Printf("Parse config json failed: %v\n", err)
		panic(err)
	}

	fmt.Printf("Config: %v\n", Conf)

	// Init Key
	cipher.PwdKey = []byte(Conf.Ciphers.PwdKey)

	// Init Log
	err = log.InitGlobalLogger(&Conf.LogConfig, zap.AddCallerSkip(1))
	if err != nil {
		fmt.Printf("Init log module failed: %v\n", err)
		panic(err)
	}

	util.InitErrorMap()

	// Init DB
	db.InitDatabase(Conf.DBConfig.User, Conf.DBConfig.Password, Conf.DBConfig.IP)

	addr := fmt.Sprintf("%s:%d", "localhost", Conf.ListenPort)
	err = controller.RunServer(Conf.ListenIP, addr)
	fmt.Printf("The tcp server is running faild: %v\n", err)
	panic(err)

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// <-c
	// os.Exit(1)
}
