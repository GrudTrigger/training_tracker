import { useState } from 'react';
import { Plus, Trash2, Save } from 'lucide-react';
import { Exercise, Workout, WorkoutExercise, Set } from '../types';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Label } from './ui/label';
import { Textarea } from './ui/textarea';
import { Card } from './ui/card';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from './ui/select';

interface AddWorkoutProps {
  exercises: Exercise[];
  onSave: (workout: Omit<Workout, 'id'>) => void;
  onCancel: () => void;
}

export function AddWorkout({ exercises, onSave, onCancel }: AddWorkoutProps) {
  const [name, setName] = useState('');
  const [notes, setNotes] = useState('');
  const [duration, setDuration] = useState('');
  const [workoutExercises, setWorkoutExercises] = useState<WorkoutExercise[]>([]);

  const addExercise = () => {
    if (exercises.length === 0) return;
    
    setWorkoutExercises([
      ...workoutExercises,
      {
        exerciseId: exercises[0].id,
        sets: [{ id: crypto.randomUUID(), weight: 0, reps: 0 }],
      },
    ]);
  };

  const removeExercise = (index: number) => {
    setWorkoutExercises(workoutExercises.filter((_, i) => i !== index));
  };

  const updateExercise = (index: number, exerciseId: string) => {
    const updated = [...workoutExercises];
    updated[index].exerciseId = exerciseId;
    setWorkoutExercises(updated);
  };

  const addSet = (exerciseIndex: number) => {
    const updated = [...workoutExercises];
    updated[exerciseIndex].sets.push({
      id: crypto.randomUUID(),
      weight: 0,
      reps: 0,
    });
    setWorkoutExercises(updated);
  };

  const removeSet = (exerciseIndex: number, setIndex: number) => {
    const updated = [...workoutExercises];
    updated[exerciseIndex].sets = updated[exerciseIndex].sets.filter(
      (_, i) => i !== setIndex
    );
    setWorkoutExercises(updated);
  };

  const updateSet = (
    exerciseIndex: number,
    setIndex: number,
    field: keyof Set,
    value: number
  ) => {
    const updated = [...workoutExercises];
    updated[exerciseIndex].sets[setIndex][field] = value;
    setWorkoutExercises(updated);
  };

  const handleSave = () => {
    if (!name.trim()) {
      alert('Введите название тренировки');
      return;
    }

    if (workoutExercises.length === 0) {
      alert('Добавьте хотя бы одно упражнение');
      return;
    }

    const workout: Omit<Workout, 'id'> = {
      name: name.trim(),
      date: new Date().toISOString(),
      exercises: workoutExercises,
      duration: duration ? parseInt(duration) : undefined,
      notes: notes.trim() || undefined,
    };

    onSave(workout);
  };

  const getExerciseName = (exerciseId: string) => {
    return exercises.find((e) => e.id === exerciseId)?.name || '';
  };

  return (
    <div className="flex flex-col h-full bg-[var(--tg-theme-bg-color,#ffffff)]">
      <div className="flex items-center justify-between p-4 border-b border-[var(--tg-theme-hint-color,#e0e0e0)]">
        <h1 className="text-[var(--tg-theme-text-color,#000000)]">Новая тренировка</h1>
        <div className="flex gap-2">
          <Button
            variant="outline"
            size="sm"
            onClick={onCancel}
            className="border-[var(--tg-theme-hint-color,#e0e0e0)]"
          >
            Отмена
          </Button>
          <Button
            size="sm"
            onClick={handleSave}
            className="bg-[var(--tg-theme-button-color,#3390ec)] text-[var(--tg-theme-button-text-color,#ffffff)]"
          >
            <Save className="w-4 h-4 mr-2" />
            Сохранить
          </Button>
        </div>
      </div>

      <div className="flex-1 overflow-y-auto p-4 space-y-4">
        <div className="space-y-2">
          <Label htmlFor="name" className="text-[var(--tg-theme-text-color,#000000)]">
            Название тренировки
          </Label>
          <Input
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Например: Грудь и трицепс"
            className="border-[var(--tg-theme-hint-color,#e0e0e0)]"
          />
        </div>

        <div className="space-y-2">
          <Label htmlFor="duration" className="text-[var(--tg-theme-text-color,#000000)]">
            Длительность (минуты)
          </Label>
          <Input
            id="duration"
            type="number"
            value={duration}
            onChange={(e) => setDuration(e.target.value)}
            placeholder="60"
            className="border-[var(--tg-theme-hint-color,#e0e0e0)]"
          />
        </div>

        <div className="space-y-3">
          <div className="flex items-center justify-between">
            <Label className="text-[var(--tg-theme-text-color,#000000)]">Упражнения</Label>
            <Button
              size="sm"
              onClick={addExercise}
              className="bg-[var(--tg-theme-button-color,#3390ec)] text-[var(--tg-theme-button-text-color,#ffffff)]"
            >
              <Plus className="w-4 h-4 mr-2" />
              Добавить упражнение
            </Button>
          </div>

          {workoutExercises.map((workoutEx, exIndex) => (
            <Card
              key={exIndex}
              className="p-4 space-y-3 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]"
            >
              <div className="flex items-center justify-between gap-2">
                <Select
                  value={workoutEx.exerciseId}
                  onValueChange={(value) => updateExercise(exIndex, value)}
                >
                  <SelectTrigger className="flex-1">
                    <SelectValue>
                      {getExerciseName(workoutEx.exerciseId)}
                    </SelectValue>
                  </SelectTrigger>
                  <SelectContent>
                    {exercises.map((exercise) => (
                      <SelectItem key={exercise.id} value={exercise.id}>
                        {exercise.name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={() => removeExercise(exIndex)}
                  className="text-red-500"
                >
                  <Trash2 className="w-4 h-4" />
                </Button>
              </div>

              <div className="space-y-2">
                <div className="flex items-center justify-between">
                  <span className="text-sm text-[var(--tg-theme-hint-color,#999999)]">
                    Подходы
                  </span>
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => addSet(exIndex)}
                  >
                    <Plus className="w-3 h-3 mr-1" />
                    Подход
                  </Button>
                </div>

                {workoutEx.sets.map((set, setIndex) => (
                  <div key={set.id} className="flex items-center gap-2">
                    <span className="text-sm text-[var(--tg-theme-hint-color,#999999)] w-8">
                      {setIndex + 1}.
                    </span>
                    <Input
                      type="number"
                      value={set.weight || ''}
                      onChange={(e) =>
                        updateSet(
                          exIndex,
                          setIndex,
                          'weight',
                          parseFloat(e.target.value) || 0
                        )
                      }
                      placeholder="Вес (кг)"
                      className="flex-1"
                    />
                    <Input
                      type="number"
                      value={set.reps || ''}
                      onChange={(e) =>
                        updateSet(
                          exIndex,
                          setIndex,
                          'reps',
                          parseInt(e.target.value) || 0
                        )
                      }
                      placeholder="Повт."
                      className="flex-1"
                    />
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={() => removeSet(exIndex, setIndex)}
                      className="text-red-500"
                    >
                      <Trash2 className="w-4 h-4" />
                    </Button>
                  </div>
                ))}
              </div>
            </Card>
          ))}
        </div>

        <div className="space-y-2">
          <Label htmlFor="notes" className="text-[var(--tg-theme-text-color,#000000)]">
            Заметки
          </Label>
          <Textarea
            id="notes"
            value={notes}
            onChange={(e) => setNotes(e.target.value)}
            placeholder="Дополнительные заметки..."
            className="border-[var(--tg-theme-hint-color,#e0e0e0)] min-h-[80px]"
          />
        </div>
      </div>
    </div>
  );
}
