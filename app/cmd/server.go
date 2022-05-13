// Content managed by Project Forge, see [projectforge.md] for details.
package cmd

import (
	"context"
	"fmt"
	"runtime"

	"github.com/muesli/coral"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller"
	"projectforge.dev/projectforge/app/lib/filesystem"
	"projectforge.dev/projectforge/app/lib/telemetry"
	"projectforge.dev/projectforge/app/util"
)

const keyServer = "server"

func serverCmd() *coral.Command {
	short := fmt.Sprintf("Starts the http server on port %d (by default)", util.AppPort)
	f := func(*coral.Command, []string) error { return startServer(_flags) }
	ret := &coral.Command{Use: keyServer, Short: short, RunE: f}
	return ret
}

func startServer(flags *Flags) error {
	if err := initIfNeeded(); err != nil {
		return errors.Wrap(err, "error initializing application")
	}

	st, r, _, err := loadServer(flags, _logger)
	if err != nil {
		return err
	}

	_, err = listenandserve(util.AppName, flags.Address, flags.Port, r)
	if err != nil {
		return err
	}

	err = st.Close(context.Background())
	if err != nil {
		return errors.Wrap(err, "unable to close application")
	}

	return nil
}

func loadServer(flags *Flags, logger util.Logger) (*app.State, fasthttp.RequestHandler, util.Logger, error) {
	r := controller.AppRoutes()
	f := filesystem.NewFileSystem(flags.ConfigDir, logger)
	telemetryDisabled := util.GetEnvBool("disable_telemetry", false)
	st, err := app.NewState(flags.Debug, _buildInfo, f, !telemetryDisabled, logger)
	if err != nil {
		return nil, nil, logger, err
	}

	ctx, span, logger := telemetry.StartSpan(context.Background(), "app:init", logger)
	defer span.Complete()

	svcs, err := app.NewServices(ctx, st)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "error creating services")
	}
	st.Services = svcs

	controller.SetAppState(st)

	logger.Infof("started %s v%s using address [%s:%d] on %s:%s", util.AppName, _buildInfo.Version, flags.Address, flags.Port, runtime.GOOS, runtime.GOARCH)

	return st, r, logger, nil
}
