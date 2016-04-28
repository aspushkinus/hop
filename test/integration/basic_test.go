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
		// page, err = agoutiDriver.NewPage()

		page, err = agouti.NewPage("http://webdriver:4444/wd/hub")
		// page, err = agouti.NewPage("http://127.0.0.1:4444/wd/hub")

		// Cannot use SauceLabs as it needs to test app in localhost
		// needs SauceLabs Connect for this which is not supported
		// page, err = agouti.SauceLabs("firefox", "Linux", "firefox", "33", "sporto", "xxx")

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

	It("show about", func() {
		Expect(page.Navigate(host + "/Main.elm#/about")).To(Succeed())
		title := page.Find(".title")
		Eventually(title).Should(HaveText("About"))
	})
})
