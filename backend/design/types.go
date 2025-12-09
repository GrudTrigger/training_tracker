package design

import (
	. "goa.design/goa/v3/dsl"
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
		Example("2025-12-08") //YYYY-MM-DD
	})

	Attribute("duration", Int, "Продолжительность тренировки в секундах", func() {
		Minimum(60)
		Maximum(14400)
		Example("5440")
		Description("На фронте нужно будет преобразовывать из секунд в дату, так удобнее хранить в БД")
	})

	Attribute("exercises", ArrayOf(ExerciseModel), "Упражнения выполненные за тренировку", func() {
		MinLength(1)
	})
})

// Упражнение
var ExerciseModel = Type("ExerciseModel", func() {
	Attribute("title", String, "Название упражнения", func() {
		MinLength(1)
		MaxLength(50)
		Example("Жим лежа на скамейке")
	})

	Attribute("muscle_group", Int32, "Группа мыщц указывается в виде числа, нужна для получения упражнений по группе мыщц", func() {
		Minimum(0)
		Maximum(10) //TODO: посчитать общее количество групп и указать валидное значение, пока стоит рандом
		Example(1)
	})

	Attribute("sets", Int32, "Количество подходов", func() {
		Minimum(1)
		Maximum(100)
		Example(4)
	})

	Attribute("reps", Int32, "Количество повторений в подходе", func() {
		Minimum(1)
		Maximum(30)
		Example(10)
	})

	Attribute("weight", Float32, "Вес в кг", func() {
		Minimum(1)
		Maximum(500)
		Example(102.5)
	})
})
