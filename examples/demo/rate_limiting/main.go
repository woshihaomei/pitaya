package main

import (
	"flag"
	"fmt"
	"time"

	"strings"

	"github.com/spf13/viper"
	"github.com/woshihaomei/pitaya"
	"github.com/woshihaomei/pitaya/acceptor"
	"github.com/woshihaomei/pitaya/acceptorwrapper"
	"github.com/woshihaomei/pitaya/component"
	"github.com/woshihaomei/pitaya/config"
	"github.com/woshihaomei/pitaya/examples/demo/rate_limiting/services"
	"github.com/woshihaomei/pitaya/serialize/json"
)

func configureFrontend(port int) {
	room := services.NewRoom()
	pitaya.Register(room,
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower))

	// 5 requests in 1 minute. Doesn't make sense, just to test
	// rate limiting
	vConfig := viper.New()
	vConfig.Set("pitaya.conn.ratelimiting.limit", 5)
	vConfig.Set("pitaya.conn.ratelimiting.interval", time.Minute)
	pConfig := config.NewConfig(vConfig)

	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", port))
	wrapped := acceptorwrapper.WithWrappers(
		tcp,
		acceptorwrapper.NewRateLimitingWrapper(pConfig))
	pitaya.AddAcceptor(wrapped)
}

func main() {
	defer pitaya.Shutdown()

	port := flag.Int("port", 3250, "the port to listen")
	svType := "room"

	flag.Parse()

	pitaya.SetSerializer(json.NewSerializer())
	configureFrontend(*port)

	pitaya.Configure(true, svType, pitaya.Cluster, map[string]string{})
	pitaya.Start()
}
