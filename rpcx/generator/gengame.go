package generator

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gps-gaming-gs/goleaf/rpcx/parser"
	"github.com/zeromicro/go-zero/core/collection"
	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const (
	gameInternalLogicFuncTemplate = `
		func {{.method}}(args []interface{}) {
			// 收到的消息
			{{if .hasReq}}
			m := args[0].({{.request}})
			{{end}}
			// 消息的发送者
			a := args[1].(gate.Agent)
		
			// 输出收到的消息的内容
			log.Debug("{{.logicName}} %v", m)
		
			// 给发送者回应一个消息
			{{if .hasReply}}
			a.WriteMsg({{.responseType}}{
				//Code: "200",
			})
			{{end}}
		}`
)

//go:embed game_external.tpl
var gameExternalTemplate string

//go:embed game_internal_chanrpc.tpl
var gameInternalChanRPCTemplate string

//go:embed game_internal_handler.tpl
var gameInternalHandlerTemplate string

//go:embed game_internal_module.tpl
var gameInternalModuleTemplate string

//go:embed game_internal_logic.tpl
var gameInternalLogicTemplate string

// GenGame generates the configuration structure definition file of the rpc service,
// which contains the zrpc.RpcServerConf configuration item by default.
// You can specify the naming style of the target file name through config.Config. For details,
// see https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/config.go
func (g *Generator) GenGame(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {

	err := g.gameExternal(ctx, cfg)
	if err != nil {
		return err
	}

	err = g.gameChanRPC(ctx, cfg)
	if err != nil {
		return err
	}

	err = g.gameHandler(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.gameModule(ctx, cfg)
	if err != nil {
		return err
	}

	err = g.gameLogic(ctx, proto, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) gameExternal(ctx DirContext, cfg *conf.Config) error {
	dir := ctx.GetGame()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "external")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, gameExternalTemplateFile, gameExternalTemplate)
	if err != nil {
		return err
	}

	//return os.WriteFile(fileName, []byte(text), os.ModePerm)
	imports := make([]string, 0)
	gameImport := fmt.Sprintf(`"%v"`, ctx.GetGameInternal().Package)
	imports = append(imports, gameImport)

	return util.With("external").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports": strings.Join(imports, pathx.NL),
	}, fileName, false)
}

func (g *Generator) gameChanRPC(ctx DirContext, cfg *conf.Config) error {
	dir := ctx.GetGameInternal()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "chanrpc")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, gameInternalChanRPCTemplateFile, gameInternalChanRPCTemplate)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, []byte(text), os.ModePerm)
}

func (g *Generator) gameHandler(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {
	imports := make([]string, 0)
	alias := collection.NewSet()
	for _, rpc := range proto.Service[0].RPC {
		alias.AddStr(fmt.Sprintf("handler(%s{}, logic.%s)",
			fmt.Sprintf("%s.%s", proto.Package.Name, parser.CamelCase(rpc.RequestType)),
			parser.CamelCase(rpc.Name)))
	}

	gameInternalLogicImport := fmt.Sprintf(`"%v"`, ctx.GetGameInternalLogic().Package)
	msgModuleImport := fmt.Sprintf(`"%v"`, ctx.GetMsg().Package)
	imports = append(imports, gameInternalLogicImport, msgModuleImport, `"reflect"`)

	dir := ctx.GetGameInternal()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "handler")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		os.Remove(fileName)
	}

	text, err := pathx.LoadTemplate(category, gameInternalHandlerTemplateFile, gameInternalHandlerTemplate)
	if err != nil {
		return err
	}

	aliasKeys := alias.KeysStr()
	sort.Strings(aliasKeys)
	return util.With("gameInternalHandler").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports":  strings.Join(imports, pathx.NL),
		"handlers": strings.Join(aliasKeys, pathx.NL),
	}, fileName, false)
}

func (g *Generator) gameModule(ctx DirContext, cfg *conf.Config) error {
	dir := ctx.GetGameInternal()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "module")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, gameInternalModuleTemplateFile, gameInternalModuleTemplate)
	if err != nil {
		return err
	}

	imports := make([]string, 0)
	gameImport := fmt.Sprintf(`"%v"`, ctx.GetBase().Package)
	imports = append(imports, gameImport)

	return util.With("gameInternalModule").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports": strings.Join(imports, pathx.NL),
	}, fileName, false)
}

func (g *Generator) gameLogic(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetGameInternalLogic()
	for _, item := range proto.Service {
		serviceName := item.Name
		for _, rpc := range item.RPC {
			var (
				err           error
				filename      string
				logicName     string
				logicFilename string
				packageName   string
			)

			logicName = fmt.Sprintf("%sLogic", stringx.From(rpc.Name).ToCamel())

			nameJoin := fmt.Sprintf("%s_logic", serviceName)
			packageName = strings.ToLower(stringx.From(nameJoin).ToCamel())
			logicFilename, err = format.FileNamingFormat(cfg.NamingFormat, rpc.Name+"_logic")
			if err != nil {
				console.Info("FileNamingFormat")
				return err
			}

			filename = filepath.Join(dir.Filename, logicFilename+".go")
			functions, err := g.genLogicFunction(serviceName, proto.Package.Name, logicName, rpc)
			if err != nil {
				console.Info("genLogicFunction")
				return err
			}

			imports := collection.NewSet()
			imports.AddStr(fmt.Sprintf(`"%v"`, ctx.GetPb().Package))
			text, err := pathx.LoadTemplate(category, gameInternalLogicTemplateFile, gameInternalLogicTemplate)
			if err != nil {
				console.Info("LoadTemplate")
				return err
			}

			if err = util.With("gameInternalLogic").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
				"logicName":   logicName,
				"functions":   functions,
				"packageName": packageName,
				"imports":     strings.Join(imports.KeysStr(), pathx.NL),
			}, filename, false); err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Generator) genLogicFunction(serviceName string, goPackage, logicName string, rpc *parser.RPC) (string,
	error) {
	functions := make([]string, 0)
	text, err := pathx.LoadTemplate(category, gameInternalLogicFuncTemplateFile, gameInternalLogicFuncTemplate)
	if err != nil {
		return "", err
	}
	comment := parser.GetComment(rpc.Doc())
	streamServer := fmt.Sprintf("%s.%s_%s%s", goPackage, parser.CamelCase(serviceName),
		parser.CamelCase(rpc.Name), "Server")
	buffer, err := util.With("fun").Parse(text).Execute(map[string]interface{}{
		"logicName":    logicName,
		"method":       parser.CamelCase(rpc.Name),
		"hasReq":       !rpc.StreamsRequest,
		"request":      fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.RequestType)),
		"hasReply":     !rpc.StreamsRequest && !rpc.StreamsReturns,
		"response":     fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.ReturnsType)),
		"responseType": fmt.Sprintf("&%s.%s", goPackage, parser.CamelCase(rpc.ReturnsType)),
		"stream":       rpc.StreamsRequest || rpc.StreamsReturns,
		"streamBody":   streamServer,
		"hasComment":   len(comment) > 0,
		"comment":      comment,
	})
	if err != nil {
		return "", err
	}

	functions = append(functions, buffer.String())
	return strings.Join(functions, pathx.NL), nil
}

func titleCase(str string) string {
	capitalized := strings.Title(str)
	return strings.ReplaceAll(capitalized, " ", "")
}
