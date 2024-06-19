### README.md

# Weather App

A CLI application to fetch current weather or convert speed using the OpenWeatherMap API.

## Author

Weather App by Gooooaha Agaga Inc. (crappyqr@example.com)

Pet project for Mindinet

## Features

- Convert speed from meters per second (m/s) to kilometers per hour (km/h)
- Fetch current weather information for a specified city

## Requirements

- Go 1.16 or higher
- OpenWeatherMap API Key

## Installation

### Clone the repository

```sh
git clone https://hadronp@bitbucket.org/hadronp/openweather-cli.git
cd weather_app
```

### Initialize Go modules

```sh
go mod init weather_app
go get github.com/urfave/cli/v2
```

### Set your OpenWeatherMap API Key

Edit the `utils/fetch.go` file and replace `"YOUR_API_KEY"` with your actual OpenWeatherMap API key.

```go
const apiKey = "YOUR_API_KEY"
```

### Build the application

```sh
go build -o weather_app
```

## Usage

### Linux

1. **Set the correct permissions**:

    ```sh
    chmod +x weather_app
    ```

2. **Run the executable**:

    To fetch weather data for a city:

    ```sh
    ./weather_app --city "London"
    ```

    To convert speed from m/s to km/h:

    ```sh
    ./weather_app --speed 10
    ```

    To view the help message:

    ```sh
    ./weather_app --help
    ```

### Windows

1. **Run the executable**:

    To fetch weather data for a city:

    ```sh
    weather_app.exe --city "Cagayan de Oro City"
    ```

    To convert speed from m/s to km/h:

    ```sh
    weather_app.exe --speed 10
    ```

    To view the help message:

    ```sh
    weather_app.exe --help
    ```

## License
This project is licensed under the MIT License.
