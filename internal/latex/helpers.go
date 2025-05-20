package latex

import (
	"ScanEvalApp/internal/database/models"
	"ScanEvalApp/internal/logging"

	"log/slog"

	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
	"gorm.io/gorm"
)

func removeDiacritics(input string) string {
	t := norm.NFD.String(input)
	t = strings.Map(func(r rune) rune {
		if unicode.IsMark(r) {
			return -1
		}
		return r
	}, t)
	t = strings.ReplaceAll(t, " ", "_")
	return t
}

func FindStudentByRegistrationNumber(db *gorm.DB, registrationNumber int) (*models.Student, error) {
	errorLogger := logging.GetErrorLogger()
	var student models.Student
	if err := db.Where("registration_number = ?", registrationNumber).First(&student).Error; err != nil {
		errorLogger.Error("Student not found with ", slog.Uint64("registration_number", uint64(registrationNumber)), slog.String("error", err.Error()))
		return nil, err
	}
	return &student, nil
}
