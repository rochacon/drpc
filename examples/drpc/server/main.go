// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
	"storj.io/drpc/drpcmux"
	"storj.io/drpc/drpcserver"

	"storj.io/drpc/examples/drpc/pb"
)

func main() {
	var tlsCrtPath, tlsKeyPath string
	flag.StringVar(&tlsCrtPath, "tc", "", "tls certificate path")
	flag.StringVar(&tlsKeyPath, "tk", "", "tls private key path")
	flag.Parse()

	tlsEnabled := tlsCrtPath != "" && tlsKeyPath != ""

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	cookieMonster := &CookieMonsterServer{}
	mux := drpcmux.New()
	err := pb.DRPCRegisterCookieMonster(mux, cookieMonster)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Printf("failed to register cookieMonster service")
		os.Exit(1)
	}
	s := drpcserver.New(mux)
	var lis net.Listener
	addr := "0.0.0.0:8080"
	log.WithFields(log.Fields{
		"addr":        addr,
		"tls_enabled": tlsEnabled,
		"tls_crt":     tlsCrtPath,
		"tls_key":     tlsKeyPath,
	}).Printf("opening listen socket")
	if tlsEnabled {
		var tlsCrt tls.Certificate
		tlsCrt, err = tls.LoadX509KeyPair(tlsCrtPath, tlsKeyPath)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Printf("failed to read tls certificates")
			os.Exit(1)
		}
		lis, err = tls.Listen("tcp", addr, &tls.Config{
			Certificates: []tls.Certificate{tlsCrt},
		})
	} else {
		lis, err = net.Listen("tcp", addr)
	}
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Printf("failed to open listen socket")
		os.Exit(1)
	}
	log.Printf("starting server")
	ctx := context.Background()
	err = s.Serve(ctx, lis)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Printf("failed to serve DRPC")
		os.Exit(1)
	}
}

type CookieMonsterServer struct {
	pb.DRPCCookieMonsterUnimplementedServer
}

// EatCookie turns a cookie into crumbs.
func (s *CookieMonsterServer) EatCookie(ctx context.Context, cookie *pb.Cookie) (*pb.Crumbs, error) {
	return &pb.Crumbs{
		Cookie: cookie,
	}, nil
}
