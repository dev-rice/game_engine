package resources

type Resources struct {
	gold int64
}

func NewResources(gold int64) Resources {
	return Resources{gold: gold}
}

func (r Resources) GetGold() int64 {
	return r.gold
}

func (r *Resources) SetGold(g int64) {
	r.gold = g
}

func (r *Resources) AddGold(g int64) {
	r.gold += g
}

func (r *Resources) TakeAwayGold(g int64) {
	r.gold -= g
}