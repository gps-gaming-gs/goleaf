package generator

import (
	"archive/zip"
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gps-gaming-gs/goleaf/rpcx/parser"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed cocos/NewProject.zip
var ccNewProjectTemplateZipFile string

//go:embed cocos/NewProject.zip
var ccNewProjectTemplate []byte

//go:embed cocos/ccDefaultFolder.tpl
var ccDefaultFolder string

//go:embed cocos/assets/scripts/bootstrap/ccGameApp.tpl
var ccGameAppTemplate string

//go:embed cocos/assets/scripts/controllers/ccBaseController.tpl
var ccBaseControllerTemplate string

//go:embed cocos/assets/scripts/controllers/ccLoginUIController.tpl
var ccLoginUIControllerTemplate string

//go:embed cocos/assets/scripts/lib/lib.md
var ccLibTemplate string

//go:embed cocos/assets/scripts/lib/howler.js
var ccLibHowlerTemplate string

//go:embed cocos/assets/scripts/lib/protobuf/protobuf.js
var ccLibProtobufTemplate string

//go:embed cocos/assets/scripts/lib/protobuf/protobuf.js.map
var ccLibProtobufMapTemplate string

//go:embed cocos/assets/scripts/lib/protobuf/protobuf.d.ts
var ccLibProtobufIndexTemplate string

//go:embed cocos/assets/scripts/models/ccCmdType.tpl
var ccCmdTypeTemplate string

//go:embed cocos/assets/scripts/models/ccGameModel.tpl
var ccGameModelTemplate string

//go:embed cocos/assets/scripts/services/audio/ccAudioConstant.tpl
var ccAudioConstantTemplate string

//go:embed cocos/assets/scripts/services/audio/ccAudioManager.tpl
var ccAudioManagerTemplate string

//go:embed cocos/assets/scripts/services/audio/ccBGMAudioManager.tpl
var ccBGMAudioManagerTemplate string

//go:embed cocos/assets/scripts/services/audio/ccGeneralAudioManager.tpl
var ccGeneralAudioManagerTemplate string

//go:embed cocos/assets/scripts/services/audio/ccMainGameAudioManager.tpl
var ccMainGameAudioManagerTemplate string

//go:embed cocos/assets/scripts/services/net/ccNetEvent.tpl
var ccNetEventTemplate string

//go:embed cocos/assets/scripts/services/net/ccNetEventDispatcher.tpl
var ccNetEventDispatcherTemplate string

//go:embed cocos/assets/scripts/services/net/ccNetManager.tpl
var ccNetManagerTemplate string

//go:embed cocos/assets/scripts/services/net/ccProtoManager.tpl
var ccProtoManagerTemplate string

//go:embed cocos/assets/scripts/services/proxy/ccAuthProxy.tpl
var ccAuthProxyTemplate string

//go:embed cocos/assets/scripts/services/ccEventManager.tpl
var ccEventManagerTemplate string

//go:embed cocos/assets/scripts/services/ccResManager.tpl
var ccResManagerTemplate string

//go:embed cocos/assets/scripts/services/ccTimerManager.tpl
var ccTimerManagerTemplate string

//go:embed cocos/assets/scripts/services/ccUIManager.tpl
var ccUIManagerTemplate string

//go:embed cocos/assets/scripts/views/ccBaseUIView.tpl
var ccBaseUIViewTemplate string

//go:embed cocos/assets/scripts/views/ccLoginUIView.tpl
var ccLoginUIViewTemplate string

//go:embed cocos/assets/scripts/views/ccUIEvent.tpl
var ccUIEventTemplate string

//go:embed cocos/assets/scripts/ccGameConfig.tpl
var ccGameConfigTemplate string

//go:embed cocos/assets/scripts/ccGameLaunch.tpl
var ccGameLaunchTemplate string

//go:embed cocos/tools/ccToolPackage.tpl
var ccToolPackageTemplate string

//go:embed cocos/tools/ccToolWrapPBjs.tpl
var ccToolWrapPBJSTemplate string

//go:embed cocos/tools/ccToolWrapPBts.tpl
var ccToolWrapPBTSTemplate string

//go:embed cocos/ccMakefile.tpl
var ccMakefileTemplate string

//go:embed cocos/resources/config/ccApp.tpl
var ccResourceConfigTemplate string

// GenCocos generates the configuration structure definition file of the rpc service,
// which contains the zrpc.RpcServerConf configuration item by default.
// You can specify the naming style of the target file name through config.Config. For details,
// see https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/config.go
func (g *Generator) GenCocos(ctx CocosDirContext, proto parser.Proto, cfg *conf.Config) error {
	err := g.ccNewProject(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccDefaultFolders(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccMakefile(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccMain(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccBootstrap(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccController(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccLib(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccModel(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccService(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccView(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccTool(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccResource(ctx, proto, cfg)
	if err != nil {
		return err
	}

	err = g.ccCopyProtoFile(ctx, proto, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) ccDefaultFolders(ctx CocosDirContext, _ parser.Proto, _ *conf.Config) error {
	text, _ := pathx.LoadTemplate(category, ccDefaultFolderFile, ccDefaultFolder)
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		err := os.MkdirAll(filepath.Join(ctx.GetCCAsset().Filename, scanner.Text()), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) ccNewProject(ctx CocosDirContext, _ parser.Proto, _ *conf.Config) error {
	dir := ctx.GetCocos()

	// 打开嵌入的 Zip 文件
	zipReader, err := zip.NewReader(bytes.NewReader(ccNewProjectTemplate), int64(len(ccNewProjectTemplate)))
	if err != nil {
		log.Fatal(err)
	}

	// 遍历 Zip 文件中的每个文件/目录
	for _, file := range zipReader.File {
		// 构建新的输出路径，将旧的文件夹名替换为新的名称
		outputPath := strings.Replace(file.Name, "NewProject", dir.Base, 1)
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(outputPath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}

		outputFile, err := os.Create(outputPath)
		if err != nil {
			log.Fatal(err)
		}
		defer outputFile.Close()

		// 打开 Zip 文件中的文件
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		// 将 Zip 文件中的内容拷贝到输出文件
		_, err = io.Copy(outputFile, zippedFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (g *Generator) ccMain(ctx CocosDirContext, proto parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetCCScript()

	// GameConfig.ts
	fileName := filepath.Join(dir.Filename, "GameConfig.ts")
	if pathx.FileExists(fileName) {
		return nil
	}
	text, err := pathx.LoadTemplate(category, ccGameConfigTemplateFile, ccGameConfigTemplate)
	if err != nil {
		return err
	}
	if err = util.With("ccGameConfig").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		//
	}, fileName, true); err != nil {
		return err
	}

	// GameLaunch.ts
	fileName = filepath.Join(dir.Filename, "GameLaunch.ts")
	if pathx.FileExists(fileName) {
		return nil
	}
	text, err = pathx.LoadTemplate(category, ccGameLaunchTemplateFile, ccGameLaunchTemplate)
	if err != nil {
		return err
	}
	if err = util.With("ccGameLaunch").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		//
	}, fileName, true); err != nil {
		return err
	}

	return nil
}

func (g *Generator) ccLib(ctx CocosDirContext, proto parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetCCScriptLib()

	// Lib.md
	Lib := func() error {
		fileName := filepath.Join(dir.Filename, "lib.md")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccLibTemplateFile, ccLibTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccLib").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			//
		}, fileName, true); err != nil {
			return err
		}
		return nil
	}

	// Howler.js
	Howler := func() error {
		fileName := filepath.Join(dir.Filename, "howler.js")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccLibHowlerTemplateFile, ccLibHowlerTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccLibHowler").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			//
		}, fileName, true); err != nil {
			return err
		}
		return nil
	}

	Protobuf := func() error {
		dir := ctx.GetCCScriptLibProtobuf()
		// Protobuf.js
		ProtobufJS := func() error {
			fileName := filepath.Join(dir.Filename, "protobuf.js")
			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccLibProtobufTemplateFile, ccLibProtobufTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccLibProtobuf").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}

		// Protobuf.js.map
		ProtobufJSMap := func() error {
			fileName := filepath.Join(dir.Filename, "protobuf.js.map")
			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccLibProtobufMapTemplateFile, ccLibProtobufMapTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccLibProtobufMap").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}

		// Protobuf.js.map
		ProtobufJSIndex := func() error {
			fileName := filepath.Join(dir.Filename, "index.d.ts")

			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccLibProtobufIndexTemplateFile, ccLibProtobufIndexTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccLibProtobufIndex").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}

		err := ProtobufJS()
		if err != nil {
			return err
		}

		err = ProtobufJSMap()
		if err != nil {
			return err
		}

		err = ProtobufJSIndex()
		if err != nil {
			return err
		}

		return nil
	}

	err := Lib()
	if err != nil {
		return err
	}

	err = Howler()
	if err != nil {
		return err
	}

	err = Protobuf()
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) ccModel(ctx CocosDirContext, proto parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetCCScriptModel()

	// CmdType.ts
	cmdType := func() error {
		alias := collection.NewSet()
		for _, item := range proto.Message {
			msgName := getMessageName(*item.Message)
			alias.AddStr(fmt.Sprintf("%s", parser.CamelCase(msgName)))
		}
		aliasKeys := alias.KeysStr()
		sort.Strings(aliasKeys)
		for i, key := range aliasKeys {
			aliasKeys[i] = fmt.Sprintf(`%s = %d`, key, i)
		}

		fileName := filepath.Join(dir.Filename, "CmdType.ts")
		if pathx.FileExists(fileName) {
			os.Remove(fileName)
		}

		text, err := pathx.LoadTemplate(category, ccCmdTypeTemplateFile, ccCmdTypeTemplate)
		if err != nil {
			return err
		}

		if err = util.With("ccCmdType").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			"CmdTypes": strings.Join(aliasKeys, ",\n\t"),
		}, fileName, true); err != nil {
			return err
		}

		return nil
	}

	// GameModel.ts
	gameModel := func() error {
		fileName := filepath.Join(dir.Filename, "GameModel.ts")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccGameModelTemplateFile, ccGameModelTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccGameModel").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			//
		}, fileName, true); err != nil {
			return err
		}

		return nil
	}

	err := cmdType()
	if err != nil {
		return err
	}

	err = gameModel()
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) ccBootstrap(ctx CocosDirContext, proto parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetCCScriptBootstrap()

	// GameApp.ts
	fileName := filepath.Join(dir.Filename, "GameApp.ts")
	if pathx.FileExists(fileName) {
		return nil
	}
	text, err := pathx.LoadTemplate(category, ccGameAppTemplateFile, ccGameAppTemplate)
	if err != nil {
		return err
	}
	if err = util.With("ccGameApp").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		//
	}, fileName, true); err != nil {
		return err
	}

	return nil
}

func (g *Generator) ccController(ctx CocosDirContext, proto parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetCCScriptController()

	// BaseController.ts
	fileName := filepath.Join(dir.Filename, "BaseController.ts")
	if pathx.FileExists(fileName) {
		return nil
	}
	text, err := pathx.LoadTemplate(category, ccBaseControllerTemplateFile, ccBaseControllerTemplate)
	if err != nil {
		return err
	}
	if err = util.With("ccBaseController").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		//
	}, fileName, true); err != nil {
		return err
	}

	// LoginUIController.ts
	fileName = filepath.Join(dir.Filename, "LoginUIController.ts")
	if pathx.FileExists(fileName) {
		return nil
	}
	text, err = pathx.LoadTemplate(category, ccLoginUIControllerTemplateFile, ccLoginUIControllerTemplate)
	if err != nil {
		return err
	}
	if err = util.With("ccLoginUIController").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		//
	}, fileName, true); err != nil {
		return err
	}

	return nil
}

func (g *Generator) ccService(ctx CocosDirContext, proto parser.Proto, cfg *conf.Config) error {

	ccAudio := func() error {
		dirService := ctx.GetCCScriptServiceAudio()
		// AudioConstant.ts
		AudioConstant := func() error {
			fileName := filepath.Join(dirService.Filename, "AudioConstant.ts")
			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccAudioConstantTemplateFile, ccAudioConstantTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccAudioConstant").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}

		// AudioManager.ts
		AudioManager := func() error {
			fileName := filepath.Join(dirService.Filename, "AudioManager.ts")
			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccAudioManagerTemplateFile, ccAudioManagerTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccAudioManager").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}

		// GeneralAudioManager.ts
		GeneralAudioManager := func() error {
			fileName := filepath.Join(dirService.Filename, "GeneralAudioManager.ts")
			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccGeneralAudioManagerTemplateFile, ccGeneralAudioManagerTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccGeneralAudioManager").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}

		// MainGameAudioManager.ts
		MainGameAudioManager := func() error {
			fileName := filepath.Join(dirService.Filename, "MainGameAudioManager.ts")
			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccMainGameAudioManagerTemplateFile, ccMainGameAudioManagerTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccMainGameAudioManager").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}
		// BGMAudioManager.ts
		BGMAudioManager := func() error {
			fileName := filepath.Join(dirService.Filename, "BGMAudioManager.ts")
			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccBGMAudioManagerTemplateFile, ccBGMAudioManagerTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccBGMAudioManager").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}

		err := AudioConstant()
		if err != nil {
			return err
		}
		err = AudioManager()
		if err != nil {
			return err
		}
		err = GeneralAudioManager()
		if err != nil {
			return err
		}
		err = MainGameAudioManager()
		if err != nil {
			return err
		}
		err = BGMAudioManager()
		if err != nil {
			return err
		}
		return nil
	}

	ccNet := func() error {
		dirService := ctx.GetCCScriptServiceNet()
		// NetEvent.ts
		NetEvent := func() error {
			fileName := filepath.Join(dirService.Filename, "NetEvent.ts")
			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccNetEventTemplateFile, ccNetEventTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccNetEvent").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}

		// NetEventDispatcher.ts
		NetEventDispatcher := func() error {
			fileName := filepath.Join(dirService.Filename, "NetEventDispatcher.ts")
			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccNetEventDispatcherTemplateFile, ccNetEventDispatcherTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccNetEventDispatcher").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}
		// NetManager.ts
		NetManager := func() error {
			fileName := filepath.Join(dirService.Filename, "NetManager.ts")
			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccNetManagerTemplateFile, ccNetManagerTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccNetManager").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}

		// ProtoManager.ts
		ProtoManager := func() error {
			fileName := filepath.Join(dirService.Filename, "ProtoManager.ts")
			if pathx.FileExists(fileName) {
				return nil
			}
			text, err := pathx.LoadTemplate(category, ccProtoManagerTemplateFile, ccProtoManagerTemplate)
			if err != nil {
				return err
			}
			if err = util.With("ccProtoManager").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
				//
			}, fileName, true); err != nil {
				return err
			}
			return nil
		}

		err := NetEvent()
		if err != nil {
			return err
		}
		err = NetEventDispatcher()
		if err != nil {
			return err
		}
		err = NetManager()
		if err != nil {
			return err
		}
		err = ProtoManager()
		if err != nil {
			return err
		}

		return nil
	}

	ccProxy := func() error {
		dirService := ctx.GetCCScriptServiceProxy()
		// AuthProxy.ts
		fileName := filepath.Join(dirService.Filename, "AuthProxy.ts")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccAuthProxyTemplateFile, ccAuthProxyTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccAuthProxy").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			//
		}, fileName, true); err != nil {
			return err
		}
		return nil
	}

	dirService := ctx.GetCCScriptService()

	ccEventManager := func() error {
		fileName := filepath.Join(dirService.Filename, "EventManager.ts")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccEventManagerTemplateFile, ccEventManagerTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccEventManager").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			//
		}, fileName, true); err != nil {
			return err
		}
		return nil
	}

	ccResManager := func() error {
		fileName := filepath.Join(dirService.Filename, "ResManager.ts")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccResManagerTemplateFile, ccResManagerTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccResManager").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			//
		}, fileName, true); err != nil {
			return err
		}
		return nil
	}

	ccTimerManager := func() error {
		fileName := filepath.Join(dirService.Filename, "TimerManager.ts")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccTimerManagerTemplateFile, ccTimerManagerTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccTimerManager").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			//
		}, fileName, true); err != nil {
			return err
		}
		return nil
	}

	ccUIManager := func() error {
		fileName := filepath.Join(dirService.Filename, "UIManager.ts")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccUIManagerTemplateFile, ccUIManagerTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccUIManager").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			//
		}, fileName, true); err != nil {
			return err
		}
		return nil
	}

	err := ccAudio()
	if err != nil {
		return err
	}

	err = ccNet()
	if err != nil {
		return err
	}

	err = ccProxy()
	if err != nil {
		return err
	}

	err = ccEventManager()
	if err != nil {
		return err
	}

	err = ccResManager()
	if err != nil {
		return err
	}

	err = ccTimerManager()
	if err != nil {
		return err
	}

	err = ccUIManager()
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) ccView(ctx CocosDirContext, proto parser.Proto, cfg *conf.Config) error {

	dir := ctx.GetCCScriptView()

	// BaseUIView.ts
	fileName := filepath.Join(dir.Filename, "BaseUIView.ts")
	if pathx.FileExists(fileName) {
		return nil
	}
	text, err := pathx.LoadTemplate(category, ccBaseUIViewTemplateFile, ccBaseUIViewTemplate)
	if err != nil {
		return err
	}
	if err = util.With("ccBaseUIView").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		//
	}, fileName, true); err != nil {
		return err
	}

	// LoginUIView.ts
	fileName = filepath.Join(dir.Filename, "LoginUIView.ts")
	if pathx.FileExists(fileName) {
		return nil
	}
	text, err = pathx.LoadTemplate(category, ccLoginUIViewTemplateFile, ccLoginUIViewTemplate)
	if err != nil {
		return err
	}
	if err = util.With("ccLoginUIView").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		//
	}, fileName, true); err != nil {
		return err
	}

	// UIEvent.ts
	fileName = filepath.Join(dir.Filename, "UIEvent.ts")
	if pathx.FileExists(fileName) {
		return nil
	}
	text, err = pathx.LoadTemplate(category, ccUIEventTemplateFile, ccUIEventTemplate)
	if err != nil {
		return err
	}
	if err = util.With("ccUIEvent").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		//
	}, fileName, true); err != nil {
		return err
	}

	return nil
}

func (g *Generator) ccTool(ctx CocosDirContext, proto parser.Proto, cfg *conf.Config) error {

	dir := ctx.GetCCToolProtoCompile()

	toolPackage := func() error {
		fileName := filepath.Join(dir.Filename, "package.json")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccToolPackageTemplateFile, ccToolPackageTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccToolPackage").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			"model": proto.Name,
		}, fileName, true); err != nil {
			return err
		}
		return nil
	}

	toolPbTS := func() error {
		fileName := filepath.Join(dir.Filename, "wrap-pbts.js")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccToolWrapPBTSTemplateFile, ccToolWrapPBTSTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccToolWrapPBTS").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			//
		}, fileName, true); err != nil {
			return err
		}
		return nil
	}

	toolPbJS := func() error {
		fileName := filepath.Join(dir.Filename, "wrap-pbjs.js")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccToolWrapPBJSTemplateFile, ccToolWrapPBJSTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccToolWrapPBJS").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			//
		}, fileName, true); err != nil {
			return err
		}
		return nil
	}

	err := toolPackage()
	if err != nil {
		return err
	}

	err = toolPbJS()
	if err != nil {
		return err
	}

	err = toolPbTS()
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) ccMakefile(ctx CocosDirContext, proto parser.Proto, cfg *conf.Config) error {

	dir := ctx.GetCocos()

	fileName := filepath.Join(dir.Filename, "Makefile")
	if pathx.FileExists(fileName) {
		return nil
	}
	text, err := pathx.LoadTemplate(category, ccMakefileTemplateFile, ccMakefileTemplate)
	if err != nil {
		return err
	}
	if err = util.With("ccMakefile").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
		//
	}, fileName, true); err != nil {
		return err
	}

	return nil
}

func (g *Generator) ccResource(ctx CocosDirContext, proto parser.Proto, cfg *conf.Config) error {

	resConfig := func() error {
		dir := ctx.GetCCResourceConfig()

		fileName := filepath.Join(dir.Filename, "app.json")
		if pathx.FileExists(fileName) {
			return nil
		}
		text, err := pathx.LoadTemplate(category, ccResourceConfigTemplateFile, ccResourceConfigTemplate)
		if err != nil {
			return err
		}
		if err = util.With("ccResourceConfig").GoFmt(false).Parse(text).SaveTo(map[string]interface{}{
			//
		}, fileName, true); err != nil {
			return err
		}
		return nil
	}

	err := resConfig()
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) ccCopyProtoFile(ctx CocosDirContext, proto parser.Proto, _ *conf.Config) error {

	dest := filepath.Join(ctx.GetCCAsset().Filename, proto.Name)

	bytesRead, err := os.ReadFile(proto.Src)
	if err != nil {
		return err
	}

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dest, bytesRead, 0644)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
