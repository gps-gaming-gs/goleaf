package generator

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/gps-gaming-gs/goleaf/rpcx/parser"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed conf.tpl
var confTemplate string

//go:embed json.tpl
var jsonTemplate string

//go:embed conf_server.tpl
var confServerTemplate string

// GenConf generates the configuration structure definition file of the rpc service,
// which contains the zrpc.RpcServerConf configuration item by default.
// You can specify the naming style of the target file name through config.Config. For details,
// see https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/config.go
func (g *Generator) GenConf(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {

	err := g.conf(ctx, cfg)
	if err != nil {
		return err
	}

	err = g.json(ctx, cfg)
	if err != nil {
		return err
	}

	err = g.confServer(ctx, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) conf(ctx DirContext, cfg *conf.Config) error {
	dir := ctx.GetConfig()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "conf")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, confTemplateFile, confTemplate)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, []byte(text), os.ModePerm)
}

func (g *Generator) json(ctx DirContext, cfg *conf.Config) error {
	dir := ctx.GetConfig()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "json")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, jsonTemplateFile, jsonTemplate)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, []byte(text), os.ModePerm)
}

func (g *Generator) confServer(ctx DirContext, cfg *conf.Config) error {

	dir := ctx.GetConf()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "server")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".json")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, confServerTemplateFile, confServerTemplate)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, []byte(text), os.ModePerm)
}
