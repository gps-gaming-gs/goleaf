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

//go:embed reader.tpl
var readerTemplate string

// GenGameData generates the configuration structure definition file of the rpc service,
// which contains the zrpc.RpcServerConf configuration item by default.
// You can specify the naming style of the target file name through config.Config. For details,
// see https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/config.go
func (g *Generator) GenGameData(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetGameData()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "reader")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, readerTemplateFile, readerTemplate)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, []byte(text), os.ModePerm)
}
