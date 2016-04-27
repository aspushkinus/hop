package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var host = "http://localhost:8000"

var _ = Describe("Basic", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("works", func() {
		By("showing content", func() {
			Expect(page.Navigate(host + "/Main.elm")).To(Succeed())
			Expect(page).To(HaveURL(host + "/Main.elm"))

			title := page.Find(".title")
			// Eventually(title.Should(HaveText("Main")))
			Eventually(title).Should(HaveText("Main"))
		})
	})
})
