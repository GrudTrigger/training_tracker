<script setup lang="ts">
definePageMeta({
  layout: 'app',
  middleware: 'auth',
})

const { currentUserWorkouts, exerciseMap } = useTrainingTrackerDemo()

const formatDate = (value: string) =>
  new Date(value).toLocaleDateString('ru-RU', {
    month: 'long',
    day: 'numeric',
    year: 'numeric',
  })

const formatVolume = (value: number) => `${Math.round(value).toLocaleString('ru-RU')} кг`

const getWorkoutVolume = (workout: (typeof currentUserWorkouts.value)[number]) =>
  workout.exercises.reduce(
    (sum, exercise) =>
      sum + exercise.sets.reduce((setsSum, set) => setsSum + set.reps * set.weightKg, 0),
    0,
  )

const getWorkoutSetsCount = (workout: (typeof currentUserWorkouts.value)[number]) =>
  workout.exercises.reduce((sum, exercise) => sum + exercise.sets.length, 0)
</script>

<template>
  <div class="page-body">
    <AppPageHeader eyebrow="Журнал" title="История тренировок" :meta="`${currentUserWorkouts.length} записей`" />

    <section class="surface">
      <div class="history-compact-list">
        <NuxtLink
          v-for="workout in currentUserWorkouts"
          :key="workout.id"
          :to="`/app/workouts/${workout.id}`"
          class="workout-row"
        >
          <div class="workout-row__main">
            <div class="workout-row__title">
              <strong>{{ workout.title }}</strong>
              <span>{{ formatDate(workout.performedAt) }}</span>
            </div>
            <p>{{ workout.note || 'Без заметки к этой тренировке.' }}</p>
          </div>

          <div class="workout-row__meta">
            <span>{{ workout.exercises.length }} упражнений</span>
            <span>{{ getWorkoutSetsCount(workout) }} подходов</span>
            <strong>{{ formatVolume(getWorkoutVolume(workout)) }}</strong>
          </div>
        </NuxtLink>
      </div>
    </section>
  </div>
</template>
