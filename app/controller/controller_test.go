package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/liam-lai/service-show-filter/app/controller"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Controller", func() {
	Context("DrmEnabled", func() {
		It("should return 400 with application/json type if payload is invalid", func() {
			//Initial Response Object
			resp := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(resp)

			//Initial Request Connection Content
			body := bytes.NewBufferString("dummy string")
			c.Request, _ = http.NewRequest("POST", "/", body)
			c.Request.Header.Add("Content-Type", gin.MIMEJSON)

			//Test Function
			controller.DrmEnabled(c)

			//Assert Response
			Expect(resp.Code).To(Equal(400))
			Expect(resp.HeaderMap["Content-Type"][0]).To(Equal("application/json; charset=utf-8"))
			Expect(resp.Body.String()).To(MatchJSON(`
			{
				"error": "Could not decode request: JSON parsing failed"
			}`))
		})
		It("should return 200 with application/json type if payload is valid json", func() {
			//Initial Response Object
			postBody := map[string]interface{}{}
			body, err := json.Marshal(postBody)
			Expect(err).To(BeNil())

			resp := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(resp)

			//Initial Request Connection Content
			c.Request, _ = http.NewRequest("POST", "/", bytes.NewReader(body))
			c.Request.Header.Add("Content-Type", gin.MIMEJSON)

			//Test Function
			controller.DrmEnabled(c)

			//Assert Response
			Expect(resp.Code).To(Equal(200))
			Expect(resp.HeaderMap["Content-Type"][0]).To(Equal("application/json; charset=utf-8"))
			Expect(resp.Body.String()).To(MatchJSON(`{}`))
		})
	})

})
