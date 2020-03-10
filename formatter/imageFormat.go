package formatter

import "github.com/DualVectorFoil/stem/model"

func ImageFormat(models []*model.ImageDetailModel) []*ImageDetailResp {
	if models == nil {
		return nil
	}

	resp := []*ImageDetailResp{}
	for _, m := range models {
		if m == nil {
			continue
		}
		resp = append(resp, &ImageDetailResp{
			Url:    m.Url,
			Width:  m.Width,
			Height: m.Height,
		})
	}
	return resp
}

type ImageDetailResp struct {
	Url    string  `json:"url"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
}
