package errors

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func MW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		start := time.Now()
		if err = next(c); err != nil {
			c.Error(err)
		}
		stop := time.Now()
		buf := config.pool.Get().(*bytes.Buffer)
		buf.Reset()
		defer config.pool.Put(buf)

		if _, err = config.template.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
			switch tag {
			case "custom":
				if config.CustomTagFunc == nil {
					return 0, nil
				}
				return config.CustomTagFunc(c, buf)
			case "time_unix":
				return buf.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
			case "time_unix_milli":
				// go 1.17 or later, it supports time#UnixMilli()
				return buf.WriteString(strconv.FormatInt(time.Now().UnixNano()/1000000, 10))
			case "time_unix_micro":
				// go 1.17 or later, it supports time#UnixMicro()
				return buf.WriteString(strconv.FormatInt(time.Now().UnixNano()/1000, 10))
			case "time_unix_nano":
				return buf.WriteString(strconv.FormatInt(time.Now().UnixNano(), 10))
			case "time_rfc3339":
				return buf.WriteString(time.Now().Format(time.RFC3339))
			case "time_rfc3339_nano":
				return buf.WriteString(time.Now().Format(time.RFC3339Nano))
			case "time_custom":
				return buf.WriteString(time.Now().Format(config.CustomTimeFormat))
			case "id":
				id := req.Header.Get(echo.HeaderXRequestID)
				if id == "" {
					id = res.Header().Get(echo.HeaderXRequestID)
				}
				return buf.WriteString(id)
			case "remote_ip":
				return buf.WriteString(c.RealIP())
			case "host":
				return buf.WriteString(req.Host)
			case "uri":
				return buf.WriteString(req.RequestURI)
			case "method":
				return buf.WriteString(req.Method)
			case "path":
				p := req.URL.Path
				if p == "" {
					p = "/"
				}
				return buf.WriteString(p)
			case "route":
				return buf.WriteString(c.Path())
			case "protocol":
				return buf.WriteString(req.Proto)
			case "referer":
				return buf.WriteString(req.Referer())
			case "user_agent":
				return buf.WriteString(req.UserAgent())
			case "status":
				n := res.Status
				s := config.colorer.Green(n)
				switch {
				case n >= 500:
					s = config.colorer.Red(n)
				case n >= 400:
					s = config.colorer.Yellow(n)
				case n >= 300:
					s = config.colorer.Cyan(n)
				}
				return buf.WriteString(s)
			case "error":
				if err != nil {
					// Error may contain invalid JSON e.g. `"`
					b, _ := json.Marshal(err.Error())
					b = b[1 : len(b)-1]
					return buf.Write(b)
				}
			case "latency":
				l := stop.Sub(start)
				return buf.WriteString(strconv.FormatInt(int64(l), 10))
			case "latency_human":
				return buf.WriteString(stop.Sub(start).String())
			case "bytes_in":
				cl := req.Header.Get(echo.HeaderContentLength)
				if cl == "" {
					cl = "0"
				}
				return buf.WriteString(cl)
			case "bytes_out":
				return buf.WriteString(strconv.FormatInt(res.Size, 10))
			default:
				switch {
				case strings.HasPrefix(tag, "header:"):
					return buf.Write([]byte(c.Request().Header.Get(tag[7:])))
				case strings.HasPrefix(tag, "query:"):
					return buf.Write([]byte(c.QueryParam(tag[6:])))
				case strings.HasPrefix(tag, "form:"):
					return buf.Write([]byte(c.FormValue(tag[5:])))
				case strings.HasPrefix(tag, "cookie:"):
					cookie, err := c.Cookie(tag[7:])
					if err == nil {
						return buf.Write([]byte(cookie.Value))
					}
				}
			}
			return 0, nil
		}); err != nil {
			return
		}

		if config.Output == nil {
			_, err = c.Logger().Output().Write(buf.Bytes())
			return
		}
		_, err = config.Output.Write(buf.Bytes())
		return
	}
}
