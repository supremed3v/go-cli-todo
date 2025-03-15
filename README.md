# CLI To-Do List

This is a simple command-line interface (CLI) to-do list application written in Go. It allows you to add, list, mark as completed, and remove tasks.

## Features

- Add a new task
- List all tasks
- Mark a task as completed
- Remove a task
- Save tasks to a JSON file

## Installation

1. Make sure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).
2. Clone the repository:
   ```sh
   git clone https://github.com/supremed3v/cli-todo.git
   ```
3. Navigate to the project directory:
   ```sh
   cd cli-todo
   ```
4. Build the project:
   ```sh
   go build
   ```

## Usage

Run the application:

```sh
./cli-todo
```

Follow the on-screen instructions to manage your tasks.

## Code Overview

### Main Function

The `main` function initializes the tasks and provides a menu for the user to interact with the to-do list.

### Task Struct

The `Task` struct represents a task with a description and a completion status.

### Functions

- `loadTasks()`: Loads tasks from `tasks.json`.
- `saveTasks(tasks []Task)`: Saves tasks to `tasks.json`.
- `addTask(tasks *[]Task)`: Adds a new task to the list.
- `listTasks(tasks []Task)`: Lists all tasks.
- `markTaskAsCompleted(tasks *[]Task)`: Marks a task as completed.
- `removeTask(tasks *[]Task)`: Removes a task from the list.

### Take Aways

- You will learn why we use pointers (in our case to update variables)
- How to get user input
- How to write data to a file
- Working with JSON in Go

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Contact

For any questions or suggestions, please contact [saadsiddiqui-programmer@protonmail.com](mailto:saadsiddiqui-programmer@protonmail.com).
