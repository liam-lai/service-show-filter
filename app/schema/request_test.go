package schema_test

import (
	"encoding/json"

	"github.com/liam-lai/service-show-filter/app/schema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Schema", func() {
	Context("Payload", func() {
		It("should umarshal and marshal as expected", func() {
			input := `{
				"payload": [
					{
						"slug": "show/seapatrol",
						"title": "Sea Patrol",
						"tvChannel": "Channel 9"
					}
				],
				"skip": 0,
				"take": 10,
				"totalRecords": 75
			}`
			seriesinfo := &schema.Request{}
			err := json.Unmarshal([]byte(input), &seriesinfo)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
	Context("Series", func() {
		It("should umarshal and marshal as expected", func() {
			input := `{
				"country": "UK",
				"description": "What's life like when you have enough children to field your own football team?",
				"drm": true,
				"episodeCount": 3,
				"genre": "Reality",
				"image": {
					"showImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg"
				},
				"language": "English",
				"nextEpisode": {
					"channel": null,
					"channelLogo": "http://catchup.ninemsn.com.au/img/player/logo_go.gif",
					"date": null,
					"html": "<br><span class=\"visit\">Visit the Official Website</span></span>",
					"url": "http://go.ninemsn.com.au/"
				},
				"primaryColour": "#ff7800",
				"seasons": [
					{
						"slug": "show/16kidsandcounting/season/1"
					}
				],
				"slug": "show/16kidsandcounting",
				"title": "16 Kids and Counting",
				"tvChannel": "GEM"
			}`

			seriesinfo := &schema.Series{}
			err := json.Unmarshal([]byte(input), &seriesinfo)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should umarshal and marshal as required fields", func() {
			input := `{
				"image": {
					"showImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg"
				},
				"slug": "show/16kidsandcounting",
				"title": "16 Kids and Counting"
			}`
			seriesinfo := &schema.Series{}
			err := json.Unmarshal([]byte(input), &seriesinfo)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(seriesinfo.Image).NotTo(BeNil())
			Expect(seriesinfo.Image.ShowImage).NotTo(BeEmpty())
			Expect(seriesinfo.Slug).NotTo(BeEmpty())
			Expect(seriesinfo.Title).NotTo(BeEmpty())

		})
	})

})
