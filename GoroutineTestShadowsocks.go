package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Jigsaw-Code/outline-sdk/x/config"
	"github.com/aminghafoory/shadowTester/internal/database"
)

func startTesting(
	db *database.Queries,
	concurrency int,
	timeBetweenRequests time.Duration) {

	log.Printf("Testing on %v goroutine every %s duration", concurrency, timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		sss, err := db.GetNextSSToTest(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error testing ss", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, ss := range sss {
			wg.Add(1)
			go TestShadowsocks(wg, db, ss)
		}
		wg.Wait()

	}

}

func TestShadowsocks(wg *sync.WaitGroup, db *database.Queries, ss database.Ss) {

	defer wg.Done()

	start := time.Now()

	dialer, err := config.NewStreamDialer(ss.Sslink)
	if err != nil {
		db.UpdateDestinationIP(context.Background(), database.UpdateDestinationIPParams{
			ApiReportedIp: sql.NullString{
				String: ss.Ip,
				Valid:  false,
			},
			ID: ss.ID,
		})
		db.TestSS(context.Background(), database.TestSSParams{
			Source:       ss.Ip,
			SsID:         ss.ID,
			ResponseTime: -1,
			IsSuccessful: false,
		})
		return
	}

	dialContext := func(ctx context.Context, network, addr string) (net.Conn, error) {
		if !strings.HasPrefix(network, "tcp") {
			return nil, fmt.Errorf("protocol not supported: %v", network)
		}
		return dialer.Dial(ctx, addr)
	}
	httpClient := &http.Client{Transport: &http.Transport{DialContext: dialContext}}

	resp, err := httpClient.Get("https://api64.ipify.org/")
	if err != nil {
		db.UpdateDestinationIP(context.Background(), database.UpdateDestinationIPParams{
			ApiReportedIp: sql.NullString{
				String: ss.Ip,
				Valid:  false,
			},
			ID: ss.ID,
		})
		db.TestSS(context.Background(), database.TestSSParams{
			Source:       ss.Ip,
			SsID:         ss.ID,
			ResponseTime: -1,
			IsSuccessful: false,
		})
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		db.UpdateDestinationIP(context.Background(), database.UpdateDestinationIPParams{
			ApiReportedIp: sql.NullString{
				String: "",
				Valid:  false,
			},
			ID: ss.ID,
		})
		db.TestSS(context.Background(), database.TestSSParams{
			Source:       ss.Ip,
			SsID:         ss.ID,
			ResponseTime: -1,
			IsSuccessful: false,
		})
		return
	}
	if resp.StatusCode != 200 {

		db.UpdateDestinationIP(context.Background(), database.UpdateDestinationIPParams{
			ApiReportedIp: sql.NullString{
				String: "",
				Valid:  false,
			},
			ID: ss.ID,
		})
		db.TestSS(context.Background(), database.TestSSParams{
			Source:       ss.Ip,
			SsID:         ss.ID,
			ResponseTime: -1,
			IsSuccessful: false,
		})

		return
	}
	ip := string(body)

	responseTime := time.Since(start)
	log.Println(responseTime)
	_, err = db.TestSS(context.Background(), database.TestSSParams{
		Source:       ss.Ip,
		SsID:         ss.ID,
		ResponseTime: int32(responseTime.Milliseconds()),
		IsSuccessful: true,
	})

	if err != nil {
		log.Printf("error %+v", err)
		return
	}

	_, err = db.UpdateDestinationIP(context.Background(), database.UpdateDestinationIPParams{
		ApiReportedIp: sql.NullString{
			String: ip,
			Valid:  true,
		},
		ID: ss.ID,
	})

	if err != nil {
		log.Printf("error %+v", err)
		return
	}

}
