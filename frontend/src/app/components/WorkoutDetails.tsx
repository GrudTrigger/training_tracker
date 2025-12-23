import { ArrowLeft, Calendar, Clock, Dumbbell } from 'lucide-react';
import { Workout, Exercise } from '../types';
import { Button } from './ui/button';
import { Card } from './ui/card';
import { format } from 'date-fns';
import { ru } from 'date-fns/locale';

interface WorkoutDetailsProps {
  workout: Workout;
  exercises: Exercise[];
  onClose: () => void;
}

export function WorkoutDetails({ workout, exercises, onClose }: WorkoutDetailsProps) {
  const getExerciseName = (exerciseId: string) => {
    return exercises.find((e) => e.id === exerciseId)?.name || 'Неизвестное упражнение';
  };

  const getTotalVolume = () => {
    return workout.exercises.reduce((total, ex) => {
      return (
        total +
        ex.sets.reduce((sum, set) => sum + set.weight * set.reps, 0)
      );
    }, 0);
  };

  return (
    <div className="flex flex-col h-full bg-[var(--tg-theme-bg-color,#ffffff)]">
      <div className="flex items-center gap-3 p-4 border-b border-[var(--tg-theme-hint-color,#e0e0e0)]">
        <Button
          variant="ghost"
          size="sm"
          onClick={onClose}
          className="text-[var(--tg-theme-button-color,#3390ec)]"
        >
          <ArrowLeft className="w-4 h-4 mr-2" />
          Назад
        </Button>
      </div>

      <div className="flex-1 overflow-y-auto p-4 space-y-4">
        <div>
          <h1 className="text-[var(--tg-theme-text-color,#000000)] mb-4">
            {workout.name}
          </h1>

          <div className="flex flex-wrap gap-4 text-[var(--tg-theme-hint-color,#999999)]">
            <div className="flex items-center gap-2">
              <Calendar className="w-4 h-4" />
              <span>
                {format(new Date(workout.date), 'd MMMM yyyy, HH:mm', {
                  locale: ru,
                })}
              </span>
            </div>

            {workout.duration && (
              <div className="flex items-center gap-2">
                <Clock className="w-4 h-4" />
                <span>{workout.duration} минут</span>
              </div>
            )}

            <div className="flex items-center gap-2">
              <Dumbbell className="w-4 h-4" />
              <span>{getTotalVolume()} кг (общий объем)</span>
            </div>
          </div>
        </div>

        {workout.notes && (
          <Card className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-yellow-50">
            <p className="text-sm text-[var(--tg-theme-text-color,#000000)]">
              <strong>Заметки:</strong> {workout.notes}
            </p>
          </Card>
        )}

        <div className="space-y-3">
          <h2 className="text-[var(--tg-theme-text-color,#000000)]">Упражнения</h2>

          {workout.exercises.map((workoutEx, exIndex) => {
            const totalSets = workoutEx.sets.length;
            const totalReps = workoutEx.sets.reduce((sum, set) => sum + set.reps, 0);
            const totalWeight = workoutEx.sets.reduce(
              (sum, set) => sum + set.weight * set.reps,
              0
            );

            return (
              <Card
                key={exIndex}
                className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]"
              >
                <div className="mb-3">
                  <h3 className="text-[var(--tg-theme-text-color,#000000)] mb-2">
                    {getExerciseName(workoutEx.exerciseId)}
                  </h3>
                  <div className="flex gap-4 text-sm text-[var(--tg-theme-hint-color,#999999)]">
                    <span>{totalSets} подходов</span>
                    <span>{totalReps} повторений</span>
                    <span>{totalWeight} кг</span>
                  </div>
                </div>

                <div className="space-y-2">
                  <div className="grid grid-cols-3 gap-2 text-sm text-[var(--tg-theme-hint-color,#999999)] pb-2 border-b border-[var(--tg-theme-hint-color,#e0e0e0)]">
                    <span>Подход</span>
                    <span>Вес (кг)</span>
                    <span>Повторения</span>
                  </div>

                  {workoutEx.sets.map((set, setIndex) => (
                    <div
                      key={set.id}
                      className="grid grid-cols-3 gap-2 text-sm text-[var(--tg-theme-text-color,#000000)]"
                    >
                      <span className="text-[var(--tg-theme-hint-color,#999999)]">
                        {setIndex + 1}
                      </span>
                      <span>{set.weight}</span>
                      <span>{set.reps}</span>
                    </div>
                  ))}
                </div>
              </Card>
            );
          })}
        </div>
      </div>
    </div>
  );
}
