package main

import (
	"context"
	"flag"
	"go-basic-final/internal/structure"
	"log"
	"os"
	"time"
)

const LOG = "./log.txt"

func main() {
	var s structure.SynchronizationImp
	fo, err := os.OpenFile(LOG, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer fo.Close()
	if err != nil {
		panic(err)
	}
	sLog := structure.NewSynchronizationWithLog(&s, fo, fo)

	ctx, cancel := context.WithCancel(context.Background())
	mainDir, subDir := dirs()
	go start(cancel, &sLog, mainDir, subDir)

	for {
		select {
		case <-ctx.Done():
			ctx, cancel = context.WithCancel(context.Background())
			time.Sleep(time.Second * 5)
			go start(cancel, &sLog, mainDir, subDir)
		default:
			time.Sleep(time.Second * 5)
		}
	}
}

func start(cancel context.CancelFunc, s structure.Synchronization, mainDir, subDir string) {
	err := s.Synchronize(mainDir, subDir)
	if err != nil {

	}

	cancel()
}

func dirs() (string, string) {
	mainDir := flag.String("main-dir", "", "main dir for synchronization")
	subDir := flag.String("sub-dir", "", "sub dir for synchronization")
	flag.Parse()

	if *mainDir == "" || *subDir == "" {
		log.Fatal("dirs are not correct")
	}

	return *mainDir, *subDir
}
