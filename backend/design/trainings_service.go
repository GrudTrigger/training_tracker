package design

import (
	. "goa.design/goa/v3/dsl" //nolint:staticcheck
)

var _ = Service("trainings", func() {
	Description("Сервис для CRUD операция с моделью Training(Тренировки)")

	Method("create", func() {
		Description("Создание тренировки с упражнениями и подходами")
		Meta("openapi:summary", "Создание тренировки")

		Payload(CreateTrainingPayload)
		Result(Training)

		Error("bad_request", ErrorResult)
		Error("not_found", ErrorResult, "Упражнение не найдено")

		HTTP(func() {
			POST("/trainings")
			Response(StatusCreated)
			Response("bad_request", StatusBadRequest)
			Response("not_found", StatusNotFound)
		})
	})

	Method("all", func() {
		Description("Получение всех своих тренировок")
		Meta("openapi:summary", "Получение всех тренировок")

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

		Result(ArrayOf(TrainingAll))

		Error("bad_request", ErrorResult)

		HTTP(func() {
			GET("trainings/all")
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

	Method("delete", func() {
		Description("Удаление тренировки")
		Meta("openapi:summary", "Удаление тренировки по uuid")

		Payload(func() {
			Attribute("uuid", String, func() {
				Format(FormatUUID)
				Example("550e8400-e29b-41d4-a716-446655440000")
			})
			Required("uuid")
		})

		Error("not_found", ErrorResult, "Тренировка не найдена")

		HTTP(func() {
			DELETE("trainings/{uuid}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
	})
})
