package main

import (
	"github.com/d2r2/go-dht"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type (
	Data struct {
		Farenheit float32 `json:"farenheit"`
		Celsius float32 `json:"celsius"`
		Humidity float32 `json:"humidity"`
		Retries int `json:"retries"`
	}
)

func main() {
	var mutex = &sync.Mutex{}
	var data Data


	go func() {
		var tmp Data
		for {
			mutex.Lock()
			tmp.Celsius,tmp.Humidity, tmp.Retries , _ = dht.ReadDHTxxWithRetry(dht.DHT22, 18, true, 10)
			tmp.Farenheit=tmp.Celsius * 1.8 + 32
			data = tmp
			mutex.Unlock()
			time.Sleep(time.Second * 10)
		}
	}()

	r := gin.Default()
	r.GET("/data", func(c *gin.Context) {
		mutex.Lock()
		c.JSON(200, data)
		mutex.Unlock()
	})
	r.Run(":80")
}
