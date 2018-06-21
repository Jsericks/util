package util

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

const (
	NotFoundInEnv = `%s not found in ENV`
)

func EnsureEnv(envVars ...string) error {
	ok := true
	missing := []string{}
	for _, key := range envVars {
		if _, k := os.LookupEnv(key); !k {
			ok = false
			missing = append(missing, key)
		}
	}
	if !ok {
		return fmt.Errorf(NotFoundInEnv, strings.Join(missing, ", "))
	}
	return nil
}

func FetchEnvStringWithDefault(k, d string) string {
	if s, ok := os.LookupEnv(k); ok {
		return s
	}
	return d
}

func FetchEnvIntWithDefault(k string, d int) int {
	if v, ok := os.LookupEnv(k); ok {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return d
}

func FetchEnvStringAndSplitWithDefault(k, d string) []string {
	s := FetchEnvStringWithDefault(k, d)
	return strings.Split(s, ",")
}

func MakeCancelChan() chan struct{} {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	return func(sigCh chan os.Signal) chan struct{} {
		cancel := make(chan struct{})
		go func() {
			<-sigCh
			close(cancel)
		}()
		return cancel
	}(sig)
}
