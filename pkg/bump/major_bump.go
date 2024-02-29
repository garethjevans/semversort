package bump

import "github.com/carvel-dev/semver/v4"

type MajorBump struct{}

func (MajorBump) Apply(v semver.Version) semver.Version {
	v.Major++
	v.Minor = 0
	v.Patch = 0
	v.Pre = nil
	return v
}
