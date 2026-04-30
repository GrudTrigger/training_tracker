import { computed } from 'vue'

export type Role = 'athlete' | 'admin'

export type UserProfile = {
  id: string
  name: string
  email: string
  role: Role
  joinedAt: string
}

export type Exercise = {
  id: string
  name: string
  muscleGroup: string
}

export type WorkoutSet = {
  id: string
  reps: number
  weightKg: number
}

export type WorkoutExercise = {
  id: string
  exerciseId: string
  sets: WorkoutSet[]
}

export type Workout = {
  id: string
  userId: string
  title: string
  note: string
  performedAt: string
  exercises: WorkoutExercise[]
}

type DraftSet = {
  id: string
  reps: number
  weightKg: number
}

type DraftExercise = {
  id: string
  exerciseId: string
  sets: DraftSet[]
}

type DraftWorkout = {
  title: string
  note: string
  performedAt: string
  exercises: DraftExercise[]
}

const todayISODate = () => new Date().toISOString().slice(0, 10)

const makeId = () => Math.random().toString(36).slice(2, 10)

const defaultDraftExercise = (exerciseId = 'pullup') => ({
  id: makeId(),
  exerciseId,
  sets: [
    {
      id: makeId(),
      reps: 8,
      weightKg: 0,
    },
  ],
})

const defaultDraftWorkout = (): DraftWorkout => ({
  title: 'Верх тела',
  note: '',
  performedAt: todayISODate(),
  exercises: [defaultDraftExercise()],
})

export const useTrainingTrackerDemo = () => {
  const users = useState<UserProfile[]>('demo-users', () => [
    {
      id: 'user-athlete',
      name: 'George',
      email: 'george@example.com',
      role: 'athlete',
      joinedAt: '2026-03-14',
    },
    {
      id: 'user-admin',
      name: 'Алекс',
      email: 'admin@tracker.app',
      role: 'admin',
      joinedAt: '2026-02-02',
    },
    {
      id: 'user-athlete-2',
      name: 'Майя',
      email: 'maya@example.com',
      role: 'athlete',
      joinedAt: '2026-04-10',
    },
  ])

  const exercises = useState<Exercise[]>('demo-exercises', () => [
    { id: 'pullup', name: 'Подтягивания', muscleGroup: 'Спина' },
    { id: 'bench-press', name: 'Жим лежа', muscleGroup: 'Грудь' },
    { id: 'squat', name: 'Приседания со штангой', muscleGroup: 'Ноги' },
    { id: 'row', name: 'Тяга штанги в наклоне', muscleGroup: 'Спина' },
    { id: 'ohp', name: 'Жим стоя', muscleGroup: 'Плечи' },
    { id: 'rdl', name: 'Румынская тяга', muscleGroup: 'Задняя цепь' },
  ])

  const workouts = useState<Workout[]>('demo-workouts', () => [
    {
      id: 'w-1',
      userId: 'user-athlete',
      title: 'День тяги',
      note: 'Короткий отдых между подходами и акцент на чистую технику.',
      performedAt: '2026-04-29',
      exercises: [
        {
          id: 'we-1',
          exerciseId: 'pullup',
          sets: [
            { id: 's-1', reps: 10, weightKg: 0 },
            { id: 's-2', reps: 8, weightKg: 5 },
            { id: 's-3', reps: 6, weightKg: 10 },
          ],
        },
        {
          id: 'we-2',
          exerciseId: 'row',
          sets: [
            { id: 's-4', reps: 10, weightKg: 60 },
            { id: 's-5', reps: 10, weightKg: 60 },
            { id: 's-6', reps: 8, weightKg: 70 },
          ],
        },
      ],
    },
    {
      id: 'w-2',
      userId: 'user-athlete',
      title: 'День ног',
      note: 'Медленная негативная фаза в приседаниях.',
      performedAt: '2026-04-26',
      exercises: [
        {
          id: 'we-3',
          exerciseId: 'squat',
          sets: [
            { id: 's-7', reps: 8, weightKg: 80 },
            { id: 's-8', reps: 8, weightKg: 85 },
            { id: 's-9', reps: 6, weightKg: 90 },
          ],
        },
        {
          id: 'we-4',
          exerciseId: 'rdl',
          sets: [
            { id: 's-10', reps: 10, weightKg: 70 },
            { id: 's-11', reps: 10, weightKg: 75 },
          ],
        },
      ],
    },
    {
      id: 'w-3',
      userId: 'user-athlete-2',
      title: 'День жима',
      note: 'Объемный блок.',
      performedAt: '2026-04-28',
      exercises: [
        {
          id: 'we-5',
          exerciseId: 'bench-press',
          sets: [
            { id: 's-12', reps: 10, weightKg: 50 },
            { id: 's-13', reps: 8, weightKg: 55 },
          ],
        },
        {
          id: 'we-6',
          exerciseId: 'ohp',
          sets: [
            { id: 's-14', reps: 8, weightKg: 30 },
            { id: 's-15', reps: 8, weightKg: 32.5 },
          ],
        },
      ],
    },
  ])

  const currentUserId = useState<string | null>('demo-current-user-id', () => null)
  const currentView = useState<'overview' | 'create' | 'history' | 'admin'>(
    'demo-current-view',
    () => 'overview',
  )
  const selectedExerciseId = useState<string>('demo-selected-exercise-id', () => 'pullup')
  const draftWorkout = useState<DraftWorkout>('demo-draft-workout', defaultDraftWorkout)

  const currentUser = computed(() =>
    users.value.find((user) => user.id === currentUserId.value) ?? null,
  )

  const athleteUsers = computed(() => users.value.filter((user) => user.role === 'athlete'))
  const adminUsers = computed(() => users.value.filter((user) => user.role === 'admin'))

  const currentUserWorkouts = computed(() => {
    if (!currentUser.value) {
      return []
    }

    if (currentUser.value.role === 'admin') {
      return workouts.value
    }

    return workouts.value.filter((workout) => workout.userId === currentUser.value?.id)
  })

  const exerciseMap = computed(() =>
    new Map(exercises.value.map((exercise) => [exercise.id, exercise])),
  )

  const allCurrentSets = computed(() =>
    currentUserWorkouts.value.flatMap((workout) =>
      workout.exercises.flatMap((exercise) =>
        exercise.sets.map((set) => ({
          workout,
          exercise,
          set,
        })),
      ),
    ),
  )

  const overallStats = computed(() => {
    const totalSets = allCurrentSets.value.length
    const totalVolumeKg = allCurrentSets.value.reduce(
      (sum, item) => sum + item.set.reps * item.set.weightKg,
      0,
    )
    const workoutsCount = currentUserWorkouts.value.length
    const lastWorkoutAt = currentUserWorkouts.value[0]?.performedAt ?? null

    return {
      workoutsCount,
      totalSets,
      totalVolumeKg,
      lastWorkoutAt,
    }
  })

  const selectedExerciseStats = computed(() => {
    const matching = currentUserWorkouts.value.flatMap((workout) =>
      workout.exercises
        .filter((exercise) => exercise.exerciseId === selectedExerciseId.value)
        .flatMap((exercise) =>
          exercise.sets.map((set) => ({
            workoutId: workout.id,
            workoutTitle: workout.title,
            performedAt: workout.performedAt,
            reps: set.reps,
            weightKg: set.weightKg,
            volumeKg: set.reps * set.weightKg,
          })),
        ),
    )

    const totalVolumeKg = matching.reduce((sum, set) => sum + set.volumeKg, 0)
    const bestWeightKg = matching.reduce((max, set) => Math.max(max, set.weightKg), 0)
    const bestVolumeKg = matching.reduce((max, set) => Math.max(max, set.volumeKg), 0)

    return {
      entries: matching,
      totalVolumeKg,
      bestWeightKg,
      bestVolumeKg,
      workoutsCount: new Set(matching.map((item) => item.workoutId)).size,
    }
  })

  const adminStats = computed(() => {
    const allSets = workouts.value.flatMap((workout) =>
      workout.exercises.flatMap((exercise) => exercise.sets),
    )

    const popularExercises = exercises.value
      .map((exercise) => {
        const usages = workouts.value.reduce((count, workout) => {
          return count + workout.exercises.filter((item) => item.exerciseId === exercise.id).length
        }, 0)

        return {
          ...exercise,
          usages,
        }
      })
      .sort((a, b) => b.usages - a.usages)

    return {
      usersCount: users.value.length,
      athleteCount: athleteUsers.value.length,
      adminCount: adminUsers.value.length,
      workoutsCount: workouts.value.length,
      totalSets: allSets.length,
      popularExercises,
    }
  })

  const login = (email: string) => {
    const found = users.value.find((user) => user.email.toLowerCase() === email.trim().toLowerCase())
    if (!found) {
      return { ok: false as const, message: 'Пользователь не найден. Попробуй демо-вход или зарегистрируй новый аккаунт.' }
    }

    currentUserId.value = found.id
    currentView.value = found.role === 'admin' ? 'admin' : 'overview'

    return { ok: true as const }
  }

  const loginAsDemo = (role: Role) => {
    const found = users.value.find((user) => user.role === role)
    if (!found) {
      return
    }

    currentUserId.value = found.id
    currentView.value = role === 'admin' ? 'admin' : 'overview'
  }

  const register = (name: string, email: string, password: string) => {
    if (!name.trim() || !email.trim() || password.trim().length < 6) {
      return { ok: false as const, message: 'Заполни имя, email и пароль не короче 6 символов.' }
    }

    const exists = users.value.some((user) => user.email.toLowerCase() === email.trim().toLowerCase())
    if (exists) {
      return { ok: false as const, message: 'Аккаунт с таким email уже существует.' }
    }

    const created: UserProfile = {
      id: makeId(),
      name: name.trim(),
      email: email.trim(),
      role: 'athlete',
      joinedAt: todayISODate(),
    }

    users.value = [created, ...users.value]
    currentUserId.value = created.id
    currentView.value = 'overview'

    return { ok: true as const }
  }

  const logout = () => {
    currentUserId.value = null
    currentView.value = 'overview'
  }

  const addDraftExercise = () => {
    const firstExerciseId = exercises.value[0]?.id ?? ''
    draftWorkout.value = {
      ...draftWorkout.value,
      exercises: [...draftWorkout.value.exercises, defaultDraftExercise(firstExerciseId)],
    }
  }

  const removeDraftExercise = (draftExerciseId: string) => {
    if (draftWorkout.value.exercises.length === 1) {
      return
    }

    draftWorkout.value = {
      ...draftWorkout.value,
      exercises: draftWorkout.value.exercises.filter((exercise) => exercise.id !== draftExerciseId),
    }
  }

  const updateDraftExercise = (
    draftExerciseId: string,
    updater: (exercise: DraftExercise) => DraftExercise,
  ) => {
    draftWorkout.value = {
      ...draftWorkout.value,
      exercises: draftWorkout.value.exercises.map((exercise) =>
        exercise.id === draftExerciseId ? updater(exercise) : exercise,
      ),
    }
  }

  const addSet = (draftExerciseId: string) => {
    updateDraftExercise(draftExerciseId, (exercise) => ({
      ...exercise,
      sets: [
        ...exercise.sets,
        {
          id: makeId(),
          reps: 8,
          weightKg: 0,
        },
      ],
    }))
  }

  const removeSet = (draftExerciseId: string, setId: string) => {
    updateDraftExercise(draftExerciseId, (exercise) => {
      if (exercise.sets.length === 1) {
        return exercise
      }

      return {
        ...exercise,
        sets: exercise.sets.filter((set) => set.id !== setId),
      }
    })
  }

  const submitWorkout = () => {
    if (!currentUser.value) {
      return { ok: false as const, message: 'Перед добавлением тренировки нужно войти в аккаунт.' }
    }

    if (!draftWorkout.value.title.trim()) {
      return { ok: false as const, message: 'Укажи название тренировки.' }
    }

    const invalidExercise = draftWorkout.value.exercises.some(
      (exercise) =>
        !exercise.exerciseId ||
        exercise.sets.some((set) => set.reps <= 0 || set.weightKg < 0),
    )

    if (invalidExercise) {
      return { ok: false as const, message: 'Для каждого упражнения нужны корректные подходы с повторениями и весом.' }
    }

    const createdWorkout: Workout = {
      id: makeId(),
      userId: currentUser.value.id,
      title: draftWorkout.value.title.trim(),
      note: draftWorkout.value.note.trim(),
      performedAt: draftWorkout.value.performedAt,
      exercises: draftWorkout.value.exercises.map((exercise) => ({
        id: makeId(),
        exerciseId: exercise.exerciseId,
        sets: exercise.sets.map((set) => ({
          id: makeId(),
          reps: set.reps,
          weightKg: set.weightKg,
        })),
      })),
    }

    workouts.value = [createdWorkout, ...workouts.value]
    draftWorkout.value = defaultDraftWorkout()
    currentView.value = 'history'

    return { ok: true as const }
  }

  return {
    users,
    exercises,
    workouts,
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
  }
}
