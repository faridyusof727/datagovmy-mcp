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

	births := mcp.NewTool("births",
		mcp.WithDescription("Number of people born daily in Malaysia, based on registrations with JPN from 1920 to the present"),
		mcp.WithString("date",
			mcp.Description("Date of birth in YYYY-MM-DD format; note that this date represents the actual date of birth and NOT the date of registration with JPN"),
		),
		mcp.WithNumber("births",
			mcp.Description("Number of births for the date"),
		),
	)

	fuelprice := mcp.NewTool("fuelprice",
		mcp.WithDescription("Weekly retail prices of RON95 petrol, RON97 petrol, and diesel in Malaysia"),
		mcp.WithString("date",
			mcp.Description("The date of effect of the price, in YYYY-MM-DD format. You can omit this field if not found any data"),
		),
		mcp.WithString("series_type",
			mcp.Description("Price in RM (level), or weekly change in RM (change_weekly)."),
		),
	)
	registrationTransactionsCar := mcp.NewTool("registration_transactions_car",
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

	return map[*mcp.Tool]server.ToolHandlerFunc{
		&populationState:             populationStateHandler,
		&populationMalaysia:          populationMalaysiaHandler,
		&births:                      birthHandler,
		&fuelprice:                   fuelpriceHandler,
		&registrationTransactionsCar: registrationTransactionsCarHandler,
	}
}
