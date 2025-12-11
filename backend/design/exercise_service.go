package design

import (
	. "goa.design/goa/v3/dsl" //nolint:staticcheck
)

var _ = Service("exercise", func() {
	Description("Сервис для CRUD операция с моделью ExerciseList")

	Method("create", func() {
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

	Method("update", func() {
		Description("Редактирование упражнения")
		Meta("openapi:summary", "Редактирование упражнения")

		Payload(func() {
			Extend(ExerciseListPayload)
			Attribute("exerciseId", String, func() {
				Format(FormatUUID)
				Example("550e8400-e29b-41d4-a716-446655440000")
			})
			Required("exerciseId")
		})

		Result(ExerciseList)
		Error("not_found", ErrorResult, "Упражнение не найдено, проверьте UUID")
		Error("bad_request", ErrorResult, "Invalid update date exercise")

		HTTP(func() {
			PUT("exercise/update")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("bad_request", StatusBadRequest)
		})
	})

	Method("delete", func() {
		Description("Удаление упражнения по uuid")
		Meta("openapi:summary", "Удаление упражнения по uuid")

		Payload(func() {
			Attribute("exerciseId", String, func() {
				Format(FormatUUID)
				Example("550e8400-e29b-41d4-a716-446655440000")
			})
			Required("exerciseId")
		})

		Error("not_found", ErrorResult, "Упражнение не найдено")

		HTTP(func() {
			DELETE("exercise/{exerciseId}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
	})
})
