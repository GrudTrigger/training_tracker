import { Dumbbell, Calendar, Clock, ChevronRight, Plus } from 'lucide-react';
import { Workout, Exercise } from '../types';
import { Card } from './ui/card';
import { Button } from './ui/button';
import { format } from 'date-fns';
import { ru } from 'date-fns/locale';
import {useEffect} from "react";

interface WorkoutsListProps {
  workouts: Workout[];
  exercises: Exercise[];
  onWorkoutClick: (workout: Workout) => void;
  onAddWorkout: () => void;
}

export function WorkoutsList({ workouts, exercises, onWorkoutClick, onAddWorkout }: WorkoutsListProps) {

  useEffect(()=>{
      fetch("http://localhost:8080/trainings/all").then(res => console.log(res.json()))
  })

  const getExerciseName = (exerciseId: string) => {
    return exercises.find(e => e.id === exerciseId)?.name || 'Неизвестное упражнение';
  };

  const getTotalSets = (workout: Workout) => {
    return workout.exercises.reduce((total, ex) => total + ex.sets.length, 0);
  };

  return (
    <div className="flex flex-col h-full bg-[var(--tg-theme-bg-color,#ffffff)]">
      <div className="flex items-center justify-between p-4 border-b border-[var(--tg-theme-hint-color,#e0e0e0)]">
        <h1 className="text-[var(--tg-theme-text-color,#000000)]">Мои тренировки</h1>
        <Button 
          onClick={onAddWorkout}
          className="bg-[var(--tg-theme-button-color,#3390ec)] text-[var(--tg-theme-button-text-color,#ffffff)] hover:opacity-90"
          size="sm"
        >
          <Plus className="w-4 h-4 mr-2" />
          Добавить
        </Button>
      </div>

      <div className="flex-1 overflow-y-auto p-4 space-y-3">
        {workouts.length === 0 ? (
          <div className="flex flex-col items-center justify-center h-full text-center py-12">
            <Dumbbell className="w-16 h-16 text-[var(--tg-theme-hint-color,#999999)] mb-4" />
            <p className="text-[var(--tg-theme-hint-color,#999999)] mb-4">
              У вас пока нет тренировок
            </p>
            <Button 
              onClick={onAddWorkout}
              className="bg-[var(--tg-theme-button-color,#3390ec)] text-[var(--tg-theme-button-text-color,#ffffff)]"
            >
              <Plus className="w-4 h-4 mr-2" />
              Добавить первую тренировку
            </Button>
          </div>
        ) : (
          workouts.map((workout) => (
            <Card
              key={workout.id}
              className="p-4 cursor-pointer hover:shadow-md transition-shadow border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]"
              onClick={() => onWorkoutClick(workout)}
            >
              <div className="flex items-start justify-between">
                <div className="flex-1">
                  <h3 className="text-[var(--tg-theme-text-color,#000000)] mb-2">
                    {workout.name}
                  </h3>
                  
                  <div className="flex items-center gap-4 mb-2 text-[var(--tg-theme-hint-color,#999999)]">
                    <div className="flex items-center gap-1">
                      <Calendar className="w-3 h-3" />
                      <span className="text-sm">
                        {format(new Date(workout.date), 'd MMM', { locale: ru })}
                      </span>
                    </div>
                    
                    {workout.duration && (
                      <div className="flex items-center gap-1">
                        <Clock className="w-3 h-3" />
                        <span className="text-sm">{workout.duration} мин</span>
                      </div>
                    )}
                    
                    <div className="flex items-center gap-1">
                      <Dumbbell className="w-3 h-3" />
                      <span className="text-sm">{getTotalSets(workout)} подходов</span>
                    </div>
                  </div>

                  <div className="text-sm text-[var(--tg-theme-hint-color,#999999)]">
                    {workout.exercises.map((ex, idx) => (
                      <span key={ex.exerciseId}>
                        {getExerciseName(ex.exerciseId)}
                        {idx < workout.exercises.length - 1 ? ', ' : ''}
                      </span>
                    ))}
                  </div>
                </div>

                <ChevronRight className="w-5 h-5 text-[var(--tg-theme-hint-color,#999999)] flex-shrink-0 ml-2" />
              </div>
            </Card>
          ))
        )}
      </div>
    </div>
  );
}