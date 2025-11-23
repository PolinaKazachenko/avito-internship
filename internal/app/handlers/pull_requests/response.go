package pull_requests

import "avito-internship/internal/app/dto"

type CreateResponse struct {
	PR *PullRequest `json:"pr"`
}

type ReassignResponse struct {
	PR         *PullRequest `json:"pr"`
	ReplacedBy string       `json:"replaced_by"`
}

type MergeResponse struct {
	PR *FullPullRequest `json:"pr"`
}

func convertAddResponse(pr *dto.PullRequest) *CreateResponse {
	if pr == nil {
		return nil
	}
	return &CreateResponse{
		PR: convertPullRequest(pr),
	}
}

func convertMergeResponse(pr *dto.PullRequest) *MergeResponse {
	if pr == nil {
		return nil
	}
	return &MergeResponse{
		PR: convertFullPullRequest(pr),
	}
}

func convertReassignResponse(newReviewer string, pr *dto.PullRequest) *ReassignResponse {
	if pr == nil {
		return nil
	}
	return &ReassignResponse{
		PR:         convertPullRequest(pr),
		ReplacedBy: newReviewer,
	}
}
