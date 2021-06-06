package mysql

const (
	sqlCourseTable = "courses"
)

// sqlCourse is the port between course model and mysql adapter. This struct drives course data from course model to mysql repository adapter, which is course_repository.go.
type sqlCourse struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Duration string `db:"duration"`
	Price    string `db:"price"`
}
