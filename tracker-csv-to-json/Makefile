test: vet install-ginkgo
	ginkgo --race --randomize-all -r .

vet:
	go vet ./...

install-ginkgo:
	go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
