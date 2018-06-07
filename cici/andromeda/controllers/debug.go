package controllers

import (
	"cici/andromeda/models"
)

type DebugController struct {
	Controller
}

func (c *DebugController) Stamps() {
	type Response struct {
		*models.CacheStamp
		SliceLenth int
		MapLenth   int
	}
	var resp = new(Response)
	resp.CacheStamp = models.CacheStampApp
	resp.SliceLenth = len(resp.CacheStamp.Slice)
	resp.MapLenth = len(resp.CacheStamp.Map)
	c.WriteSuccessResult(resp)
}
