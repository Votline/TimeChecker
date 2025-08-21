# TimeChecker

TimeChecker is a minimalist time display application built with OpenGL and GLFW.

## Description

The application displays current time in a separate window.

### Key Features:
- **Time display**: Current time in HH:MM:SS format
- **Lightweight**: Minimal resource consumption
- **StopWatch mode**: Simple stopwatch, controlled by one button
- **On Top**: Over all windows mode

## Technologies
- **Go** - Primary programming language
- **[Go-gl](https://github.com/go-gl/gl)** - Graphics rendering
- **[Go-gl/glfw](https://github.com/go-gl/glfw)** - Window management

## Installation & Building

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/Votline/TimeCheck
    cd TimeCheck
    ```

2.  **Download dependencies**:
    ```bash
    go mod download
    ```

3.  **Build the application**:
    ```bash
    go build -o timecheck
    ```

4.  **Run**:
    ```bash
    ./timecheck
    ```

## Licenses
This project is licensed under [GNU AGPL v3](LICENSE).

The full license texts are available in the [licenses directory](licenses/)
