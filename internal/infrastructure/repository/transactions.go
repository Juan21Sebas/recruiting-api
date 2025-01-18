package repository

import (
	"RECRUITING-API/internal/pkg/entity/model"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (p *BDRepository) CreateCandidate(ctx *gin.Context, request *model.CreateCandidate) (int64, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now().UTC()

	query := `
        INSERT INTO candidates (
            name,
            email,
            gender,
            salary_expected,
            phone,
            status,
            experience_years,
            created_at,
            updated_at
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	result, err := p.db.Exec(
		query,
		request.Name,
		request.Email,
		request.Gender,
		request.SalaryExpected,
		request.Phone,
		"PENDING", // Estado inicial por defecto
		request.ExperienceYears,
		now,
		now,
	)
	if err != nil {
		return 0, fmt.Errorf("error creating candidate: %w", err)
	}

	// Obtener el ID generado automáticamente
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting inserted id: %w", err)
	}

	return id, nil
}
