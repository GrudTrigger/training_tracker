package design

import (
	. "goa.design/goa/v3/dsl" //nolint:staticcheck
)

var _ = Service("trainings", func(){
	Description("Сервис для CRUD операция с моделью Training(Тренировки)")

	Method("create", func(){
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
})