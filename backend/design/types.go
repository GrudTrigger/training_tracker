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

	Attribute("create_at", String, "Дата создания", func() {
		Format(FormatDateTime)
	})
})

var Training = Type("Training", func() {
	Attribute("id", String, FormatUUID)
	Attribute("title", String)
	Attribute("date", String, func() {
		Format(FormatDate)
	})
	Attribute("duration", Int)
	Attribute("created_at", String, func() {
		Format(FormatDateTime)
	})
	Required("id", "title", "date", "duration")
})

var ExerciseSetPayload = Type("ExerciseSetPayload", func() {
	Attribute("reps", Int, func() {
		Minimum(1)
	})
	Attribute("weight", Float64)
	Required("reps")
})

var TrainingExercisePayload = Type("TrainingExercisePayload", func() {
	Attribute("exercise_id", String, func() {
		Format(FormatUUID)
	})
	Attribute("sets", ArrayOf(ExerciseSetPayload))
	Required("exercise_id", "sets")
})

var CreateTrainingPayload = Type("CreateTrainingPayload", func() {
	Attribute("title", String, func() {
		MaxLength(50)
	})
	Attribute("date", String, func() {
		Format(FormatDate)
		Example("2025-12-25")
	})
	Attribute("duration", Int, func() {
		Minimum(1)
	})
	Attribute("exercises", ArrayOf(TrainingExercisePayload))
	Required("title", "date", "duration", "exercises")
})

//---------------------------

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

var ExerciseSet = Type("ExerciseSet", func() {
	Attribute("id", String, "Unique concert identifier", func() {
		Format(FormatUUID)
		Example("550e8400-e29b-41d4-a716-446655440000")
		Description("System-generated unique identifier")
	})
	Attribute("reps", Int, func() {
		Minimum(1)
	})
	Attribute("weight", Float64)
	Required("reps")
})

var ExercisesWithTraining = Type("ExercisesWithTraining", func() {
	Extend(Exercises)
	Attribute("sets", ArrayOf(ExerciseSet))
})

var TrainingAll = Type("TrainingAll", func() {
	Description("Модель Списка Тренировок с UUID")

	Extend(Training)
	Attribute("exercises", ArrayOf(ExercisesWithTraining))
})

var TrainingsStatistics = Type("TrainingsStatistics", func() {
	Description("Модель статистики")
	Attribute("trainings_count", Int, "Всего тренировок", func() {
		Minimum(0)
	})
	Attribute("sets_count", Int, "Всего подходов", func() {
		Minimum(0)
	})
	Attribute("reps_count", Int, "Всего повторений", func() {
		Minimum(0)
	})
	Attribute("average_duration", Int, "Средяя продолжительность тренировок", func() {
		Minimum(0)
	})
	Required("trainings_count", "sets_count", "reps_count", "average_duration")
})
