import { useState } from 'react';
import { Dumbbell, Plus, Search, Pencil, Trash2 } from 'lucide-react';
import { Exercise } from '../types';
import { Card } from './ui/card';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Badge } from './ui/badge';

interface ExercisesListProps {
  exercises: Exercise[];
  onAddExercise: () => void;
  onEditExercise: (exercise: Exercise) => void;
  onDeleteExercise: (exerciseId: string) => void;
}

export function ExercisesList({
  exercises,
  onAddExercise,
  onEditExercise,
  onDeleteExercise,
}: ExercisesListProps) {
  const [searchQuery, setSearchQuery] = useState('');
  const [selectedMuscleGroup, setSelectedMuscleGroup] = useState<string | null>(null);

  const muscleGroups = Array.from(new Set(exercises.map((e) => e.muscleGroup)));

  const filteredExercises = exercises.filter((exercise) => {
    const matchesSearch = exercise.name
      .toLowerCase()
      .includes(searchQuery.toLowerCase());
    const matchesMuscleGroup =
      !selectedMuscleGroup || exercise.muscleGroup === selectedMuscleGroup;
    return matchesSearch && matchesMuscleGroup;
  });

  return (
    <div className="flex flex-col h-full bg-[var(--tg-theme-bg-color,#ffffff)]">
      <div className="flex items-center justify-between p-4 border-b border-[var(--tg-theme-hint-color,#e0e0e0)]">
        <h1 className="text-[var(--tg-theme-text-color,#000000)]">Упражнения</h1>
        <Button
          onClick={onAddExercise}
          className="bg-[var(--tg-theme-button-color,#3390ec)] text-[var(--tg-theme-button-text-color,#ffffff)] hover:opacity-90"
          size="sm"
        >
          <Plus className="w-4 h-4 mr-2" />
          Добавить
        </Button>
      </div>

      <div className="p-4 space-y-3 border-b border-[var(--tg-theme-hint-color,#e0e0e0)]">
        <div className="relative">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-[var(--tg-theme-hint-color,#999999)]" />
          <Input
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            placeholder="Поиск упражнений..."
            className="pl-9 border-[var(--tg-theme-hint-color,#e0e0e0)]"
          />
        </div>

        <div className="flex gap-2 overflow-x-auto pb-2">
          <Button
            variant={selectedMuscleGroup === null ? 'default' : 'outline'}
            size="sm"
            onClick={() => setSelectedMuscleGroup(null)}
            className={
              selectedMuscleGroup === null
                ? 'bg-[var(--tg-theme-button-color,#3390ec)] text-[var(--tg-theme-button-text-color,#ffffff)]'
                : ''
            }
          >
            Все
          </Button>
          {muscleGroups.map((group) => (
            <Button
              key={group}
              variant={selectedMuscleGroup === group ? 'default' : 'outline'}
              size="sm"
              onClick={() => setSelectedMuscleGroup(group)}
              className={
                selectedMuscleGroup === group
                  ? 'bg-[var(--tg-theme-button-color,#3390ec)] text-[var(--tg-theme-button-text-color,#ffffff)] whitespace-nowrap'
                  : 'whitespace-nowrap'
              }
            >
              {group}
            </Button>
          ))}
        </div>
      </div>

      <div className="flex-1 overflow-y-auto p-4 space-y-3">
        {filteredExercises.length === 0 ? (
          <div className="flex flex-col items-center justify-center h-full text-center py-12">
            <Dumbbell className="w-16 h-16 text-[var(--tg-theme-hint-color,#999999)] mb-4" />
            <p className="text-[var(--tg-theme-hint-color,#999999)] mb-4">
              {searchQuery || selectedMuscleGroup
                ? 'Упражнения не найдены'
                : 'У вас пока нет упражнений'}
            </p>
            {!searchQuery && !selectedMuscleGroup && (
              <Button
                onClick={onAddExercise}
                className="bg-[var(--tg-theme-button-color,#3390ec)] text-[var(--tg-theme-button-text-color,#ffffff)]"
              >
                <Plus className="w-4 h-4 mr-2" />
                Добавить упражнение
              </Button>
            )}
          </div>
        ) : (
          filteredExercises.map((exercise) => (
            <Card
              key={exercise.id}
              className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]"
            >
              <div className="flex items-start justify-between gap-3">
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 mb-2">
                    <h3 className="text-[var(--tg-theme-text-color,#000000)] truncate">
                      {exercise.name}
                    </h3>
                    <Badge
                      variant="secondary"
                      className="bg-[var(--tg-theme-button-color,#3390ec)]/10 text-[var(--tg-theme-button-color,#3390ec)] whitespace-nowrap"
                    >
                      {exercise.muscleGroup}
                    </Badge>
                  </div>
                  {exercise.description && (
                    <p className="text-sm text-[var(--tg-theme-hint-color,#999999)]">
                      {exercise.description}
                    </p>
                  )}
                </div>

                <div className="flex gap-2">
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={() => onEditExercise(exercise)}
                    className="text-[var(--tg-theme-button-color,#3390ec)]"
                  >
                    <Pencil className="w-4 h-4" />
                  </Button>
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={() => onDeleteExercise(exercise.id)}
                    className="text-red-500"
                  >
                    <Trash2 className="w-4 h-4" />
                  </Button>
                </div>
              </div>
            </Card>
          ))
        )}
      </div>
    </div>
  );
}