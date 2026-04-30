<script setup lang="ts">
import {
  Activity,
  BarChart3,
  ClipboardList,
  Crown,
  Dumbbell,
  LogOut,
  Plus,
  Shield,
  UserPlus,
} from 'lucide-vue-next'

const {
  users,
  exercises,
  currentUser,
  currentView,
  selectedExerciseId,
  draftWorkout,
  currentUserWorkouts,
  exerciseMap,
  overallStats,
  selectedExerciseStats,
  adminStats,
  login,
  loginAsDemo,
  register,
  logout,
  addDraftExercise,
  removeDraftExercise,
  updateDraftExercise,
  addSet,
  removeSet,
  submitWorkout,
} = useTrainingTrackerDemo()

type AuthMode = 'login' | 'register'

const authMode = ref<AuthMode>('login')
const authMessage = ref('')

const loginForm = reactive({
  email: 'george@example.com',
  password: 'password123',
})

const registerForm = reactive({
  name: '',
  email: '',
  password: '',
})

const navItems = computed(() => {
  const base = [
    { key: 'overview', label: 'Обзор', icon: Activity },
    { key: 'create', label: 'Новая', icon: Plus },
    { key: 'history', label: 'История', icon: ClipboardList },
  ] as const

  if (currentUser.value?.role === 'admin') {
    return [...base, { key: 'admin', label: 'Админка', icon: Shield }] as const
  }

  return base
})

const selectedExerciseName = computed(() => {
  return exerciseMap.value.get(selectedExerciseId.value)?.name ?? 'Упражнение'
})

const roleLabel = (role: 'athlete' | 'admin') => (role === 'admin' ? 'Админ' : 'Спортсмен')

const recentUsers = computed(() =>
  [...users.value]
    .sort((a, b) => b.joinedAt.localeCompare(a.joinedAt))
    .slice(0, 5),
)

const handleLogin = () => {
  const result = login(loginForm.email)
  authMessage.value = result.ok ? '' : result.message
}

const handleRegister = () => {
  const result = register(registerForm.name, registerForm.email, registerForm.password)
  authMessage.value = result.ok ? '' : result.message

  if (result.ok) {
    registerForm.name = ''
    registerForm.email = ''
    registerForm.password = ''
  }
}

const handleSubmitWorkout = () => {
  const result = submitWorkout()
  authMessage.value = result.ok ? '' : result.message
}

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
      sum +
      exercise.sets.reduce((setsSum, set) => setsSum + set.reps * set.weightKg, 0),
    0,
  )
</script>

<template>
  <main class="app-shell">
    <section v-if="!currentUser" class="auth-layout">
      <div class="auth-copy">
        <p class="section-kicker">Трекер тренировок</p>
        <h1>Отслеживай тренировки в зале в интерфейсе, удобном и на телефоне, и на десктопе.</h1>
        <p class="section-lead">
          Входи в аккаунт, записывай тренировки, смотри статистику по упражнениям
          и переключайся в админку, когда нужны метрики по приложению.
        </p>

        <div class="feature-strip">
          <div class="feature-chip">
            <Dumbbell :size="16" />
            <span>Тренировки с упражнениями и подходами</span>
          </div>
          <div class="feature-chip">
            <BarChart3 :size="16" />
            <span>Рекорды и тоннаж по упражнениям</span>
          </div>
          <div class="feature-chip">
            <Shield :size="16" />
            <span>Админка с аналитикой</span>
          </div>
        </div>
      </div>

      <div class="auth-panel">
        <div class="auth-toggle">
          <button
            class="segmented-button"
            :class="{ 'segmented-button--active': authMode === 'login' }"
            type="button"
            @click="authMode = 'login'"
          >
            <Activity :size="16" />
            <span>Вход</span>
          </button>
          <button
            class="segmented-button"
            :class="{ 'segmented-button--active': authMode === 'register' }"
            type="button"
            @click="authMode = 'register'"
          >
            <UserPlus :size="16" />
            <span>Регистрация</span>
          </button>
        </div>

        <form v-if="authMode === 'login'" class="form-stack" @submit.prevent="handleLogin">
          <label class="field">
            <span>Email</span>
            <input v-model="loginForm.email" type="email" placeholder="george@example.com" />
          </label>
          <label class="field">
            <span>Пароль</span>
            <input v-model="loginForm.password" type="password" placeholder="••••••••" />
          </label>
          <button class="primary-button" type="submit">
            <Activity :size="16" />
            <span>Войти</span>
          </button>
        </form>

        <form v-else class="form-stack" @submit.prevent="handleRegister">
          <label class="field">
            <span>Имя</span>
            <input v-model="registerForm.name" type="text" placeholder="Георгий" />
          </label>
          <label class="field">
            <span>Email</span>
            <input v-model="registerForm.email" type="email" placeholder="george@example.com" />
          </label>
          <label class="field">
            <span>Пароль</span>
            <input v-model="registerForm.password" type="password" placeholder="Минимум 6 символов" />
          </label>
          <button class="primary-button" type="submit">
            <UserPlus :size="16" />
            <span>Создать аккаунт</span>
          </button>
        </form>

        <p v-if="authMessage" class="status-note">{{ authMessage }}</p>

        <div class="demo-actions">
          <button class="secondary-button" type="button" @click="loginAsDemo('athlete')">
            <Dumbbell :size="16" />
            <span>Демо спортсмена</span>
          </button>
          <button class="secondary-button" type="button" @click="loginAsDemo('admin')">
            <Crown :size="16" />
            <span>Демо админа</span>
          </button>
        </div>
      </div>
    </section>

    <section v-else class="workspace">
      <aside class="sidebar">
        <div class="sidebar__top">
          <div class="brand-lockup">
            <div class="brand-badge">
              <Dumbbell :size="18" />
            </div>
            <div>
              <strong>Трекер тренировок</strong>
              <p>{{ currentUser.role === 'admin' ? 'Панель администратора' : 'Кабинет спортсмена' }}</p>
            </div>
          </div>

          <div class="profile-card">
            <div>
              <strong>{{ currentUser.name }}</strong>
              <p>{{ currentUser.email }}</p>
            </div>
            <span class="role-pill" :class="`role-pill--${currentUser.role}`">
              {{ roleLabel(currentUser.role) }}
            </span>
          </div>

          <nav class="nav-list" aria-label="Primary">
            <button
              v-for="item in navItems"
              :key="item.key"
              class="nav-button"
              :class="{ 'nav-button--active': currentView === item.key }"
              type="button"
              @click="currentView = item.key"
            >
              <component :is="item.icon" :size="18" />
              <span>{{ item.label }}</span>
            </button>
          </nav>
        </div>

        <button class="ghost-button" type="button" @click="logout">
          <LogOut :size="16" />
          <span>Выйти</span>
        </button>
      </aside>

      <div class="mobile-nav">
        <button
          v-for="item in navItems"
          :key="`mobile-${item.key}`"
          class="mobile-nav__button"
          :class="{ 'mobile-nav__button--active': currentView === item.key }"
          type="button"
          @click="currentView = item.key"
        >
          <component :is="item.icon" :size="18" />
          <span>{{ item.label }}</span>
        </button>
      </div>

      <section class="content">
        <header class="page-header">
          <div>
            <p class="section-kicker">Сегодня</p>
            <h2>
              {{
                currentView === 'overview'
                  ? 'Обзор'
                  : currentView === 'create'
                    ? 'Новая тренировка'
                    : currentView === 'history'
                      ? 'История тренировок'
                      : 'Админка'
              }}
            </h2>
          </div>
          <div class="header-meta">
            <span>{{ formatDate(new Date().toISOString()) }}</span>
          </div>
        </header>

        <div v-if="currentView === 'overview'" class="page-body">
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
                  v-for="entry in selectedExerciseStats.entries.slice(0, 6)"
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

            <article class="surface">
              <div class="surface-head">
                <div>
                  <h3>Последние тренировки</h3>
                  <p>Недавние сессии с упражнениями и общим тоннажем.</p>
                </div>
              </div>

              <ul class="history-list">
                <li
                  v-for="workout in currentUserWorkouts.slice(0, 5)"
                  :key="workout.id"
                  class="history-item"
                >
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
          </section>
        </div>

        <div v-else-if="currentView === 'create'" class="page-body">
          <section class="surface">
            <div class="surface-head">
              <div>
                <h3>Новая тренировка</h3>
                <p>Собери сессию из упражнений, подходов, повторений и веса.</p>
              </div>
            </div>

            <div class="form-grid">
              <label class="field">
                <span>Название тренировки</span>
                <input v-model="draftWorkout.title" type="text" placeholder="Верх тела" />
              </label>
              <label class="field">
                <span>Дата</span>
                <input v-model="draftWorkout.performedAt" type="date" />
              </label>
              <label class="field field--full">
                <span>Заметка</span>
                <textarea
                  v-model="draftWorkout.note"
                  rows="3"
                  placeholder="Что стоит запомнить по этой тренировке?"
                />
              </label>
            </div>

            <div class="exercise-stack">
              <section
                v-for="exerciseDraft in draftWorkout.exercises"
                :key="exerciseDraft.id"
                class="exercise-editor"
              >
                <div class="exercise-editor__header">
                  <label class="field">
                    <span>Упражнение</span>
                    <select
                      class="select-input"
                      :value="exerciseDraft.exerciseId"
                      @change="
                        updateDraftExercise(exerciseDraft.id, (exercise) => ({
                          ...exercise,
                          exerciseId: ($event.target as HTMLSelectElement).value,
                        }))
                      "
                    "
                    >
                      <option v-for="exercise in exercises" :key="exercise.id" :value="exercise.id">
                        {{ exercise.name }}
                      </option>
                    </select>
                  </label>

                  <button class="icon-button" type="button" @click="removeDraftExercise(exerciseDraft.id)">
                    <span>×</span>
                  </button>
                </div>

                <div class="set-table">
                  <div class="set-table__head">
                    <span>Подход</span>
                    <span>Повторы</span>
                    <span>Вес</span>
                    <span></span>
                  </div>

                  <div
                    v-for="(set, setIndex) in exerciseDraft.sets"
                    :key="set.id"
                    class="set-table__row"
                  >
                    <span class="set-index">{{ setIndex + 1 }}</span>
                    <input
                      class="compact-input"
                      :value="set.reps"
                      type="number"
                      min="1"
                      @input="
                        updateDraftExercise(exerciseDraft.id, (exercise) => ({
                          ...exercise,
                          sets: exercise.sets.map((item) =>
                            item.id === set.id
                              ? { ...item, reps: Number(($event.target as HTMLInputElement).value) }
                              : item,
                          ),
                        }))
                      "
                    />
                    <input
                      class="compact-input"
                      :value="set.weightKg"
                      type="number"
                      min="0"
                      step="0.5"
                      @input="
                        updateDraftExercise(exerciseDraft.id, (exercise) => ({
                          ...exercise,
                          sets: exercise.sets.map((item) =>
                            item.id === set.id
                              ? { ...item, weightKg: Number(($event.target as HTMLInputElement).value) }
                              : item,
                          ),
                        }))
                      "
                    />
                    <button class="icon-button" type="button" @click="removeSet(exerciseDraft.id, set.id)">
                      <span>×</span>
                    </button>
                  </div>
                </div>

                <button class="secondary-button secondary-button--small" type="button" @click="addSet(exerciseDraft.id)">
                  <Plus :size="16" />
                  <span>Add set</span>
                </button>
              </section>
            </div>

            <div class="exercise-actions">
              <button class="secondary-button secondary-button--small" type="button" @click="addDraftExercise">
                <Plus :size="16" />
                <span>Добавить упражнение</span>
              </button>
            </div>

            <div class="action-row">
              <button class="primary-button" type="button" @click="handleSubmitWorkout">
                <Dumbbell :size="16" />
                <span>Сохранить тренировку</span>
              </button>
            </div>
          </section>
        </div>

        <div v-else-if="currentView === 'history'" class="page-body">
          <section class="surface">
            <div class="surface-head">
              <div>
                <h3>История тренировок</h3>
                <p>Все прошлые тренировки с упражнениями, подходами и общим тоннажем.</p>
              </div>
            </div>

            <div class="history-stack">
              <article
                v-for="workout in currentUserWorkouts"
                :key="workout.id"
                class="history-card"
              >
                <div class="history-card__head">
                  <div>
                    <h4>{{ workout.title }}</h4>
                    <p>{{ formatDate(workout.performedAt) }}</p>
                  </div>
                  <span class="history-card__volume">{{ formatVolume(getWorkoutVolume(workout)) }}</span>
                </div>

                <p class="history-card__note">{{ workout.note || 'Без заметки к этой сессии.' }}</p>

                <div class="history-card__exercises">
                  <div
                    v-for="exercise in workout.exercises"
                    :key="exercise.id"
                    class="history-exercise"
                  >
                    <div class="history-exercise__head">
                      <strong>{{ exerciseMap.get(exercise.exerciseId)?.name }}</strong>
                      <span>{{ exercise.sets.length }} sets</span>
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
              </article>
            </div>
          </section>
        </div>

        <div v-else class="page-body">
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

        <footer class="footer-strip">
          <span>Выбрано упражнение: {{ selectedExerciseName }}</span>
          <span>Доступно тренировок: {{ currentUserWorkouts.length }}</span>
        </footer>
      </section>
    </section>
  </main>
</template>
