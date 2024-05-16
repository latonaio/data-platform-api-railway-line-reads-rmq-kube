package requests

type Header struct {
	RailwayLine							int		`json:"RailwayLine"`
	RailwayLineType						string	`json:"RailwayLineType"`
	RailwayLineOwner					int		`json:"RailwayLineOwner"`
	RailwayLineOwnerBusinessPartnerRole	string	`json:"RailwayLineOwnerBusinessPartnerRole"`
	Brand								*int	`json:"Brand"`
	PersonResponsible					*string	`json:"PersonResponsible"`
	URL									*string	`json:"URL"`
	ValidityStartDate					string	`json:"ValidityStartDate"`
	ValidityEndDate						string	`json:"ValidityEndDate"`
	DepartureStationOfOutboundLine		int		`json:"DepartureStationOfOutboundLine"`
	DestinationStationOfOutboundLine	int		`json:"DestinationStationOfOutboundLine"`
	Description							string	`json:"Description"`
	LongText							string	`json:"LongText"`
	Introduction						*string	`json:"Introduction"`
	Project								*int	`json:"Project"`
	WBSElement							*int	`json:"WBSElement"`
	Tag1								*string	`json:"Tag1"`
	Tag2								*string	`json:"Tag2"`
	Tag3								*string	`json:"Tag3"`
	Tag4								*string	`json:"Tag4"`
	CreationDate						string	`json:"CreationDate"`
	CreationTime						string	`json:"CreationTime"`
	LastChangeDate						string	`json:"LastChangeDate"`
	LastChangeTime						string	`json:"LastChangeTime"`
	CreateUser							int		`json:"CreateUser"`
	LastChangeUser						int		`json:"LastChangeUser"`
	IsReleased							*bool	`json:"IsReleased"`
	IsMarkedForDeletion					*bool	`json:"IsMarkedForDeletion"`
}
