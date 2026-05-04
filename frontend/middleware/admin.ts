export default defineNuxtRouteMiddleware(() => {
  const { currentUser, isAdmin } = useTrainingTrackerDemo()

  if (!currentUser.value) {
    return navigateTo('/login')
  }

  if (!isAdmin.value) {
    return navigateTo('/app')
  }
})
