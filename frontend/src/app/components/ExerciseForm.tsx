import { useState, useEffect } from 'react';
import { Save } from 'lucide-react';
import { Exercise } from '../types';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Label } from './ui/label';
import { Textarea } from './ui/textarea';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from './ui/select';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from './ui/dialog';

interface ExerciseFormProps {
  exercise?: Exercise;
  isOpen: boolean;
  onClose: () => void;
  onSave: (exercise: Omit<Exercise, 'id'> & { id?: string }) => void;
}

const muscleGroups = [
  'Грудь',
  'Спина',
  'Плечи',
  'Руки',
  'Ноги',
  'Пресс',
  'Кардио',
];

export function ExerciseForm({ exercise, isOpen, onClose, onSave }: ExerciseFormProps) {
  const [name, setName] = useState('');
  const [muscleGroup, setMuscleGroup] = useState(muscleGroups[0]);
  const [description, setDescription] = useState('');

  useEffect(() => {
    if (exercise) {
      setName(exercise.name);
      setMuscleGroup(exercise.muscleGroup);
      setDescription(exercise.description || '');
    } else {
      setName('');
      setMuscleGroup(muscleGroups[0]);
      setDescription('');
    }
  }, [exercise, isOpen]);

  const handleSubmit = () => {
    if (!name.trim()) {
      alert('Введите название упражнения');
      return;
    }

    onSave({
      id: exercise?.id,
      name: name.trim(),
      muscleGroup,
      description: description.trim() || undefined,
    });

    onClose();
  };

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="bg-[var(--tg-theme-bg-color,#ffffff)]">
        <DialogHeader>
          <DialogTitle className="text-[var(--tg-theme-text-color,#000000)]">
            {exercise ? 'Редактировать упражнение' : 'Новое упражнение'}
          </DialogTitle>
        </DialogHeader>

        <div className="space-y-4 py-4">
          <div className="space-y-2">
            <Label htmlFor="exercise-name" className="text-[var(--tg-theme-text-color,#000000)]">
              Название упражнения
            </Label>
            <Input
              id="exercise-name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="Например: Жим штанги лежа"
              className="border-[var(--tg-theme-hint-color,#e0e0e0)]"
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="muscle-group" className="text-[var(--tg-theme-text-color,#000000)]">
              Группа мышц
            </Label>
            <Select value={muscleGroup} onValueChange={setMuscleGroup}>
              <SelectTrigger id="muscle-group" className="border-[var(--tg-theme-hint-color,#e0e0e0)]">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                {muscleGroups.map((group) => (
                  <SelectItem key={group} value={group}>
                    {group}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div className="space-y-2">
            <Label htmlFor="description" className="text-[var(--tg-theme-text-color,#000000)]">
              Описание (опционально)
            </Label>
            <Textarea
              id="description"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              placeholder="Краткое описание упражнения..."
              className="border-[var(--tg-theme-hint-color,#e0e0e0)] min-h-[80px]"
            />
          </div>

          <div className="flex gap-2 pt-2">
            <Button
              variant="outline"
              onClick={onClose}
              className="flex-1 border-[var(--tg-theme-hint-color,#e0e0e0)]"
            >
              Отмена
            </Button>
            <Button
              onClick={handleSubmit}
              className="flex-1 bg-[var(--tg-theme-button-color,#3390ec)] text-[var(--tg-theme-button-text-color,#ffffff)]"
            >
              <Save className="w-4 h-4 mr-2" />
              Сохранить
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
}
