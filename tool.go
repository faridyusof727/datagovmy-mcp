package main

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func LoadTools() map[*mcp.Tool]server.ToolHandlerFunc {
	populationMalaysia := mcp.NewTool("population_malaysia",
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

	populationState := mcp.NewTool("population_state",
		mcp.WithDescription("Population at state level from 1970 to 2024, by sex, age group and ethnicity"),
		mcp.WithString("age",
			mcp.Description("Either all age groups ('overall') or five-year age groups e.g. 0-4, 5-9, 10-14, etc. 85+ is the oldest category"),
		),
		mcp.WithString("sex",
			mcp.Description("Male ('male') or female ('female'). Can't use ('both')"),
		),
		mcp.WithString("ethnicity",
			mcp.Description("All ethnic groups ('overall'), Malay ('bumi_malay'), other Bumiputera ('bumi_other'), Chinese ('chinese'), Indian ('indian'), other citizens ('other_citizen'), or non-citizen residents ('other_noncitizen')"),
		),
		mcp.WithString("state",
			mcp.Description("One of 16 states"),
		),
		mcp.WithString("date",
			mcp.Description("The date in YYYY-MM-DD format, with MM-DD set to 01-01 as this is annual data"),
		),
	)

	return map[*mcp.Tool]server.ToolHandlerFunc{
		&populationState:    populationStateHandler,
		&populationMalaysia: populationMalaysiaHandler,
	}
}
