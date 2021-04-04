all: test
	go build
	./service-show-filter

test:
	ginkgo -r .