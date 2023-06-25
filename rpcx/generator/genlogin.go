package generator

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gps-gaming-gs/goleaf/rpcx/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed login_external.tpl
var loginExternalTemplate string

//go:embed login_internal_handler.tpl
var loginInternalHandlerTemplate string

//go:embed login_internal_module.tpl
var loginInternalModuleTemplate string

// GenLogin generates the configuration structure definition file of the rpc service,
// which contains the zrpc.RpcServerConf configuration item by default.
// You can specify the naming style of the target file name through config.Config. For details,
// see https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/config.go
func (g *Generator) GenLogin(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	err := g.genExternal(ctx, cfg)
	if err != nil {
		return err
	}

	err = g.genHandler(ctx, cfg)
	if err != nil {
		return err
	}

	err = g.genModule(ctx, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) genExternal(ctx DirContext, cfg *conf.Config) error {
	dir := ctx.GetLogin()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "external")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, loginExternalTemplateFile, loginExternalTemplate)
	if err != nil {
		return err
	}

	imports := make([]string, 0)
	loginImport := fmt.Sprintf(`"%v"`, ctx.GetLoginInternal().Package)
	imports = append(imports, loginImport)

	return util.With("external").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports": strings.Join(imports, pathx.NL),
	}, fileName, false)
}

func (g *Generator) genHandler(ctx DirContext, cfg *conf.Config) error {
	dir := ctx.GetLoginInternal()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "handler")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, loginInternalHandlerTemplateFile, loginInternalHandlerTemplate)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, []byte(text), os.ModePerm)
}

func (g *Generator) genModule(ctx DirContext, cfg *conf.Config) error {
	dir := ctx.GetLoginInternal()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "module")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, loginInternalModuleTemplateFile, loginInternalModuleTemplate)
	if err != nil {
		return err
	}

	imports := make([]string, 0)
	baseImport := fmt.Sprintf(`"%v"`, ctx.GetBase().Package)
	imports = append(imports, baseImport)

	return util.With("module").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports": strings.Join(imports, pathx.NL),
	}, fileName, false)
}
