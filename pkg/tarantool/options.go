package tarantool

type Option func(*options)

type options struct {
	host     string
	port     int
	user     string
	password string
}

func Host(host string) Option {
	return func(c *options) {
		c.host = host
	}
}

func Port(port int) Option {
	return func(c *options) {
		c.port = port
	}
}

func User(user string) Option {
	return func(c *options) {
		c.user = user
	}
}

func Password(password string) Option {
	return func(c *options) {
		c.password = password
	}
}
