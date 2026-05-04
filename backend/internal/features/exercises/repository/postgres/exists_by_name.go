package exercises_postgres_repository

import "context"

func (r *Repository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(
		ctx,
		`select exists(select 1 from exercises where lower(name) = lower($1))`,
		name,
	).Scan(&exists)

	return exists, err
}
