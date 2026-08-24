package action

import (
	"github.com/ghf-go/goapp/apppackage/app/models"
	"github.com/ghf-go/goapp/apppackage/app/utils"
	"github.com/gin-gonic/gin"
)

type projectSaveReq struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name" binding:"required"`
	GitUrl     string `json:"gitUrl" binding:"required"`
	Type       string `json:"type" binding:"required"`
	TestBranch string `json:"testBranch"`
	ProdBranch string `json:"prodBranch"`
}

// ProjectSaveAction 新增或更新工程（id>0 为更新）
func ProjectSaveAction(c *gin.Context) {
	var req projectSaveReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "参数错误: "+err.Error())
		return
	}
	if !validProjectType(req.Type) {
		utils.Fail(c, "工程类型不合法，可选: mobile/desktop/linux/web")
		return
	}
	if req.ID > 0 {
		updateProject(c, req)
		return
	}
	project := models.ProjectModel{
		Name:        req.Name,
		GitUrl:      req.GitUrl,
		Type:        req.Type,
		TestBranch:  req.TestBranch,
		ProdBranch:  req.ProdBranch,
		TestVersion: "1.0.0",
		ProdVersion: "1.0.0",
	}
	if err := utils.DB.Create(&project).Error; err != nil {
		utils.Fail(c, "保存失败: "+err.Error())
		return
	}
	utils.DB.First(&project, project.ID)
	utils.Success(c, project)
}

func updateProject(c *gin.Context, req projectSaveReq) {
	var project models.ProjectModel
	if err := utils.DB.First(&project, req.ID).Error; err != nil {
		utils.Fail(c, "工程不存在")
		return
	}
	updates := map[string]any{
		"name":        req.Name,
		"git_url":     req.GitUrl,
		"type":        req.Type,
		"test_branch": req.TestBranch,
		"prod_branch": req.ProdBranch,
	}
	if err := utils.DB.Model(&project).Updates(updates).Error; err != nil {
		utils.Fail(c, "更新失败: "+err.Error())
		return
	}
	utils.Success(c, nil)
}

func validProjectType(t string) bool {
	switch t {
	case models.ProjectTypeMobile, models.ProjectTypeDesktop,
		models.ProjectTypeLinux, models.ProjectTypeWeb:
		return true
	}
	return false
}
