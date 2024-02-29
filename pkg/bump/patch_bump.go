package bump

import "github.com/carvel-dev/semver/v4"

type PatchBump struct{}

func (PatchBump) Apply(v semver.Version) semver.Version {
	v.Patch++
	v.Pre = nil
	return v
}
