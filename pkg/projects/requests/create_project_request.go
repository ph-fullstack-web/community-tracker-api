package projects

type CreateProjectRequest struct {
	ProjectName string `validate:"required" json:"project"`
	ProjectLead int    ` validate:"required" json:"lead"`
}
