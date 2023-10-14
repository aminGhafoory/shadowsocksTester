package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aminghafoory/shadowTester/internal/database"
)

func splitShadowsocks(ss string) (Shadowsocks, error) {

	secretSlice := strings.Split(ss, "@")
	if len(secretSlice) != 2 {
		return Shadowsocks{}, fmt.Errorf("error in parsing ss url")
	}
	secret := secretSlice[0]
	secret = strings.TrimPrefix(secret, "ss://")
	other := secretSlice[1]
	IPSlice := strings.Split(other, ":")
	if len(IPSlice) != 2 {
		return Shadowsocks{}, fmt.Errorf("error in parsing ss url")
	}
	ip := IPSlice[0]

	other = IPSlice[1]
	portSlice := strings.Split(other, "#")
	if len(portSlice) != 2 {
		return Shadowsocks{}, fmt.Errorf("error in parsing ss url")
	}
	portstr := portSlice[0]
	port, err := strconv.Atoi(portstr)
	if err != nil {
		return Shadowsocks{}, err
	}
	nameEncoded := portSlice[1]
	nameDecoded, err := url.QueryUnescape(nameEncoded)
	if err != nil {
		return Shadowsocks{}, err
	}

	if err != nil {
		return Shadowsocks{}, err
	}

	return Shadowsocks{
		Sslink: ss,
		secret: secret,
		IP:     ip,
		Port:   port,
		Name:   nameDecoded,
	}, nil

}

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequests time.Duration) {

	log.Printf("Scraping on %v goroutine every %s duration", concurrency, timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		subs, err := db.GetNextSubsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching subs", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, sub := range subs {
			wg.Add(1)
			go scrapeSubs(wg, db, sub)
		}
		wg.Wait()

	}

}

func scrapeSubs(wg *sync.WaitGroup, db *database.Queries, sub database.Sub) {

	resp, err := http.Get(sub.Url)
	if err != nil {
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	shadowsockss := strings.Split(string(body), "\n")
	for _, v := range shadowsockss {
		ss, err := splitShadowsocks(v)
		if err != nil {
			continue
		}
		db.CreateShadowsocks(context.Background(), database.CreateShadowsocksParams{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Sslink:    ss.Sslink,
			Ip:        ss.IP,
			SubID:     sub.ID,
			Port:      int32(ss.Port),
			Secret:    ss.secret,
			Name:      ss.Name,
		})

	}

	wg.Done()

}
