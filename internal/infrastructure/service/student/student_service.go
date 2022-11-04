package service

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	. "hexrestapi1/internal/infrastructure/domain/student"
	. "hexrestapi1/internal/infrastructure/port/student"
)

type StudentService interface {
	GetAllStudents(ctx context.Context) (*[]Student, error)
	GetStudent(ctx context.Context, id string) (*Student, error)
	CreateStudent(ctx context.Context, student *Student) (int64, error)
	UpdateStudent(ctx context.Context, student *Student) (int64, error)
	DeleteStudent(ctx context.Context, id string) (int64, error)
}

type studentService struct {
	Collection         *mongo.Collection
	Repository StudentRepository
}

func NewStudentService(db *mongo.Database, repos StudentRepository) StudentService {
	collectionName := "student"
	return &studentService{Collection: db.Collection(collectionName), Repository: repos}
}

// GetAllStudent implements StudentService
func (s *studentService) GetAllStudents(ctx context.Context) (*[]Student, error) {
	return s.Repository.GetAllStudents(ctx)
}

// GetStudent implements StudentService
func (s *studentService) GetStudent(ctx context.Context, id string) (*Student, error) {
	return s.Repository.GetStudent(ctx, id)
}

// InsertStudent implements StudentService
func (s *studentService) CreateStudent(ctx context.Context, student *Student) (int64, error) {
	return s.Repository.CreateStudent(ctx, student)
}

// UpdateStudent implements StudentService
func (s *studentService) UpdateStudent(ctx context.Context, student *Student) (int64, error) {
	return s.Repository.UpdateStudent(ctx, student)
}

// DeleteStudent implements StudentService
func (s *studentService) DeleteStudent(ctx context.Context, id string) (int64, error) {
	return s.Repository.DeleteStudent(ctx, id)
}



