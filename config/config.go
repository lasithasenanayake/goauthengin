package config

import (
	"os"
)

type config struct {
	m make(map[])
}

func (c *config)Add(Key, Value string) {
	c.m[key]=Value
	writetoFile()
}

func (c *config)Get(Key string) string {
	return c.m[key]
}
