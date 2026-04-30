export const useApiFetch = <T>(path: string) => {
  const config = useRuntimeConfig()

  return $fetch<T>(path, {
    baseURL: config.public.apiBase,
  })
}
