package service_test

import (
	"github.com/liam-lai/service-show-filter/app/schema"
	"github.com/liam-lai/service-show-filter/app/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	Context("FilterDrmEnabled", func() {
		It("should return nil if Drm is false", func() {
			input := &schema.Series{
				Drm: false,
			}
			output := service.FilterDrmEnabled(input)
			Expect(output).To(BeNil())
		})
		It("should return nil if EpisodeCount is 0", func() {
			input := &schema.Series{
				Drm:          true,
				EpisodeCount: 0,
			}
			output := service.FilterDrmEnabled(input)
			Expect(output).To(BeNil())
		})
		It("should return FilteredDrm", func() {
			input := &schema.Series{
				Drm:          true,
				EpisodeCount: 1,
				Title:        "The Taste",
			}
			output := service.FilterDrmEnabled(input)
			Expect(output.Title).To(Equal("The Taste"))
		})
	})

})
