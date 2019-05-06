package cmd

type Config struct {
	Owner     string `mapstructure:"owner"`
	Repo      string `mapstructure:"repo"`
	Token     string `mapstructure:"token"`
	Milestone int    `mapstructure:"milestone"`
	State     string `mapstructure:"state"`
	Groups    []struct {
		Labels []string `mapstructure:"labels"`
		Title  string   `mapstructure:"title"`
	} `mapstructure:"groups"`
	Template string `mapstructure:"template"`
}

func (c *Config) AllLabels() map[int][]string {
	allLabels := make(map[int][]string)
	for i, group := range c.Groups {
		if allLabels[i] == nil {
			allLabels[i] = make([]string, 0)
		}
		allLabels[i] = append(allLabels[i], group.Labels...)
	}
	return allLabels
}
