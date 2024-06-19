package main

import (
	"fmt"
	"log"
	"os"

	"weather_app/utils"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Weather App",
		Usage: "Fetch current weather or convert speed",
		Authors: []*cli.Author{
			{
				Name:  "Gooooaha Agaga Inc.",
				Email: "crappyqr@example.com",
			},
		},
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "speed",
				Usage: "Speed in meters per second (m/s) to convert to km/h",
			},
			&cli.StringFlag{
				Name:  "city",
				Usage: "City name to fetch current weather from OpenWeatherMap API",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Weather CLI tool by Gooooahaaa Agaga Inc (crappyqr@example.com)")
			fmt.Println("---------------------------------------------------------------")
			fmt.Println("                                                               ")

			speed := c.Float64("speed")
			city := c.String("city")

			if speed != 0 {
				speedKmh := utils.ConvertSpeed(speed)
				fmt.Printf("%.2f m/s is equal to %.2f km/h\n", speed, speedKmh)
			}

			if city != "" {
				err := utils.FetchWeather(city)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
