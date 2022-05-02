// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"net"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"storj.io/drpc/drpcconn"

	"storj.io/drpc/examples/drpc/pb"
)

func main() {
	var addr string
	var tlsEnabled, verbose bool
	flag.BoolVar(&verbose, "v", false, "print result for each call")
	flag.BoolVar(&tlsEnabled, "tls", false, "use tls connection")
	flag.StringVar(&addr, "a", "127.0.0.1:8080", "server address")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	log.WithFields(log.Fields{"addr": addr, "tls": tlsEnabled}).Print("dialing")
	var rawconn net.Conn
	var err error
	dialStart := time.Now()
	if tlsEnabled {
		rawconn, err = tls.Dial("tcp", addr, nil)
	} else {
		rawconn, err = net.Dial("tcp", addr)
	}
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Println("failed to connect to server")
		os.Exit(1)
	}
	log.WithFields(log.Fields{"elapsed": time.Since(dialStart)}).Print("dialed")

	conn := drpcconn.New(rawconn)
	defer conn.Close()

	client := pb.NewDRPCCookieMonsterClient(conn)

	if flag.NArg() == 0 || flag.Arg(0) == "list" {
		for k := range pb.Cookie_Type_name {
			checkJar(client, pb.Cookie_Type(k))
		}
		return
	}

	cookieName := flag.Arg(1)
	cookieType, ok := pb.Cookie_Type_value[cookieName]
	if !ok {
		log.Printf("could not find cookie flavored %q", cookieName)
		os.Exit(1)
	}
	switch flag.Arg(0) {
	case "eat":
		amountStr := flag.Arg(2)
		if amountStr == "" {
			amountStr = "1"
		}
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Println("failed to parse amount of cookies to eat")
			os.Exit(1)
		}
		for n := 0; n < amount; n++ {
			eatCookie(client, pb.Cookie_Type(cookieType))
		}
	case "get":
		checkJar(client, pb.Cookie_Type(cookieType))
	}
}

func checkJar(client pb.DRPCCookieMonsterClient, cookieType pb.Cookie_Type) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	start := time.Now()

	jar, err := client.CheckJar(ctx, &pb.Cookie{
		Type: cookieType,
	})
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Println("failed")
		os.Exit(1)
	}
	log.WithFields(log.Fields{
		"cookie":    pb.Cookie_Type_name[int32(jar.Cookie.Type)],
		"elapsed":   time.Since(start),
		"remaining": jar.Remaining,
	}).Printf("CheckJar")
}

func eatCookie(client pb.DRPCCookieMonsterClient, cookieType pb.Cookie_Type) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	start := time.Now()

	jar, err := client.EatCookie(ctx, &pb.Cookie{
		Type: cookieType,
	})
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Println("failed")
		os.Exit(1)
	}
	log.WithFields(log.Fields{
		"cookie":    pb.Cookie_Type_name[int32(jar.Cookie.Type)],
		"elapsed":   time.Since(start),
		"remaining": jar.Remaining,
	}).Printf("EatCookie")
}
