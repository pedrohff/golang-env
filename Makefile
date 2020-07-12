env:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0
	cp  ./resources/pre-commit ./.git/hooks/
	chmod +x ./.git/hooks/pre-commit

hello:
	echo "ok!"
	echo "ok!"
