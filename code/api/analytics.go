package api

import (
	"bytes"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strings"
)

var (
	AnalyticsController IAnalyticsController
)

type EventType string

const (
	EVENT_SHARE = "Share"
	EVENT_SUBMIT = "Submit"
)

type IAnalyticsController interface {
	RegisterEvent(EventType, *http.Request)
}

type STDAnalytics struct {

}

func (std STDAnalytics) RegisterEvent(event_type EventType, r *http.Request) {
    slog.Info(fmt.Sprintf("event triggered: %s", event_type))
}


type PlausibleAnalytics struct {
	PlausibleURL string
	AppDomain string
	AppURL string
}

func NewPlausibleAnalytics(
	plaus_url string,
	app_domain string,
	app_url string,
) PlausibleAnalytics {
	return PlausibleAnalytics{
		PlausibleURL: plaus_url,
		AppDomain: app_domain,
		AppURL: app_url,
	}
}

func (pa PlausibleAnalytics) RegisterEvent(event_type EventType, r *http.Request) {
    payload := []byte(fmt.Sprintf(`{"name":"%s","url":"%s","domain":"%s"}`, event_type, pa.AppURL, pa.AppDomain))

    req, err := http.NewRequest("POST", pa.PlausibleURL, bytes.NewBuffer(payload))
    if err != nil {
		slog.Error("could not create analytics request: "+ err.Error())
		return
    }

    req.Header.Set("User-Agent", r.UserAgent())
    req.Header.Set("X-Forwarded-For", getIPAddress(r))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    _, err = client.Do(req)
    if err != nil {
		slog.Error("could not do analytics request: "+ err.Error())
		return
    }
}

func getIPAddress(r *http.Request) string {
    // First, check the X-Forwarded-For header
    xForwardedFor := r.Header.Get("X-Forwarded-For")
    if xForwardedFor != "" {
        // The header can contain a comma-separated list of IP addresses. The client's IP is typically the first one.
        ips := strings.Split(xForwardedFor, ",")
        for _, ip := range ips {
            ip = strings.TrimSpace(ip)
            if net.ParseIP(ip) != nil {
                return ip
            }
        }
    }

    // If X-Forwarded-For is not set or doesn't contain a valid IP, fall back to using RemoteAddr
    remoteAddr := strings.Split(r.RemoteAddr, ":")
    if len(remoteAddr) > 0 {
        return remoteAddr[0]
    }

    // Return an empty string if both methods fail
    return ""
}
