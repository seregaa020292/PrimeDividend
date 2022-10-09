package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func Fatalln(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Println(err error) {
	if err != nil {
		log.Println(err)
	}
}

func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func DoWithAttempts(fn func() error, maxAttempts int, delay time.Duration) error {
	var err error

	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			maxAttempts--

			continue
		}

		return nil
	}

	return err
}

func Context() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-sigs:
			cancel()
		case <-ctx.Done():
			return
		}
	}()

	return ctx
}

func WaitForService(host string) {
	log.Printf("Waiting for %s\n", host)

	for {
		log.Printf("Testing connection to %s\n", host)

		conn, err := net.Dial("tcp", host)
		if err == nil {
			_ = conn.Close()
			log.Printf("%s is up!\n", host)
			return
		}

		time.Sleep(time.Millisecond * 500)
	}
}

func OpenOrCreateFile(fileName string) (*os.File, error) {
	if err := os.MkdirAll(path.Dir(fileName), os.ModePerm); err != nil {
		return nil, err
	}

	return os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
}

func GenCookie(key string, value string, opts *http.Cookie) *http.Cookie {
	if opts == nil {
		opts = &http.Cookie{}
	}

	if opts.Path == "" {
		opts.Path = "/"
	}
	if opts.SameSite == http.SameSiteDefaultMode || opts.SameSite == 0 {
		opts.SameSite = http.SameSiteNoneMode
	}

	return &http.Cookie{
		Name:     key,
		Value:    value,
		Domain:   opts.Domain,
		HttpOnly: opts.HttpOnly,
		Expires:  opts.Expires,
		MaxAge:   opts.MaxAge,
		Path:     opts.Path,
		SameSite: opts.SameSite,
		Secure:   opts.Secure,
	}
}

// GenerateRandomBytes returns securely generated random bytes.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
