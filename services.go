package main

import (
    "flag"
    "fmt"
    discoverd "github.com/flynn/go-discoverd"
    "log"
)

type services struct {
    clientCmd
    onlyOne   *bool
    printAttr *bool
}

func (cmd *services) Name() string {
    return "services"
}

func (cmd *services) DefineFlags(fs *flag.FlagSet) {
    cmd.onlyOne = fs.Bool("1", false, "only show one service")
    cmd.printAttr = fs.Bool("a", false, "output attributes if any")
}

func (cmd *services) Run(fs *flag.FlagSet) {
    cmd.InitClient(false)
    services, err := cmd.client.Services(fs.Arg(0), discoverd.DefaultTimeout)
    if err != nil {
        log.Fatal(err)
    }
    if *cmd.onlyOne {
        if len(services) > 0 {
            cmd.PrintService(services[0])
        }
        return
    }
    for _, service := range services {
        cmd.PrintService(service)
    }
}

func (cmd *services) PrintService(srv *discoverd.Service) {
    fmt.Printf("%s", srv.Addr)
    fmt.Printf(";name=%s:host=%s:port=%s", srv.Name, srv.Host, srv.Port)
    if *cmd.printAttr {
        fmt.Printf(";")
        for k, v := range srv.Attrs {
            fmt.Printf("%s=%s", k, v)
        }
    }
    fmt.Printf("\n")

}
