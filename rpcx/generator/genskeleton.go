package generator

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gps-gaming-gs/goleaf/rpcx/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed skeleton.tpl
var skeletonTemplate string

// GenSkeleton generates the configuration structure definition file of the rpc service,
// which contains the zrpc.RpcServerConf configuration item by default.
// You can specify the naming style of the target file name through config.Config. For details,
// see https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/config.go
func (g *Generator) GenSkeleton(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetBase()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "skeleton")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, skeletonTemplateFile, skeletonTemplate)
	if err != nil {
		return err
	}

	imports := make([]string, 0)
	configImport := fmt.Sprintf(`"%v"`, ctx.GetConfig().Package)
	imports = append(imports, configImport)

	return util.With("skeleton").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports": strings.Join(imports, pathx.NL),
	}, fileName, false)
}
