package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

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
