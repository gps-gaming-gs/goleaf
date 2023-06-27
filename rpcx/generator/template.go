package generator

import (
	"fmt"

	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const (
	category                          = "rpc"
	skeletonTemplateFile              = "skeleton.tpl"
	confTemplateFile                  = "conf.tpl"
	jsonTemplateFile                  = "json.tpl"
	gameExternalTemplateFile          = "game_external.tpl"
	gameInternalChanRPCTemplateFile   = "game_internal_chanrpc.tpl"
	gameInternalHandlerTemplateFile   = "game_internal_handler.tpl"
	gameInternalModuleTemplateFile    = "game_internal_module.tpl"
	readerTemplateFile                = "reader.tpl"
	gateExternalTemplateFile          = "gate_external.tpl"
	gateRouterTemplateFile            = "gate_router.tpl"
	gateInternalModuleTemplateFile    = "gate_internal_module.tpl"
	loginExternalTemplateFile         = "login_external.tpl"
	loginInternalHandlerTemplateFile  = "login_internal_handler.tpl"
	loginInternalModuleTemplateFile   = "login_internal_module.tpl"
	mainTemplateFile                  = "main.tpl"
	makefileTemplateFile              = "makefile.tpl"
	msgTemplateFile                   = "msg.tpl"
	rpcTemplateFile                   = "template.tpl"
	confServerTemplateFile            = "conf_server.tpl"
	gameInternalLogicTemplateFile     = "game_internal_logic.tpl"
	gameInternalLogicFuncTemplateFile = "game_internal_logic_func.tpl"
	//cc
	ccMakefileTemplateFile               = "ccMakefile.tpl"
	ccGameAppTemplateFile                = "ccGameApp.tpl"
	ccBaseControllerTemplateFile         = "ccBaseController.tpl"
	ccLoginUIControllerTemplateFile      = "ccLoginUIController.tpl"
	ccLibTemplateFile                    = "lib.md"
	ccLibHowlerTemplateFile              = "howler.js"
	ccLibProtobufTemplateFile            = "protobuf.js"
	ccLibProtobufMapTemplateFile         = "protobuf.js.map"
	ccLibProtobufIndexTemplateFile       = "protobuf.d.ts"
	ccLibBezierCurveTemplateFile         = "BezierCurve.ts"
	ccLibBezierCurveAnimTemplateFile     = "BezierCurveAnimation.ts"
	ccLibBezierCurveMoveBaseTemplateFile = "BezierCurveMoveBase.ts"
	ccCmdTypeTemplateFile                = "ccCmdType.tpl"
	ccGameModelTemplateFile              = "ccGameModel.tpl"
	ccAudioConstantTemplateFile          = "ccAudioConstant.tpl"
	ccAudioManagerTemplateFile           = "ccAudioManager.tpl"
	ccBGMAudioManagerTemplateFile        = "ccBGMAudioManager.tpl"
	ccGeneralAudioManagerTemplateFile    = "ccGeneralAudioManager.tpl"
	ccMainGameAudioManagerTemplateFile   = "ccMainGameAudioManager.tpl"
	ccNetEventTemplateFile               = "ccNetEvent.tpl"
	ccNetEventDispatcherTemplateFile     = "ccNetEventDispatcher.tpl"
	ccNetManagerTemplateFile             = "ccNetManager.tpl"
	ccProtoManagerTemplateFile           = "ccProtoManager.tpl"
	ccAuthProxyTemplateFile              = "ccAuthProxy.tpl"
	ccEventManagerTemplateFile           = "ccEventManager.tpl"
	ccResManagerTemplateFile             = "ccResManager.tpl"
	ccTimerManagerTemplateFile           = "ccTimerManager.tpl"
	ccUIManagerTemplateFile              = "ccUIManager.tpl"
	ccBaseUIViewTemplateFile             = "ccBaseUIView.tpl"
	ccLoginUIViewTemplateFile            = "ccLoginUIView.tpl"
	ccUIEventTemplateFile                = "ccUIEvent.tpl"
	ccGameConfigTemplateFile             = "ccGameConfig.tpl"
	ccGameLaunchTemplateFile             = "ccGameLaunch.tpl"
	ccToolPackageTemplateFile            = "ccToolPackage.tpl"
	ccToolWrapPBJSTemplateFile           = "ccToolWrapPBjs.tpl"
	ccToolWrapPBTSTemplateFile           = "ccToolWrapPBts.tpl"
	// cc resource
	ccResourceConfigTemplateFile = "ccApp.tpl"
	// cc template
	ccNewProjectZipFile = "NewProject.zip"
	ccCtrlTemplateFile  = "ccControllerTemplate.tpl"
	ccViewTemplateFile  = "ccViewUITemplate.tpl"
	ccProxyTemplateFile = "ccProxyTemplate.tpl"
	// default folder
	ccDefaultFolderFile = "ccDefaultFolder.tpl"
)

var templates = map[string]string{
	skeletonTemplateFile:             skeletonTemplate,
	confTemplateFile:                 confTemplate,
	jsonTemplateFile:                 jsonTemplate,
	gameExternalTemplateFile:         gameExternalTemplate,
	gameInternalChanRPCTemplateFile:  gameInternalChanRPCTemplate,
	gameInternalHandlerTemplateFile:  gameInternalHandlerTemplate,
	gameInternalModuleTemplateFile:   gameInternalModuleTemplate,
	readerTemplateFile:               readerTemplate,
	gateExternalTemplateFile:         gateExternalTemplate,
	gateInternalModuleTemplateFile:   gateInternalModuleTemplate,
	gateRouterTemplateFile:           gateRouterTemplate,
	loginExternalTemplateFile:        loginExternalTemplate,
	loginInternalModuleTemplateFile:  loginInternalModuleTemplate,
	loginInternalHandlerTemplateFile: loginInternalHandlerTemplate,
	mainTemplateFile:                 mainTemplate,
	makefileTemplateFile:             makefileTemplate,
	msgTemplateFile:                  msgTemplate,
	rpcTemplateFile:                  rpcTemplateText,
	confServerTemplateFile:           confServerTemplate,
	gameInternalLogicTemplateFile:    gameInternalLogicTemplate,
	// cc
	ccMakefileTemplateFile:               ccMakefileTemplate,
	ccGameAppTemplateFile:                ccGameAppTemplate,
	ccBaseControllerTemplateFile:         ccBaseControllerTemplate,
	ccLoginUIControllerTemplateFile:      ccLoginUIControllerTemplate,
	ccLibTemplateFile:                    ccLibTemplate,
	ccLibHowlerTemplateFile:              ccLibHowlerTemplate,
	ccLibProtobufTemplateFile:            ccLibProtobufTemplate,
	ccLibProtobufMapTemplateFile:         ccLibProtobufMapTemplate,
	ccLibProtobufIndexTemplateFile:       ccLibProtobufIndexTemplate,
	ccLibBezierCurveTemplateFile:         ccLibBezierCurveTemplate,
	ccLibBezierCurveAnimTemplateFile:     ccLibBezierCurveAnimTemplate,
	ccLibBezierCurveMoveBaseTemplateFile: ccLibBezierCurveMoveBaseTemplate,

	ccCmdTypeTemplateFile:              ccCmdTypeTemplate,
	ccGameModelTemplateFile:            ccGameModelTemplate,
	ccAudioConstantTemplateFile:        ccAudioConstantTemplate,
	ccAudioManagerTemplateFile:         ccAudioManagerTemplate,
	ccBGMAudioManagerTemplateFile:      ccBGMAudioManagerTemplate,
	ccGeneralAudioManagerTemplateFile:  ccGeneralAudioManagerTemplate,
	ccMainGameAudioManagerTemplateFile: ccMainGameAudioManagerTemplate,
	ccNetEventTemplateFile:             ccNetEventTemplate,
	ccNetEventDispatcherTemplateFile:   ccNetEventDispatcherTemplate,
	ccNetManagerTemplateFile:           ccNetManagerTemplate,
	ccProtoManagerTemplateFile:         ccProtoManagerTemplate,
	ccAuthProxyTemplateFile:            ccAuthProxyTemplate,
	ccEventManagerTemplateFile:         ccEventManagerTemplate,
	ccResManagerTemplateFile:           ccResManagerTemplate,
	ccTimerManagerTemplateFile:         ccTimerManagerTemplate,
	ccUIManagerTemplateFile:            ccUIManagerTemplate,
	ccBaseUIViewTemplateFile:           ccBaseUIViewTemplate,
	ccLoginUIViewTemplateFile:          ccLoginUIViewTemplate,
	ccUIEventTemplateFile:              ccUIEventTemplate,
	ccGameConfigTemplateFile:           ccGameConfigTemplate,
	ccGameLaunchTemplateFile:           ccGameLaunchTemplate,
	ccToolPackageTemplateFile:          ccToolPackageTemplate,
	ccToolWrapPBJSTemplateFile:         ccToolWrapPBJSTemplate,
	ccToolWrapPBTSTemplateFile:         ccToolWrapPBTSTemplate,
	// cc resource
	ccResourceConfigTemplateFile: ccResourceConfigTemplate,
	// cc template
	ccNewProjectZipFile: ccNewProjectTemplateZipFile,
	ccCtrlTemplateFile:  ccCtrlTemplate,
	ccViewTemplateFile:  ccViewTemplate,
	ccProxyTemplateFile: ccProxyTemplate,
	// default folder
	ccDefaultFolderFile: ccDefaultFolder,
}

// GenTemplates is the entry for command goctl template,
// it will create the specified category
func GenTemplates() error {
	return pathx.InitTemplates(category, templates)
}

// RevertTemplate restores the deleted template files
func RevertTemplate(name string) error {
	content, ok := templates[name]
	if !ok {
		return fmt.Errorf("%s: no such file name", name)
	}
	return pathx.CreateTemplate(category, name, content)
}

// Clean deletes all template files
func Clean() error {
	return pathx.Clean(category)
}

// Update is used to update the template files, it will delete the existing old templates at first,
// and then create the latest template files
func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return pathx.InitTemplates(category, templates)
}

// Category returns a const string value for rpc template category
func Category() string {
	return category
}
