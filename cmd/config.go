package cmd

type Config struct {
	Owner     string `yaml:"owner"`
	Repo      string `yaml:"repo"`
	Milestone int    `yaml:"milestone"`
	State     string `yaml:"state"`
	Groups    []struct {
		Labels []string `yaml:"labels"`
		Title  string   `yaml:"title"`
	} `yaml:"groups"`
	Template string `yaml:"template"`
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
