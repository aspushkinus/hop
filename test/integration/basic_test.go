package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var host = "http://example_basic:8000"

var _ = Describe("Basic", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error

		capabilities := agouti.NewCapabilities().Browser("chrome").Platform("linux").With("javascriptEnabled")
		page, err = agouti.NewPage("http://webdriver:4444/wd/hub", agouti.Desired(capabilities))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("works", func() {
		Expect(page.Navigate(host + "/Main.elm")).To(Succeed())
		Expect(page).To(HaveURL(host + "/Main.elm"))

		title := page.Find(".title")
		Eventually(title).Should(HaveText("Main"))
	})

	It("shows main", func() {
		Expect(page.Navigate(host + "/Main.elm#/")).To(Succeed())

		title := page.Find(".title")
		Eventually(title).Should(HaveText("Main"))
	})

	It("show about", func() {
		Expect(page.Navigate(host + "/Main.elm#/about")).To(Succeed())
		title := page.Find(".title")
		Eventually(title).Should(HaveText("About"))
	})
})
