package model

import (
	"freelance/database"
	"log"

	"gorm.io/gorm"
)

type ProjectType string

const (
	PUBLIC  ProjectType = "PUBLIC"
	PRIVATE ProjectType = "PRIVATE"
)

type Project struct {
	gorm.Model
	ID          int64       `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL; column:id" `
	Title       string      `gorm:"size:255;not null;column:title" json:"title"`
	Description string      `gorm:"size:255;not null;column:description" json:"description"`
	Cost        int64       `gorm:"NOT NULL;column:cost" json:"cost"`
	Duration    int32       `gorm:"column:duration" json:"duration"`
	ProjectType ProjectType `gorm:"column:project_type" json:"projectType"`
	OwnerId     int64       `gorm:"NOT NULL;" json:"ownerId"`
	Owner       User        `gorm:"foreignKey:OwnerId;references:ID" json:"owner"`
}

func (project *Project) Save() (*Project, error) {
	err := database.Database.Create(&project).Error
	if err != nil {
		return &Project{}, err
	}
	return project, nil
}

func GetAllProjects() []Project {
	var projects []Project
	result := database.Database.Preload("Owner").Find(&projects)
	if result.Error != nil {
		log.Fatal("Error Getting projects")
	}
	return projects
}

func FindProjectById(id uint) (Project, error) {
	var project Project
	err := database.Database.Where("ID=?", id).Find(&project).Error
	if err != nil {
		return Project{}, err
	}
	return project, nil
}
