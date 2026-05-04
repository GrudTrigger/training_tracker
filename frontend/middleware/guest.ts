export default defineNuxtRouteMiddleware(() => {
  const { currentUser, isAdmin } = useTrainingTrackerDemo()

  if (!currentUser.value) {
    return
  }

  return navigateTo(isAdmin.value ? '/admin' : '/app')
})
