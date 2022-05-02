// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"storj.io/drpc/drpcmux"
	"storj.io/drpc/drpcserver"

	"storj.io/drpc/examples/drpc/pb"
)

func main() {
	var tlsCrt, tlsKey string
	flag.StringVar(&tlsCrt, "tc", "", "tls certificate file")
	flag.StringVar(&tlsKey, "tk", "", "tls private key")
	flag.Parse()
	tlsEnabled := tlsCrt != "" && tlsKey != ""

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	cookieMonster := NewCookieMonsterServer()
	mux := drpcmux.New()
	err := pb.DRPCRegisterCookieMonster(mux, cookieMonster)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Printf("failed to register cookieMonster service")
		os.Exit(1)
	}
	addr := "0.0.0.0:8080"
	log.WithFields(log.Fields{"addr": addr}).Printf("opening listen socket")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Printf("failed to open listen socket")
		os.Exit(1)
	}
	if tlsEnabled {
		log.WithFields(log.Fields{"cert": tlsCrt, "key": tlsKey}).Printf("encrypting socket")
		cert, err := tls.LoadX509KeyPair(tlsCrt, tlsKey)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Print("failed to load certificate")
			os.Exit(1)
		}
		lis = tls.NewListener(lis, &tls.Config{
			Certificates: []tls.Certificate{cert},
			MinVersion:   tls.VersionTLS13,
		})
	}

	log.Printf("starting server")
	ctx := context.Background()
	srv := drpcserver.New(mux)
	err = srv.Serve(ctx, lis)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Printf("failed to serve DRPC")
		os.Exit(1)
	}
}

type CookieMonsterServer struct {
	pb.DRPCCookieMonsterUnimplementedServer

	jarMutex sync.RWMutex
	jar      map[pb.Cookie_Type]int64
}

func NewCookieMonsterServer() *CookieMonsterServer {
	jar := make(map[pb.Cookie_Type]int64)
	for k := range pb.Cookie_Type_name {
		jar[pb.Cookie_Type(k)] = 1000
	}
	return &CookieMonsterServer{
		jarMutex: sync.RWMutex{},
		jar:      jar,
	}
}

func (s *CookieMonsterServer) CheckJar(ctx context.Context, cookie *pb.Cookie) (*pb.CookieJar, error) {
	s.jarMutex.RLock()
	defer s.jarMutex.RUnlock()

	if _, ok := s.jar[cookie.Type]; !ok {
		return nil, fmt.Errorf("no cookie jar found for cookie type %v", cookie.Type)
	}
	return &pb.CookieJar{
		Cookie:    cookie,
		Remaining: s.jar[cookie.Type],
	}, nil
}

func (s *CookieMonsterServer) EatCookie(ctx context.Context, cookie *pb.Cookie) (*pb.CookieJar, error) {
	s.jarMutex.Lock()
	defer s.jarMutex.Unlock()

	if _, ok := s.jar[cookie.Type]; !ok {
		return nil, fmt.Errorf("no cookie jar found for cookie type %v", cookie)
	}
	if s.jar[cookie.Type] == 0 {
		return nil, fmt.Errorf("not enough cookies to be eaten")
	}
	s.jar[cookie.Type] -= 1
	return &pb.CookieJar{
		Cookie:    cookie,
		Remaining: s.jar[cookie.Type],
	}, nil
}
