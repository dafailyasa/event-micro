package util

import (
	"github.com/gofiber/fiber/v2"
)

type PaginationParamsStruct struct {
	Page   int
	Limit  int
	Search string
	Sort   int
	Skip   int
}

func QueryPaginationParams(ctx *fiber.Ctx) *PaginationParamsStruct {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	search := ctx.Query("search")
	sort := ctx.QueryInt("sort", -1)

	skip := (page - 1) * limit

	return &PaginationParamsStruct{
		Page:   page,
		Limit:  limit,
		Search: search,
		Sort:   sort,
		Skip:   skip,
	}
}
