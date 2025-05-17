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

var fuelprice = mcp.NewTool("fuelprice",
	mcp.WithDescription("Weekly retail prices of RON95 petrol, RON97 petrol, and diesel in Malaysia"),
	mcp.WithString("date",
		mcp.Description("The date of effect of the price, in YYYY-MM-DD format. You can omit this field if not found any data"),
	),
	mcp.WithString("series_type",
		mcp.Description("Price in RM (level), or weekly change in RM (change_weekly)."),
	),
)

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
