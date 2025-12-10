package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("exercise", func() {
	Description("Сервис для CRUD операция с моделью ExerciseList")

	Method("exercise/create", func() {
		Description("Создание нового упражнения")
		Meta("openapi:summary", "Создание нового упражнения")

		Payload(ExerciseListPayload)
		Result(ExerciseList)

		Error("bad_request", ErrorResult, "Invalid input data provided")

		HTTP(func() {
			POST("/exercise")
			Response(StatusCreated)
			Response("bad_request", StatusBadRequest)
		})
	})

	Method("all", func() {
		Description("Получение всех упражнений с пагинацией")
		Meta("openapi:summary", "Получение всех упражнений с пагинацией")
		Payload(func() {
			Attribute("limit", Int, "Количество запрашиваемых элементов за один offset", func() {
				Minimum(1)
				Default(1)
				Example(20)
			})
			Attribute("offset", Int, "Сколько делать шаг для пагинации", func() {
				Minimum(0)
				Default(0)
				Example(0)
			})
		})

		Result(ArrayOf(ExerciseList))

		HTTP(func() {
			GET("exercise/all")
			Param("limit", Int, "limit", func() {
				Minimum(1)
				Example(20)
			})
			Param("offset", Int, "offset", func() {
				Minimum(0)
				Example(0)
			})
			Response(StatusOK)
		})
	})
})
