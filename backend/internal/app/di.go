package app

import (
	"context"
	"fmt"

	"github.com/GrudTrigger/training_tracker/backend/gen/exercises"
	genexercise "github.com/GrudTrigger/training_tracker/backend/gen/exercises"
	genstatistics "github.com/GrudTrigger/training_tracker/backend/gen/statistics"
	gentrainings "github.com/GrudTrigger/training_tracker/backend/gen/trainings"
	"github.com/GrudTrigger/training_tracker/backend/internal/config"
	"github.com/GrudTrigger/training_tracker/backend/internal/migrator"
	"github.com/GrudTrigger/training_tracker/backend/internal/repository"
	repoExercise "github.com/GrudTrigger/training_tracker/backend/internal/repository/exercise"
	repoTrainings "github.com/GrudTrigger/training_tracker/backend/internal/repository/trainings"
	"github.com/GrudTrigger/training_tracker/backend/internal/service/exercise"
	"github.com/GrudTrigger/training_tracker/backend/internal/service/statistics"
	"github.com/GrudTrigger/training_tracker/backend/internal/service/trainings"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

type DiContainer struct {
	postgres *pgx.Conn

	exerciseRepo      repository.ExerciseRepo
	exerciseSvc       *exercise.Service
	exerciseEndpoints *exercises.Endpoints

	trainingsRepo      repository.TrainingsRepo
	trainingsSvc       *trainings.Service
	trainingsEndpoints *gentrainings.Endpoints

	statisticsSvc       *statistics.Service
	statisticsEndpoints *genstatistics.Endpoints
}

func NewDiContainer() *DiContainer {
	return &DiContainer{}
}

func (d *DiContainer) StatisticsEndpoints(ctx context.Context) *genstatistics.Endpoints {
	if d.statisticsEndpoints == nil {
		d.statisticsEndpoints = genstatistics.NewEndpoints(d.StatisticsService(ctx))
	}
	return d.statisticsEndpoints
}

func (d *DiContainer) StatisticsService(ctx context.Context) *statistics.Service {
	if d.statisticsSvc == nil && d.trainingsRepo == nil {
		d.trainingsRepo = repoTrainings.NewRepository(d.Postgres(ctx))
		d.statisticsSvc = statistics.NewService(d.trainingsRepo)
	}
	if d.statisticsSvc == nil && d.trainingsRepo != nil {
		d.statisticsSvc = statistics.NewService(d.trainingsRepo)
	}
	return d.statisticsSvc
}

func (d *DiContainer) TrainingsEndpoints(ctx context.Context) *gentrainings.Endpoints {
	if d.trainingsEndpoints == nil {
		d.trainingsEndpoints = gentrainings.NewEndpoints(d.TrainingsService(ctx))
	}
	return d.trainingsEndpoints
}

// TODO: тут может быть баг, проверить!
func (d *DiContainer) TrainingsService(ctx context.Context) *trainings.Service {
	if d.trainingsSvc == nil && d.exerciseRepo != nil {
		d.trainingsSvc = trainings.NewService(d.TrainingsRepository(ctx), d.exerciseRepo)
	}
	if d.trainingsSvc == nil && d.exerciseRepo == nil {
		d.exerciseRepo = repoExercise.NewExerciseRepository(d.Postgres(ctx))
		d.trainingsSvc = trainings.NewService(d.TrainingsRepository(ctx), d.exerciseRepo)
	}

	return d.trainingsSvc
}

func (d *DiContainer) TrainingsRepository(ctx context.Context) repository.TrainingsRepo {
	if d.trainingsRepo == nil {
		d.trainingsRepo = repoTrainings.NewRepository(d.Postgres(ctx))
	}
	return d.trainingsRepo
}

func (d *DiContainer) ExerciseEndpoints(ctx context.Context) *exercises.Endpoints {
	if d.exerciseEndpoints == nil {
		d.exerciseEndpoints = genexercise.NewEndpoints(d.ExerciseService(ctx))
	}
	return d.exerciseEndpoints
}

func (d *DiContainer) ExerciseService(ctx context.Context) *exercise.Service {
	if d.exerciseSvc == nil {
		d.exerciseSvc = exercise.NewExerciseService(d.ExerciseRepository(ctx))
	}
	return d.exerciseSvc
}

func (d *DiContainer) ExerciseRepository(ctx context.Context) repository.ExerciseRepo {
	if d.exerciseRepo == nil {
		d.exerciseRepo = repoExercise.NewExerciseRepository(d.Postgres(ctx))
	}
	return d.exerciseRepo
}

func (d *DiContainer) Postgres(ctx context.Context) *pgx.Conn {
	if d.postgres == nil {
		conn, err := pgx.Connect(ctx, config.AppConfig().Postgres.URI())
		if err != nil {
			panic(fmt.Errorf("failed to connect to database: %w", err))
		}

		err = conn.Ping(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to connect to database in ping: %w", err))
		}

		m := migrator.NewMigrator(stdlib.OpenDB(*conn.Config().Copy()), "../migrations")
		err = m.Up()
		if err != nil {
			panic(fmt.Errorf("failed migrations: %w", err))
		}
		d.postgres = conn
	}
	return d.postgres
}
