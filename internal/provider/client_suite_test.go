package provider

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCostradar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Costradar Suite")
}
