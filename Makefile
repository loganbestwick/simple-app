fix:
	git ls-files -- \*.go | xargs gofmt -s -w
	git ls-files -- \*.go | xargs goimports -w
