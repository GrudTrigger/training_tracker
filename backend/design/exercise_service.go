package design

import (
	. "goa.design/goa/v3/dsl"
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
			POST("/exercise-list")
			Response(StatusCreated)
			Response("bad_request", StatusBadRequest)
		})
	})
})
