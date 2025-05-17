package tools

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
)

var registrationTransactionsCar = mcp.NewTool("registration_transactions_car",
	mcp.WithDescription("Car registration transactions from 2000 to the present"),
	mcp.WithString("date_reg", mcp.Description("The date of registration of the vehicle in YYYY-MM-DD format; please note that this date may not be the same as the date of purchase")),
	mcp.WithString("date_start", mcp.Description("For date range and the start date of registration of the vehicle in YYYY-MM-DD format; please note that this date may not be the same as the date of purchase")),
	mcp.WithString("date_end", mcp.Description("For date range and the end date of registration of the vehicle in YYYY-MM-DD format; please note that this date may not be the same as the date of purchase")),
	mcp.WithString("date_reg", mcp.Description("The date of registration of the vehicle in YYYY-MM-DD format; please note that this date may not be the same as the date of purchase")),
	mcp.WithString("type", mcp.Description("One of 5 vehicle types classed under as Cars for the purpose of analysis, namely motorcars ('motokar'), MPVs ('motokar_pelbagai_utiliti'), jeeps ('jip'), pick-up trucks ('pick_up') and window vans ('window_van')")),
	mcp.WithString("maker", mcp.Description("Maker of the vehicle (e.g. Perodua, Proton, Toyota) in upper-case text")),
	mcp.WithString("model", mcp.Description("Model of the vehicle (e.g. Bezza (Perodua), Saga (Proton), City (Honda)) in upper-case text")),
	mcp.WithString("colour", mcp.Description("Colour of the car in lower-case English text; please note these are broadly defined colours which do not distinguish between shades of the same colour (e.g. light blue and dark blue are both classed as blue)")),
	mcp.WithString("fuel", mcp.Description("Fuel used by the car's engine(s); there are 7 types, namely petrol ('petrol'), diesel ('diesel'), green diesel ('greendiesel'), natural gas ('ng'), liquefied natural gas ('lng'), hydrogen ('hydrogen'), and electricity ('electric'). Cars which can run on electricity or fuel are classed as hybrid (either 'hybrid_petrol' or 'hybrid_diesel'). Combinations of two fuels indicate that the car's engine can use more than one type of fuel; for instance, 'diesel_ng' means the car's engine can use both diesel and natural gas as fuel.")),
	mcp.WithString("state", mcp.Description("One of 16 states, or 'Rakan Niaga'; this either indicates the state of the JPJ office the car was registered at, or that the car was registered through an official JPJ partner portal ('Rakan Niaga'). Please note that this data field has no relation to the car's number plate, which may be freely chosen with no dependence on where the car was registered.")),
)

func registrationTransactionsCarHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	const csvURL = "https://storage.data.gov.my/transportation/cars_2025.csv"

	// Download the CSV file
	resp, err := http.Get(csvURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download CSV file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code when downloading CSV: %d", resp.StatusCode)
	}

	// Parse the CSV in a streaming fashion (row by row)
	reader := csv.NewReader(resp.Body)
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %v", err)
	}
	colIdx := make(map[string]int)
	for i, col := range header {
		colIdx[col] = i
	}

	// Get filter params
	args := request.Params.Arguments
	getStr := func(key string) string {
		if v, ok := args[key].(string); ok {
			return v
		}
		return ""
	}
	dateReg := getStr("date_reg")
	dateStart := getStr("date_start")
	dateEnd := getStr("date_end")
	typ := getStr("type")
	maker := getStr("maker")
	model := getStr("model")
	colour := getStr("colour")
	fuel := getStr("fuel")
	state := getStr("state")

	// Helper for date range
	inDateRange := func(date string) bool {
		if dateStart != "" && date < dateStart {
			return false
		}
		if dateEnd != "" && date > dateEnd {
			return false
		}
		return true
	}

	// Stream and filter rows
	var filtered [][]string
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read CSV row: %v", err)
		}
		if dateReg != "" && row[colIdx["date_reg"]] != dateReg {
			continue
		}
		if typ != "" && row[colIdx["type"]] != typ {
			continue
		}
		if maker != "" && row[colIdx["maker"]] != maker {
			continue
		}
		if model != "" && row[colIdx["model"]] != model {
			continue
		}
		if colour != "" && row[colIdx["colour"]] != colour {
			continue
		}
		if fuel != "" && row[colIdx["fuel"]] != fuel {
			continue
		}
		if state != "" && row[colIdx["state"]] != state {
			continue
		}
		if (dateStart != "" || dateEnd != "") && !inDateRange(row[colIdx["date_reg"]]) {
			continue
		}
		filtered = append(filtered, row)
	}

	// Prepend header if any results
	var result [][]string
	if len(filtered) > 0 {
		result = append(result, header)
		result = append(result, filtered...)
	} else {
		result = [][]string{header}
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal rows to JSON: %v", err)
	}

	return mcp.NewToolResultText(string(jsonBytes)), nil
}
