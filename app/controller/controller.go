package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liam-lai/service-show-filter/app/schema"
	s "github.com/liam-lai/service-show-filter/app/service"
	"github.com/rs/zerolog/log"
)

func Status(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
func DrmEnabled(c *gin.Context) {
	seriesInfo := schema.Request{}
	err := c.ShouldBindJSON(&seriesInfo)
	if err != nil {
		decodeErr := schema.DecodeErr{Err: err}
		log.Warn().Err(decodeErr.Err).Msg(decodeErr.ResError())
		c.JSON(http.StatusBadRequest, gin.H{"error": decodeErr.ResError()})
		return
	}

	res := schema.Response{}
	for _, serie := range seriesInfo.Payload {
		filtered := s.FilterDrmEnabled(&serie)
		if filtered != nil {
			res.Response = append(res.Response, filtered)
		}
	}

	c.JSON(200, res)
}
