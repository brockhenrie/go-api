package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
)

func Paginate(page int,entity Entity, db *gorm.DB) fiber.Map {

	//entity can be user or product
	limit := 15
	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset)
	total := entity.Count(db)
	
	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"page":       page,
			"total": total,
			"lastPage":   math.Ceil(float64(int(total) / limit)),
		},
	}
}
