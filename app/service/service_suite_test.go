package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDbSelector(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Service Suite")
}
