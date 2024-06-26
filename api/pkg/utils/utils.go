package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	mathRand "math/rand"
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

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

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[mathRand.Intn(len(letters))]
	}
	return string(s)
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
