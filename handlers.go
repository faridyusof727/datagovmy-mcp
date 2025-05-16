package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"encoding/csv"
	"github.com/mark3labs/mcp-go/mcp"
)

const BaseURL = "https://api.data.gov.my/data-catalogue"

func populationMalaysiaHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	url, err := url.Parse(BaseURL)
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("id", "population_malaysia")

	filters := make([]string, 0)
	age, ok := request.Params.Arguments["age"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@age", age))
	}

	sex, ok := request.Params.Arguments["sex"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@sex", sex))
	}

	ethnicity, ok := request.Params.Arguments["ethnicity"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@ethnicity", ethnicity))
	}

	date, ok := request.Params.Arguments["date"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@date", date))

	}

	if len(filters) > 0 {
		q.Add("filter", strings.Join(filters, ","))
	}

	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return mcp.NewToolResultText(string(body)), nil
}

func populationStateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	url, err := url.Parse(BaseURL)
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("id", "population_state")

	filters := make([]string, 0)
	age, ok := request.Params.Arguments["age"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@age", age))
	}

	sex, ok := request.Params.Arguments["sex"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@sex", sex))
	}

	ethnicity, ok := request.Params.Arguments["ethnicity"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@ethnicity", ethnicity))
	}

	date, ok := request.Params.Arguments["date"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@date", date))

	}

	state, ok := request.Params.Arguments["state"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@state", state))
	}

	if len(filters) > 0 {
		q.Add("filter", strings.Join(filters, ","))
	}

	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return mcp.NewToolResultText(string(body)), nil
}

func birthHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	url, err := url.Parse(BaseURL)
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("id", "births")

	filters := make([]string, 0)
	date, ok := request.Params.Arguments["date"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@date", date))
	}

	births, ok := request.Params.Arguments["births"].(int64)
	if ok {
		filters = append(filters, fmt.Sprintf("%d@births", births))
	}

	if len(filters) > 0 {
		q.Add("filter", strings.Join(filters, ","))
	}

	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return mcp.NewToolResultText(string(body)), nil
}

func fuelpriceHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	url, err := url.Parse(BaseURL)
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("id", "fuelprice")

	filters := make([]string, 0)
	date, ok := request.Params.Arguments["date"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@date", date))
	}

	seriesType, ok := request.Params.Arguments["series_type"].(string)
	if ok {
		filters = append(filters, fmt.Sprintf("%s@series_type", seriesType))
	}

	if len(filters) > 0 {
		q.Add("filter", strings.Join(filters, ","))
	}

	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return mcp.NewToolResultText(string(body)), nil
}

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
