package model

import (
	"freelance/database"
	"log"

	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	ID   uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL; column:id" `
	Name string `gorm:"size:255;NOT NULL;column:name" json:"name"`
	Slug string `gorm:"size:255;NOT NULL;column:slug" json:"slug"`
}

type UpdateSkill struct {
	Id   uint64 `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}

func (skill *Skill) Save() (*Skill, error) {
	err := database.Database.Create(&skill).Error
	if err != nil {
		return &Skill{}, err
	}
	return skill, nil
}

func GetAllSkills() []Skill {
	var skills []Skill
	result := database.Database.Find(&skills)
	if result.Error != nil {
		log.Fatal("Error Getting skills")
	}
	return skills
}

func FindSkillById(id uint64) (*Skill, error) {
	var skill Skill
	err := database.Database.Where("ID=?", id).Find(&skill).Error
	if err != nil {
		return &Skill{}, err
	}
	return &skill, nil
}

func UpdateSkillById(updateSkill UpdateSkill) (*Skill, error) {

	skill, err := FindSkillById(updateSkill.Id)
	if err != nil {
		log.Fatal("Skill With id", updateSkill.Id, " not found")
	}

	saveErr := database.Database.Save(&updateSkill).Error
	if saveErr != nil {
		return &Skill{}, err
	}
	return skill, nil
}
