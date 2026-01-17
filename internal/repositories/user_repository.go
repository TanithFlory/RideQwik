package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rideqwik/api/internal/db"
	"github.com/rideqwik/api/internal/models"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		queries: db.New(pool),
	}
}

func (r *UserRepository) Create(user *models.User) error {
	ctx := context.Background()

	dbUser, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Name:         user.Name,
		Phone:        pgtype.Text{String: user.Phone, Valid: user.Phone != ""},
	})

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	user.ID = int(dbUser.ID)
	user.CreatedAt = dbUser.CreatedAt.Time
	user.UpdatedAt = dbUser.UpdatedAt.Time

	return nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	ctx := context.Background()

	dbUser, err := r.queries.GetUserByEmail(ctx, email)
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &models.User{
		ID:           int(dbUser.ID),
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		Name:         dbUser.Name,
		Phone:        nullStringToGo(&dbUser.Phone.String),
		CreatedAt:    dbUser.CreatedAt.Time,
		UpdatedAt:    dbUser.UpdatedAt.Time,
	}, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	ctx := context.Background()

	dbUser, err := r.queries.GetUserByID(ctx, int32(id))
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &models.User{
		ID:           int(dbUser.ID),
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		Name:         dbUser.Name,
		Phone:        nullStringToGo(&dbUser.Phone.String),
		CreatedAt:    dbUser.CreatedAt.Time,
		UpdatedAt:    dbUser.UpdatedAt.Time,
	}, nil
}

func pgxToNullString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func nullStringToGo(ns *string) string {
	if ns == nil {
		return ""
	}
	return *ns
}
