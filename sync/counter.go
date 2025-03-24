package main

import ()

type Counter struct {
	value int
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
	c.value++
}
