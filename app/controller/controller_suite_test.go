package controller_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDbSelector(t *testing.T) {
	RegisterFailHandler(Fail)
	gin.SetMode("test")
	RunSpecs(t, "Controller Suite")
}
