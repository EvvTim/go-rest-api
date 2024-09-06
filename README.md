This is a boilerplate project for building a REST API using Go (Golang). The project is structured in a modular way, with different internal packages handling configuration, logging, and user-related functionalities. This setup promotes clean code organization and separation of concerns.

**Project Structure**

```bash
go-rest-api/
├── cmd/
│   └── main.go                # Main entry point for the application
├── internal/
│   ├── config/
│   │   └── config.go          # Configuration loader
│   ├── handlers/
│   │   └── handler.go         # General HTTP handlers for the API
│   └── user/
│       ├── handler.go         # User-related HTTP handlers
│       ├── model.go           # User model and related structures
│       └── storage.go         # Storage layer for users (DB interactions)
├── pkg/
│   └── logging/
│       └── logging.go         # Logging utilities using logrus
├── .gitignore                 # Git ignore file
├── config.yml                 # Configuration file
├── go.mod                     # Go module dependencies
└── go.sum                     # Go dependencies lock file
```

**Features**

• Modular project structure.
• User management (example).
• Logging utility for tracking and debugging (using logrus).
• Configuration management using YAML.
• Ready-to-use boilerplate for REST API development.

  **Prerequisites**
  
• [Go](https://golang.org/doc/install) 1.16 or higher.
• Make sure you have git installed and configured.
  

**Getting Started**

1. Clone the repository:
```bash
git clone https://github.com/your-username/go-rest-api.git
cd go-rest-api
```
2. Install dependencies:
```bash
go mod tidy
```
3. Configure the application by editing the config.yml file.

Example config.yml:

```yml
is_debug: true
listen:
  type: tcp           # Options: 'tcp' or 'sock'
  bind_ip: 0.0.0.0    # Set to '0.0.0.0' to listen on all interfaces
  port: 10000         # The port to listen on if type is 'tcp'
```
4. Run the application:
```bash
go run cmd/main.go
```
5. The API will start on http://0.0.0.0:10000.


**Configuration Details**

The application configuration is managed through the config.yml file. Here are the key configuration parameters:

• is_debug: Boolean flag to enable or disable debug mode.
• listen.type: Determines whether the application listens on a TCP port or Unix socket.
• Use tcp to bind to an IP address and port (default mode).
• Use sock for Unix socket mode (creates a socket file app.sock).
• listen.bind_ip: The IP address to bind when in TCP mode (e.g., 127.0.0.1 or 0.0.0.0).
• listen.port: The port number to listen on (e.g., 8080).

**Logging**

The project uses the logrus package for logging. All logs are written to both the console and a log file (logs/all.log). Here’s how the logging is structured:

• **File output:** All logs are saved to logs/all.log.

• **Console output:** Logs are also printed to the console (stdout).

• **Log levels:** The logging level is set to TraceLevel, which means it captures all log levels, including Debug, Info, Warn, and Error.

• **Caller information:** Log entries include the file name and line number from where the log was generated.

**Log setup**

In the pkg/logging/logging.go, a new logger is created during initialization:

1. **Create logs directory**: If the logs directory does not exist, it will be created.
2. **Set log output**: Log messages are written both to the logs/all.log file and the standard output.
3. **Set log level**: The logger is configured to log at the TraceLevel, which captures the most detailed logs.

  
**Example log output:**
```plaintext
INFO[0000] create router                              main.main() main.go:13
INFO[0000] register user handler                      main.main() main.go:18
INFO[0000] start application                          main.start() main.go:22
INFO[0000] listen tcp                                 main.start() main.go:34
INFO[0000] server is listening port: 10000            main.start() main.go:36
```

This shows a log entry with the timestamp, log level (INFO), message, function name, and the file/line number of the log entry.

**Contributing**

Feel free to open issues or submit pull requests if you have suggestions or improvements.
