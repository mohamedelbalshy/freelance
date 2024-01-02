package model

import (
	"freelance/database"

	"gorm.io/gorm"
)

type ProposalStatus string

const (
	ACTIVE    ProposalStatus = "ACTIVE"
	SUBMITTED ProposalStatus = "SUBMITTED"
)

type Proposal struct {
	gorm.Model
	ID          uint64         `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL; column:id" `
	CreatedById int64          `gorm:"NOT NULL;" json:"createdById"`
	CreatedBy   User           `gorm:"foreignKey:CreatedById;references:ID" json:"createdBy"`
	ProjectId   int64          `gorm:"NOT NULL;" json:"projectId"`
	Project     Project        `gorm:"foreignKey:ProjectId;references:ID" json:"project"`
	Status      ProposalStatus `gorm:"column:proposal_status" json:"proposalStatus"`
	CoverLetter string         `gorm:"size:255;NOT NULL;column:cover_letter" json:"coverLetter"`
	HorlyRate   float32        `gorm:"NOT NULL;column:hourly_rate" json:"hourlyRate"`
}

func (proposal *Proposal) Save() (*Proposal, error) {
	err := database.Database.Create(&proposal).Error
	if err != nil {
		return &Proposal{}, err
	}
	return proposal, nil
}
