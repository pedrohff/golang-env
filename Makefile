env:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0
	cp  ./resources/pre-commit ./.git/hooks/

hello:
	echo "ok!"
