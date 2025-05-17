package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const BaseURL = "https://api.data.gov.my/data-catalogue"

// LoadTools returns a map of predefined MCP tools with their corresponding handler functions.
// Each tool represents a specific dataset or service with configurable parameters.
// The returned map can be used to register and manage different data retrieval tools.
func LoadTools() map[*mcp.Tool]server.ToolHandlerFunc {
	return map[*mcp.Tool]server.ToolHandlerFunc{
		&populationState:             populationStateHandler,
		&populationMalaysia:          populationMalaysiaHandler,
		&births:                      birthHandler,
		&fuelprice:                   fuelpriceHandler,
		&registrationTransactionsCar: registrationTransactionsCarHandler,
		&hhIncome:                    hhIncomeHandler,
	}
}
