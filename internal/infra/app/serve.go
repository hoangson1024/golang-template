package app

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"go.opencensus.io/plugin/ochttp"
)

func (a *ApplicationContext) Serve() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "API",
		Action: func(c *cli.Context) error {

			r := gin.Default()
			r.Use(healthcheck.Default())
			r.Use(Logger())

			// New group v1
			v1 := r.Group("/v1")
			v1.POST("/test", func(cc *gin.Context) {
				res, _ := a.wrapper.Load(cc)
				cc.JSON(http.StatusOK, res)
			})

			// Listen and Server in 0.0.0.0:80
			http.ListenAndServe(
				fmt.Sprintf(":%d", a.cfg.Port),
				&ochttp.Handler{
					Handler: r,
				},
			)
			return nil
		},
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before Request
		t := time.Now()

		// Resquest
		bufRes, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(bufRes))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(bufRes)) //We have to create a new Buffer, because rdr1 will be read.

		// Response
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Request.Body = rdr2

		c.Next()

		// after request
		latency := time.Since(t)
		status := c.Writer.Status()

		log.Printf("| %d | %s | %s | %s | %s \nResquest body: %s\nResponse body:%s",
			status,
			latency,
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL,
			readBody(rdr1),
			blw.body.String(),
		)
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
