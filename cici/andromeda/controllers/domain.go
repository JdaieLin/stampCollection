package controllers

var DomainInfo = map[string]interface{}{
	"title":  "andromeda",
	"author": "liuwei0932@gmail.com",
}

type MainController struct {
	Controller
}

func (this *MainController) Get() {
	// this.Data["domain"] = conf.Conf().App.Domain
	this.WriteResult(DomainInfo, nil)
}
