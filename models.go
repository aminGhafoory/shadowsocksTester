package main

import "github.com/aminghafoory/shadowTester/internal/database"

type Shadowsocks struct {
	Sslink string `json:"link"`
	secret string
	IP     string `json:"host"`
	Port   int    `json:"port"`
	Name   string `json:"name"`
}

func DbSStoShadowsocks(dbSS database.Ss) Shadowsocks {
	return Shadowsocks{
		Sslink: dbSS.Sslink,
		secret: dbSS.Secret,
		IP:     dbSS.Ip,
		Port:   int(dbSS.Port),
		Name:   dbSS.Name,
	}
}

// func (ss Shadowsocks) prettyPrint() {
// 	fmt.Println("IP: ", ss.ip)
// 	fmt.Println("PORT: ", ss.port)
// 	fmt.Println("SECRET: ", ss.secret)
// 	fmt.Println("NAME: ", ss.name)
// 	fmt.Println()
// }

type Sub struct {
	URL string `json:"url"`
}

func dbSubToSub(dbsub database.Sub) Sub {
	return Sub{
		URL: dbsub.Url,
	}
}

type Bestlink struct {
	ID              int    `json:"id"`
	Secret          string `json:"secret"`
	IP              string `json:"IP"`
	Port            int    `json:"port"`
	Name            string `json:"name"`
	Sslink          string `json:"sslink"`
	AvgResponseTime int    `json:"avg_response_time"`
	SuccessfulCount int    `json:"successful_count"`
	FailureCount    int    `json:"failure_count"`
}

func dbBestLinkToBestlink(dbBestlink database.GetBestListRow) Bestlink {
	return Bestlink{
		ID:              int(dbBestlink.ID),
		Secret:          dbBestlink.Secret,
		IP:              dbBestlink.Ip,
		Port:            int(dbBestlink.Port),
		Name:            dbBestlink.Name,
		Sslink:          dbBestlink.Sslink,
		AvgResponseTime: int(dbBestlink.AvgResponseTime),
		SuccessfulCount: int(dbBestlink.SuccessfulCount),
		FailureCount:    int(dbBestlink.FailureCount),
	}
}
