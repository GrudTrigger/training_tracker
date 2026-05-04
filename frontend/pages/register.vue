<script setup lang="ts">
import { Activity, UserPlus } from 'lucide-vue-next'

definePageMeta({
  middleware: 'guest',
})

const router = useRouter()
const { register } = useTrainingTrackerDemo()

const authMessage = ref('')
const registerForm = reactive({
  name: '',
  email: '',
  password: '',
})

const handleRegister = async () => {
  const result = register(registerForm.name, registerForm.email, registerForm.password)
  authMessage.value = result.ok ? '' : result.message

  if (!result.ok) {
    return
  }

  await router.push('/app')
}
</script>

<template>
  <main class="app-shell">
    <section class="auth-layout">
      <div class="auth-copy">
        <p class="section-kicker">Трекер тренировок</p>
        <h1>Создай аккаунт и начни вести тренировки без лишнего шума.</h1>
        <p class="section-lead">
          После регистрации сразу попадешь в рабочую часть приложения: журнал,
          статистика, история тренировок и прогресс по упражнениям.
        </p>
      </div>

      <div class="auth-panel">
        <div class="auth-toggle">
          <NuxtLink class="segmented-button" to="/login">
            <Activity :size="16" />
            <span>Вход</span>
          </NuxtLink>
          <NuxtLink class="segmented-button segmented-button--active" to="/register">
            <UserPlus :size="16" />
            <span>Регистрация</span>
          </NuxtLink>
        </div>

        <form class="form-stack" @submit.prevent="handleRegister">
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
      </div>
    </section>
  </main>
</template>
