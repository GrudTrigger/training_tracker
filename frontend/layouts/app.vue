<script setup lang="ts">
import { Activity, BarChart3, ClipboardList, Dumbbell, LogOut, Plus, Shield } from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const { currentUser, isAdmin, logout } = useTrainingTrackerDemo()

const navItems = computed(() => {
  const items = [
    { to: '/app', label: 'Обзор', icon: Activity },
    { to: '/app/workouts/new', label: 'Новая', icon: Plus },
    { to: '/app/workouts', label: 'История', icon: ClipboardList },
    { to: '/app/stats', label: 'Статистика', icon: BarChart3 },
  ]

  if (isAdmin.value) {
    items.push({ to: '/admin', label: 'Админка', icon: Shield })
  }

  return items
})

const roleLabel = computed(() => (isAdmin.value ? 'Админ' : 'Спортсмен'))

const handleLogout = async () => {
  logout()
  await router.push('/login')
}

const activePath = computed(() => {
  const matches = navItems.value
    .map((item) => item.to)
    .filter((path) => route.path === path || route.path.startsWith(`${path}/`))
    .sort((left, right) => right.length - left.length)

  return matches[0] ?? ''
})

const isActive = (path: string) => activePath.value === path
</script>

<template>
  <div class="workspace">
    <aside class="sidebar">
      <div class="sidebar__top">
        <div class="brand-lockup">
          <div class="brand-badge">
            <Dumbbell :size="18" />
          </div>
          <div>
            <strong>Трекер тренировок</strong>
            <p>{{ isAdmin ? 'Панель администратора' : 'Кабинет спортсмена' }}</p>
          </div>
        </div>

        <div v-if="currentUser" class="profile-card">
          <div>
            <strong>{{ currentUser.name }}</strong>
            <p>{{ currentUser.email }}</p>
          </div>
          <span class="role-pill" :class="`role-pill--${currentUser.role}`">
            {{ roleLabel }}
          </span>
        </div>

        <nav class="nav-list" aria-label="Основная навигация">
          <NuxtLink
            v-for="item in navItems"
            :key="item.to"
            :to="item.to"
            class="nav-button"
            :class="{ 'nav-button--active': isActive(item.to) }"
          >
            <component :is="item.icon" :size="18" />
            <span>{{ item.label }}</span>
          </NuxtLink>
        </nav>
      </div>

      <button class="ghost-button" type="button" @click="handleLogout">
        <LogOut :size="16" />
        <span>Выйти</span>
      </button>
    </aside>

    <div class="mobile-nav" :style="{ gridTemplateColumns: `repeat(${navItems.length}, minmax(0, 1fr))` }">
      <NuxtLink
        v-for="item in navItems"
        :key="`mobile-${item.to}`"
        :to="item.to"
        class="mobile-nav__button"
        :class="{ 'mobile-nav__button--active': isActive(item.to) }"
      >
        <component :is="item.icon" :size="18" />
        <span>{{ item.label }}</span>
      </NuxtLink>
    </div>

    <section class="content">
      <slot />
    </section>
  </div>
</template>
