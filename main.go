package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	targetUrl := "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"
	req, _ := http.NewRequest("GET", targetUrl, nil)
	var (
		total float64
		cnt   int
	)
	for {
		res, _ := http.DefaultClient.Do(req)
		body, _ := io.ReadAll(res.Body)
		split := strings.Split(string(body), " ")
		for _, v := range split {
			float, _ := strconv.ParseFloat(v, 64)
			total += float
			cnt++
		}
		fmt.Printf(" 总共有%d组数据，平均数为%.03f\n ", cnt, total)
		res.Body.Close()
		time.Sleep(2 * time.Second)
		total = 0
	}
}
