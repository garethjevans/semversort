package bump

import "github.com/carvel-dev/semver/v4"

type MinorBump struct{}

func (MinorBump) Apply(v semver.Version) semver.Version {
	v.Minor++
	v.Patch = 0
	v.Pre = nil
	return v
}
