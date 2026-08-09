package models

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
func (a *aiagentskillruleMdoel) Run() error {
	return nil
}
