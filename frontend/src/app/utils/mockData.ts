import { Exercise, Workout } from '../types';

export const mockExercises: Exercise[] = [
  { id: '1', name: 'Жим штанги лежа', muscleGroup: 'Грудь', description: 'Базовое упражнение для груди' },
  { id: '2', name: 'Приседания со штангой', muscleGroup: 'Ноги', description: 'Базовое упражнение для ног' },
  { id: '3', name: 'Становая тяга', muscleGroup: 'Спина', description: 'Базовое упражнение для спины' },
  { id: '4', name: 'Жим гантелей на наклонной', muscleGroup: 'Грудь', description: 'Упражнение для верхней части груди' },
  { id: '5', name: 'Подтягивания', muscleGroup: 'Спина', description: 'Упражнение с собственным весом' },
  { id: '6', name: 'Жим штанги стоя', muscleGroup: 'Плечи', description: 'Базовое упражнение для плеч' },
  { id: '7', name: 'Сгибания на бицепс', muscleGroup: 'Руки', description: 'Изолированное упражнение на бицепс' },
  { id: '8', name: 'Французский жим', muscleGroup: 'Руки', description: 'Упражнение на трицепс' },
  { id: '9', name: 'Выпады с гантелями', muscleGroup: 'Ноги', description: 'Упражнение для ног и ягодиц' },
  { id: '10', name: 'Тяга штанги в наклоне', muscleGroup: 'Спина', description: 'Базовое упражнение для спины' },
];

export const mockWorkouts: Workout[] = [
  {
    id: '1',
    name: 'Грудь и трицепс',
    date: '2024-12-20T10:00:00',
    duration: 75,
    exercises: [
      {
        exerciseId: '1',
        sets: [
          { id: 's1', weight: 80, reps: 10 },
          { id: 's2', weight: 80, reps: 8 },
          { id: 's3', weight: 75, reps: 10 },
        ],
      },
      {
        exerciseId: '4',
        sets: [
          { id: 's4', weight: 30, reps: 12 },
          { id: 's5', weight: 30, reps: 10 },
          { id: 's6', weight: 28, reps: 12 },
        ],
      },
    ],
    notes: 'Отличная тренировка!',
  },
  {
    id: '2',
    name: 'Спина и бицепс',
    date: '2024-12-18T09:30:00',
    duration: 80,
    exercises: [
      {
        exerciseId: '3',
        sets: [
          { id: 's7', weight: 100, reps: 8 },
          { id: 's8', weight: 100, reps: 8 },
          { id: 's9', weight: 100, reps: 6 },
        ],
      },
      {
        exerciseId: '5',
        sets: [
          { id: 's10', weight: 0, reps: 12 },
          { id: 's11', weight: 0, reps: 10 },
          { id: 's12', weight: 0, reps: 8 },
        ],
      },
    ],
  },
  {
    id: '3',
    name: 'Ноги',
    date: '2024-12-16T11:00:00',
    duration: 90,
    exercises: [
      {
        exerciseId: '2',
        sets: [
          { id: 's13', weight: 120, reps: 10 },
          { id: 's14', weight: 120, reps: 10 },
          { id: 's15', weight: 120, reps: 8 },
          { id: 's16', weight: 110, reps: 10 },
        ],
      },
    ],
  },
];
