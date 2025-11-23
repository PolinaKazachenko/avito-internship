package teams

import "avito-internship/internal/app/dto"

type Member struct {
	UserID   string `json:"user_id" binding:"required,min=1"`
	Username string `json:"username" binding:"required,min=1"`
	IsActive bool   `json:"is_active"`
}

type Team struct {
	TeamName string   `json:"team_name" binding:"required,min=1"`
	Members  []Member `json:"members" binding:"required,min=2"`
}

func convertTeam(team *dto.Team) *Team {
	if team == nil {
		return nil
	}
	members := make([]Member, 0, len(team.Members))
	for _, m := range team.Members {
		members = append(members, Member{
			UserID:   m.ID,
			Username: m.UserName,
			IsActive: m.IsActive,
		})
	}
	return &Team{
		TeamName: team.Name,
		Members:  members,
	}
}
