<script setup lang="ts">
import { Dumbbell, Plus } from 'lucide-vue-next'

definePageMeta({
  layout: 'app',
  middleware: 'auth',
})

const router = useRouter()
const {
  exercises,
  draftWorkout,
  addDraftExercise,
  removeDraftExercise,
  updateDraftExercise,
  addSet,
  removeSet,
  submitWorkout,
} = useTrainingTrackerDemo()

const formMessage = ref('')

const handleSubmitWorkout = async () => {
  const result = submitWorkout()
  formMessage.value = result.ok ? '' : result.message

  if (!result.ok) {
    return
  }

  await router.push('/app/workouts')
}
</script>

<template>
  <div class="page-body">
    <AppPageHeader eyebrow="Тренировка" title="Новая тренировка" :meta="draftWorkout.performedAt" />

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

            <div v-for="(set, setIndex) in exerciseDraft.sets" :key="set.id" class="set-table__row">
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
            <span>Добавить подход</span>
          </button>
        </section>
      </div>

      <div class="exercise-actions">
        <button class="secondary-button secondary-button--small" type="button" @click="addDraftExercise">
          <Plus :size="16" />
          <span>Добавить упражнение</span>
        </button>
      </div>

      <p v-if="formMessage" class="status-note">{{ formMessage }}</p>

      <div class="action-row">
        <button class="primary-button" type="button" @click="handleSubmitWorkout">
          <Dumbbell :size="16" />
          <span>Сохранить тренировку</span>
        </button>
      </div>
    </section>
  </div>
</template>
