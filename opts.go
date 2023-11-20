package dextk

type config struct {
	readSlots int
}

type Opt func(*config)

func WithReadCache(slots int) Opt {
	return func(c *config) {
		c.readSlots = slots
	}
}
