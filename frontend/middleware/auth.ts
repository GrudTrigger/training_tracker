export default defineNuxtRouteMiddleware(() => {
  const { currentUser } = useTrainingTrackerDemo()

  if (!currentUser.value) {
    return navigateTo('/login')
  }
})
