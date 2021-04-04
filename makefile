test:
	go get github.com/onsi/ginkgo/ginkgo
	ginkgo -r .

dev: test
	go build
	./service-show-filter

dependency: 
	go mod tidy
	
release: dependency test
	go build
