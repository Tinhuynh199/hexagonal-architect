package port

import (
	"context"

	. "hexrestapi1/internal/infrastructure/domain/student"
)


// Driven Actor -- Core -> MongoDB
type StudentRepository interface {
	GetAllStudents(ctx context.Context) (*[]Student, error)
	GetStudent(ctx context.Context, id string) (*Student, error)
	CreateStudent(ctx context.Context, student *Student) (int64, error)
	UpdateStudent(ctx context.Context, student *Student) (int64, error)
	DeleteStudent(ctx context.Context, id string) (int64, error)
}