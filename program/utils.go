package program

type PersonCountPair struct {
	Key   Person
	Value int
}

type PersonCountPairList []PersonCountPair

func (p PersonCountPairList) Len() int           { return len(p) }
func (p PersonCountPairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PersonCountPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type EpisodeCountPair struct {
	Key   string
	Value int
}

type EpisodeCountPairList []EpisodeCountPair

func (p EpisodeCountPairList) Len() int           { return len(p) }
func (p EpisodeCountPairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p EpisodeCountPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
