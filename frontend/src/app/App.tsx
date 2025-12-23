import { useState, useEffect } from 'react';
import { House, Plus, Dumbbell, ChartBar, Settings } from 'lucide-react';
import { Workout, Exercise } from './types';
import { mockExercises, mockWorkouts } from './utils/mockData';
import { WorkoutsList } from './components/WorkoutsList';
import { AddWorkout } from './components/AddWorkout';
import { ExercisesList } from './components/ExercisesList';
import { Statistics } from './components/Statistics';
import { Admin } from './components/Admin';
import { WorkoutDetails } from './components/WorkoutDetails';
import { ExerciseForm } from './components/ExerciseForm';
import { Toaster } from './components/ui/sonner';
import { toast } from 'sonner';

type Tab = 'workouts' | 'add' | 'exercises' | 'stats' | 'admin';

export default function App() {
  const [activeTab, setActiveTab] = useState<Tab>('workouts');
  const [workouts, setWorkouts] = useState<Workout[]>([]);
  const [exercises, setExercises] = useState<Exercise[]>([]);
  const [selectedWorkout, setSelectedWorkout] = useState<Workout | null>(null);
  const [editingExercise, setEditingExercise] = useState<Exercise | null>(null);
  const [isExerciseFormOpen, setIsExerciseFormOpen] = useState(false);

  // Инициализация с mock данными
  useEffect(() => {
    setWorkouts(mockWorkouts);
    setExercises(mockExercises);
  }, []);

  // Обработчики для тренировок
  const handleAddWorkout = (workout: Omit<Workout, 'id'>) => {
    const newWorkout: Workout = {
      ...workout,
      id: crypto.randomUUID(),
    };
    setWorkouts([newWorkout, ...workouts]);
    setActiveTab('workouts');
    toast.success('Тренировка добавлена!');
  };

  const handleWorkoutClick = (workout: Workout) => {
    setSelectedWorkout(workout);
  };

  const handleCloseWorkoutDetails = () => {
    setSelectedWorkout(null);
  };

  // Обработчики для упражнений
  const handleSaveExercise = (exercise: Omit<Exercise, 'id'> & { id?: string }) => {
    if (exercise.id) {
      // Редактирование существующего
      setExercises(
        exercises.map((e) =>
          e.id === exercise.id ? { ...exercise, id: exercise.id } : e
        )
      );
      toast.success('Упражнение обновлено!');
    } else {
      // Добавление нового
      const newExercise: Exercise = {
        ...exercise,
        id: crypto.randomUUID(),
      };
      setExercises([...exercises, newExercise]);
      toast.success('Упражнение добавлено!');
    }
    setEditingExercise(null);
  };

  const handleDeleteExercise = (exerciseId: string) => {
    if (confirm('Вы уверены, что хотите удалить это упражнение?')) {
      setExercises(exercises.filter((e) => e.id !== exerciseId));
      toast.success('Упражнение удалено');
    }
  };

  const handleEditExercise = (exercise: Exercise) => {
    setEditingExercise(exercise);
    setIsExerciseFormOpen(true);
  };

  const handleAddExercise = () => {
    setEditingExercise(null);
    setIsExerciseFormOpen(true);
  };

  // Обработчики для админки
  const handleClearAllData = () => {
    setWorkouts([]);
    setExercises([]);
    toast.success('Все данные удалены');
  };

  const handleExportData = () => {
    const data = {
      workouts,
      exercises,
      exportDate: new Date().toISOString(),
    };
    const blob = new Blob([JSON.stringify(data, null, 2)], {
      type: 'application/json',
    });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `training-tracker-backup-${new Date().toISOString().split('T')[0]}.json`;
    a.click();
    URL.revokeObjectURL(url);
    toast.success('Данные экспортированы!');
  };

  const handleImportData = (data: { workouts: Workout[]; exercises: Exercise[] }) => {
    setWorkouts(data.workouts);
    setExercises(data.exercises);
  };

  // Рендер активной страницы
  const renderContent = () => {
    if (selectedWorkout) {
      return (
        <WorkoutDetails
          workout={selectedWorkout}
          exercises={exercises}
          onClose={handleCloseWorkoutDetails}
        />
      );
    }

    switch (activeTab) {
      case 'workouts':
        return (
          <WorkoutsList
            workouts={workouts}
            exercises={exercises}
            onWorkoutClick={handleWorkoutClick}
            onAddWorkout={() => setActiveTab('add')}
          />
        );
      case 'add':
        return (
          <AddWorkout
            exercises={exercises}
            onSave={handleAddWorkout}
            onCancel={() => setActiveTab('workouts')}
          />
        );
      case 'exercises':
        return (
          <ExercisesList
            exercises={exercises}
            onAddExercise={handleAddExercise}
            onEditExercise={handleEditExercise}
            onDeleteExercise={handleDeleteExercise}
          />
        );
      case 'stats':
        return <Statistics workouts={workouts} exercises={exercises} />;
      case 'admin':
        return (
          <Admin
            workouts={workouts}
            exercises={exercises}
            onClearAllData={handleClearAllData}
            onExportData={handleExportData}
            onImportData={handleImportData}
          />
        );
      default:
        return null;
    }
  };

  const tabs = [
    { id: 'workouts' as Tab, label: 'Главная', icon: House },
    { id: 'add' as Tab, label: 'Добавить', icon: Plus },
    { id: 'exercises' as Tab, label: 'Упражнения', icon: Dumbbell },
    { id: 'stats' as Tab, label: 'Статистика', icon: ChartBar },
    { id: 'admin' as Tab, label: 'Админка', icon: Settings },
  ];

  return (
    <div className="h-screen flex flex-col bg-[var(--tg-theme-bg-color,#ffffff)]">
      {/* Основной контент */}
      <div className="flex-1 overflow-hidden">{renderContent()}</div>

      {/* Нижняя навигация - скрываем при просмотре деталей */}
      {!selectedWorkout && (
        <div className="border-t border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-bg-color,#ffffff)]">
          <div className="flex justify-around items-center h-16">
            {tabs.map((tab) => {
              const Icon = tab.icon;
              const isActive = activeTab === tab.id;
              return (
                <button
                  key={tab.id}
                  onClick={() => setActiveTab(tab.id)}
                  className={`flex flex-col items-center justify-center gap-1 px-3 py-2 transition-colors ${
                    isActive
                      ? 'text-[var(--tg-theme-button-color,#3390ec)]'
                      : 'text-[var(--tg-theme-hint-color,#999999)]'
                  }`}
                >
                  <Icon className="w-5 h-5" />
                  <span className="text-xs">{tab.label}</span>
                </button>
              );
            })}
          </div>
        </div>
      )}

      {/* Exercise Form Dialog */}
      <ExerciseForm
        exercise={editingExercise || undefined}
        isOpen={isExerciseFormOpen}
        onClose={() => {
          setIsExerciseFormOpen(false);
          setEditingExercise(null);
        }}
        onSave={handleSaveExercise}
      />

      {/* Toast notifications */}
      <Toaster position="top-center" />
    </div>
  );
}