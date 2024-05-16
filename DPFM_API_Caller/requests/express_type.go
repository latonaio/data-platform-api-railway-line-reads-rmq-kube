package requests

type ExpressType struct {
	RailwayLine			int		`json:"RailwayLine"`
	ExpressType			string	`json:"ExpressType"`
	ValidityStartDate	string	`json:"ValidityStartDate"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	Description			string	`json:"Description"`
	LongText			string	`json:"LongText"`
	Introduction		*string	`json:"Introduction"`
	Project				*int	`json:"Project"`
	WBSElement			*int	`json:"WBSElement"`
	Tag1				*string	`json:"Tag1"`
	Tag2				*string	`json:"Tag2"`
	Tag3				*string	`json:"Tag3"`
	Tag4				*string	`json:"Tag4"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	CreateUser			int		`json:"CreateUser"`
	LastChangeUser		int		`json:"LastChangeUser"`
	IsReleased			*bool	`json:"IsReleased"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
