package mooc

// Course is the data structure that represents a course. Its part of the business logic/models/domain.
type Course struct {
	id       string
	name     string
	duration string
	price    string
}

// NewCourse creates a new course.
func NewCourse(id, name, duration, price string) Course {
	return Course{
		id:       id,
		name:     name,
		duration: duration,
		price:    price,
	}
}

// ID returns the course unique identifier.
func (c Course) ID() string {
	return c.id
}

// Name returns the course name.
func (c Course) Name() string {
	return c.name
}

// Duration returns the course duration.
func (c Course) Duration() string {
	return c.duration
}

// Price returns the course price.
func (c Course) Price() string {
	return c.price
}
