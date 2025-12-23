export interface Exercise {
  id: string;
  name: string;
  muscleGroup: string;
  description?: string;
}

export interface WorkoutExercise {
  exerciseId: string;
  sets: Set[];
}

export interface Set {
  id: string;
  weight: number;
  reps: number;
}

export interface Workout {
  id: string;
  name: string;
  date: string;
  exercises: WorkoutExercise[];
  duration?: number;
  notes?: string;
}

export interface Stats {
  totalWorkouts: number;
  totalSets: number;
  totalReps: number;
  favoriteExercises: { name: string; count: number }[];
}
