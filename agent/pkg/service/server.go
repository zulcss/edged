package service

import (
//	"fmt"
	"log"
//	"net"

//	"google.golang.org/grpc"
        "github.com/spf13/viper"

//        pb "github.com/zulcss/edged/proto"
)

func Run() {
	host := viper.GetString("agent.host")
	log.Printf("Starting server on %s", host)
}
