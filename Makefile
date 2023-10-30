KUBECONFIG=$(HOME)/.kube/dev
level=info

test:
	go fmt ./...
	go vet ./...
	go test ./...
	go mod tidy
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run -v

run:
	go run ./cmd -kubeconfig=$(KUBECONFIG) -log-level=$(level)