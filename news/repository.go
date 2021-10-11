package news

import (
	"sort"
	"strings"
)

type Repository struct {
	news []News
}

func CreateMockRepository() Repository {
	return Repository{
		news: getMockNews(),
	}
}

func filter(news News, search string, tags []string) bool {
	if len(search) > 0 {
		if !strings.Contains(news.Title, search) && !strings.Contains(news.Content, search) {
			return false
		}
	}

	if len(tags) > 0 {
		hit := false
		for _, newsTag := range news.Tags {
			for _, searchTag := range tags {
				if newsTag == searchTag {
					hit = true
				}
			}
		}

		return hit
	}

	return true
}

func (repo Repository) Get(from int, pageSize int, search string, tags []string) ([]News, error) {
	var filteredNews []News

	for _, n := range repo.news {
		if filter(n, search, tags) {
			filteredNews = append(filteredNews, n)
		}
	}

	to := from + pageSize

	if from >= len(filteredNews) {
		return []News{}, nil
	}

	if to >= len(filteredNews) {
		to = len(filteredNews)
	}

	return filteredNews[from:to], nil
}

func (repo Repository) GetRelated(newsId string) ([]News, bool, error) {
	newsMap := map[string]News{}
	relatedNews := map[string]int{}
	var news News

	for _, n := range repo.news {
		newsMap[n.Id] = n
		if n.Id == newsId {
			news = n
		}
	}

	if len(news.Id) == 0 {
		return []News{}, false, nil
	}

	for _, n := range repo.news {
		if n.Id == newsId {
			continue
		}

		for _, tag := range n.Tags {
			for _, originalTag := range news.Tags {
				if tag == originalTag {
					relatedNews[n.Id] += 1
				}
			}
		}
	}

	relatedNewsPairs := make(NewsCountPairList, len(relatedNews))
	i := 0
	for k, v := range relatedNews {
		relatedNewsPairs[i] = NewsCountPair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(relatedNewsPairs))

	to := 10

	if to >= len(relatedNewsPairs) {
		to = len(relatedNewsPairs)
	}

	result := make([]News, to)
	for i = 0; i < to; i++ {
		result[i] = newsMap[relatedNewsPairs[i].Key]
	}

	return result, true, nil

}
