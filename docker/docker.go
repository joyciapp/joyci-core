package docker

// Docker represents a docker container
type Docker struct {
	VolumeDir  string
	WorkDir    string
	Image      string
	Executable string
	Arguments  []string
}
