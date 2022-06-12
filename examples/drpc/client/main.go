// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"net"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"storj.io/drpc"
	"storj.io/drpc/drpcconn"

	"storj.io/drpc/examples/drpc/pb"
)

func main() {
	var addr string
	var concurrency, requests int
	var tlsEnabled, verbose bool
	flag.StringVar(&addr, "a", "127.0.0.1:8080", "address to dial to")
	flag.IntVar(&concurrency, "c", 1, "worker concurrency")
	flag.IntVar(&requests, "n", 1, "number of requests to be made")
	flag.BoolVar(&tlsEnabled, "tls", false, "enable TLS")
	flag.BoolVar(&verbose, "v", false, "print result for each call")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	latency := make(chan int64, concurrency*requests)
	stats := Stats{
		Start: time.Now(),
	}
	go func() {
		for l := range latency {
			stats.Calls += 1
			stats.LatencySum += l
		}
	}()

	log.WithFields(log.Fields{
		"addr":        addr,
		"concurrency": concurrency,
		"requests":    requests,
	}).Printf("starting workers")
	wg := &sync.WaitGroup{}
	for id := 0; id < concurrency; id++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, requests int, id int, latency chan<- int64) {
			defer wg.Done()
			var err error
			var rawconn drpc.Transport
			if tlsEnabled {
				rawconn, err = tls.Dial("tcp", addr, nil)
			} else {
				rawconn, err = net.Dial("tcp", addr)
			}
			if err != nil {
				log.WithFields(log.Fields{"error": err}).Println("failed to connect to server")
				os.Exit(1)
			}
			defer rawconn.Close()

			conn := drpcconn.New(rawconn)
			defer conn.Close()

			client := pb.NewDRPCCookieMonsterClient(conn)

			for rc := 0; rc < requests; rc++ {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				start := time.Now()
				crumbs, err := client.EatCookie(ctx, &pb.Cookie{
					Type: pb.Cookie_Chocolate,
				})
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
					}).Println("EatCookie")
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
