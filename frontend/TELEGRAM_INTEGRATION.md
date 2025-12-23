# Интеграция с Telegram Mini Apps

Это руководство по интеграции вашего Training Tracker с Telegram Mini Apps.

## Подготовка

### 1. Установка Telegram Web App SDK

Добавьте скрипт в `index.html` (если его еще нет):

```html
<script src="https://telegram.org/js/telegram-web-app.js"></script>
```

### 2. Инициализация Telegram Web App

Создайте файл `src/app/utils/telegram.ts`:

```typescript
// Получаем объект Telegram Web App
export const tg = window.Telegram?.WebApp;

// Инициализация приложения
export function initTelegramApp() {
  if (tg) {
    // Расширяем приложение на весь экран
    tg.expand();
    
    // Включаем кнопку закрытия
    tg.enableClosingConfirmation();
    
    // Устанавливаем цвет заголовка
    tg.setHeaderColor('secondary_bg_color');
    
    // Готово к использованию
    tg.ready();
  }
}

// Получение данных пользователя
export function getTelegramUser() {
  return tg?.initDataUnsafe?.user;
}

// Отправка данных обратно в бот
export function sendDataToBot(data: any) {
  if (tg) {
    tg.sendData(JSON.stringify(data));
  }
}

// Показать главную кнопку
export function showMainButton(text: string, onClick: () => void) {
  if (tg) {
    tg.MainButton.setText(text);
    tg.MainButton.show();
    tg.MainButton.onClick(onClick);
  }
}

// Скрыть главную кнопку
export function hideMainButton() {
  if (tg) {
    tg.MainButton.hide();
  }
}

// Показать кнопку назад
export function showBackButton(onClick: () => void) {
  if (tg) {
    tg.BackButton.show();
    tg.BackButton.onClick(onClick);
  }
}

// Скрыть кнопку назад
export function hideBackButton() {
  if (tg) {
    tg.BackButton.hide();
  }
}

// Haptic feedback
export function hapticFeedback(type: 'light' | 'medium' | 'heavy' | 'success' | 'warning' | 'error') {
  if (tg?.HapticFeedback) {
    if (type === 'success' || type === 'warning' || type === 'error') {
      tg.HapticFeedback.notificationOccurred(type);
    } else {
      tg.HapticFeedback.impactOccurred(type);
    }
  }
}
```

### 3. Использование в App.tsx

Добавьте в начало компонента:

```typescript
import { initTelegramApp, getTelegramUser } from './utils/telegram';

// В useEffect
useEffect(() => {
  initTelegramApp();
  const user = getTelegramUser();
  console.log('Telegram User:', user);
}, []);
```

## Создание Telegram Bot

### 1. Создайте бота через @BotFather

1. Откройте Telegram и найдите @BotFather
2. Отправьте `/newbot`
3. Следуйте инструкциям для создания бота
4. Сохраните токен бота

### 2. Создайте Mini App

1. Отправьте `/newapp` в @BotFather
2. Выберите вашего бота
3. Введите название, описание и загрузите иконку
4. Укажите URL вашего развернутого приложения
5. Загрузите GIF/демонстрацию (опционально)

### 3. Настройте Web App URL

```
/setmenubutton - установить кнопку в меню
```

Укажите URL вашего приложения (например, с Vercel, Netlify, GitHub Pages)

## Деплой приложения

### Vercel (Рекомендуется)

1. Установите Vercel CLI:
```bash
npm i -g vercel
```

2. Деплой:
```bash
vercel
```

3. Скопируйте URL и используйте его в настройках Mini App

### Netlify

1. Установите Netlify CLI:
```bash
npm i -g netlify-cli
```

2. Деплой:
```bash
netlify deploy --prod
```

### GitHub Pages

1. Добавьте в `vite.config.ts`:
```typescript
base: '/your-repo-name/'
```

2. Создайте workflow `.github/workflows/deploy.yml`
3. Включите GitHub Pages в настройках репозитория

## Персистентность данных

Для сохранения данных между сессиями используйте:

### Cloud Storage от Telegram

```typescript
// Сохранение
tg.CloudStorage.setItem('workouts', JSON.stringify(workouts));

// Загрузка
tg.CloudStorage.getItem('workouts', (error, value) => {
  if (!error && value) {
    setWorkouts(JSON.parse(value));
  }
});
```

### Или используйте Supabase

Для более продвинутой функциональности с синхронизацией между устройствами, используйте Supabase (см. интеграцию ниже).

## Полезные функции

### Уведомления при действиях

Добавьте haptic feedback при нажатиях:

```typescript
import { hapticFeedback } from './utils/telegram';

// При добавлении тренировки
hapticFeedback('success');

// При удалении
hapticFeedback('warning');
```

### Кнопка "Назад" в Telegram

```typescript
import { showBackButton, hideBackButton } from './utils/telegram';

// Показать при просмотре деталей
useEffect(() => {
  if (selectedWorkout) {
    showBackButton(() => {
      setSelectedWorkout(null);
    });
  } else {
    hideBackButton();
  }
}, [selectedWorkout]);
```

## Тестирование

### Локальное тестирование

1. Используйте ngrok для туннеля:
```bash
npx ngrok http 5173
```

2. Используйте HTTPS URL в настройках Mini App

### Проверка темы

Приложение автоматически использует цвета из Telegram темы через CSS переменные:
- `--tg-theme-bg-color`
- `--tg-theme-text-color`
- `--tg-theme-hint-color`
- `--tg-theme-button-color`
- `--tg-theme-button-text-color`

## Дополнительно

### TypeScript типы

Установите типы для Telegram Web App:

```bash
npm install --save-dev @types/telegram-web-app
```

Добавьте в `vite-env.d.ts`:

```typescript
interface Window {
  Telegram?: {
    WebApp: any;
  };
}
```

## Ресурсы

- [Telegram Mini Apps Documentation](https://core.telegram.org/bots/webapps)
- [Telegram Bot API](https://core.telegram.org/bots/api)
- [Examples on GitHub](https://github.com/Telegram-Mini-Apps)
