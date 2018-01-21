test:
	go test -covermode=atomic
cover:
	go test -coverprofile=coverage.txt -covermode=atomic
	go tool cover -html=coverage.txt -o coverage.html
	open coverage.html
changelog:
	github_changelog_generator --no-issues suzuki-shunsuke/go-chatwords
