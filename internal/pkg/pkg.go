package pkg

import (
	"sort"
	"sync"

	"go.mattglei.ch/vanityprox/internal/github"
)

type Packages struct {
	packages []github.Repository
	mutex    sync.Mutex
}

func (p *Packages) All() []github.Repository {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	sort.Slice(
		p.packages,
		func(i, j int) bool { return p.packages[j].Updated.Before(p.packages[i].Updated) },
	)
	return p.packages
}
