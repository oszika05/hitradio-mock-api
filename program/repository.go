package program

import (
	"gopkg.in/guregu/null.v4"
	"sort"
	"strings"
)

type Repository struct {
	programs          []Program
	episodes          []Episode
	people            []Person
	premiumEpisodeIds map[string]bool
}

func CreateMockRepository() Repository {
	return Repository{
		programs: mockPrograms(),
		episodes: mockEpisodes(),
		people:   mockPerson(),
		premiumEpisodeIds: map[string]bool{
			"kozeppont-1": true,
			"kozeppont-2": true,
		},
	}
}

func (repo Repository) GetPrograms(from int, pageSize int, search null.String) ([]Program, error) {

	if pageSize == 0 {
		return []Program{}, nil
	}

	var filtered []Program

	for _, program := range repo.programs {
		if search.Valid && (strings.Contains(strings.ToLower(program.Name), strings.ToLower(search.String)) || strings.Contains(strings.ToLower(program.Description.String), strings.ToLower(search.String))) {
			filtered = append(filtered, program)
		}

		if !search.Valid {
			filtered = append(filtered, program)
		}
	}

	if from >= len(filtered) {
		return []Program{}, nil
	}

	to := from + pageSize

	if to >= len(filtered) {
		to = len(filtered)
	}

	return filtered[from:to], nil
}

func (repo Repository) GetProgram(id string) (Program, bool, error) {
	for _, program := range repo.programs {
		if program.Id == id {
			return program, true, nil
		}
	}

	return Program{}, false, nil
}

func (repo Repository) GetEpisodes(from int, pageSize int, programId null.String, search null.String, tags []string, people []string, isUserLoggedIn bool) ([]Episode, error) {
	if pageSize == 0 {
		return []Episode{}, nil
	}

	var filtered []Episode

	targetMatches := 0

	if programId.Valid {
		targetMatches += 1
	}

	if search.Valid {
		targetMatches += 1
	}

	if len(tags) > 0 {
		targetMatches += 1
	}

	if len(people) > 0 {
		targetMatches += 1
	}

	for _, episode := range repo.episodes {
		matches := 0

		if programId.Valid && programId.String == episode.Program.Id {
			matches += 1
		}

		if search.Valid && (strings.Contains(strings.ToLower(episode.Title), strings.ToLower(search.String)) || strings.Contains(strings.ToLower(episode.Description.String), strings.ToLower(search.String))) {
			matches += 1
		}

		tagMatch := false
		for _, searchTag := range tags {
			for _, eposideTag := range episode.Tags {
				if eposideTag == searchTag {
					tagMatch = true
					break
				}
			}
		}

		if tagMatch {
			matches += 1
		}

		personMatch := false
		for _, searchPerson := range people {
			for _, episodePerson := range episode.Hosts {
				if episodePerson.Id == searchPerson {
					personMatch = true
					break
				}
			}

			for _, episodePerson := range episode.Guests {
				if episodePerson.Id == searchPerson {
					personMatch = true
					break
				}
			}
		}

		if personMatch {
			matches += 1
		}

		if matches == targetMatches {
			filtered = append(filtered, episode)
		}
	}

	if from >= len(filtered) {
		return []Episode{}, nil
	}

	to := from + pageSize

	if to >= len(filtered) {
		to = len(filtered)
	}

	authFiltered := make([]Episode, len(filtered))
	for i, item := range filtered {
		if !isUserLoggedIn && repo.premiumEpisodeIds[filtered[i].Id] {
			item.AudioUrl = ""
			authFiltered[i] = item
		} else {
			authFiltered[i] = filtered[i]
		}
	}

	return authFiltered[from:to], nil
}

func (repo Repository) GetEpisode(episodeId string, isUserLoggedIn bool) (Episode, bool, error) {
	for _, episode := range repo.episodes {
		if episode.Id == episodeId {
			if !isUserLoggedIn && repo.premiumEpisodeIds[episode.Id] {
				episode.AudioUrl = ""
			}

			return episode, true, nil
		}
	}

	return Episode{}, false, nil
}

func (repo Repository) GetPeople(from int, pageSize int, search null.String, personType PersonType) ([]Person, error) {
	if pageSize == 0 {
		return []Person{}, nil
	}

	var filtered []Person

	targetMatches := 0
	if search.Valid {
		targetMatches += 1
	}
	if personType == PersonTypeHost || personType == PersonTypeGuest {
		targetMatches += 1
	}

	for _, person := range repo.people {
		matches := 0

		if search.Valid && (strings.Contains(strings.ToLower(person.Name), strings.ToLower(search.String)) || strings.Contains(strings.ToLower(person.Description.String), strings.ToLower(search.String))) {
			matches += 1
		}

		if (personType == PersonTypeHost || personType == PersonTypeGuest) && person.Type == personType {
			matches += 1
		}

		if matches == targetMatches {
			filtered = append(filtered, person)
		}
	}

	if from >= len(filtered) {
		return []Person{}, nil
	}

	to := from + pageSize

	if to >= len(filtered) {
		to = len(filtered)
	}

	return filtered[from:to], nil
}

func (repo Repository) GetPerson(personId string) (Person, bool, error) {
	for _, person := range repo.people {
		if person.Id == personId {
			return person, true, nil
		}
	}

	return Person{}, false, nil
}

func (repo Repository) GetRelatedPeople(personId string) ([]Person, bool, error) {
	relatedPeople := map[Person]int{}

	personType := PersonType_Empty

	for _, p := range repo.people {
		if p.Id == personId {
			personType = p.Type
		}
	}

	if personType == PersonType_Empty {
		return []Person{}, false, nil
	}

	for _, episode := range repo.episodes {
		if personType == PersonTypeHost {
			found := false
			for _, host := range episode.Hosts {
				if host.Id == personId {
					found = true
				}
			}

			if !found {
				continue
			}

			for _, host := range episode.Hosts {
				if host.Id != personId {
					relatedPeople[host] += 1
				}
			}
		} else {
			found := false
			for _, guest := range episode.Guests {
				if guest.Id == personId {
					found = true
				}
			}

			if !found {
				continue
			}

			for _, guest := range episode.Guests {
				if guest.Id != personId {
					relatedPeople[guest] += 1
				}
			}
		}
	}

	relatedPeoplePairs := make(PersonCountPairList, len(relatedPeople))
	i := 0
	for k, v := range relatedPeople {
		relatedPeoplePairs[i] = PersonCountPair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(relatedPeoplePairs))

	to := 10

	if to >= len(relatedPeoplePairs) {
		to = len(relatedPeoplePairs)
	}

	result := make([]Person, to)
	for i = 0; i < to; i++ {
		result[i] = relatedPeoplePairs[i].Key
	}

	return result, true, nil
}

func (repo Repository) GetRelatedEpisodes(episodeId string, isUserLoggedIn bool) ([]Episode, bool, error) {
	episodes := map[string]Episode{}
	relatedEpisodes := map[string]int{}

	var episode Episode

	for _, e := range repo.episodes {
		episodes[e.Id] = e
		if e.Id == episodeId {
			episode = e
		}
	}

	if len(episode.Id) == 0 {
		return []Episode{}, false, nil
	}

	for _, e := range repo.episodes {
		points := 0

		if e.Id == episode.Id {
			continue
		}

		for _, guest := range e.Guests {
			for _, g := range episode.Guests {
				if g.Id == guest.Id {
					points += 20
				}
			}
		}

		for _, host := range e.Hosts {
			for _, h := range episode.Hosts {
				if h.Id == host.Id {
					points += 5
				}
			}
		}

		for _, tag := range e.Tags {
			for _, t := range episode.Tags {
				if tag == t {
					points += 15
				}
			}
		}

		relatedEpisodes[e.Id] = points
	}

	relatedEpisodePairs := make(EpisodeCountPairList, len(relatedEpisodes))
	i := 0
	for k, v := range relatedEpisodes {
		relatedEpisodePairs[i] = EpisodeCountPair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(relatedEpisodePairs))

	to := 10

	if to >= len(relatedEpisodePairs) {
		to = len(relatedEpisodePairs)
	}

	result := make([]Episode, to)
	for i = 0; i < to; i++ {
		result[i] = episodes[relatedEpisodePairs[i].Key]
	}

	authFiltered := make([]Episode, len(result))
	for i, item := range result {
		if !isUserLoggedIn && repo.premiumEpisodeIds[result[i].Id] {
			item.AudioUrl = ""
			authFiltered[i] = item
		} else {
			authFiltered[i] = result[i]
		}
	}

	return authFiltered, true, nil
}
