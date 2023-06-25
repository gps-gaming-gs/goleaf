package generator

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/emicklei/proto"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util"

	"github.com/gps-gaming-gs/goleaf/rpcx/parser"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed msg.tpl
var msgTemplate string

// GenMsg generates the configuration structure definition file of the rpc service,
// which contains the zrpc.RpcServerConf configuration item by default.
// You can specify the naming style of the target file name through config.Config. For details,
// see https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/config.go
func (g *Generator) GenMsg(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {

	err := g.msg(ctx, proto, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) msg(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {

	alias := collection.NewSet()
	for _, item := range proto.Message {
		msgName := getMessageName(*item.Message)
		alias.AddStr(fmt.Sprintf("Processor.Register(&%s{})", parser.CamelCase(msgName)))
	}

	dir := ctx.GetMsg()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "msg")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		os.Remove(fileName)
	}

	text, err := pathx.LoadTemplate(category, msgTemplateFile, msgTemplate)
	if err != nil {
		return err
	}

	aliasKeys := alias.KeysStr()
	sort.Strings(aliasKeys)
	if err = util.With("msg").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"messages": strings.Join(aliasKeys, pathx.NL),
	}, fileName, true); err != nil {
		return err
	}

	return nil
}

func getMessageName(msg proto.Message) string {
	list := []string{msg.Name}

	for {
		parent := msg.Parent
		if parent == nil {
			break
		}

		parentMsg, ok := parent.(*proto.Message)
		if !ok {
			break
		}

		tmp := []string{parentMsg.Name}
		list = append(tmp, list...)
		msg = *parentMsg
	}

	return strings.Join(list, "_")
}
