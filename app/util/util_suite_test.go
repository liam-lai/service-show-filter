package util_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDbSelector(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Util Suite")
}
