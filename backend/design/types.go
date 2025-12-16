package design

import (
	. "goa.design/goa/v3/dsl" //nolint:staticcheck
)

// Тренировка
var WorkoutModel = Type("WorkoutModel", func() {
	Description("Модель тренировок")

	Attribute("title", String, "Название тренировки", func() {
		MinLength(1)
		MaxLength(50)
		Example("Тренировка верхней поверхности бедра")
	})

	Attribute("date", String, "Дата проведение тренировки", func() {
		Format(FormatDate)
		Example("2025-12-08") // YYYY-MM-DD
	})

	Attribute("duration", Int, "Продолжительность тренировки в секундах", func() {
		Minimum(60)
		Maximum(14400)
		Example("5440")
		Description("На фронте нужно будет преобразовывать из секунд в дату, так удобнее хранить в БД")
	})

	// Attribute("exercises", ArrayOf(ExerciseModel), "Упражнения выполненные за тренировку", func() {
	// 	MinLength(1)
	// })
})

var ExercisesModel = Type("ExercisesModel", func() {
	Description("Список доступных упражнений")

	Attribute("title", String, "Название упражнения", func() {
		MinLength(1)
		MaxLength(50) // TODO: подобрать значение потом
		Example("Жим лежа на скамье")
	})

	Attribute("muscle_group", Int32, "Группа мыщц указывается в виде числа, нужна для получения упражнений по группе мыщц", func() {
		Minimum(0)
		Maximum(10) // TODO: посчитать общее количество групп и указать валидное значение, пока стоит рандом
		Example(1)
	})
})

var ExercisesPayload = Type("ExercisePayload", func() {
	Description("Необходимые поля для создания упражнения")

	Extend(ExercisesModel)
	Required("title", "muscle_group")
})

var Exercises = Type("Exercises", func() {
	Description("Модель Списка Упражнений с UUID")

	Attribute("id", String, "Unique concert identifier", func() {
		Format(FormatUUID)
		Example("550e8400-e29b-41d4-a716-446655440000")
		Description("System-generated unique identifier")
	})

	Extend(ExercisesModel)
	Required("id", "title", "muscle_group")
})
