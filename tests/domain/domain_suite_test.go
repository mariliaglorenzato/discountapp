package domain_test

import (
	"testing"

	"discountapp/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDomain(t *testing.T) {
	RegisterFailHandler(Fail)
	config.SetConfigs("test")
	RunSpecs(t, "Domain Suite")
}
