package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/vntchain/vnt-explorer/common"
	"github.com/vntchain/vnt-explorer/models"
)

type NodeController struct {
	BaseController
}

func (this *NodeController) Post() {
	node := &models.Node{}
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, node)
	if err != nil {
		this.ReturnErrorMsg("Wrong format of Node: %s", err.Error())
		return
	}

	err = node.Insert()
	if err != nil {
		this.ReturnErrorMsg("Failed to create Node: %s", err.Error())
	} else {
		this.ReturnData(node)
	}
}

func (this *NodeController) List() {
	offset, err := this.GetInt("offset")
	if err != nil {
		beego.Warn("Failed to read offset: ", err.Error())
		offset = common.DefaultOffset
	}

	limit, err := this.GetInt("limit")
	if err != nil {
		beego.Warn("Failed to read limit: ", err.Error())
		limit = common.DefaultPageSize
	}

	node := &models.Node{}
	nodes, err := node.List(offset, limit)
	if err != nil {
		this.ReturnErrorMsg("Failed to list nodes: %s", err.Error())
	} else {
		this.ReturnData(nodes)
	}

}

func (this *NodeController) Get() {
	//beego.Info("params", this.Ctx.Input.Params())
	address := this.Ctx.Input.Param(":address")
	if len(address) == 0 {
		this.ReturnErrorMsg("Failed to get address", "")
		return
	}

	node := &models.Node{}
	dbItem, err := node.Get(address)
	if err != nil {
		this.ReturnErrorMsg("Failed to read node: %s", err.Error())
	} else {
		this.ReturnData(dbItem)
	}
}