package tools

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

var populationMalaysia = mcp.NewTool("population_malaysia",
	mcp.WithDescription("Population at national level from 1970 to 2024, by sex, age group and ethnicity"),
	mcp.WithString("age",
		mcp.Description("Either all age groups ('overall') or five-year age groups e.g. 0-4, 5-9, 10-14, etc. 85+ is the oldest category"),
	),
	mcp.WithString("sex",
		mcp.Description("Either both sexes ('both'), male ('male') or female ('female')"),
	),
	mcp.WithString("ethnicity",
		mcp.Description("All ethnic groups ('overall'), Malay ('bumi_malay'), other Bumiputera ('bumi_other'), Chinese ('chinese'), Indian ('indian'), other citizens ('other_citizen'), or non-citizen residents ('other_noncitizen')"),
	),
	mcp.WithString("date",
		mcp.Description("The date in YYYY-MM-DD format, with MM-DD set to 01-01 as this is annual data"),
	),
)

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
