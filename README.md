# MCP Server for DataGovMy

This project is an [MCP (Model Context Protocol)](https://github.com/modelcontextprotocol) server written in Go that exposes tools and resources from [DataGovMy](https://github.com/datagovmy/data) — Malaysia's open government data platform. It allows you to access Malaysian government datasets (such as population statistics) programmatically via Cline (an AI agent) and Cursor (an AI-powered code editor).

## Features

- Exposes DataGovMy datasets (e.g., population by state, age, ethnicity, etc.) as MCP tools
- Enables AI-powered access to Malaysian government data via Cline and Cursor
- Written in Go for performance and portability
- Easily extensible to support more DataGovMy datasets

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

     ```
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

- Add new DataGovMy tools by editing `tool.go` and `handlers.go`.
- Register new endpoints and update the MCP tool schema as needed.
- Rebuild and restart the server to apply changes.

## Troubleshooting

- **Port already in use:** Change the port in the code or stop the conflicting process.
- **Cline can't connect:** Ensure the server is running and the URL is correct in Cline's settings.
- **Go build errors:** Make sure you have the correct Go version and all dependencies installed.

## License

MIT License

---

For more information, see the [Model Context Protocol documentation](https://github.com/modelcontextprotocol) and [DataGovMy](https://github.com/datagovmy/data).
