// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"flag"
	"net"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"storj.io/drpc/drpcconn"

	"storj.io/drpc/examples/drpc/pb"
)

func main() {
	var concurrency, requests int
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "print result for each call")
	flag.IntVar(&concurrency, "c", 1, "worker concurrency")
	flag.IntVar(&requests, "n", 1, "number of requests to be made")
	flag.Parse()

	ctx := context.Background()

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	latency := make(chan int64, concurrency)
	stats := Stats{
		Start: time.Now(),
	}
	go func() {
		for {
			select {
			case l := <-latency:
				stats.Calls += 1
				stats.LatencySum += l
			}
		}
		log.Println("latency channel closed")
	}()

	log.WithFields(log.Fields{
		"concurrency": concurrency,
		"requests":    requests,
	}).Printf("starting workers")
	wg := &sync.WaitGroup{}
	for id := 0; id < concurrency; id++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, requests int, id int, latency chan<- int64) {
			defer wg.Done()

			rawconn, err := net.Dial("tcp", "127.0.0.1:8080")
			if err != nil {
				log.WithFields(log.Fields{"error": err}).Println("failed to connect to server")
				os.Exit(1)
			}

			conn := drpcconn.New(rawconn)
			defer conn.Close()

			client := pb.NewDRPCCookieMonsterClient(conn)

			for rc := 0; rc < requests; rc++ {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				start := time.Now()
				crumbs, err := client.EatCookie(ctx, &pb.Cookie{
					Type: pb.Cookie_Oatmeal,
				})
				if err != nil {
					log.WithFields(log.Fields{"worker": id, "error": err}).Println("failed")
					continue
				}

				if err != nil {
					log.WithFields(log.Fields{"worker": id, "error": err}).Println("failed")
					continue
				}
				elapsed := time.Since(start)
				latency <- int64(elapsed)
				if verbose {
					log.WithFields(log.Fields{
						"worker":  id,
						"cookie":  crumbs.Cookie.Type.String(),
						"latency": elapsed.String(),
					}).Println("eaten")
				}
			}
		}(wg, requests, id, latency)
	}
	wg.Wait()
	close(latency)
	log.WithFields(log.Fields{
		"calls":        stats.Calls,
		"latency_mean": stats.LatencyMean().String(),
		"elapsed":      time.Since(stats.Start).String(),
	}).Println("stats")
}

type Stats struct {
	Calls      int64
	LatencySum int64
	Start      time.Time
}

func (s Stats) LatencyMean() time.Duration {
	return time.Duration(s.LatencySum / s.Calls)
}
