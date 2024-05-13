package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
)

func main() {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		wg          sync.WaitGroup
		winner      string
		resultChan  = make(chan string)
		server      = []string{"Yandex taxi", "Uber", "Grab", "Lyft", "Didi Chuxing"}
	)

	defer cancel()
	for i := range server {
		s := server[i]
		wg.Add(1)
		go func() {
			requestRide(ctx, s, resultChan, cancel)
			wg.Done()
		}()
	}

	go func() {
		winner = <-resultChan
		cancel()
	}()
	wg.Wait()
	log.Fatal("Winner:", winner)

}

func requestRide(ctx context.Context, serviceName string, resultChan chan string, cansel context.CancelFunc) {

	for {
		select {
		case <-ctx.Done():
			log.Println("Stop searching taxi  on server: ", serviceName, ctx.Err())
			return
		default:
			if rand.Int()%1000 == 33 {
				resultChan <- serviceName
				cansel()
				return
			}
			continue
		}
	}
}
