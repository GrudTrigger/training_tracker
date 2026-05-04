<script setup lang="ts">
import { Activity, BarChart3, Crown, Dumbbell, Shield } from 'lucide-vue-next'

definePageMeta({
  middleware: 'guest',
})

const router = useRouter()
const { login, loginAsDemo } = useTrainingTrackerDemo()

const authMessage = ref('')
const loginForm = reactive({
  email: 'george@example.com',
  password: 'password123',
})

const handleLogin = async () => {
  const result = login(loginForm.email)
  authMessage.value = result.ok ? '' : result.message

  if (!result.ok) {
    return
  }

  await router.push(result.user.role === 'admin' ? '/admin' : '/app')
}

const handleDemoLogin = async (role: 'athlete' | 'admin') => {
  const user = loginAsDemo(role)
  if (!user) {
    return
  }

  await router.push(user.role === 'admin' ? '/admin' : '/app')
}
</script>

<template>
  <main class="app-shell">
    <section class="auth-layout">
      <div class="auth-copy">
        <p class="section-kicker">Трекер тренировок</p>
        <h1>Тренировки, прогресс и рекорды — в одном месте.</h1>
        <p class="section-lead">
          Входи в аккаунт, записывай тренировки, смотри статистику по упражнениям
          и открывай админку, если работаешь с системой целиком.
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
          <NuxtLink class="segmented-button segmented-button--active" to="/login">
            <Activity :size="16" />
            <span>Вход</span>
          </NuxtLink>
          <NuxtLink class="segmented-button" to="/register">
            <span>Регистрация</span>
          </NuxtLink>
        </div>

        <form class="form-stack" @submit.prevent="handleLogin">
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

        <p v-if="authMessage" class="status-note">{{ authMessage }}</p>

        <div class="demo-actions">
          <button class="secondary-button" type="button" @click="handleDemoLogin('athlete')">
            <Dumbbell :size="16" />
            <span>Демо спортсмена</span>
          </button>
          <button class="secondary-button" type="button" @click="handleDemoLogin('admin')">
            <Crown :size="16" />
            <span>Демо админа</span>
          </button>
        </div>
      </div>
    </section>
  </main>
</template>
