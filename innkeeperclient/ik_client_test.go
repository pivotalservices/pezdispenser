package innkeeperclient_test

import (
	"net/http"

	. "github.com/pivotal-pez/pezdispenser/innkeeperclient"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Given IkClient", func() {
	Describe("Given .ProvisionHost()", func() {
		Context("When called with valid sku and auth", func() {
			var (
				err               error
				res               *ProvisionHostResponse
				server            *ghttp.Server
				innkeeperUser     = "admin"
				innkeeperPassword = "pass"
			)
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyBasicAuth(innkeeperUser, innkeeperPassword),
						ghttp.RespondWith(http.StatusOK, `{ "status": "success", "data": [{"requestid": "28ac758e-a02c-11e5-9531-0050569b9b57"}], "message": "ok" }`),
					),
				)
				ikClient := New(server.URL(), innkeeperUser, innkeeperPassword)
				res, err = ikClient.ProvisionHost("m1.small", "pez-owner")
			})
			AfterEach(func() {
				server.Close()
			})
			It("Then it should respond with success status and a request id", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(res.Status).Should(Equal(StatusSuccess))
				Ω(res.Data[0].RequestID).ShouldNot(Equal(""))
			})
		})
	})
	Describe("Given .GetStatus() method", func() {

		Context("When called with a valid requestid and that requestid has a status 'success'", func() {
			var (
				err               error
				res               *GetStatusResponse
				server            *ghttp.Server
				innkeeperUser     = "admin"
				innkeeperPassword = "pass"
			)
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyBasicAuth(innkeeperUser, innkeeperPassword),
						ghttp.RespondWith(http.StatusOK, `{
							"status": "success",
							"data": {
								"status": "ready",
								"credentials": {"name": "host-07-25", "oob_ip": "10.65.70.125", "oob_user": "user", "oob_pw": "xxxx"}
							},
							"message": "ok" 
						}`),
					),
				)
				ikClient := New(server.URL(), innkeeperUser, innkeeperPassword)
				res, err = ikClient.GetStatus("requestid")
			})
			AfterEach(func() {
				server.Close()
			})
			It("Then it should return an object containing current state from innkeeper", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(res).ShouldNot(BeNil())
				Ω(res.Data.Status).Should(Equal(StatusReady))
				Ω(res.Data.Credentials).ShouldNot(BeNil())
			})
		})
		Context("When called with a valid requestid that is 'running'", func() {
			var (
				err               error
				res               *GetStatusResponse
				server            *ghttp.Server
				innkeeperUser     = "admin"
				innkeeperPassword = "pass"
			)
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyBasicAuth(innkeeperUser, innkeeperPassword),
						ghttp.RespondWith(http.StatusOK, `{
							"status": "success",
							"data": {
								"status": "running",
								"credentials": {"name": "host-07-25", "oob_ip": "10.65.70.125", "oob_user": "user", "oob_pw": "xxxx"}
							},
							"message": "ok" 
						}`),
					),
				)
				ikClient := New(server.URL(), innkeeperUser, innkeeperPassword)
				res, err = ikClient.GetStatus("requestid")
			})
			AfterEach(func() {
				server.Close()
			})
			It("Then it should return an object containing current state from innkeeper", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(res).ShouldNot(BeNil())
				Ω(res.Data.Status).Should(Equal(StatusRunning))
			})
		})
	})
})
