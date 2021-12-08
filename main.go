package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"hitradio-mock-api/auth"
	"hitradio-mock-api/news"
	"hitradio-mock-api/program"
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

	r.GET("/news/:id", func(c *gin.Context) {

		newsId := c.Param("id")

		newsItem, found, err := repo.GetNewsItem(newsId)

		if err != nil {
			handleErr(c, 500, err)
			return
		}

		if !found {
			handleErr(c, 404, errors.New("not found"))
			return
		}

		c.JSON(200, newsItem)
	})

	r.GET("/news/:id/related", func(c *gin.Context) {

		newsId := c.Param("id")

		newsList, found, err := repo.GetRelated(newsId)

		if err != nil {
			handleErr(c, 500, err)
			return
		}

		if !found {
			handleErr(c, 404, errors.New("not found"))
			return
		}

		c.JSON(200, newsList)
	})
}

func setupProgramEndpoint(r *gin.Engine) {
	repo := program.CreateMockRepository()

	r.GET("/program", func(c *gin.Context) {
		fromStr := c.Query("from")
		pageSizeStr := c.Query("pageSize")
		searchStr := c.Query("search")

		if len(fromStr) == 0 {
			fromStr = "0"
		}

		if len(pageSizeStr) == 0 {
			pageSizeStr = "15"
		}

		search := null.StringFrom(searchStr)
		if len(searchStr) == 0 {
			search.Valid = false
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

		programs, err := repo.GetPrograms(int(from), int(pageSize), search)

		if err != nil {
			handleErr(c, 500, err)
			return
		}

		c.JSON(200, programs)
	})

	r.GET("/program/:id", func(c *gin.Context) {

		id := c.Param("id")

		if len(id) == 0 {
			handleErr(c, 400, errors.New("No id provided"))
			return
		}

		programs, found, err := repo.GetProgram(id)

		if err != nil {
			handleErr(c, 500, err)
			return
		}

		if !found {
			handleErr(c, 404, errors.New("not found"))
			return
		}

		c.JSON(200, programs)
	})

	r.GET("/episode", func(c *gin.Context) {

		token := c.GetHeader("token")
		isUserLoggedIn, err := auth.IsUserLoggedIn(context.Background(), token)

		if err != nil {
			handleErr(c, 400, err)
			return
		}

		fromStr := c.Query("from")
		pageSizeStr := c.Query("pageSize")
		searchStr := c.Query("search")
		programIdStr := c.Query("programId")
		tags := c.QueryArray("tag")
		people := c.QueryArray("person")

		if len(fromStr) == 0 {
			fromStr = "0"
		}

		if len(pageSizeStr) == 0 {
			pageSizeStr = "15"
		}

		search := null.StringFrom(searchStr)
		if len(searchStr) == 0 {
			search.Valid = false
		}

		programId := null.StringFrom(programIdStr)
		if len(programIdStr) == 0 {
			programId.Valid = false
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

		episodes, err := repo.GetEpisodes(int(from), int(pageSize), programId, search, tags, people, isUserLoggedIn)

		if err != nil {
			handleErr(c, 500, err)
			return
		}

		c.JSON(200, episodes)
	})

	r.GET("/person", func(c *gin.Context) {

		fromStr := c.Query("from")
		pageSizeStr := c.Query("pageSize")
		searchStr := c.Query("search")
		personTypeStr := c.Query("type")

		if len(fromStr) == 0 {
			fromStr = "0"
		}

		if len(pageSizeStr) == 0 {
			pageSizeStr = "15"
		}

		search := null.StringFrom(searchStr)
		if len(searchStr) == 0 {
			search.Valid = false
		}

		personType := program.PersonType(personTypeStr)

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

		people, err := repo.GetPeople(int(from), int(pageSize), search, personType)

		if err != nil {
			handleErr(c, 500, err)
			return
		}

		c.JSON(200, people)
	})

	r.GET("/person/:id/related", func(c *gin.Context) {

		personId := c.Param("id")

		people, found, err := repo.GetRelatedPeople(personId)

		if err != nil {
			handleErr(c, 500, err)
			return
		}

		if !found {
			handleErr(c, 404, errors.New("not found"))
			return
		}

		c.JSON(200, people)
	})

	r.GET("/person/:id", func(c *gin.Context) {

		personId := c.Param("id")

		people, found, err := repo.GetPerson(personId)

		if err != nil {
			handleErr(c, 500, err)
			return
		}

		if !found {
			handleErr(c, 404, errors.New("not found"))
			return
		}

		c.JSON(200, people)
	})

	r.GET("/episode/:id/related", func(c *gin.Context) {

		token := c.GetHeader("token")
		isUserLoggedIn, err := auth.IsUserLoggedIn(context.Background(), token)

		fmt.Printf("isUserLoggedIn: %v, err: %v\n", isUserLoggedIn, err)

		if err != nil {
			handleErr(c, 400, err)
			return
		}

		episodeId := c.Param("id")

		episodes, found, err := repo.GetRelatedEpisodes(episodeId, isUserLoggedIn)

		if err != nil {
			handleErr(c, 500, err)
			return
		}

		if !found {
			handleErr(c, 404, errors.New("not found"))
			return
		}

		c.JSON(200, episodes)
	})

	r.GET("/episode/:id", func(c *gin.Context) {

		token := c.GetHeader("token")
		isUserLoggedIn, err := auth.IsUserLoggedIn(context.Background(), token)

		if err != nil {
			handleErr(c, 400, err)
			return
		}

		episodeId := c.Param("id")

		episodes, found, err := repo.GetEpisode(episodeId, isUserLoggedIn)

		if err != nil {
			handleErr(c, 500, err)
			return
		}

		if !found {
			handleErr(c, 404, errors.New("not found"))
			return
		}

		c.JSON(200, episodes)
	})

}

func main() {
	r := gin.Default()

	setupNewsEndpoint(r)
	setupProgramEndpoint(r)

	r.Run()
}
