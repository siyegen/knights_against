package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
)

var _ = Describe("Server", func() {
	// BeforeSuite(func() {
	// 	kaServer := server.Server{}
	// 	kaServer.Run()
	// })
	Describe("/index", func() {
		Context("when a user navigates to the page", func() {
			It("returns a page with a link", func() {
				response, err := http.Get("localhost:7789/")
				Expect(err).To(BeNil())

				body, err := ioutil.ReadAll(response.Body)
				Expect(err).To(BeNil())

				Expect(string(body)).To(MatchRegexp("<a>"))
			})
			It("returns 200 for the status code", func() {
				response, err := http.Get("localhost:7789/")
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
			})
		})
	})
})
