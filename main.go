package main

import (
	"errors"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"golang.org/x/xerrors"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
	if err := doSomething1(); err != nil {
		sentry.CaptureException(err)
	}
}

func doSomething1() error {
	if err := doSomething2(); err != nil {
		return xerrors.Errorf("doSomething1: %w", err)
	}
	return nil
}

func doSomething2() error {
	if err := doSomething3(); err != nil {
		return xerrors.Errorf("doSomething2: %w", err)
	}
	return nil
}

func doSomething3() error {
	if err := somethingWentWrong(); err != nil {
		return xerrors.Errorf("doSomething3: %w", err)
	}
	return nil
}

func somethingWentWrong() error {
	err := errors.New("internal error: param1: c")
	return xerrors.Errorf("oops: %w", err)
}
