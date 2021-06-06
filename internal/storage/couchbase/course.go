package couchbase

const (
	cbCourseBucket = "courses"
)

// sqlCourse is the port between course model and mysql adapter. This struct drives course data from course model to mysql repository adapter, which is course_repository.go.
type cbCourse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"`
	Price    string `json:"price"`
}
