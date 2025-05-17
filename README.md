# MCP Server for DataGovMy

This project is an [MCP (Model Context Protocol)](https://github.com/modelcontextprotocol) server written in Go that exposes tools and resources from [DataGovMy](https://github.com/data-gov-my) — Malaysia's open government data platform. It allows you to access Malaysian government datasets (such as population statistics) programmatically via Cline (an AI agent) and Cursor (an AI-powered code editor).

![Show](show.gif)

## Features

- Exposes DataGovMy datasets (e.g., population by state, age, ethnicity, etc.) as MCP tools
- Enables AI-powered access to Malaysian government data via Cline and Cursor
- Written in Go for performance and portability
- Easily extensible to support more DataGovMy datasets

## About This Project

This repository provides an MCP (Model Context Protocol) server that acts as an interface to connect to the [data.gov.my](https://data.gov.my/) API (where available). It is designed for MCP integration and is not specific to any single assistant or application.

- **Why Go?**  
  You might be wondering why this project uses Go instead of the [official MCP package](https://github.com/modelcontextprotocol/typescript-sdk) in Node.js. The answer is simple: I know Go better, it has fewer dependencies, and I can easily release binaries for different operating systems.

- **No vector database, custom LLM training, or RAG:**  
  This project does not use any kind of vector database, custom large language model training, or retrieval-augmented generation (RAG). It is purely an MCP server that exposes DataGovMy datasets as tools for AI agents and developer tools.

- **Upcoming DataGovMy Assistant:**  
  The DataGovMy team is working on a ChatGPT-style assistant, [MyDataGPT](https://github.com/data-gov-my/datagovmy-ai?tab=readme-ov-file#mydatagpt-assistant-coming-soon-), which will provide a conversational interface to Malaysian government data. This MCP server is a separate project focused on MCP integration.

## Available Tools

This MCP server currently exposes the following tools:

- **births**: Number of people born daily in Malaysia, based on registrations with JPN from 1920 to the present.
- **fuelprice**: Weekly retail prices of RON95 petrol, RON97 petrol, and diesel in Malaysia.
- **population_malaysia**: Population at national level from 1970 to 2024, by sex, age group and ethnicity.
- **population_state**: Population at state level from 1970 to 2024, by sex, age group and ethnicity.
- **registration_transactions_car**: Car registration transactions from 2000 to the present. Filter by registration date, type (motokar, MPV, jeep, pick-up, window van), maker, model, colour, fuel type (petrol, diesel, green diesel, natural gas, LNG, hydrogen, electric, hybrid), and state of registration.
- **hh_income**: Mean and median monthly gross household income in Malaysia from 1970 to 2022.

Want to add more tools? Contributions are welcome! Please open an issue or submit a pull request if you'd like to help extend the server with additional DataGovMy datasets or features.

## Download

You can download the latest release for Windows, macOS, or Linux from the [GitHub Releases page](https://github.com/faridyusof727/datagovmy-mcp/releases).

1. Go to the [Releases page](https://github.com/faridyusof727/datagovmy-mcp/releases).
2. Find the version you want (e.g., `v1.0.0`).
3. Download the binary for your operating system:
   - `*_windows_amd64.zip` for Windows
   - `*_darwin_amd64.tar.gz` for macOS
   - `*_linux_amd64.tar.gz` for Linux
4. Extract the downloaded file and follow the usage instructions.

## What is DataGovMy?

[DataGovMy](https://github.com/datagovmy/data) is Malaysia's open data platform, providing datasets on demographics, economics, health, and more. This MCP server makes those datasets accessible to AI agents and developer tools.

## Prerequisites

- [Go](https://golang.org/dl/) 1.18 or newer
- [Cline](https://cline.modelcontext.com/) (installed and configured)
- [Cursor](https://www.cursor.so/) (optional, for AI-powered code editing)
- (Optional) [Node.js](https://nodejs.org/) if you want to use Cline's browser tools

## Installation

1. **Clone this repository:**

   ```sh
   git clone <your-repo-url>
   cd custom
   ```

2. **Install Go dependencies:**

   ```sh
   go mod tidy
   ```

3. **Build the server:**

   ```sh
   go build -o mcp-datagovmy
   ```

4. **Run the server:**

   ```sh
   ./mcp-datagovmy
   ```

   By default, the server will start on `localhost:8080`. You can change the port by editing the code or using environment variables if supported.

## Using with Cline or Cursor (via stdio)

You can configure Cline (and Cursor) to launch this MCP server as a subprocess and communicate via stdio, which is more robust for local development than HTTP.

1. **Build the server:**

   ```sh
   go build -o mcp-datagovmy
   ```

2. **Update your Cline (or Cursor) MCP config file** (usually `cline_mcp_settings.json`). Add an entry like this:

   ```json
   {
     "mcpServers": {
       "datagovmy": {
         "disabled": false,
         "timeout": 60,
         "command": "/absolute/path/to/mcp-datagovmy",
         "transportType": "stdio",
         "autoApprove": [
           "population_malaysia",
           "population_state"
         ]
       }
     }
   }
   ```

   - **command**: Absolute path to your built server binary (e.g., `/Users/youruser/Documents/Codes/mcp/custom/mcp-datagovmy`)
   - **transportType**: Must be `"stdio"` for subprocess communication
   - **autoApprove**: (Optional) List of tool names to auto-approve without prompting
   - **timeout**: (Optional) How long to wait for a response (in seconds)
   - **disabled**: Set to `false` to enable the server

3. **Save and reload Cline or Cursor.** The server will be launched automatically as needed.

4. **Use Cline or Cursor to access DataGovMy tools.**
   - Example usage:

     ```text
     Use the "population_malaysia" tool from the "datagovmy" server with age="overall", date="2024-01-01", ethnicity="overall", sex="both".
     ```

## Using with Cursor

1. **Open your project in Cursor.**
2. **Ensure Cline is connected and your MCP server is running.**
3. **Use Cursor's AI features to interact with DataGovMy tools.**
   - For example, you can ask Cursor to "use the population_state tool from the datagovmy server with state='Selangor', age='0-4', date='2024-01-01', ethnicity='chinese', sex='male'".

## Example Workflow

1. Start your MCP server:

   ```sh
   ./mcp-datagovmy
   ```

2. Connect Cline to your server as described above.
3. In Cline or Cursor, run a command like:

   ```text
   Use the "population_state" tool from the "datagovmy" server with state="Selangor", age="0-4", date="2024-01-01", ethnicity="chinese", sex="male".
   ```

## Extending the Server

The server is now modularized for easier extension. To add a new DataGovMy tool:

1. **Create a new file in the `tools/` directory** (e.g., `tools/my_new_tool.go`).
   - Define your tool using `mcp.NewTool` and implement its handler function in this file.
   - See existing files in `tools/` (like `births.go`, `population_malaysia.go`) for examples.

2. **Register your tool and handler** in `tools/tool.go`:
   - Import your new tool and handler at the top if needed.
   - Add an entry to the map returned by `LoadTools()`:

     ```go
     &myNewTool: myNewToolHandler,
     ```

3. **Rebuild and restart the server** to apply your changes.

This modular structure makes it easy to add, update, or debug individual tools without affecting others.

**Tip for contributors:**  
When implementing new tools or handlers, refer to the [DataGovMy Developer Portal](https://developer.data.gov.my/) for guidance on API request structure, available endpoints, and best practices.

## Troubleshooting

- **Port already in use:** Change the port in the code or stop the conflicting process.
- **Cline can't connect:** Ensure the server is running and the URL is correct in Cline's settings.
- **Go build errors:** Make sure you have the correct Go version and all dependencies installed.

## License

MIT License

---

For more information, see the [Model Context Protocol documentation](https://github.com/modelcontextprotocol) and [DataGovMy](https://github.com/datagovmy/data).
