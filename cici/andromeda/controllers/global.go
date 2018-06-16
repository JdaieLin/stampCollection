package controllers

import (
	"cici/andromeda/constant"
	"cici/andromeda/models"
	"encoding/json"
)

type BaseController struct {
	Controller
}

func (c *BaseController) All() {
	type All struct {
		models.Serial
		Atlas []models.Atlas
	}
	type Reponse struct {
		All []All
	}
	var req = new(BaseRequest)
	var err = json.Unmarshal(c.PostBody(), req)
	if err != nil {
		c.WriteError(constant.ERR_REQ_BODY, string(c.PostBody()))
		return
	}
	seriales, err := models.GetAllSerials()
	if err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
		return
	}
	atlases, err := models.GetAllAtlas()
	if err != nil {
		c.WriteError(constant.ERR_REQ_SERVER, err.Error())
		return
	}
	var atlasMap = map[int64][]models.Atlas{}
	for _, atlas := range atlases {
		if objs, ok := atlasMap[atlas.SerialID]; ok {
			atlasMap[atlas.SerialID] = append(objs, atlas)
		} else {
			atlasMap[atlas.SerialID] = []models.Atlas{atlas}
		}
	}
	var resp Reponse
	for _, serial := range seriales {
		var all = All{Serial: serial}
		if objs, ok := atlasMap[serial.ID]; ok {
			all.Atlas = objs
		}
		resp.All = append(resp.All, all)
	}
	c.WriteSuccessResult(resp)
}
