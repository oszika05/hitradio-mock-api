package news

import "strings"

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
