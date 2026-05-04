<script setup lang="ts">
definePageMeta({
  layout: 'app',
  middleware: 'auth',
})

const { currentUserWorkouts, overallStats } = useTrainingTrackerDemo()

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
</script>

<template>
  <div class="page-body">
    <AppPageHeader eyebrow="Сегодня" title="Обзор" :meta="formatDate(new Date().toISOString())" />

    <section class="stats-band">
      <article class="metric-panel">
        <span>Тренировки</span>
        <strong>{{ overallStats.workoutsCount }}</strong>
      </article>
      <article class="metric-panel">
        <span>Подходы</span>
        <strong>{{ overallStats.totalSets }}</strong>
      </article>
      <article class="metric-panel">
        <span>Общий тоннаж</span>
        <strong>{{ formatVolume(overallStats.totalVolumeKg) }}</strong>
      </article>
      <article class="metric-panel">
        <span>Последняя тренировка</span>
        <strong>{{ overallStats.lastWorkoutAt ? formatDate(overallStats.lastWorkoutAt) : '—' }}</strong>
      </article>
    </section>

    <section class="dashboard-grid">
      <article class="surface">
        <div class="surface-head">
          <div>
            <h3>Последние тренировки</h3>
            <p>Недавние сессии с упражнениями и общим тоннажем.</p>
          </div>
        </div>

        <ul class="history-list">
          <li v-for="workout in currentUserWorkouts.slice(0, 5)" :key="workout.id" class="history-item">
            <div class="history-item__title">
              <strong>{{ workout.title }}</strong>
              <span>{{ formatDate(workout.performedAt) }}</span>
            </div>
            <p>{{ workout.note || 'Без заметки к этой тренировке.' }}</p>
            <div class="history-item__footer">
              <span>{{ workout.exercises.length }} упражнений</span>
              <span>{{ formatVolume(getWorkoutVolume(workout)) }}</span>
            </div>
          </li>
        </ul>
      </article>

      <article class="surface">
        <div class="surface-head">
          <div>
            <h3>Быстрые действия</h3>
            <p>Переходи к основным сценариям без лишних кликов.</p>
          </div>
        </div>

        <div class="feature-strip">
          <NuxtLink class="feature-chip" to="/app/workouts/new">
            <span>Добавить новую тренировку</span>
          </NuxtLink>
          <NuxtLink class="feature-chip" to="/app/stats">
            <span>Открыть статистику по упражнениям</span>
          </NuxtLink>
          <NuxtLink class="feature-chip" to="/app/workouts">
            <span>Посмотреть историю тренировок</span>
          </NuxtLink>
        </div>
      </article>
    </section>
  </div>
</template>
