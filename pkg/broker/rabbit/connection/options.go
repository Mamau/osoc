package connection

import "time"

type Option func(*options)

type options struct {
	protocol  string
	user      string
	password  string
	host      string
	locale    string
	port      int
	heartbeat time.Duration
}

func Locale(l string) Option {
	return func(c *options) {
		c.locale = l
	}
}
func Heartbeat(h time.Duration) Option {
	return func(c *options) {
		c.heartbeat = h
	}
}

func Port(p int) Option {
	return func(c *options) {
		c.port = p
	}
}

func Host(h string) Option {
	return func(c *options) {
		c.host = h
	}
}

func Password(p string) Option {
	return func(c *options) {
		c.password = p
	}
}

func User(l string) Option {
	return func(c *options) {
		c.user = l
	}
}

func Protocol(p string) Option {
	return func(c *options) {
		c.protocol = p
	}
}
