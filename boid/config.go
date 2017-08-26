package boid

type Config interface {
	Avoidance() float64
}

type config struct {
	avoidance float64
}

func (c *config) Avoidance() float64 {
	return c.avoidance
}

type ConfigBuilder interface {
	Avoidance(value float64) ConfigBuilder
}

type configBuilder struct {
	*config
}

func BuildConfig() ConfigBuilder {
	return &configBuilder{
		config: &config{
			avoidance: 0.0,
		},
	}
}

func (b *configBuilder) Avoidance(value float64) ConfigBuilder {
	b.avoidance = value
	return b
}
