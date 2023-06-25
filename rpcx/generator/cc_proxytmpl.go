package generator

import (
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"path/filepath"
)

//go:embed cocos/assets/scripts/services/proxy/ccProxyTemplate.tpl
var ccProxyTemplate string

// CCProxyTmpl returns a sample of a proto file
func CCProxyTmpl(ProxyName string) error {
	zctx := &ZRpcContext{}
	zctx.Src = "./"
	zctx.GoOutput = filepath.Dir(zctx.Src)
	zctx.GrpcOutput = filepath.Dir(zctx.Src)

	abs, err := filepath.Abs(zctx.Output)
	if err != nil {
		return err
	}

	projectCtx, err := ctx.Prepare(abs)
	if err != nil {
		return err
	}

	cocosDirCtx, err := mkdircocos(projectCtx, zctx)
	if err != nil {
		return err
	}

	err = genCCProxy(cocosDirCtx, ProxyName)
	if err != nil {
		return err
	}

	return nil
}

func genCCProxy(ctx CocosDirContext, proxyName string) error {
	dir := ctx.GetCCScriptServiceProxy()

	fileName := filepath.Join(dir.Filename, proxyName+"Proxy.ts")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, ccProxyTemplateFile, ccProxyTemplate)
	if err != nil {
		return err
	}

	return util.With("ccProxy").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		"ProxyName": proxyName,
	}, fileName, false)
}
