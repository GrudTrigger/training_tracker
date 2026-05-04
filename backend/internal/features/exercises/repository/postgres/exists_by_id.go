package exercises_postgres_repository

import "context"

func (r *Repository) ExistsByID(ctx context.Context, id string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(
		ctx,
		`select exists(select 1 from exercises where id = $1)`,
		id,
	).Scan(&exists)

	return exists, err
}
