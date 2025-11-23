package teams

import "avito-internship/internal/app/dto"

type AddResponse struct {
	Team *Team `json:"team"`
}

func convertAddResponse(team *dto.Team) *AddResponse {
	if team == nil {
		return nil
	}
	return &AddResponse{
		Team: convertTeam(team),
	}
}
