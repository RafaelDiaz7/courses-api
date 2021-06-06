package couchbase

import (
	mooc "courses-api-mysql-and-cb/internal"
	"fmt"
	"github.com/couchbase/gocb/v2"
	"log"
	"time"
)

type CbCourseRepository struct {
	cluster *gocb.Cluster
}

func NewCbRepository(cluster *gocb.Cluster) *CbCourseRepository {
	return &CbCourseRepository{
		cluster: cluster,
	}
}

//aca posiblemente solo necesite pasar la struct del request porque ya tiene los json bindings
func (r *CbCourseRepository) Save(course mooc.Course) error {

	bucket := r.cluster.Bucket(cbCourseBucket)
	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError)
		return fmt.Errorf("error verifying if bucket is ready: %v", err)
	}

	collection := bucket.DefaultCollection()

	couchbaseCourse := cbCourse{
		ID: course.ID(),
		Name: course.Name(),
		Duration: course.Duration(),
		Price: course.Price(),
	}

	//the course data needs to be passed to cbCourse struct that have json bindings
	//uuid, err := gocb.UUIDIdGeneratorFunction(couchbaseCourse)
	//if err != nil {
	//	return fmt.Errorf("error trying on call to Id generator: %v", err)
	//}

	result, err := collection.Insert(couchbaseCourse.ID, &couchbaseCourse, nil)
	if err != nil {
                return fmt.Errorf("error trying to persist course on database: %v", err)
        }
	log.Println(result.Cas())
	return nil
}
