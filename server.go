package main

import (
	"fmt"
	"log"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/pborman/uuid"
	"github.com/valyala/fasthttp"
)

func redirectHandler(c *fasthttp.RequestCtx) {
	// Get referer
	referer := c.Referer()
	log.Printf("redirect from %s", string(referer))

	// Set tracking cookie
	cookie := fasthttp.Cookie{}
	cookie.SetKey("ADUID")
	cookie.SetValue(uuid.NewRandom().String())
	cookie.SetExpire(time.Now().AddDate(1, 0, 0))
	cookie.SetDomain(".ad-vendor1.com")
	cookie.SetPath("/")
	c.Response.Header.SetCookie(&cookie)

	// Redirect to designated page
	redirectTo := string(c.QueryArgs().Peek("to"))
	if redirectTo != "" {
		// log.Printf("redirect to %s", redirectTo)
		c.Redirect(redirectTo, 302) //use 302 to avoid browser cache
	} else {
		fmt.Fprintf(c, "error. not redirect")
		c.SetStatusCode(fasthttp.StatusNotFound)
	}
}

func retargetHandler(c *fasthttp.RequestCtx) {
	// Get referer
	referer := c.Referer()
	log.Printf("user views: %s ", string(referer))
	c.SetStatusCode(fasthttp.StatusOK)
}

func main() {
	router := fasthttprouter.New()

	router.GET("/redirect", redirectHandler)
	router.GET("/retarget", retargetHandler)

	log.Printf("Start HTTP server")
	panic(fasthttp.ListenAndServe(":8888", router.Handler))
}
