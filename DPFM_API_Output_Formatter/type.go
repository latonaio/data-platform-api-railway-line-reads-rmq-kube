package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header             	*[]Header				`json:"Header"`
	Partner            	*[]Partner				`json:"Partner"`
	ExpressType        	*[]ExpressType			`json:"ExpressType"`
	StopStation        	*[]StopStation			`json:"StopStation"`
	DepartureStation   	*[]DepartureStation		`json:"DepartureStation"`
	DestinationStation 	*[]DestinationStation	`json:"DestinationStation"`
	DstDptStation	 	*[]DstDptStation		`json:"DstDptStation"`
}

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

type Partner struct {
	Station                 	int     `json:"Station"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}

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

type StopStation struct {
	RailwayLine					int		`json:"RailwayLine"`
	ExpressType					string	`json:"ExpressType"`
	RailwayLineStationID		int		`json:"RailwayLineStationID"`
	StopStation					int		`json:"StopStation"`
	OrderNumberOnOutboundLine	int		`json:"OrderNumberOnOutboundLine"`
	ValidityStartDate			string	`json:"ValidityStartDate"`
	ValidityEndDate				string	`json:"ValidityEndDate"`
	Description					string	`json:"Description"`
	LongText					string	`json:"LongText"`
	Introduction				*string	`json:"Introduction"`
	Project						*int	`json:"Project"`
	WBSElement					*int	`json:"WBSElement"`
	Tag1						*string	`json:"Tag1"`
	Tag2						*string	`json:"Tag2"`
	Tag3						*string	`json:"Tag3"`
	Tag4						*string	`json:"Tag4"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	CreateUser					int		`json:"CreateUser"`
	LastChangeUser				int		`json:"LastChangeUser"`
	IsReleased					*bool	`json:"IsReleased"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}

type DepartureStation struct {
	RailwayLine					int		`json:"RailwayLine"`
	ExpressType					string	`json:"ExpressType"`
	DepartureStation			int		`json:"DepartureStation"`
	RailwayLineStationID		int		`json:"RailwayLineStationID"`
	ValidityStartDate			string	`json:"ValidityStartDate"`
	ValidityEndDate				string	`json:"ValidityEndDate"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	CreateUser					int		`json:"CreateUser"`
	LastChangeUser				int		`json:"LastChangeUser"`
	IsReleased					*bool	`json:"IsReleased"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}

type DestinationStation struct {
	RailwayLine					int		`json:"RailwayLine"`
	ExpressType					string	`json:"ExpressType"`
	DestinationStation			int		`json:"DestinationStation"`
	RailwayLineStationID		int		`json:"RailwayLineStationID"`
	ValidityStartDate			string	`json:"ValidityStartDate"`
	ValidityEndDate				string	`json:"ValidityEndDate"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	CreateUser					int		`json:"CreateUser"`
	LastChangeUser				int		`json:"LastChangeUser"`
	IsReleased					*bool	`json:"IsReleased"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}

type DptDstStation struct {
	RailwayLine					int		`json:"RailwayLine"`
	ExpressType					string	`json:"ExpressType"`
	DepartureStation			int		`json:"DepartureStation"`
	DestinationStation			int		`json:"DestinationStation"`
	ValidityStartDate			string	`json:"ValidityStartDate"`
	ValidityEndDate				string	`json:"ValidityEndDate"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	CreateUser					int		`json:"CreateUser"`
	LastChangeUser				int		`json:"LastChangeUser"`
	IsReleased					*bool	`json:"IsReleased"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
