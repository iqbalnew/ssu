package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	addonsLogger "bitbucket.bri.co.id/scm/addons/addons-task-service/server/logger"
	"github.com/felixge/httpsnoop"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
)

type generatedCode struct {
	refCode   string
	entryCode string
}

func setHeaderHandler(h http.Handler, sid *shortid.Shortid) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")

		if allowedOrigin(r.Header.Get("Origin")) {
			// w.Header().Set("Content-Security-Policy", "default-src 'self'")
			w.Header().Set("Access-Control-Allow-Origin", strings.Join(config.CorsAllowedOrigins, ", "))
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(config.CorsAllowedMethods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(config.CorsAllowedHeaders, ", "))

			timeCode := time.Now().Format("20060102150405")
			code, err := sid.Generate()
			if err != nil {
				logrus.Errorln("Failed to generate refCode")
			}
			refCode := fmt.Sprintf("HTTP%s%s", timeCode, code)

			w.Header().Set("App-Reference-Code", refCode)
			w.Header().Set("App-Entry-Code", config.ServiceName)

			logrus.Println(w.Header().Get("App-Time-Code"))
			logrus.Println(w.Header().Get("App-Reference-Code"))
			logrus.Println(w.Header().Get("App-Entry-Code"))

		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

type HTTPReqInfo struct {
	// GET etc.
	method  string
	uri     string
	referer string
	ipaddr  string
	// response code, like 200, 404
	code int
	// number of bytes of the response sent
	size int64
	// how long did it take to
	duration  time.Duration
	userAgent string
	refCode   string
	entryCode string
}

func logRequestHandler(h http.Handler, logger *addonsLogger.Logger) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ri := &HTTPReqInfo{
			method:    r.Method,
			uri:       r.URL.String(),
			referer:   r.Header.Get("Referer"),
			userAgent: r.Header.Get("User-Agent"),
		}

		ri.ipaddr = requestGetRemoteAddress(r)

		// this runs handler h and captures information about
		// HTTP request
		m := httpsnoop.CaptureMetrics(h, w, r)

		ri.code = m.Code
		ri.size = m.Written
		ri.duration = m.Duration
		ri.refCode = w.Header().Get("App-Reference-Code")
		ri.entryCode = w.Header().Get("App-Entry-Code")

		data := map[string]interface{}{
			"method":    ri.method,
			"uri":       ri.uri,
			"referer":   ri.referer,
			"ipaddr":    ri.ipaddr,
			"code":      strconv.Itoa(ri.code),
			"size":      strconv.FormatInt(ri.size, 10),
			"duration":  ri.duration.String(),
			"userAgent": ri.userAgent,
			"refCode":   ri.refCode,
			"entryCode": ri.entryCode,
		}

		logger.InfoWithDataMap("Httplog", data)
	}
	return http.HandlerFunc(fn)
}

// Request.RemoteAddress contains port, which we want to remove i.e.:
// "[::1]:58292" => "[::1]"
func ipAddrFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

// requestGetRemoteAddress returns ip address of the client making the request,
// taking into account http proxies
func requestGetRemoteAddress(r *http.Request) string {
	hdr := r.Header
	hdrRealIP := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")
	if hdrRealIP == "" && hdrForwardedFor == "" {
		return ipAddrFromRemoteAddr(r.RemoteAddr)
	}
	if hdrForwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		// TODO: should return first non-local address
		return parts[0]
	}
	return hdrRealIP
}
