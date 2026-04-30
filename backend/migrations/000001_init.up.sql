create table if not exists exercises (
    id uuid primary key,
    name text not null unique,
    muscle_group text not null,
    created_at timestamptz not null default now()
);

create table if not exists workouts (
    id uuid primary key,
    performed_at timestamptz not null,
    note text not null default ''
);

create table if not exists workout_sets (
    id uuid primary key,
    workout_id uuid not null references workouts(id) on delete cascade,
    exercise_id uuid not null references exercises(id),
    reps integer not null,
    weight_kg numeric(10,2) not null,
    set_order integer not null
);

create index if not exists workouts_performed_at_idx on workouts (performed_at desc);
create index if not exists workout_sets_workout_id_idx on workout_sets (workout_id);
create index if not exists workout_sets_exercise_id_idx on workout_sets (exercise_id);
