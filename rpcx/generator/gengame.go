package generator

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gps-gaming-gs/goleaf/rpcx/parser"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const (
	gameInternalLogicFuncTemplate = `
		func {{.logicName}}(args []interface{}) {
			// 收到的消息
			m := args[0].(*msg.{{.reqType}})
			// 消息的发送者
			a := args[1].(gate.Agent)
		
			// 输出收到的消息的内容
			log.Debug("{{.logicName}} %v", m)
		
			// 给发送者回应一个消息
			a.WriteMsg(&msg.{{.respType}}{
				//Code: "200",
			})
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
	for _, item := range proto.Message {
		msgName := getMessageName(*item.Message)
		if strings.Contains(strings.ToLower(msgName), "req") == false {
			continue
		}
		logicName := fmt.Sprintf("%sLogic", stringx.From(msgName).ToCamel())
		alias.AddStr(fmt.Sprintf("handler(&msg.%s{}, logic.%s)",
			parser.CamelCase(msgName), logicName))
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
		return nil
	}

	text, err := pathx.LoadTemplate(category, gameInternalHandlerTemplateFile, gameInternalHandlerTemplate)
	if err != nil {
		return err
	}

	aliasKeys := alias.KeysStr()
	sort.Strings(aliasKeys)
	return util.With("module").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
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

	return util.With("module").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports": strings.Join(imports, pathx.NL),
	}, fileName, false)
}

func (g *Generator) gameLogic(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetGameInternalLogic()
	service := proto.Service[0].Service.Name
	for _, message := range proto.Message {
		if strings.Contains(strings.ToLower(message.Name), "req") == false {
			continue
		}
		logicName := fmt.Sprintf("%sLogic", stringx.From(message.Name).ToCamel())
		logicFilename, err := format.FileNamingFormat(cfg.NamingFormat, message.Name+"_logic")
		if err != nil {
			return err
		}

		filename := filepath.Join(dir.Filename, logicFilename+".go")
		functions, err := g.genLogicFunction(service, proto.PbPackage, logicName, message)
		if err != nil {
			return err
		}

		imports := collection.NewSet()
		imports.AddStr(fmt.Sprintf(`"%v"`, ctx.GetPb().Package))
		text, err := pathx.LoadTemplate(category, gameInternalLogicTemplateFile, gameInternalLogicTemplate)
		if err != nil {
			return err
		}
		err = util.With("logic").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
			"logicName":   fmt.Sprintf("%sLogic", stringx.From(message.Name).ToCamel()),
			"functions":   functions,
			"packageName": "logic",
			"imports":     strings.Join(imports.KeysStr(), pathx.NL),
		}, filename, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) genLogicFunction(serviceName, goPackage, logicName string,
	message parser.Message) (string,
	error) {
	functions := make([]string, 0)
	text, err := pathx.LoadTemplate(category, gameInternalLogicFuncTemplateFile, gameInternalLogicFuncTemplate)
	if err != nil {
		return "", err
	}
	typeName := strings.ReplaceAll(strings.ToLower(message.Name), "req", "")
	comment := parser.GetComment(message.Doc())
	streamServer := fmt.Sprintf("%s.%s_%s%s", goPackage, parser.CamelCase(serviceName),
		parser.CamelCase(message.Name), "Server")
	buffer, err := util.With("fun").Parse(text).Execute(map[string]interface{}{
		"logicName":  logicName,
		"streamBody": streamServer,
		"hasComment": len(comment) > 0,
		"comment":    comment,
		"reqType":    parser.CamelCase(message.Name),
		"respType":   parser.CamelCase(typeName) + "Resp",
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
