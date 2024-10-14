package engine

import (
	"embed"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type ScoreReport struct {
	Points       int      `json:"points"`
	PointsTotal  int      `json:"points_total"`
	Modules      []Module `json:"modules"`
	TotalModules int      `json:"total_modules"`
	MachineName  string   `json:"machine_name"`
}

//go:embed assets
var assets embed.FS

func runScoreReport() error {
	e := echo.New()
	fs := echo.MustSubFS(&assets, "assets")
	e.StaticFS("/", fs)
	e.GET("/report-data", func(c echo.Context) error {
		mods, err := getModules()
		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		// return the score report as JSON
		hname, err := os.Hostname()
		if err != nil {
			hname = "unknown"
		}
		report := ScoreReport{
			Points:       0,
			PointsTotal:  0,
			Modules:      mods,
			TotalModules: len(mods),
			MachineName:  hname,
		}

		for _, mod := range mods {
			if mod.Points == nil {
				continue
			}
			report.PointsTotal += *mod.Points
		}
		c.JSON(http.StatusOK, report)
		return nil
	})

	return e.Start(":8080")
}
