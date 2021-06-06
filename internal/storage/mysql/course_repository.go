package mysql

import (
	"context"
	mooc "courses-api-mysql-and-cb/internal"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

// course_repository.go have the code for mysql repository adapter.

// CourseRepository is a MySQL mooc.CourseRepository implementation.
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.CourseRepository.
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Save implements the mooc.CourseRepository interface.
func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID(),
		Name:     course.Name(),
		Duration: course.Duration(),
		Price:    course.Price(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

// ScanTableRows implements the mooc.CourseRepository interface to scan table rows and return a slice of them.
func (r *CourseRepository) ScanTableRows() ([]sqlCourse,error) {
	queryString := "SELECT * FROM "+sqlCourseTable
	results, err := r.db.Query(queryString)
	if err != nil {
		fmt.Errorf("error trying to fetching table rows: %v",err)
	}
	defer results.Close()
	var resultRows = make([]sqlCourse, 0)
	for results.Next() {
		var resultRow = sqlCourse{}
		err = results.Scan(&resultRow.ID, &resultRow.Name, &resultRow.Duration, &resultRow.Price)
		if err != nil {
			fmt.Errorf("error trying to scan table row to struct: %v",err)
		}
		resultRows = append(resultRows, resultRow)
	}
	return resultRows,nil
}
