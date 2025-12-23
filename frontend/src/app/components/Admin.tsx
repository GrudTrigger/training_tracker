import { TriangleAlert, Trash2, Download, Upload } from 'lucide-react';
import { Workout, Exercise } from '../types';
import { Card } from './ui/card';
import { Button } from './ui/button';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from './ui/alert-dialog';

interface AdminProps {
  workouts: Workout[];
  exercises: Exercise[];
  onClearAllData: () => void;
  onExportData: () => void;
  onImportData: (data: { workouts: Workout[]; exercises: Exercise[] }) => void;
}

export function Admin({
  workouts,
  exercises,
  onClearAllData,
  onExportData,
  onImportData,
}: AdminProps) {
  const handleImport = () => {
    const input = document.createElement('input');
    input.type = 'file';
    input.accept = '.json';
    input.onchange = (e) => {
      const file = (e.target as HTMLInputElement).files?.[0];
      if (!file) return;

      const reader = new FileReader();
      reader.onload = (event) => {
        try {
          const data = JSON.parse(event.target?.result as string);
          if (data.workouts && data.exercises) {
            onImportData(data);
            alert('Данные успешно импортированы!');
          } else {
            alert('Неверный формат файла');
          }
        } catch (error) {
          alert('Ошибка при чтении файла');
        }
      };
      reader.readAsText(file);
    };
    input.click();
  };

  return (
    <div className="flex flex-col h-full bg-[var(--tg-theme-bg-color,#ffffff)]">
      <div className="p-4 border-b border-[var(--tg-theme-hint-color,#e0e0e0)]">
        <h1 className="text-[var(--tg-theme-text-color,#000000)]">Администрирование</h1>
      </div>

      <div className="flex-1 overflow-y-auto p-4 space-y-4">
        {/* Статистика данных */}
        <Card className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]">
          <h3 className="text-[var(--tg-theme-text-color,#000000)] mb-3">
            Обзор данных
          </h3>
          <div className="space-y-2">
            <div className="flex justify-between items-center">
              <span className="text-[var(--tg-theme-hint-color,#999999)]">
                Всего тренировок:
              </span>
              <span className="text-[var(--tg-theme-text-color,#000000)]">
                {workouts.length}
              </span>
            </div>
            <div className="flex justify-between items-center">
              <span className="text-[var(--tg-theme-hint-color,#999999)]">
                Всего упражнений:
              </span>
              <span className="text-[var(--tg-theme-text-color,#000000)]">
                {exercises.length}
              </span>
            </div>
            <div className="flex justify-between items-center">
              <span className="text-[var(--tg-theme-hint-color,#999999)]">
                Размер данных:
              </span>
              <span className="text-[var(--tg-theme-text-color,#000000)]">
                {Math.round(
                  (JSON.stringify({ workouts, exercises }).length / 1024) * 100
                ) / 100}{' '}
                КБ
              </span>
            </div>
          </div>
        </Card>

        {/* Экспорт/Импорт */}
        <Card className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]">
          <h3 className="text-[var(--tg-theme-text-color,#000000)] mb-3">
            Резервное копирование
          </h3>
          <div className="space-y-3">
            <Button
              onClick={onExportData}
              className="w-full bg-[var(--tg-theme-button-color,#3390ec)] text-[var(--tg-theme-button-text-color,#ffffff)]"
            >
              <Download className="w-4 h-4 mr-2" />
              Экспортировать данные
            </Button>
            <Button
              onClick={handleImport}
              variant="outline"
              className="w-full border-[var(--tg-theme-hint-color,#e0e0e0)]"
            >
              <Upload className="w-4 h-4 mr-2" />
              Импортировать данные
            </Button>
            <p className="text-sm text-[var(--tg-theme-hint-color,#999999)]">
              Сохраните свои данные в файл JSON или восстановите из резервной копии
            </p>
          </div>
        </Card>

        {/* Опасная зона */}
        <Card className="p-4 border border-red-200 bg-red-50">
          <div className="flex items-start gap-3 mb-3">
            <TriangleAlert className="w-5 h-5 text-red-500 flex-shrink-0 mt-0.5" />
            <div>
              <h3 className="text-red-700 mb-1">Опасная зона</h3>
              <p className="text-sm text-red-600">
                Будьте осторожны с этими действиями. Они необратимы!
              </p>
            </div>
          </div>

          <AlertDialog>
            <AlertDialogTrigger asChild>
              <Button variant="destructive" className="w-full">
                <Trash2 className="w-4 h-4 mr-2" />
                Удалить все данные
              </Button>
            </AlertDialogTrigger>
            <AlertDialogContent className="bg-[var(--tg-theme-bg-color,#ffffff)]">
              <AlertDialogHeader>
                <AlertDialogTitle className="text-[var(--tg-theme-text-color,#000000)]">
                  Вы абсолютно уверены?
                </AlertDialogTitle>
                <AlertDialogDescription className="text-[var(--tg-theme-hint-color,#999999)]">
                  Это действие нельзя отменить. Будут удалены все тренировки и
                  упражнения. Перед удалением рекомендуется сделать экспорт данных.
                </AlertDialogDescription>
              </AlertDialogHeader>
              <AlertDialogFooter>
                <AlertDialogCancel className="border-[var(--tg-theme-hint-color,#e0e0e0)]">
                  Отмена
                </AlertDialogCancel>
                <AlertDialogAction
                  onClick={onClearAllData}
                  className="bg-red-500 text-white hover:bg-red-600"
                >
                  Да, удалить все
                </AlertDialogAction>
              </AlertDialogFooter>
            </AlertDialogContent>
          </AlertDialog>
        </Card>

        {/* Информация */}
        <Card className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]">
          <h3 className="text-[var(--tg-theme-text-color,#000000)] mb-3">
            О приложении
          </h3>
          <div className="space-y-2 text-sm text-[var(--tg-theme-hint-color,#999999)]">
            <p>Training Tracker v1.0.0</p>
            <p>Telegram Mini App</p>
            <p>© 2024</p>
          </div>
        </Card>

        {/* Инструкции */}
        <Card className="p-4 border border-[var(--tg-theme-hint-color,#e0e0e0)] bg-[var(--tg-theme-secondary-bg-color,#f7f7f7)]">
          <h3 className="text-[var(--tg-theme-text-color,#000000)] mb-3">
            Как использовать
          </h3>
          <div className="space-y-2 text-sm text-[var(--tg-theme-hint-color,#999999)]">
            <p>
              <strong className="text-[var(--tg-theme-text-color,#000000)]">
                Экспорт данных:
              </strong>{' '}
              Сохраняет все ваши тренировки и упражнения в файл JSON
            </p>
            <p>
              <strong className="text-[var(--tg-theme-text-color,#000000)]">
                Импорт данных:
              </strong>{' '}
              Загружает данные из ранее сохраненного файла
            </p>
            <p>
              <strong className="text-[var(--tg-theme-text-color,#000000)]">
                Удалить все:
              </strong>{' '}
              Полностью очищает базу данных приложения
            </p>
          </div>
        </Card>
      </div>
    </div>
  );
}