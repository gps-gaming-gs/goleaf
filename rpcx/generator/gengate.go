package generator

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gps-gaming-gs/goleaf/rpcx/parser"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed gate_external.tpl
var gateExternalTemplate string

//go:embed gate_internal_module.tpl
var gateInternalModuleTemplate string

//go:embed gate_router.tpl
var gateRouterTemplate string

// GenGate generates the configuration structure definition file of the rpc service,
// which contains the zrpc.RpcServerConf configuration item by default.
// You can specify the naming style of the target file name through config.Config. For details,
// see https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/config.go
func (g *Generator) GenGate(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {
	err := g.gateExternal(ctx, cfg)
	if err != nil {
		return err
	}
	err = g.gateModule(ctx, cfg)
	if err != nil {
		return err
	}
	err = g.gateRouter(ctx, proto, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) gateExternal(ctx DirContext, cfg *conf.Config) error {
	dir := ctx.GetGate()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "external")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, gateExternalTemplateFile, gateExternalTemplate)
	if err != nil {
		return err
	}

	imports := make([]string, 0)
	gateImport := fmt.Sprintf(`"%v"`, ctx.GetGateInternal().Package)
	imports = append(imports, gateImport)

	return util.With("external").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports": strings.Join(imports, pathx.NL),
	}, fileName, false)
}

func (g *Generator) gateModule(ctx DirContext, cfg *conf.Config) error {
	dir := ctx.GetGateInternal()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "module")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, gateInternalModuleTemplateFile, gateInternalModuleTemplate)
	if err != nil {
		return err
	}

	imports := make([]string, 0)
	configImport := fmt.Sprintf(`"%v"`, ctx.GetConfig().Package)
	msgImport := fmt.Sprintf(`"%v"`, ctx.GetMsg().Package)
	gameImport := fmt.Sprintf(`"%v"`, ctx.GetGame().Package)
	imports = append(imports, configImport, msgImport, gameImport)

	return util.With("module").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports": strings.Join(imports, pathx.NL),
	}, fileName, false)
}

func (g *Generator) gateRouter(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {
	imports := make([]string, 0)
	alias := collection.NewSet()
	for _, item := range proto.Message {
		msgName := getMessageName(*item.Message)
		alias.AddStr(fmt.Sprintf("msg.Processor.SetRouter(&msg.%s{}, game.ChanRPC)", parser.CamelCase(msgName)))
	}

	gameModuleImport := fmt.Sprintf(`"%v"`, ctx.GetGame().Package)
	msgModuleImport := fmt.Sprintf(`"%v"`, ctx.GetMsg().Package)
	imports = append(imports, gameModuleImport, msgModuleImport)

	dir := ctx.GetGate()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "router")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, gateRouterTemplateFile, gateRouterTemplate)
	if err != nil {
		return err
	}

	aliasKeys := alias.KeysStr()
	sort.Strings(aliasKeys)
	return util.With("main").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports": strings.Join(imports, pathx.NL),
		"routers": strings.Join(aliasKeys, pathx.NL),
	}, fileName, false)
}
