package dpfm_api_output_formatter

import (
	"data-platform-api-railway-line-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.RailwayLine,
			&pm.RailwayLineType,
			&pm.RailwayLineOwner,
			&pm.RailwayLineOwnerBusinessPartnerRole,
			&pm.Brand,
			&pm.PersonResponsible,
			&pm.URL,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.DepartureStationOfOutboundLine,
			&pm.DestinationStationOfOutboundLine,
			&pm.Description,
			&pm.LongText,
			&pm.Introduction,
			&pm.Project,
			&pm.WBSElement,
			&pm.Tag1,
			&pm.Tag2,
			&pm.Tag3,
			&pm.Tag4,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.CreateUser,
			&pm.LastChangeUser,
			&pm.IsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			RailwayLine:							data.RailwayLine,
			RailwayLineType:						data.RailwayLineType,
			RailwayLineOwner:						data.RailwayLineOwner,
			RailwayLineOwnerBusinessPartnerRole:	data.RailwayLineOwnerBusinessPartnerRole,
			Brand:									data.Brand,
			PersonResponsible:						data.PersonResponsible,
			URL:									data.URL,
			ValidityStartDate:						data.ValidityStartDate,
			ValidityEndDate:						data.ValidityEndDate,
			DepartureStationOfOutboundLine:			data.DepartureStationOfOutboundLine,
			DestinationStationOfOutboundLine:		data.DestinationStationOfOutboundLine,
			Description:							data.Description,
			LongText:								data.LongText,
			Introduction:							data.Introduction,
			Project:								data.Project,
			WBSElement:								data.WBSElement,
			Tag1:									data.Tag1,
			Tag2:									data.Tag2,
			Tag3:									data.Tag3,
			Tag4:									data.Tag4,
			CreationDate:							data.CreationDate,
			CreationTime:							data.CreationTime,
			LastChangeDate:							data.LastChangeDate,
			LastChangeTime:							data.LastChangeTime,
			CreateUser:								data.CreateUser,
			LastChangeUser:							data.LastChangeUser,
			IsReleased:								data.IsReleased,
			IsMarkedForDeletion:					data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToPartner(rows *sql.Rows) (*[]Partner, error) {
	defer rows.Close()
	partner := make([]Partner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Partner{}

		err := rows.Scan(
			&pm.RailwayLine,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
			&pm.EmailAddress,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &partner, err
		}

		data := pm
		partner = append(partner, Partner{
			RailwayLine:                   data.RailwayLine,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
			EmailAddress:            data.EmailAddress,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &partner, nil
	}

	return &partner, nil
}

func ConvertToExpressType(rows *sql.Rows) (*[]ExpressType, error) {
	defer rows.Close()
	expressType := make([]ExpressType, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ExpressType{}

		err := rows.Scan(
			&pm.ExpressType,
			&pm.ExpressType,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.Description,
			&pm.LongText,
			&pm.Introduction,
			&pm.Project,
			&pm.WBSElement,
			&pm.Tag1,
			&pm.Tag2,
			&pm.Tag3,
			&pm.Tag4,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.CreateUser,
			&pm.LastChangeUser,
			&pm.IsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &expressType, err
		}

		data := pm
		expressType = append(expressType, ExpressType{
			ExpressType:			data.ExpressType,
			ExpressType:			data.ExpressType,
			ValidityStartDate:		data.ValidityStartDate,
			ValidityEndDate:		data.ValidityEndDate,
			Description:			data.Description,
			LongText:				data.LongText,
			Introduction:			data.Introduction,
			Project:				data.Project,
			WBSElement:				data.WBSElement,
			Tag1:					data.Tag1,
			Tag2:					data.Tag2,
			Tag3:					data.Tag3,
			Tag4:					data.Tag4,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
			CreateUser:				data.CreateUser,
			LastChangeUser:			data.LastChangeUser,
			IsReleased:				data.IsReleased,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &expressType, nil
	}

	return &expressType, nil
}

func ConvertToDepartureStation(rows *sql.Rows) (*[]DepartureStation, error) {
	defer rows.Close()
	departureStation := make([]DepartureStation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DepartureStation{}

		err := rows.Scan(
			&pm.ExpressType,
			&pm.DepartureStation,
			&pm.RailwayLineStationID,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.CreateUser,
			&pm.LastChangeUser,
			&pm.IsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &departureStation, err
		}

		data := pm
		departureStation = append(departureStation, DepartureStation{
			ExpressType:			data.ExpressType,
			DepartureStation:		data.DepartureStation,
			RailwayLineStationID:	data.RailwayLineStationID,
			ValidityStartDate:		data.ValidityStartDate,
			ValidityEndDate:		data.ValidityEndDate,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
			CreateUser:				data.CreateUser,
			LastChangeUser:			data.LastChangeUser,
			IsReleased:				data.IsReleased,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &departureStation, nil
	}

	return &departureStation, nil
}

func ConvertToDestinationStation(rows *sql.Rows) (*[]DestinationStation, error) {
	defer rows.Close()
	destinationStation := make([]DestinationStation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DestinationStation{}

		err := rows.Scan(
			&pm.ExpressType,
			&pm.DestinationStation,
			&pm.RailwayLineStationID,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.CreateUser,
			&pm.LastChangeUser,
			&pm.IsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &destinationStation, err
		}

		data := pm
		destinationStation = append(destinationStation, DestinationStation{
			ExpressType:			data.ExpressType,
			DestinationStation:		data.DestinationStation,
			RailwayLineStationID:	data.RailwayLineStationID,
			ValidityStartDate:		data.ValidityStartDate,
			ValidityEndDate:		data.ValidityEndDate,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
			CreateUser:				data.CreateUser,
			LastChangeUser:			data.LastChangeUser,
			IsReleased:				data.IsReleased,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &destinationStation, nil
	}

	return &destinationStation, nil
}

func ConvertToDptDstStation(rows *sql.Rows) (*[]DptDstStation, error) {
	defer rows.Close()
	dptDstStation := make([]DptDstStation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DptDstStation{}

		err := rows.Scan(
			&pm.ExpressType,
			&pm.DepartureStation,
			&pm.DestinationStation,
			&pm.RailwayLineStationID,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.CreateUser,
			&pm.LastChangeUser,
			&pm.IsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &dptDstStation, err
		}

		data := pm
		dptDstStation = append(dptDstStation, DptDstStation{
			ExpressType:			data.ExpressType,
			DepartureStation:		data.DepartureStation,
			DestinationStation:		data.DestinationStation,
			ValidityStartDate:		data.ValidityStartDate,
			ValidityEndDate:		data.ValidityEndDate,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
			CreateUser:				data.CreateUser,
			LastChangeUser:			data.LastChangeUser,
			IsReleased:				data.IsReleased,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &dptDstStation, nil
	}

	return &dptDstStation, nil
}
