package bump

import "github.com/carvel-dev/semver/v4"

type Bump interface {
	Apply(semver.Version) semver.Version
}

type IdentityBump struct{}

func (IdentityBump) Apply(v semver.Version) semver.Version {
	return v
}
