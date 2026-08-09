package models

import (
	"fmt"

	"github.com/ghf-go/goapp/gtools/app/utils"
)

type aiagentskillruleMdoel struct {
}

func NewAiAgentSkillRuleModel() *aiagentskillruleMdoel {
	return &aiagentskillruleMdoel{}
}

func (a *aiagentskillruleMdoel) GetDesc() string {
	return "自动生产ai编程的配置信息"
}
func (a *aiagentskillruleMdoel) GetUsage() string {
	return `create 分组名称 项目名称    # 创建一个项目
    update 分组名称            # 更新一个项目
    init   分组名称            # 在当前目录初始化
    list               # 列出所有分组
	`
}
func (a *aiagentskillruleMdoel) Run(args *utils.Args) error {
	switch args.GetIndex(0) {
	case "list":
		return a.listGroup()
		// case "create":
		// 	return a.create(args.GetIndex(1), args.GetIndex(2))
		// case "update":
		// 	return a.update(args.GetIndex(1))
		// case "init":
		// 	return a.init(args.GetIndex(1))
	}
	return nil
}
func (a *aiagentskillruleMdoel) update(dirname string, gname string) error {
	return nil
}
func (a *aiagentskillruleMdoel) listGroup() error {
	type Resp struct {
		List []struct {
			Name string `json:"name"`
			Desc string `json:"description"`
		} `json:"list"`
		Total int `json:"total"`
	}
	ret := &Resp{}
	e := utils.GetConfig().ApiPost("/api/open/aicode/groupList", map[string]any{}, ret)
	if e != nil {
		return e
	}
	for _, v := range ret.List {
		fmt.Printf("%s\t%s\n", v.Name, v.Desc)
	}
	return nil
}
