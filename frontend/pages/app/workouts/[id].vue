<script setup lang="ts">
definePageMeta({
  layout: 'app',
  middleware: 'auth',
})

const route = useRoute()
const { currentUserWorkouts, exerciseMap } = useTrainingTrackerDemo()

const workout = computed(() =>
  currentUserWorkouts.value.find((item) => item.id === route.params.id) ?? null,
)

if (!workout.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Тренировка не найдена',
  })
}

const formatDate = (value: string) =>
  new Date(value).toLocaleDateString('ru-RU', {
    month: 'long',
    day: 'numeric',
    year: 'numeric',
  })

const formatVolume = (value: number) => `${Math.round(value).toLocaleString('ru-RU')} кг`

const workoutVolume = computed(() =>
  workout.value!.exercises.reduce(
    (sum, exercise) =>
      sum + exercise.sets.reduce((setsSum, set) => setsSum + set.reps * set.weightKg, 0),
    0,
  ),
)

const workoutSetsCount = computed(() =>
  workout.value!.exercises.reduce((sum, exercise) => sum + exercise.sets.length, 0),
)
</script>

<template>
  <div class="page-body">
    <AppPageHeader
      eyebrow="Тренировка"
      :title="workout?.title ?? 'Тренировка'"
      :meta="workout ? formatDate(workout.performedAt) : ''"
    />

    <section v-if="workout" class="stats-band">
      <article class="metric-panel">
        <span>Упражнения</span>
        <strong>{{ workout.exercises.length }}</strong>
      </article>
      <article class="metric-panel">
        <span>Подходы</span>
        <strong>{{ workoutSetsCount }}</strong>
      </article>
      <article class="metric-panel">
        <span>Общий тоннаж</span>
        <strong>{{ formatVolume(workoutVolume) }}</strong>
      </article>
    </section>

    <section v-if="workout" class="surface">
      <div class="surface-head">
        <div>
          <h3>Состав тренировки</h3>
          <p>{{ workout.note || 'Без заметки к этой тренировке.' }}</p>
        </div>
      </div>

      <div class="history-card__exercises">
        <div v-for="exercise in workout.exercises" :key="exercise.id" class="history-exercise">
          <div class="history-exercise__head">
            <strong>{{ exerciseMap.get(exercise.exerciseId)?.name }}</strong>
            <span>{{ exercise.sets.length }} подходов</span>
          </div>
          <ul class="set-summary-list">
            <li v-for="(set, index) in exercise.sets" :key="set.id">
              <span>Подход {{ index + 1 }}</span>
              <span>{{ set.reps }} повторов</span>
              <span>{{ set.weightKg }} кг</span>
            </li>
          </ul>
        </div>
      </div>
    </section>
  </div>
</template>
