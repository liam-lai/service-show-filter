package service

import (
	"github.com/liam-lai/service-show-filter/app/schema"
	"github.com/liam-lai/service-show-filter/app/util"
)

func FilterDrmEnabled(serie *schema.Series) *schema.FilteredDrm {
	if serie.Drm == false || serie.EpisodeCount == 0 {
		return nil
	}
	title := util.RemoveBracket(serie.Title)

	return &schema.FilteredDrm{
		Image: serie.Image.ShowImage,
		Slug:  serie.Slug,
		Title: title,
	}
}
