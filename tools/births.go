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

var births = mcp.NewTool("births",
	mcp.WithDescription("Number of people born daily in Malaysia, based on registrations with JPN from 1920 to the present"),
	mcp.WithString("date",
		mcp.Description("Date of birth in YYYY-MM-DD format; note that this date represents the actual date of birth and NOT the date of registration with JPN"),
	),
	mcp.WithNumber("births",
		mcp.Description("Number of births for the date"),
	),
)

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
