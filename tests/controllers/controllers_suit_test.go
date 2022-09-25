package controllers_test

import (
	"testing"

	"discountapp/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestContollers(t *testing.T) {
	RegisterFailHandler(Fail)
	config.SetConfigs(config.TestEnv)
	RunSpecs(t, "Controllers Suite")
}
