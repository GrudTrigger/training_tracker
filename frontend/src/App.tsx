import {
	BarChart3,
	Calendar,
	CheckCircle,
	Clock,
	Dumbbell,
	Home,
	Plus,
	Settings,
	TrendingUp,
	Users,
} from 'lucide-react'
import { useState } from 'react'

const App = () => {
	const [currentPage, setCurrentPage] = useState('home')
	const [workouts, setWorkouts] = useState([
		{
			id: 1,
			name: 'Грудь и трицепс',
			date: '2025-12-10',
			duration: 60,
			completed: true,
		},
		{
			id: 2,
			name: 'Спина и бицепс',
			date: '2025-12-08',
			duration: 75,
			completed: true,
		},
		{ id: 3, name: 'Ноги', date: '2025-12-05', duration: 90, completed: false },
	])
	const [newWorkout, setNewWorkout] = useState({
		name: '',
		duration: '',
		date: '',
	})

	const handleAddWorkout = () => {
		if (newWorkout.name && newWorkout.duration && newWorkout.date) {
			const workout = {
				id: Date.now(),
				name: newWorkout.name,
				duration: parseInt(newWorkout.duration),
				date: newWorkout.date,
				completed: false,
			}
			setWorkouts([...workouts, workout])
			setNewWorkout({ name: '', duration: '', date: '' })
			setCurrentPage('home')
		}
	}

	const handleCompleteWorkout = id => {
		setWorkouts(
			workouts.map(w => (w.id === id ? { ...w, completed: true } : w))
		)
	}

	const totalWorkouts = workouts.length
	const completedWorkouts = workouts.filter(w => w.completed).length
	const totalDuration = workouts.reduce((sum, w) => sum + w.duration, 0)

	const renderHome = () => (
		<div className='space-y-4'>
			<div className='flex justify-between items-center'>
				<h1 className='text-2xl font-bold'>Мои тренировки</h1>
				<button
					onClick={() => setCurrentPage('add')}
					className='bg-blue-500 text-white p-2 rounded-full hover:bg-blue-600 transition-colors'
				>
					<Plus size={20} />
				</button>
			</div>

			{workouts.length === 0 ? (
				<div className='text-center py-8 text-gray-500'>
					<Dumbbell size={48} className='mx-auto mb-4 text-gray-400' />
					<p>Нет добавленных тренировок</p>
					<p className='text-sm'>Нажмите + чтобы добавить первую тренировку</p>
				</div>
			) : (
				<div className='space-y-3'>
					{workouts.map(workout => (
						<div
							key={workout.id}
							className='bg-white rounded-lg p-4 shadow-sm border border-gray-100'
						>
							<div className='flex justify-between items-start'>
								<div className='flex-1'>
									<h3 className='font-semibold text-lg'>{workout.name}</h3>
									<div className='flex items-center gap-2 mt-2 text-sm text-gray-600'>
										<Calendar size={14} />
										<span>{workout.date}</span>
										<Clock size={14} />
										<span>{workout.duration} мин</span>
									</div>
								</div>
								{workout.completed ? (
									<CheckCircle className='text-green-500' size={24} />
								) : (
									<button
										onClick={() => handleCompleteWorkout(workout.id)}
										className='bg-green-500 text-white px-3 py-1 rounded-full text-sm hover:bg-green-600 transition-colors'
									>
										Выполнено
									</button>
								)}
							</div>
						</div>
					))}
				</div>
			)}
		</div>
	)

	const renderAddWorkout = () => (
		<div className='space-y-6'>
			<div className='flex justify-between items-center'>
				<button
					onClick={() => setCurrentPage('home')}
					className='text-blue-500 font-medium'
				>
					Назад
				</button>
				<h1 className='text-2xl font-bold'>Новая тренировка</h1>
				<div className='w-16'></div>
			</div>

			<div className='space-y-4'>
				<div>
					<label className='block text-sm font-medium text-gray-700 mb-1'>
						Название
					</label>
					<input
						type='text'
						value={newWorkout.name}
						onChange={e =>
							setNewWorkout({ ...newWorkout, name: e.target.value })
						}
						placeholder='Например: Грудь и трицепс'
						className='w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent'
					/>
				</div>

				<div>
					<label className='block text-sm font-medium text-gray-700 mb-1'>
						Продолжительность (мин)
					</label>
					<input
						type='number'
						value={newWorkout.duration}
						onChange={e =>
							setNewWorkout({ ...newWorkout, duration: e.target.value })
						}
						placeholder='60'
						className='w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent'
					/>
				</div>

				<div>
					<label className='block text-sm font-medium text-gray-700 mb-1'>
						Дата
					</label>
					<input
						type='date'
						value={newWorkout.date}
						onChange={e =>
							setNewWorkout({ ...newWorkout, date: e.target.value })
						}
						className='w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent'
					/>
				</div>

				<button
					onClick={handleAddWorkout}
					disabled={
						!newWorkout.name || !newWorkout.duration || !newWorkout.date
					}
					className='w-full bg-blue-500 text-white py-3 rounded-lg font-medium disabled:bg-gray-300 disabled:cursor-not-allowed hover:bg-blue-600 transition-colors'
				>
					Добавить тренировку
				</button>
			</div>
		</div>
	)

	const renderStats = () => (
		<div className='space-y-6'>
			<h1 className='text-2xl font-bold'>Статистика</h1>

			<div className='grid grid-cols-2 gap-4'>
				<div className='bg-blue-50 rounded-lg p-4 text-center'>
					<div className='text-3xl font-bold text-blue-600'>
						{totalWorkouts}
					</div>
					<div className='text-sm text-gray-600'>Всего тренировок</div>
				</div>
				<div className='bg-green-50 rounded-lg p-4 text-center'>
					<div className='text-3xl font-bold text-green-600'>
						{completedWorkouts}
					</div>
					<div className='text-sm text-gray-600'>Выполнено</div>
				</div>
				<div className='bg-purple-50 rounded-lg p-4 text-center'>
					<div className='text-3xl font-bold text-purple-600'>
						{Math.round(totalDuration / 60)}ч
					</div>
					<div className='text-sm text-gray-600'>Общее время</div>
				</div>
				<div className='bg-orange-50 rounded-lg p-4 text-center'>
					<div className='text-3xl font-bold text-orange-600'>
						{totalWorkouts > 0
							? Math.round((completedWorkouts / totalWorkouts) * 100)
							: 0}
						%
					</div>
					<div className='text-sm text-gray-600'>Прогресс</div>
				</div>
			</div>

			<div className='bg-white rounded-lg p-4'>
				<div className='flex items-center gap-2 mb-3'>
					<TrendingUp className='text-blue-500' size={20} />
					<h3 className='font-semibold'>Последние тренировки</h3>
				</div>
				<div className='space-y-2'>
					{workouts.slice(0, 3).map(workout => (
						<div
							key={workout.id}
							className='flex justify-between items-center py-2 border-b border-gray-100 last:border-0'
						>
							<span className='text-sm'>{workout.name}</span>
							<span className='text-sm text-gray-600'>
								{workout.duration} мин
							</span>
						</div>
					))}
				</div>
			</div>
		</div>
	)

	const renderAdmin = () => (
		<div className='space-y-6'>
			<h1 className='text-2xl font-bold'>Админка</h1>

			<div className='bg-white rounded-lg p-4'>
				<div className='flex items-center gap-3 mb-4'>
					<Users className='text-purple-500' size={24} />
					<h3 className='font-semibold'>Управление пользователями</h3>
				</div>
				<div className='space-y-3'>
					<div className='flex justify-between items-center'>
						<span>Активные пользователи</span>
						<span className='font-semibold'>1,247</span>
					</div>
					<div className='flex justify-between items-center'>
						<span>Новые за неделю</span>
						<span className='font-semibold'>+89</span>
					</div>
					<div className='flex justify-between items-center'>
						<span>Средняя активность</span>
						<span className='font-semibold'>4.2 тр/нед</span>
					</div>
				</div>
			</div>

			<div className='bg-white rounded-lg p-4'>
				<h3 className='font-semibold mb-3'>Статистика приложения</h3>
				<div className='space-y-2'>
					<div className='flex justify-between'>
						<span>Версия приложения</span>
						<span>v1.2.3</span>
					</div>
					<div className='flex justify-between'>
						<span>Последнее обновление</span>
						<span>11.12.2025</span>
					</div>
					<div className='flex justify-between'>
						<span>Серверное время</span>
						<span>14:30</span>
					</div>
				</div>
			</div>

			<button className='w-full bg-red-500 text-white py-3 rounded-lg font-medium hover:bg-red-600 transition-colors'>
				Очистить кэш приложения
			</button>
		</div>
	)

	const renderContent = () => {
		switch (currentPage) {
			case 'home':
				return renderHome()
			case 'add':
				return renderAddWorkout()
			case 'stats':
				return renderStats()
			case 'admin':
				return renderAdmin()
			default:
				return renderHome()
		}
	}

	return (
		<div
			className='min-h-screen bg-gray-50 pb-20'
			style={{ height: '100vh', maxHeight: '100vh', overflowY: 'auto' }}
		>
			<div className='max-w-md mx-auto bg-gray-50 p-4 pb-24'>
				{renderContent()}
			</div>

			{/* Bottom Navigation */}
			<div className='fixed bottom-0 left-0 right-0 max-w-md mx-auto bg-white border-t border-gray-200'>
				<div className='flex justify-around items-center py-2'>
					<button
						onClick={() => setCurrentPage('home')}
						className={`flex flex-col items-center p-2 ${
							currentPage === 'home' ? 'text-blue-500' : 'text-gray-500'
						}`}
					>
						<Home size={20} />
						<span className='text-xs mt-1'>Главная</span>
					</button>
					<button
						onClick={() => setCurrentPage('stats')}
						className={`flex flex-col items-center p-2 ${
							currentPage === 'stats' ? 'text-blue-500' : 'text-gray-500'
						}`}
					>
						<BarChart3 size={20} />
						<span className='text-xs mt-1'>Статистика</span>
					</button>
					<button
						onClick={() => setCurrentPage('admin')}
						className={`flex flex-col items-center p-2 ${
							currentPage === 'admin' ? 'text-blue-500' : 'text-gray-500'
						}`}
					>
						<Settings size={20} />
						<span className='text-xs mt-1'>Админка</span>
					</button>
				</div>
			</div>
		</div>
	)
}

export default App
