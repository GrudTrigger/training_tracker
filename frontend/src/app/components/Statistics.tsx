import { TrendingUp, Dumbbell, Activity, Target } from 'lucide-react';
import { Workout, Exercise } from '../types';
import { Card } from './ui/card';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, LineChart, Line } from 'recharts';
import { format, subDays } from 'date-fns';
import { ru } from 'date-fns/locale';

interface StatisticsProps {
  workouts: Workout[];
  exercises: Exercise[];
}

export function Statistics({ workouts, exercises }: StatisticsProps) {
  // Подсчет общей статистики
  const totalWorkouts = workouts.length;
  const totalSets = workouts.reduce(
    (total, workout) =>
      total + workout.exercises.reduce((sum, ex) => sum + ex.sets.length, 0),
    0
  );
  const totalReps = workouts.reduce(
    (total, workout) =>
      total +
      workout.exercises.reduce(
        (sum, ex) => sum + ex.sets.reduce((s, set) => s + set.reps, 0),
        0
      ),
    0
  );

  // Топ упражнений по частоте использования
  const exerciseFrequency = new Map<string, number>();
  workouts.forEach((workout) => {
    workout.exercises.forEach((ex) => {
      const count = exerciseFrequency.get(ex.exerciseId) || 0;
      exerciseFrequency.set(ex.exerciseId, count + 1);
    });
  });

  const topExercises = Array.from(exerciseFrequency.entries())
    .map(([exerciseId, count]) => ({
      name: exercises.find((e) => e.id === exerciseId)?.name || 'Неизвестно',
      count,
    }))
    .sort((a, b) => b.count - a.count)
    .slice(0, 5);

  // График тренировок за последние 7 дней
  const last7Days = Array.from({ length: 7 }, (_, i) => {
    const date = subDays(new Date(), 6 - i);
    const dayWorkouts = workouts.filter(
      (w) => format(new Date(w.date), 'yyyy-MM-dd') === format(date, 'yyyy-MM-dd')
    );
    
    return {
      date: format(date, 'dd.MM', { locale: ru }),
      workouts: dayWorkouts.length,
      sets: dayWorkouts.reduce(
        (total, w) => total + w.exercises.reduce((sum, ex) => sum + ex.sets.length, 0),
        0
      ),
    };
  });

  // Средняя длительность тренировки
  const workoutsWithDuration = workouts.filter((w) => w.duration);
  const avgDuration = workoutsWithDuration.length
    ? Math.round(
        workoutsWithDuration.reduce((sum, w) => sum + (w.duration || 0), 0) /
          workoutsWithDuration.length
      )
    : 0;

  return (
    <div className="flex flex-col h-full bg-[var(--tg-theme-bg-color,#ffffff)]">
      <div className="p-4 border-b border-[var(--tg-theme-hint-color,#e0e0e0)]">
        <h1 className="text-[var(--tg-theme-text-color,#000000)]">Статистика</h1>
      </div>

      <div className="flex-1 overflow-y-auto p-4 space-y-4">
        {/* Основные метрики */}
        <div className="grid grid-cols-2 gap-3">
          <Card className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]">
            <div className="flex items-center gap-3">
              <div className="p-2 rounded-lg bg-[var(--tg-theme-button-color,#3390ec)]/10">
                <Dumbbell className="w-5 h-5 text-[var(--tg-theme-button-color,#3390ec)]" />
              </div>
              <div>
                <p className="text-2xl text-[var(--tg-theme-text-color,#000000)]">
                  {totalWorkouts}
                </p>
                <p className="text-sm text-[var(--tg-theme-hint-color,#999999)]">
                  Тренировок
                </p>
              </div>
            </div>
          </Card>

          <Card className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]">
            <div className="flex items-center gap-3">
              <div className="p-2 rounded-lg bg-green-500/10">
                <Activity className="w-5 h-5 text-green-500" />
              </div>
              <div>
                <p className="text-2xl text-[var(--tg-theme-text-color,#000000)]">
                  {totalSets}
                </p>
                <p className="text-sm text-[var(--tg-theme-hint-color,#999999)]">
                  Подходов
                </p>
              </div>
            </div>
          </Card>

          <Card className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]">
            <div className="flex items-center gap-3">
              <div className="p-2 rounded-lg bg-orange-500/10">
                <Target className="w-5 h-5 text-orange-500" />
              </div>
              <div>
                <p className="text-2xl text-[var(--tg-theme-text-color,#000000)]">
                  {totalReps}
                </p>
                <p className="text-sm text-[var(--tg-theme-hint-color,#999999)]">
                  Повторений
                </p>
              </div>
            </div>
          </Card>

          <Card className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]">
            <div className="flex items-center gap-3">
              <div className="p-2 rounded-lg bg-purple-500/10">
                <TrendingUp className="w-5 h-5 text-purple-500" />
              </div>
              <div>
                <p className="text-2xl text-[var(--tg-theme-text-color,#000000)]">
                  {avgDuration}
                </p>
                <p className="text-sm text-[var(--tg-theme-hint-color,#999999)]">
                  Средняя длит. (мин)
                </p>
              </div>
            </div>
          </Card>
        </div>

        {/* График активности за 7 дней */}
        <Card className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]">
          <h3 className="text-[var(--tg-theme-text-color,#000000)] mb-4">
            Активность за последние 7 дней
          </h3>
          <ResponsiveContainer width="100%" height={200}>
            <LineChart data={last7Days}>
              <CartesianGrid strokeDasharray="3 3" stroke="#e0e0e0" />
              <XAxis 
                dataKey="date" 
                tick={{ fill: 'var(--tg-theme-hint-color,#999999)', fontSize: 12 }}
              />
              <YAxis 
                tick={{ fill: 'var(--tg-theme-hint-color,#999999)', fontSize: 12 }}
              />
              <Tooltip 
                contentStyle={{ 
                  backgroundColor: 'var(--tg-theme-secondary-bg-color,#f7f7f7)',
                  border: '1px solid var(--tg-theme-hint-color,#e0e0e0)',
                  borderRadius: '8px'
                }}
              />
              <Line 
                type="monotone" 
                dataKey="workouts" 
                stroke="var(--tg-theme-button-color,#3390ec)" 
                strokeWidth={2}
                name="Тренировки"
              />
            </LineChart>
          </ResponsiveContainer>
        </Card>

        {/* Топ упражнений */}
        <Card className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]">
          <h3 className="text-[var(--tg-theme-text-color,#000000)] mb-4">
            Топ-5 упражнений
          </h3>
          {topExercises.length > 0 ? (
            <ResponsiveContainer width="100%" height={250}>
              <BarChart data={topExercises} layout="vertical">
                <CartesianGrid strokeDasharray="3 3" stroke="#e0e0e0" />
                <XAxis 
                  type="number" 
                  tick={{ fill: 'var(--tg-theme-hint-color,#999999)', fontSize: 12 }}
                />
                <YAxis 
                  dataKey="name" 
                  type="category" 
                  width={150}
                  tick={{ fill: 'var(--tg-theme-hint-color,#999999)', fontSize: 12 }}
                />
                <Tooltip 
                  contentStyle={{ 
                    backgroundColor: 'var(--tg-theme-secondary-bg-color,#f7f7f7)',
                    border: '1px solid var(--tg-theme-hint-color,#e0e0e0)',
                    borderRadius: '8px'
                  }}
                />
                <Bar 
                  dataKey="count" 
                  fill="var(--tg-theme-button-color,#3390ec)" 
                  name="Использований"
                />
              </BarChart>
            </ResponsiveContainer>
          ) : (
            <p className="text-center text-[var(--tg-theme-hint-color,#999999)] py-8">
              Недостаточно данных для отображения
            </p>
          )}
        </Card>
      </div>
    </div>
  );
}
