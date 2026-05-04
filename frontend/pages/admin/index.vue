<script setup lang="ts">
definePageMeta({
  layout: 'app',
  middleware: 'admin',
})

const { adminStats, users } = useTrainingTrackerDemo()

const recentUsers = computed(() =>
  [...users.value]
    .sort((a, b) => b.joinedAt.localeCompare(a.joinedAt))
    .slice(0, 5),
)

const formatDate = (value: string) =>
  new Date(value).toLocaleDateString('ru-RU', {
    month: 'long',
    day: 'numeric',
    year: 'numeric',
  })

const roleLabel = (role: 'athlete' | 'admin') => (role === 'admin' ? 'Админ' : 'Спортсмен')
</script>

<template>
  <div class="page-body">
    <AppPageHeader eyebrow="Управление" title="Админка" :meta="formatDate(new Date().toISOString())" />

    <section class="stats-band">
      <article class="metric-panel">
        <span>Пользователи</span>
        <strong>{{ adminStats.usersCount }}</strong>
      </article>
      <article class="metric-panel">
        <span>Спортсмены</span>
        <strong>{{ adminStats.athleteCount }}</strong>
      </article>
      <article class="metric-panel">
        <span>Тренировки</span>
        <strong>{{ adminStats.workoutsCount }}</strong>
      </article>
      <article class="metric-panel">
        <span>Всего подходов</span>
        <strong>{{ adminStats.totalSets }}</strong>
      </article>
    </section>

    <section class="dashboard-grid">
      <article class="surface">
        <div class="surface-head">
          <div>
            <h3>Новые пользователи</h3>
            <p>Недавние регистрации и текущие роли.</p>
          </div>
        </div>

        <ul class="admin-list">
          <li v-for="user in recentUsers" :key="user.id" class="admin-list__row">
            <div>
              <strong>{{ user.name }}</strong>
              <p>{{ user.email }}</p>
            </div>
            <div class="admin-list__meta">
              <span class="role-pill" :class="`role-pill--${user.role}`">{{ roleLabel(user.role) }}</span>
              <span>{{ formatDate(user.joinedAt) }}</span>
            </div>
          </li>
        </ul>
      </article>

      <article class="surface">
        <div class="surface-head">
          <div>
            <h3>Популярные упражнения</h3>
            <p>Что спортсмены логируют чаще всего.</p>
          </div>
        </div>

        <ul class="admin-list">
          <li
            v-for="exercise in adminStats.popularExercises.slice(0, 6)"
            :key="exercise.id"
            class="admin-list__row"
          >
            <div>
              <strong>{{ exercise.name }}</strong>
              <p>{{ exercise.muscleGroup }}</p>
            </div>
            <div class="admin-list__meta">
              <span>{{ exercise.usages }} тренировок</span>
            </div>
          </li>
        </ul>
      </article>
    </section>
  </div>
</template>
