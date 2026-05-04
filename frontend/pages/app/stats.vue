<script setup lang="ts">
definePageMeta({
  layout: 'app',
  middleware: 'auth',
})

const { exercises, selectedExerciseId, selectedExerciseStats } = useTrainingTrackerDemo()

const formatDate = (value: string) =>
  new Date(value).toLocaleDateString('ru-RU', {
    month: 'long',
    day: 'numeric',
    year: 'numeric',
  })

const formatVolume = (value: number) => `${Math.round(value).toLocaleString('ru-RU')} кг`
</script>

<template>
  <div class="page-body">
    <AppPageHeader eyebrow="Аналитика" title="Статистика" :meta="formatDate(new Date().toISOString())" />

    <article class="surface">
      <div class="surface-head">
        <div>
          <h3>Статистика по упражнению</h3>
          <p>Выбери упражнение и отслеживай прогресс по времени.</p>
        </div>
        <select v-model="selectedExerciseId" class="select-input">
          <option v-for="exercise in exercises" :key="exercise.id" :value="exercise.id">
            {{ exercise.name }}
          </option>
        </select>
      </div>

      <div class="detail-stats">
        <div class="detail-stat">
          <span>Тренировок с ним</span>
          <strong>{{ selectedExerciseStats.workoutsCount }}</strong>
        </div>
        <div class="detail-stat">
          <span>Лучший вес</span>
          <strong>{{ selectedExerciseStats.bestWeightKg }} кг</strong>
        </div>
        <div class="detail-stat">
          <span>Лучший объем в подходе</span>
          <strong>{{ selectedExerciseStats.bestVolumeKg }} кг</strong>
        </div>
        <div class="detail-stat">
          <span>Суммарный объем</span>
          <strong>{{ formatVolume(selectedExerciseStats.totalVolumeKg) }}</strong>
        </div>
      </div>

      <ul class="timeline-list">
        <li
          v-for="entry in selectedExerciseStats.entries.slice(0, 10)"
          :key="`${entry.workoutId}-${entry.performedAt}-${entry.reps}-${entry.weightKg}`"
          class="timeline-row"
        >
          <div>
            <strong>{{ formatDate(entry.performedAt) }}</strong>
            <p>{{ entry.workoutTitle }}</p>
          </div>
          <div class="timeline-row__metrics">
            <span>{{ entry.weightKg }} кг</span>
            <span>{{ entry.reps }} повторов</span>
            <span>{{ entry.volumeKg }} кг</span>
          </div>
        </li>
      </ul>
    </article>
  </div>
</template>
