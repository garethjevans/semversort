package bump_test

import (
	"github.com/carvel-dev/semver/v4"
	"github.com/garethjevans/semver/pkg/bump"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PatchBump", func() {
	var inputVersion semver.Version
	var b bump.Bump
	var outputVersion semver.Version

	BeforeEach(func() {
		inputVersion = semver.Version{
			Major: 1,
			Minor: 2,
			Patch: 3,
			Pre: semver.PRVersion{
				{VersionStr: "beta"},
				{VersionNum: 1, IsNum: true},
			},
		}

		b = bump.PatchBump{}
	})

	JustBeforeEach(func() {
		outputVersion = b.Apply(inputVersion)
	})

	It("bumps patch and zeroes out the subsequent segments", func() {
		Expect(outputVersion).To(Equal(semver.Version{
			Major: 1,
			Minor: 2,
			Patch: 4,
		}))
	})
})
