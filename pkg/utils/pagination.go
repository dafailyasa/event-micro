package util

import (
	"math"

	"github.com/gofiber/fiber/v2"
)

type PaginationParamsStruct struct {
	Page   int
	Limit  int
	Search string
	Sort   int
	Skip   int
}

type MetaPaginationStruct struct {
	LastPage    int   `json:"lastPage"`
	CurrentPage int   `json:"currentPage"`
	ResultCount int64 `json:"resultCount"`
	Total       int64 `json:"total"`
	Limit       int   `json:"limit"`
}
type PaginationResStruct struct {
	Data interface{}          `json:"data"`
	Meta MetaPaginationStruct `json:"meta"`
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

func TransformPaginationRes(data interface{}, resultCount int64, total int64, query *PaginationParamsStruct) *PaginationResStruct {
	lastPage := math.Ceil(float64(resultCount) / float64(query.Limit))

	return &PaginationResStruct{
		Data: data,
		Meta: MetaPaginationStruct{
			LastPage:    int(lastPage),
			CurrentPage: query.Page,
			ResultCount: resultCount,
			Total:       total,
			Limit:       query.Limit,
		},
	}
}
