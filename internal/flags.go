package internal

import (
	"os"
	"path/filepath"
	"strconv"
)

func GetListenPort() int {
	if port := os.Getenv("OFLY_PORT"); port != "" {
		return Must(strconv.Atoi(port))
	}
	return 0
}

func Keystorage() string {
	store := os.Getenv("OFLY_PATH")
	if store == "" {
		store = filepath.Join(Must(os.UserConfigDir()), "ofly")
	}
	return store
}

func getKeyName() string {
	name := os.Getenv("OFLY_KEY")
	if name == "" {
		name = filepath.Base(Must(os.Executable()))
	}
	return name
}

func ApiDomain() string {
	if domain := os.Getenv("OFLY_API"); domain != "" {
		return domain
	}
	return "ofly.opsdude.com"
}

func AuthKey() string {
	return os.Getenv("OFLY_AUTH")
}

func ServerIp() string {
	ip := os.Getenv("OFLY_IP")
	return ip
}

func SSLCertificateEmail() string {
	if email := os.Getenv("OFLY_SSL_EMAIL"); email != "" {
		return email
	}
	return "certs@ofly.opsdude.com"
}

func UseRelay() bool {
	return os.Getenv("OFLY_RELAY") != ""
}

func TestOnlyRunLocalhost() bool {
	return os.Getenv("OFLY_TEST_LOCALHOST") == "true"
}
