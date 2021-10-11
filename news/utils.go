package news

type NewsCountPair struct {
	Key   string
	Value int
}

type NewsCountPairList []NewsCountPair

func (p NewsCountPairList) Len() int           { return len(p) }
func (p NewsCountPairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p NewsCountPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
