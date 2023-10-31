package app

import (
	"celmon/log"
	"celmon/server"
	"celmon/tg"
	"context"
	"fmt"
	"sync"
	"time"
)

func Run(ctx context.Context, c *Config) {
	var consHeight uint64
	var daHeight uint64
	var err error

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		consHeight, err = c.getConsNodeHeight(ctx)
		if err != nil {
			log.Error(err)
			return
		}
	}()

	go func() {
		defer wg.Done()

		daHeight, err = c.getDANodeHeight(ctx)
		if err != nil {
			log.Error(err)
			return
		}
	}()

	wg.Wait()

	server.GlobalState.ConsensusHeight = consHeight
	server.GlobalState.BridgeHeight = daHeight

	diff := consHeight - daHeight
	if diff >= uint64(c.AlarmCriteria.HeightDiffer) {
		server.GlobalState.Status = false

		msg := fmt.Sprintf("status(cons: %d, da: %d): 🛑", consHeight, daHeight)
		log.Info(msg)
		tg.SendMsg(msg)
	} else {
		server.GlobalState.Status = true
		msg := fmt.Sprintf("status(cons: %d, da: %d): 🟢", consHeight, daHeight)
		log.Info(msg)
	}

	return
}
