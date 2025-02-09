package test

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/esnet/gdg/pkg/test_tooling"
	"github.com/esnet/gdg/pkg/test_tooling/path"
	"github.com/joho/godotenv"
)

const (
	grafana10 = "10.2.3-ubuntu"
	grafana11 = "11.1.5-ubuntu"
)

func TestMain(m *testing.M) {
	gofakeit.Seed(time.Now().Unix()) // If 0 will use crypto/rand to generate a number
	err := path.FixTestDir("test", "..")
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(".env")
	// set global log level
	slog.SetLogLoggerLevel(slog.LevelDebug) // Set global log level to Debug
	grafanaTestVersions := []string{grafana10, grafana11}
	testModes := []string{"basicAuth", "token"}
	if os.Getenv("DEVELOPER") == "1" {
		slog.Debug("Limiting to single testMode and grafana version", slog.Any("grafanaVersion", grafanaTestVersions[1]), slog.String("testMode", testModes[0]))
		grafanaTestVersions = grafanaTestVersions[1:]
		testModes = testModes[0:1]
	}

	for _, version := range grafanaTestVersions {
		for _, i := range testModes {
			os.Setenv(test_tooling.GrafanaTestVersionEnv, version)
			if i == "token" {
				os.Setenv(test_tooling.EnableTokenTestsEnv, "1")
			} else {
				os.Setenv(test_tooling.EnableTokenTestsEnv, "0")
			}
			slog.Info("Running Test suit for",
				slog.Any("grafanaVersion", version),
				slog.Any("AuthMode", i))
			exitVal := m.Run()
			if exitVal != 0 {
				panic("Failed to run test with token value of: " + i)
			}
		}
	}
}
