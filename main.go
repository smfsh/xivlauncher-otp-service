package main

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	secret := os.Getenv("XIVOTP_SECRET")
	envIps := os.Getenv("XIVLAUNCHER_IPS")
	var ips []string

	if secret == "" || envIps == "" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go <secret> <ip1> [, <ip2>, ...]")
			fmt.Println("Or set environment variables: XIVOTP_SECRET and XIVLAUNCHER_IPS (comma-separated IPs)")
			return
		}
		secret = os.Args[1]
		ips = os.Args[2:]
	} else {
		ips = strings.Split(envIps, ",")
	}

	fmt.Printf("Starting XIVLauncher OTP service for IPs: %v\n", ips)

	// Map to track cooldowns
	cooldowns := make(map[string]time.Time)

	for {
		for _, ip := range ips {
			url := fmt.Sprintf("http://%s:4646/ffxivlauncher/", ip)

			// Skip IPs on cooldown
			if cooldown, ok := cooldowns[ip]; ok && time.Now().Before(cooldown) {
				continue
			}

			if checkPortOpen(ip + ":4646") {
				otp, err := generateOTP(secret)
				if err != nil {
					fmt.Printf("Error generating OTP for IP %s: %v\n", ip, err)
					continue
				}
				fmt.Printf("Generated OTP for IP %s: %s\n", ip, otp)
				sendOTP(url, otp)
				cooldowns[ip] = time.Now().Add(1 * time.Minute)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func checkPortOpen(address string) bool {
	timeout := time.Second * 2
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		// Log errors silently
		return false
	}
	defer conn.Close()
	return true
}

func generateOTP(secret string) (string, error) {
	otp, err := totp.GenerateCode(secret, time.Now())
	return otp, err
}

func sendOTP(url, otp string) {
	responseURL := url + otp
	resp, err := http.Get(responseURL)
	if err != nil {
		fmt.Printf("Error sending OTP to %s: %v\n", url, err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body from %s: %v\n", url, err)
		return
	}

	fmt.Printf("Successful OTP submission to %s. Status: %s\n", url, resp.Status)
	fmt.Println("Response body:", string(body))
}
