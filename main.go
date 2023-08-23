package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
)

func handleAPI(c *gin.Context) {
	url := c.Query("url")
	//key := c.Query("key")

	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "URL param not set"})
		return
	}

	res, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Unable to fetch URL"})
		fmt.Println(err)
		fmt.Println(url)
		return
	}

	body, _ := io.ReadAll(res.Body)
	//strBody := string(body)

	defer res.Body.Close()
	//data := gin.H{
	//	"url": url,
	//	"key": key,
	//}

	c.Data(http.StatusOK, "text/html; charset=utf-8", body)

}

func main() {
	u, _ := url.Parse("http://192.168.2.34:7890")
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(u)}

	router := gin.Default()
	router.GET("api/data/", handleAPI)
	router.Run(":8888")
}
