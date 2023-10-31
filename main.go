package main

import (
	"celmon/app"
	"celmon/log"
	"celmon/server"
	"celmon/tg"
	"celmon/utils"
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/pelletier/go-toml/v2"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ctx := context.Background()

	cfgPath := flag.String("config", "", "Config file")
    flag.Parse()
	if *cfgPath == "" {
		panic("Error: Please input config file path with -config flag.")
	}

	f, err := os.ReadFile(*cfgPath)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	cfg := app.Config{}
	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		log.Error(err)
		panic(err)
	}


	publicIp, err := utils.GetPublicIP()
	if err != nil {
		log.Error(err)
		panic(err)
	}
	tgTitle := fmt.Sprintf("🤖 Celmon for %s 🤖\nStatus Server: [Link](%s:%d)\n", cfg.General.Network, publicIp, cfg.General.ListenPort)
	tg.SetTg(cfg.Tg.Enable, tgTitle, cfg.Tg.Token, cfg.Tg.ChatID)

	for {
		go server.Run(cfg.General.ListenPort)
		app.Run(ctx, &cfg)
		time.Sleep(time.Duration(cfg.General.Period) * time.Minute)
	}
}
