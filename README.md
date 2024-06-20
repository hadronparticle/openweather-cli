### README.md

# Weather App

A CLI application to fetch current weather or convert speed using the OpenWeatherMap API.

## Author

Weather App by 4I9KKZ. (crappyqr@example.com)

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
git clone https://github.com/hadronparticle/openweather-cli
cd openweather-cli
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
go build -o weather-app
```

## Usage

### Linux

1. **Set the correct permissions**:

    ```sh
    chmod +x weather-app
    ```

2. **Run the executable**:

    To fetch weather data for a city:

    ```sh
    ./weather-app --city "London"
    ```

    To convert speed from m/s to km/h:

    ```sh
    ./weather-app --speed 10
    ```

    To view the help message:

    ```sh
    ./weather-app --help
    ```

### Windows

1. **Run the executable**:

    To fetch weather data for a city:

    ```sh
    weather-app.exe --city "Cagayan de Oro City"
    ```

    To convert speed from m/s to km/h:

    ```sh
    weather-app.exe --speed 10
    ```

    To view the help message:

    ```sh
    weather-app.exe --help
    ```

### Installing Executable in Linux

1. **Choose Installation Directory**:

   Decide on a directory where you want to place your executable. A common choice is `/usr/local/bin`, which is typically included in the `PATH` by default.

   ```bash
   # Assuming your executable is named weather-app
   sudo cp weather-app /usr/local/bin
   ```

   Replace `weather-app` with the actual name of your executable.

2. **Set Execution Permissions**:

   Ensure the executable has the correct permissions to be executed:

   ```bash
   sudo chmod +x /usr/local/bin/weather-app
   ```

3. **Verify Installation**:

   To verify that the installation was successful, you can try running the executable from any directory:

   ```bash
   weather-app --help
   ```

   This command should display the help information for your application, indicating that it is installed correctly and accessible from the command line.

### Notes

- **Alternative Paths**: If you prefer not to use `/usr/local/bin`, you can choose another directory that is in your `PATH` environment variable. You can check your current `PATH` with `echo $PATH`.

- **Permission Issues**: If you encounter permission issues, ensure you have appropriate permissions or use `sudo` as necessary.

### Uninstallation

If you ever need to uninstall the executable:

```bash
sudo rm /usr/local/bin/weather-app
```

This approach allows you to install your executable in Linux so that it can be invoked from anywhere in the command line, leveraging a directory (`/usr/local/bin`) that is commonly used for user-installed executables. Adjust the paths and commands as needed based on your specific setup and requirements.

## License
This project is licensed under the MIT License.
