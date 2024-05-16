package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-railway-line-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-railway-line-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var partner *[]dpfm_api_output_formatter.Partner
	var expressType *[]dpfm_api_output_formatter.expressType

	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "Headers":
			func() {
				header = c.Headers(mtx, input, output, errs, log)
			}()
		case "HeadersByRailwayLines":
			func() {
				header = c.HeadersByRailwayLines(mtx, input, output, errs, log)
			}()
		case "HeadersBySite":
			func() {
				header = c.HeadersBySite(mtx, input, output, errs, log)
			}()
		case "Partner":
			func() {
				partner = c.Partner(mtx, input, output, errs, log)
			}()
		case "Partners":
			func() {
				partner = c.Partners(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "Addresses":
			func() {
				address = c.Addresses(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalSubRegion":
			func() {
				address = c.AddressesByLocalSubRegion(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalSubRegions":
			func() {
				address = c.AddressesByLocalSubRegions(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalRegion":
			func() {
				address = c.AddressesByLocalRegion(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalRegions":
			func() {
				address = c.AddressesByLocalRegions(mtx, input, output, errs, log)
			}()
		case "ExpressType":
			func() {
				expressType = c.ExpressType(mtx, input, output, errs, log)
			}()
		case "ExpressTypes":
			func() {
				expressType = c.ExpressTypes(mtx, input, output, errs, log)
			}()
		case "StopStation":
			func() {
				stopStation = c.StopStation(mtx, input, output, errs, log)
			}()
		case "StopStations":
			func() {
				stopStation = c.StopStations(mtx, input, output, errs, log)
			}()
		case "DepartureStation":
			func() {
				departureStation = c.DepartureStation(mtx, input, output, errs, log)
			}()
		case "DepartureStations":
			func() {
				departureStation = c.DepartureStations(mtx, input, output, errs, log)
			}()
		case "DestinationStation":
			func() {
				destinationStation = c.DestinationStation(mtx, input, output, errs, log)
			}()
		case "DestinationStations":
			func() {
				destinationStation = c.DestinationStations(mtx, input, output, errs, log)
			}()
		case "DptDstStation":
			func() {
				dptDstStation = c.DptDstStation(mtx, input, output, errs, log)
			}()
		case "DptDstStations":
			func() {
				dptDstStation = c.DptDstStations(mtx, input, output, errs, log)
			}()

		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                header,
		Partner:               partner,
		ExpressType:           expressType,
		StopStation:           stopStation,
		DepartureStation:      departureStation,
		DestinationStation:    destinationStation,
		DptDstStation:		   dptDstStation,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.RailwayLine = %d", input.Header.RailwayLine)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.RailwayLine ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Headers(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.RailwayLine ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByRailwayLines(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	log.Info("HeadersByRailwayLines")
	in := ""

	for iHeader, vHeader := range input.Headers {
		railway-line := vHeader.RailwayLine
		if iHeader == 0 {
			in = fmt.Sprintf(
				"( '%d' )",
				railway-line,
			)
			continue
		}
		in = fmt.Sprintf(
			"%s ,( '%d' )",
			in,
			railway-line,
		)
	}

	where := fmt.Sprintf(" WHERE ( RailwayLine ) IN ( %s ) ", in)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.RailwayLine ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByBrand(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.Brand = %d", *input.Header.Brand)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.RailwayLine ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Partner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	var args []interface{}
	railwayLine := input.Header.RailwayLine
	partner := input.Header.Partner

	cnt := 0
	for _, v := range partner {
		args = append(args, railwayLine, v.PartnerFunction, v.BusinessPartner)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_partner_data
		WHERE (RailwayLine, PartnerFunction, BusinessPartner) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, PartnerFunction ASC, BusinessPartner ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPartner(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Partners(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	var args []interface{}
	railwayLine := input.Header.RailwayLine
	partner := input.Header.Partner

	cnt := 0
	for _, _ = range partner {
		args = append(args, railwayLine)
		cnt++
	}
	repeat := strings.Repeat("(?),", cnt-1) + "(?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_partner_data
		WHERE (RailwayLine) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, PartnerFunction ASC, BusinessPartner ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPartner(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ExpressType(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ExpressType {
	var args []interface{}

	railwayLine := input.Header.RailwayLine
	expressType := input.Header.ExpressType

	cnt := 0
	for _, v := range expressType {
		args = append(args, railwayLine, v.ExpressType)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_express_type_data
		WHERE (RailwayLine, ExpressType) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, ExpressType ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToExpressType(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ExpressTypes(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ExpressType {
	var args []interface{}
	railwayLine := input.Header.RailwayLine
	expressType := input.Header.ExpressType

	cnt := 0
	for _, _ = range expressType {
		args = append(args, railwayLine)
		cnt++
	}
	repeat := strings.Repeat("(?),", cnt-1) + "(?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_express_type_data
		WHERE (RailwayLine) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, ExpressType ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToExpressType(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) StopStation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StopStation {
	var args []interface{}

	railwayLine := input.Header.RailwayLine
	expressType := input.Header.ExpressType
	stopStation := input.Header.StopStation

	cnt := 0
	for _, v := range stopStation {
		args = append(args, railwayLine, v.ExpressType, v.StopStation)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_stop_station_data
		WHERE (RailwayLine, ExpressType, StopStation) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, ExpressType ASC, OrderNumberOnOutboundLine ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToStopStation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) StopStations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StopStation {
	var args []interface{}
	railwayLine := input.Header.RailwayLine
	expressType := input.Header.ExpressType
	stopStation := input.Header.StopStation

	cnt := 0
	for _, _ = range stopStation {
		args = append(args, railwayLine, v.ExpressType)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_stop_station_data
		WHERE (RailwayLine, ExpressType) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, ExpressType ASC; OrderNumberOnOutboundLine ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToStopStation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DepartureStation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DepartureStation {
	var args []interface{}

	railwayLine := input.Header.RailwayLine
	expressType := input.Header.ExpressType
	departureStation := input.Header.DepartureStation

	cnt := 0
	for _, v := range departureStation {
		args = append(args, railwayLine, v.ExpressType, v.DepartureStation)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_departure_station_data
		WHERE (RailwayLine, ExpressType, DepartureStation) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, ExpressType ASC, DepartureStation ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDepartureStation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DepartureStations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DepartureStation {
	var args []interface{}
	railwayLine := input.Header.RailwayLine
	expressType := input.Header.ExpressType
	departureStation := input.Header.DepartureStation

	cnt := 0
	for _, _ = range departureStation {
		args = append(args, railwayLine, v.ExpressType)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_departure_station_data
		WHERE (RailwayLine, ExpressType) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, ExpressType ASC; DepartureStation ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDepartureStation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DestinationStation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DestinationStation {
	var args []interface{}

	railwayLine := input.Header.RailwayLine
	expressType := input.Header.ExpressType
	destinationStation := input.Header.DestinationStation

	cnt := 0
	for _, v := range destinationStation {
		args = append(args, railwayLine, v.ExpressType, v.DestinationStation)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_destination_station_data
		WHERE (RailwayLine, ExpressType, DestinationStation) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, ExpressType ASC, DestinationStation ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDestinationStation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DestinationStations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DestinationStation {
	var args []interface{}
	railwayLine := input.Header.RailwayLine
	expressType := input.Header.ExpressType
	destinationStation := input.Header.DestinationStation

	cnt := 0
	for _, _ = range destinationStation {
		args = append(args, railwayLine, v.ExpressType)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_destination_station_data
		WHERE (RailwayLine, ExpressType) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, ExpressType ASC; DestinationStation ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDestinationStation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DptDstStation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DptDstStation {
	var args []interface{}

	railwayLine := input.Header.RailwayLine
	expressType := input.Header.ExpressType
	dptDstStation := input.Header.DptDstStation

	cnt := 0
	for _, v := range dptDstStation {
		args = append(args, railwayLine, v.ExpressType, v.DptDstStation)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_dpt_dst_station_data
		WHERE (RailwayLine, ExpressType, DptDstStation) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, ExpressType ASC, DptDstStation ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDptDstStation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DptDstStations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DptDstStation {
	var args []interface{}
	railwayLine := input.Header.RailwayLine
	expressType := input.Header.ExpressType
	dptDstStation := input.Header.DptDstStation

	cnt := 0
	for _, _ = range dptDstStation {
		args = append(args, railwayLine, v.ExpressType)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_railway_line_dpt_dst_station_data
		WHERE (RailwayLine, ExpressType) IN ( `+repeat+` ) 
		ORDER BY RailwayLine ASC, ExpressType ASC; DptDstStation ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDptDstStation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
