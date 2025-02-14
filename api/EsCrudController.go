package api

import (
	"github.com/1340691923/ElasticView/pkg/escache"
	es2 "github.com/1340691923/ElasticView/service/es"
	"github.com/gofiber/fiber/v2"
)

// ES CRUD操作
type EsCrudController struct {
	BaseController
}

// 可视化筛选获取数据
func (this EsCrudController) GetList(ctx *fiber.Ctx) error {
	crudFilter := new(escache.CrudFilter)
	err := ctx.BodyParser(crudFilter)
	if err != nil {
		return this.Error(ctx, err)
	}
	esConnect, err := escache.GetEsClientByID(crudFilter.EsConnect)
	if err != nil {
		return this.Error(ctx, err)
	}

	esService, err := es2.NewEsService(esConnect)
	if err != nil {
		return this.Error(ctx, err)
	}

	return esService.CrudGetList(ctx, crudFilter)
}

// 可视化GetDSL
func (this EsCrudController) GetDSL(ctx *fiber.Ctx) error {
	crudFilter := new(escache.CrudFilter)
	err := ctx.BodyParser(crudFilter)
	if err != nil {
		return this.Error(ctx, err)
	}
	esConnect, err := escache.GetEsClientByID(crudFilter.EsConnect)
	if err != nil {
		return this.Error(ctx, err)
	}

	esService, err := es2.NewEsService(esConnect)
	if err != nil {
		return this.Error(ctx, err)
	}

	return esService.CrudGetDSL(ctx, crudFilter)
}

// 下载
func (this EsCrudController) Download(ctx *fiber.Ctx) error {

	crudFilter := new(escache.CrudFilter)
	err := ctx.BodyParser(crudFilter)
	if err != nil {
		return this.Error(ctx, err)
	}
	esConnect, err := escache.GetEsClientByID(crudFilter.EsConnect)
	if err != nil {
		return this.Error(ctx, err)
	}

	esService, err := es2.NewEsService(esConnect)
	if err != nil {
		return this.Error(ctx, err)
	}

	return esService.CrudDownload(ctx, crudFilter)
}
