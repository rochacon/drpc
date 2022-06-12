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
	var concurrency, requests uint
	var debug, tlsEnabled, verbose bool
	flag.StringVar(&addr, "a", "127.0.0.1:8080", "address to dial to")
	flag.UintVar(&concurrency, "c", 1, "worker concurrency")
	flag.BoolVar(&debug, "d", false, "print debug logs")
	flag.UintVar(&requests, "n", 1, "number of requests to be made")
	flag.BoolVar(&tlsEnabled, "tls", false, "enable TLS")
	flag.BoolVar(&verbose, "v", false, "print result for each call")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	// start stats aggregator
	latency := make(chan Latency, concurrency*requests)
	stats := Stats{
		CallsPerWorker:      make([]uint64, concurrency),
		Start:               time.Now(),
		LatencySumPerWorker: make([]uint64, concurrency),
	}
	go func() {
		for l := range latency {
			stats.Calls += 1
			stats.CallsPerWorker[l.Worker] += 1
			stats.LatencySum += uint64(l.Elapsed)
			stats.LatencySumPerWorker[l.Worker] += uint64(l.Elapsed)
		}
		log.Debug("latency channel closed")
	}()

	// start work queue
	work := make(chan uint)
	go func() {
		for rc := uint(0); rc < requests; rc++ {
			work <- rc
		}
		close(work)
		log.WithFields(log.Fields{
			"requests": requests,
		}).Debug("finished queuing requests")
	}()

	// start workers
	log.WithFields(log.Fields{
		"addr":        addr,
		"concurrency": concurrency,
		"requests":    requests,
	}).Printf("starting workers")

	wg := &sync.WaitGroup{}
	for id := uint(0); id < concurrency; id++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, work <-chan uint, id uint, latency chan<- Latency) {
			defer wg.Done()
			var err error
			var rawconn drpc.Transport
			connStart := time.Now()
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
			log.WithFields(log.Fields{
				"elapsed": time.Since(connStart),
				"worker":  id,
			}).Debug("connection started")

			conn := drpcconn.New(rawconn)
			defer conn.Close()

			client := pb.NewDRPCCookieMonsterClient(conn)

			for range work {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)

				start := time.Now()
				crumbs, err := client.EatCookie(ctx, &pb.Cookie{
					Type: pb.Cookie_Chocolate,
				})
				if err != nil {
					log.WithFields(log.Fields{
						"error":  err,
						"worker": id,
					}).Println("failed")
					cancel()
					continue
				}
				elapsed := time.Since(start)
				latency <- Latency{
					Elapsed: elapsed,
					Worker:  id,
				}
				if verbose {
					log.WithFields(log.Fields{
						"cookie":  crumbs.Cookie.Type.String(),
						"latency": elapsed.String(),
						"worker":  id,
					}).Println("EatCookie")
				}
				cancel()
			}
		}(wg, work, id, latency)
	}
	wg.Wait()
	close(latency)
	log.WithFields(log.Fields{
		"calls":        stats.Calls,
		"latency_mean": stats.LatencyMean().String(),
		"elapsed":      time.Since(stats.Start).String(),
	}).Println("stats")
	for i, n := range stats.CallsPerWorker {
		log.WithFields(log.Fields{
			"calls":        n,
			"latency_mean": stats.WorkerLatencyMean(uint(i)).String(),
			"worker":       i,
		}).Println("stats per worker")
	}
}

type Latency struct {
	Elapsed time.Duration
	Worker  uint
}

type Stats struct {
	Calls               uint64
	CallsPerWorker      []uint64
	LatencySum          uint64
	LatencySumPerWorker []uint64
	Start               time.Time
}

func (s Stats) LatencyMean() time.Duration {
	return time.Duration(s.LatencySum / s.Calls)
}

func (s Stats) WorkerLatencyMean(id uint) time.Duration {
	return time.Duration(s.LatencySumPerWorker[id] / s.CallsPerWorker[id])
}
