.PHONY: build
build:
	go build -o semver main.go

test: test-get test-bump test-sort test-rewrite
test-get: build
	./semver get major 1.2.3-alpha.1+123
	./semver get minor 1.2.3-alpha.1+123
	./semver get patch 1.2.3-alpha.1+123
	./semver get pre 1.2.3-alpha.1+123
	./semver get build 1.2.3-alpha.1+123

test-bump: build
	./semver bump major 1.2.3-alpha.1+123
	./semver bump minor 1.2.3-alpha.1+123
	./semver bump patch 1.2.3-alpha.1+123

test-sort: build
	./semver sort 1.2.4 1.2.2 1.2.1 2.0.0
	curl -sS https://artifactory.eng.vmware.com/artifactory/api/storage/tap-builds-generic-local/ | jq -r '.children[] | .uri' | sed 's/\///' | sed 's/\.yaml//' | ./semver sort --range ">1.5.0 <1.7.0"

test-rewrite: build
	./semver rewrite 1.2.3-alpha.1+123 'rel-{{bump_minor.Major}}.{{bump_minor.Minor}}.x'
