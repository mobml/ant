# Ant

A CLI tool for personal productivity and growth that helps you manage your personal development through a structured hierarchy: **Plans** → **Areas** → **Goals** → **Habits**.

## Overview

Ant is a command-line interface tool designed to help you create structured personal development plans, break them down into life areas, set goals within those areas, and track daily habits that contribute to achieving those goals.

## Features

- **Plan Management**: Create and manage personal development plans with duration
- **Area Management**: Organize life areas (e.g., Health, Career, Relationships)
- **Goal Management**: Set specific goals within each area
- **Habit Tracking**: Create and track daily habits linked to goals
- **Daily Notes**: Journal functionality for daily reflections
- **Weekly Reports**: Generate markdown reports for weekly progress
- **Habit Scheduling**: Schedule habits with specific timing
- **Interactive Wizards**: CLI-based interactive creation workflows
- **Today View**: See and mark today's habits as complete

## Installation

### Prerequisites

- Go 1.24 or later

### Build from Source

```bash
# Clone the repository
git clone https://github.com/mobml/ant.git
cd ant

# Build the application
go build -o ant

# Run the CLI
./ant --help
```

### Run Directly

```bash
# Run directly with go run
go run . --help
```

## Quick Start

1. **Create your first plan**:
   ```bash
   ./ant plan create
   ```

2. **Add life areas to your plan**:
   ```bash
   ./ant area create
   ```

3. **Set goals in each area**:
   ```bash
   ./ant goal create
   ```

4. **Create habits to achieve your goals**:
   ```bash
   ./ant habit create
   ```

5. **Track your daily progress**:
   ```bash
   ./ant today
   ```

## Commands

### Plan Management
```bash
./ant plan create           # Create a new plan
./ant plan list             # List all plans
./ant plan del [id plan]    # delete plan
./ant plan update [id plan] # update plan 
```

### Area Management
```bash
./ant area create             # Create a new area
./ant area list [id plan]     # List all areas
./ant area del  [id area]     # delete an area
./ant area update [id area]   # update an  area
```

### Goal Management
```bash
./ant goal create             # Create a new goal
./ant goal list [id area]     # List all goals
./ant goal del [id goal]      # Delete a goal
./ant goal update [id goal]   # Update a goal

```

### Habit Management
```bash
./ant habit create            # Create a new habit
./ant habit list [id goal]    # List all habits
./ant habit del [id habit]    # Delete habit
./ant habit update [id habit] # Update habit
./ant habit mark [id habit]   # Mark habit as complete
./ant habit today             # Show habits for today

```


## Project Structure

```
ant/
├── main.go                    # Application entry point
├── go.mod/go.sum             # Go module dependencies
├── cmd/                      # CLI command definitions
│   ├── root.go              # Root command
│   ├── plan.go              # Plan management commands
│   ├── area.go              # Area management commands
│   ├── goal.go              # Goal management commands
│   ├── habit.go             # Habit management commands
│   └── *.go                 # Individual command implementations
├── internal/                 # Internal application code
│   ├── models/              # Data models
│   ├── services/            # Business logic services
│   ├── repositories/        # Data access layer
│   ├── domain/              # Domain logic and validation
│   └── cli/                 # CLI utilities and wizards
├── database/                # Database management
│   ├── db.go               # Database connection
│   ├── migrate.go          # Migration runner
│   └── migrations/         # SQL migration files
└── .github/workflows/      # CI/CD workflows
```

## Architecture

Ant follows clean architecture principles with clear separation between:
- **CLI Commands**: Command-line interface and user interactions
- **Services**: Business logic and application services
- **Repositories**: Data access layer
- **Domain Models**: Core business entities and validation

## Database

The application uses DuckDB as an embedded SQL database with the following main tables:
- `plans` - Personal development plans
- `areas` - Life areas within plans
- `goals` - Goals within areas
- `habits` - Daily habits linked to goals
- `habit_logs` - Habit completion tracking
- `habit_schedules` - Habit scheduling
- `daily_notes` - Daily journal entries
- `weekly_reports` - Generated weekly reports

## Roadmap

Planned improvements:
- Data export (CSV / JSON)
- Habit streaks and analytics
- Better reporting and summaries

## Dependencies

- `github.com/spf13/cobra v1.9.1` - CLI framework
- `github.com/duckdb/duckdb-go/v2 v2.5.3` - DuckDB database driver
- `github.com/manifoldco/promptui v0.9.0` - Interactive CLI prompts

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

If you encounter any issues or have questions, please file an issue on the GitHub repository.
