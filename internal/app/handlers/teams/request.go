package teams

import "avito-internship/internal/app/dto"

type AddRequest struct {
	Team
}

func convertTeamAddRequest(req *AddRequest) *dto.Team {
	if req == nil {
		return nil
	}
	members := make([]*dto.User, 0, len(req.Members))
	for _, member := range req.Members {
		members = append(members, &dto.User{
			ID:       member.UserID,
			TeamName: req.TeamName,
			UserName: member.Username,
			IsActive: member.IsActive,
		})
	}
	return &dto.Team{
		Name:    req.TeamName,
		Members: members,
	}
}
