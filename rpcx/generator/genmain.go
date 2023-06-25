package generator

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"

	"github.com/gps-gaming-gs/goleaf/rpcx/parser"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed main.tpl
var mainTemplate string

//go:embed makefile.tpl
var makefileTemplate string

type MainServiceTemplateData struct {
	Service   string
	ServerPkg string
	Pkg       string
}

// GenMain generates the main file of the rpc service, which is an rpc service program call entry
func (g *Generator) GenMain(ctx DirContext, proto parser.Proto, cfg *conf.Config,
	c *ZRpcContext) error {
	//mainFilename, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	//if err != nil {
	//	return err
	//}

	fileName := filepath.Join(ctx.GetMain().Filename, fmt.Sprintf("%v.go", "main"))
	imports := make([]string, 0)
	//pbImport := fmt.Sprintf(`"%v"`, ctx.GetPb().Package)
	//svcImport := fmt.Sprintf(`"%v"`, ctx.GetSvc().Package)
	configImport := fmt.Sprintf(`"%v"`, ctx.GetConfig().Package)
	gameModuleImport := fmt.Sprintf(`"%v"`, ctx.GetGame().Package)
	gateModuleImport := fmt.Sprintf(`"%v"`, ctx.GetGate().Package)
	loginModuleImport := fmt.Sprintf(`"%v"`, ctx.GetLogin().Package)
	imports = append(imports, configImport, gameModuleImport, gateModuleImport, loginModuleImport)

	var serviceNames []MainServiceTemplateData
	for _, e := range proto.Service {
		var (
			remoteImport string
			serverPkg    string
		)
		//if !c.Multiple {
		//	serverPkg = "server"
		//	remoteImport = fmt.Sprintf(`"%v"`, ctx.GetServer().Package)
		//} else {
		//	childPkg, err := ctx.GetServer().GetChildPackage(e.Name)
		//	if err != nil {
		//		return err
		//	}
		//
		//	serverPkg = filepath.Base(childPkg + "Server")
		//	remoteImport = fmt.Sprintf(`%s "%v"`, serverPkg, childPkg)
		//}
		imports = append(imports, remoteImport)
		serviceNames = append(serviceNames, MainServiceTemplateData{
			Service:   parser.CamelCase(e.Name),
			ServerPkg: serverPkg,
			Pkg:       proto.PbPackage,
		})
	}

	text, err := pathx.LoadTemplate(category, mainTemplateFile, mainTemplate)
	if err != nil {
		return err
	}

	etcFileName, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	if err != nil {
		return err
	}

	return util.With("main").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"serviceName":  etcFileName,
		"imports":      strings.Join(imports, pathx.NL),
		"pkg":          proto.PbPackage,
		"serviceNames": serviceNames,
	}, fileName, false)
}

// GenMakefile generates the yaml configuration file of the rpc service,
// including host, port monitoring configuration items and etcd configuration
func (g *Generator) GenMakefile(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	fileName := filepath.Join(ctx.GetMain().Filename, "Makefile")

	text, err := pathx.LoadTemplate(category, makefileTemplateFile, makefileTemplate)
	if err != nil {
		return err
	}

	return util.With("makefile").Parse(text).SaveTo(map[string]interface{}{
		"serviceName": strings.ToLower(stringx.From(ctx.GetServiceName().Source()).ToCamel()),
	}, fileName, false)
}
