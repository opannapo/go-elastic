package napoelastic

import "fmt"

const msg = `- Config
	└ Mode			» %s
	└ Server.Address	» %s
	└ Context.Timeout	» %d
	└ Elastic.Address	» %s  
`

type Config struct {
	Mode    string
	Server  *Server
	Context *Context
	Elastic *Elastic
}

type Server struct {
	Address string
}

type Context struct {
	Timeout int
}

type Elastic struct {
	Address []string
}

func (instance *Config) configPrintOut() {
	fmt.Printf(msg,
		instance.Mode,
		instance.Server.Address,
		instance.Context.Timeout,
		instance.Elastic.Address,
	)
}
