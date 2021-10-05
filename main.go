package main

import (
	"github.com/gin-gonic/gin"
	"hitradio-mock-api/news"
	"strconv"
)

func handleErr(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": err.Error()})
}

func setupNewsEndpoint(r *gin.Engine) {
	repo := news.CreateMockRepository()

	r.GET("/news", func(c *gin.Context) {

		fromStr := c.Query("from")
		pageSizeStr := c.Query("pageSize")

		if len(fromStr) == 0 {
			fromStr = "0"
		}

		if len(pageSizeStr) == 0 {
			pageSizeStr = "15"
		}

		from, err := strconv.ParseInt(fromStr, 10, 32)

		if err != nil {
			handleErr(c, 400, err)
			return
		}

		pageSize, err := strconv.ParseInt(pageSizeStr, 10, 32)

		if err != nil {
			handleErr(c, 400, err)
			return
		}

		search := c.Query("search")

		tags := c.QueryArray("tag")

		newsList, err := repo.Get(int(from), int(pageSize), search, tags)

		if err != nil {
			handleErr(c, 500, err)
			return
		}

		c.JSON(200, newsList)
	})
}

func main() {
	r := gin.Default()

	setupNewsEndpoint(r)

	r.Run()
}
