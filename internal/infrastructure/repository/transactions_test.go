package repository

import (
	"errors"
	"net/http/httptest"
	"testing"

	"RECRUITING-API/internal/pkg/entity/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateCandidate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()

	repo := &BDRepository{
		db: db,
	}

	t.Run("success case", func(t *testing.T) {
		request := &model.CreateCandidate{
			Name:            "John Doe",
			Email:           "john@example.com",
			Gender:          "Male",
			SalaryExpected:  50000.00,
			Phone:           "+1234567890",
			ExperienceYears: 5,
		}

		mock.ExpectExec("INSERT INTO candidates").
			WithArgs(
				request.Name,
				request.Email,
				request.Gender,
				request.SalaryExpected,
				request.Phone,
				"PENDING",
				request.ExperienceYears,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnResult(sqlmock.NewResult(123, 1))

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		id, err := repo.CreateCandidate(c, request)

		assert.NoError(t, err)
		assert.Equal(t, int64(123), id)
	})

	t.Run("error case - db error", func(t *testing.T) {
		request := &model.CreateCandidate{
			Name: "John Doe",
			// ... otros campos
		}

		mock.ExpectExec("INSERT INTO candidates").
			WillReturnError(errors.New("db error"))

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		_, err := repo.CreateCandidate(c, request)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error creating candidate")
	})
}

func TestUpdateCandidate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()

	repo := &BDRepository{
		db: db,
	}

	t.Run("success case", func(t *testing.T) {
		request := &model.UpdateCandidate{
			Id:              "123",
			Name:            "John Updated",
			Email:           "john.updated@example.com",
			Gender:          "Male",
			SalaryExpected:  60000.00,
			Phone:           "+1234567890",
			ExperienceYears: 6,
		}

		mock.ExpectExec("UPDATE candidates").
			WithArgs(
				request.Name,
				request.Email,
				request.Gender,
				request.SalaryExpected,
				request.Phone,
				request.ExperienceYears,
				sqlmock.AnyArg(),
				request.Id,
			).
			WillReturnResult(sqlmock.NewResult(0, 1))

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		result, err := repo.UpdateCandidate(c, request)

		assert.NoError(t, err)
		assert.Equal(t, request.Name, result.Name)
		assert.Equal(t, request.Email, result.Email)
	})

	t.Run("not found case", func(t *testing.T) {
		request := &model.UpdateCandidate{Id: "123"}

		mock.ExpectExec("UPDATE candidates").
			WillReturnResult(sqlmock.NewResult(0, 0))

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		_, err := repo.UpdateCandidate(c, request)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "candidate not found")
	})
}

func TestDeleteCandidate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()

	repo := &BDRepository{
		db: db,
	}

	t.Run("success case", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM candidates").
			WithArgs("123").
			WillReturnResult(sqlmock.NewResult(0, 1))

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		err := repo.DeleteCandidate(c, &model.DeleteCandidate{Id: "123"})

		assert.NoError(t, err)
	})

	t.Run("not found case", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM candidates").
			WithArgs("123").
			WillReturnResult(sqlmock.NewResult(0, 0))

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		err := repo.DeleteCandidate(c, &model.DeleteCandidate{Id: "123"})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "candidate not found")
	})
}
