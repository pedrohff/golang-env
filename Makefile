external-deps:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0

local-env: external-deps
	cp  ./resources/pre-commit ./.git/hooks/
	chmod +x ./.git/hooks/pre-commit

prod-env: external-deps
	go env -w GO111MODULE="on"
	go env -w CGO_ENABLED="0"
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0

docker-build:
	docker build . -f resources/Dockerfile -t $TAG

fetch-deps:
	go mod download

build:
	go build -o bin/myapp -ldflags="-s -w" ./cmd/http

build-on-docker: prod-env build

test:
	go test ./...

reports:
	mkdir reports
	go test ./... -short -coverprofile=reports/cov.out
	printf '\nCreating Go Vet tool analysis file'
	go vet ./... &> reports/vet.out
	printf '\nCreating Go Lint analysis file'
	golint ./... &> reports/lint.out
	printf '\nStatic analysis done!'