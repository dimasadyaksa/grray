package grray

import (
	"fmt"
	"sort"
)

type Container = int

type Containers []Container

func NewContainers() *Containers {
	c := make(Containers, 0)
	return &c
}

func (c *Containers) Add(e Container) {
	*c = append(*c, e)
}

func (c *Containers) String() string {
	return fmt.Sprint(*c)
}

func (c *Containers) Map(f func(Container) Container) *Containers {
	nc := NewContainers()
	c.ForEach(func(i int, container Container) {
		nc.Add(f(container))
	})
	return nc
}

func (c *Containers) Find(f func(Container) bool) int {
	for i, e := range *c {
		if f(e) {
			return i
		}
	}

	return -1
}

func (c *Containers) Filter(f func(Container) bool) *Containers {
	nc := NewContainers()
	c.ForEach(func(i int, e Container) {
		if f(e) {
			nc.Add(e)
		}
	})
	return nc
}

func (c *Containers) ForEach(f func(int, Container)) {
	for i, e := range *c {
		f(i, e)
	}
}

func (c *Containers) Sort (f func(prev, next Container ) int ) *Containers  {
	nc := c.Map(func(container Container) Container {
		return container
	})

	s := &sortable{
		c : nc,
		compare: f,
	}

	sort.Sort(s)

	return s.c
}

type sortable struct {
	c *Containers
	compare func(int,int) int
}

func (s *sortable) Len() int {
	return len(*s.c)
}

func (s *sortable) Less(i, j int) bool {
	v := *s.c

	return s.compare(v[i], v[j]) < 0
}

func (s *sortable) Swap(i, j int) {
	v := *s.c
	v[i], v[j] = v[j], v[i]
}
