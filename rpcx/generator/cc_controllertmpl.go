package generator

import (
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"path/filepath"
)

//go:embed cocos/assets/scripts/controllers/ccControllerTemplate.tpl
var ccCtrlTemplate string

//go:embed cocos/assets/scripts/views/ccViewUITemplate.tpl
var ccViewTemplate string

// CCControllerTmpl returns a sample of a proto file
func CCControllerTmpl(CtrlName string) error {
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

	err = genCCController(cocosDirCtx, CtrlName)
	if err != nil {
		return err
	}

	err = genCCView(cocosDirCtx, CtrlName)
	if err != nil {
		return err
	}

	return nil
}

func genCCController(ctx CocosDirContext, ctrlName string) error {
	dir := ctx.GetCCScriptController()

	fileName := filepath.Join(dir.Filename, ctrlName+"Controller.ts")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, ccCtrlTemplateFile, ccCtrlTemplate)
	if err != nil {
		return err
	}

	return util.With("ccCtrl").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		"CtrlName": ctrlName,
	}, fileName, false)
}

func genCCView(ctx CocosDirContext, ctrlName string) error {
	dir := ctx.GetCCScriptView()

	fileName := filepath.Join(dir.Filename, ctrlName+"View.ts")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, ccViewTemplateFile, ccViewTemplate)
	if err != nil {
		return err
	}

	return util.With("ccView").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		"CtrlName": ctrlName,
	}, fileName, false)
}
