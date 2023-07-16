package pool

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type JobDao struct {
	db *gorm.DB
}

func NewJobDao() *JobDao {
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/lyndon?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	return &JobDao{
		db: db,
	}
}

func (dao *JobDao) CreateJob(ctx context.Context, job *Job) {
	dao.db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Create(&job).Error; e != nil {
			return e
		}
		return nil
	})
}

func (dao *JobDao) SelectJob(ctx context.Context, id int64) *Job {
	var job *Job
	dao.db.First(&job, id)
	return job
}
