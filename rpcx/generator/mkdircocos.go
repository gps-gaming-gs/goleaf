package generator

import (
	"google.golang.org/grpc/encoding/proto"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
)

const (
	// for cocos folder
	cocos                 = "cocos"
	ccAsset               = "asset"
	ccResource            = "resource"
	ccResourceConfig      = "config"
	ccScript              = "script"
	ccScriptLib           = "lib"
	ccScriptLibProtobuf   = "protobuf"
	ccScriptLibBezierAnim = "bezier-anim"
	ccScriptBootstrap     = "bootstrap"
	ccScriptController    = "controller"
	ccScriptModel         = "model"
	ccScriptService       = "service"
	ccScriptServiceAudio  = "audio"
	ccScriptServiceNet    = "net"
	ccScriptServiceProxy  = "proxy"
	ccScriptUtil          = "util"
	ccScriptView          = "view"
	ccTool                = "tool"
	ccToolProtoCompile    = "proto-compile"
)

type (
	// DirContext defines a rpc service directories context
	CocosDirContext interface {
		GetCocos() Dir
		GetCCAsset() Dir
		GetCCResource() Dir
		GetCCResourceConfig() Dir
		GetCCScript() Dir
		GetCCScriptLib() Dir
		GetCCScriptLibProtobuf() Dir
		GetCCScriptLibBezierAnim() Dir
		GetCCScriptBootstrap() Dir
		GetCCScriptController() Dir
		GetCCScriptModel() Dir
		GetCCScriptService() Dir
		GetCCScriptServiceAudio() Dir
		GetCCScriptServiceNet() Dir
		GetCCScriptServiceProxy() Dir
		GetCCScriptUtil() Dir
		GetCCScriptView() Dir

		GetCCTool() Dir
		GetCCToolProtoCompile() Dir
	}

	// Dir defines a directory
	cocosDirContext struct {
		inner       map[string]Dir
		serviceName stringx.String
		ctx         *ctx.ProjectContext
	}
)

func mkdircocos(ctx *ctx.ProjectContext, c *ZRpcContext) (CocosDirContext,
	error) {
	inner := make(map[string]Dir)
	// cocos folder
	cocosDir := filepath.Join(ctx.WorkDir, c.CocosProjectName)
	ccAssetDir := filepath.Join(cocosDir, "assets")
	ccResourceDir := filepath.Join(ccAssetDir, "resources")
	ccResourceConfigDir := filepath.Join(ccResourceDir, "config")
	ccScriptDir := filepath.Join(ccAssetDir, "scripts")
	ccScriptLibDir := filepath.Join(ccScriptDir, "lib")
	ccScriptLibProtobufDir := filepath.Join(ccScriptLibDir, "protobuf")
	ccScriptLibBezierAnimDir := filepath.Join(ccScriptLibDir, "bezier_anim")
	ccScriptBootstrapDir := filepath.Join(ccScriptDir, "bootstrap")
	ccScriptControllerDir := filepath.Join(ccScriptDir, "controllers")
	ccScriptModelDir := filepath.Join(ccScriptDir, "models")
	ccScriptServiceDir := filepath.Join(ccScriptDir, "services")
	ccScriptServiceAudioDir := filepath.Join(ccScriptServiceDir, "audio")
	ccScriptServiceNetDir := filepath.Join(ccScriptServiceDir, "net")
	ccScriptServiceProxyDir := filepath.Join(ccScriptServiceDir, "proxy")
	ccScriptUtilDir := filepath.Join(ccScriptDir, "utils")
	ccScriptViewDir := filepath.Join(ccScriptDir, "views")
	//
	ccToolDir := filepath.Join(cocosDir, "tools")
	ccToolProtoCompileDir := filepath.Join(ccToolDir, "proto-compile")

	getChildPackage := func(parent, childPath string) (string, error) {
		child := strings.TrimPrefix(childPath, parent)
		abs := filepath.Join(parent, strings.ToLower(child))
		if c.Multiple {
			if err := pathx.MkdirIfNotExist(abs); err != nil {
				return "", err
			}
		}
		childPath = strings.TrimPrefix(abs, ctx.Dir)
		pkg := filepath.Join(ctx.Path, childPath)
		return filepath.ToSlash(pkg), nil
	}

	inner[cocos] = Dir{
		Filename: cocosDir,
		Package: filepath.ToSlash(filepath.Join(ctx.Path,
			strings.TrimPrefix(cocosDir, ctx.Dir))),
		Base: filepath.Base(cocosDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(cocosDir, childPath)
		},
	}

	inner[ccAsset] = Dir{
		Filename: ccAssetDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccAssetDir, ctx.Dir))),
		Base:     filepath.Base(ccAssetDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccAssetDir, childPath)
		},
	}

	inner[ccResource] = Dir{
		Filename: ccResourceDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccResourceDir, ctx.Dir))),
		Base:     filepath.Base(ccResourceDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccResourceDir, childPath)
		},
	}

	inner[ccResourceConfig] = Dir{
		Filename: ccResourceConfigDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccResourceConfigDir, ctx.Dir))),
		Base:     filepath.Base(ccResourceConfigDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccResourceConfigDir, childPath)
		},
	}

	inner[ccScript] = Dir{
		Filename: ccScriptDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptDir, childPath)
		},
	}

	inner[ccScriptLib] = Dir{
		Filename: ccScriptLibDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptLibDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptLibDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptLibDir, childPath)
		},
	}

	inner[ccScriptLibProtobuf] = Dir{
		Filename: ccScriptLibProtobufDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptLibProtobufDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptLibProtobufDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptLibProtobufDir, childPath)
		},
	}

	inner[ccScriptLibBezierAnim] = Dir{
		Filename: ccScriptLibBezierAnimDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptLibBezierAnimDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptLibBezierAnimDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptLibBezierAnimDir, childPath)
		},
	}

	inner[ccScriptBootstrap] = Dir{
		Filename: ccScriptBootstrapDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptBootstrapDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptBootstrapDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptBootstrapDir, childPath)
		},
	}

	inner[ccScriptController] = Dir{
		Filename: ccScriptControllerDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptControllerDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptControllerDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptControllerDir, childPath)
		},
	}

	inner[ccScriptModel] = Dir{
		Filename: ccScriptModelDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptModelDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptModelDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptModelDir, childPath)
		},
	}

	inner[ccScriptService] = Dir{
		Filename: ccScriptServiceDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptServiceDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptServiceDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptServiceDir, childPath)
		},
	}

	inner[ccScriptServiceAudio] = Dir{
		Filename: ccScriptServiceAudioDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptServiceAudioDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptServiceAudioDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptServiceAudioDir, childPath)
		},
	}

	inner[ccScriptServiceNet] = Dir{
		Filename: ccScriptServiceNetDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptServiceNetDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptServiceNetDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptServiceNetDir, childPath)
		},
	}

	inner[ccScriptServiceProxy] = Dir{
		Filename: ccScriptServiceProxyDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptServiceProxyDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptServiceProxyDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptServiceProxyDir, childPath)
		},
	}

	inner[ccScriptUtil] = Dir{
		Filename: ccScriptUtilDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptUtilDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptUtilDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptUtilDir, childPath)
		},
	}

	inner[ccScriptView] = Dir{
		Filename: ccScriptViewDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccScriptViewDir, ctx.Dir))),
		Base:     filepath.Base(ccScriptViewDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccScriptViewDir, childPath)
		},
	}

	inner[ccTool] = Dir{
		Filename: ccToolDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccToolDir, ctx.Dir))),
		Base:     filepath.Base(ccToolDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccToolDir, childPath)
		},
	}

	inner[ccToolProtoCompile] = Dir{
		Filename: ccToolProtoCompileDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(ccToolProtoCompileDir, ctx.Dir))),
		Base:     filepath.Base(ccToolProtoCompileDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ccToolProtoCompileDir, childPath)
		},
	}

	for _, v := range inner {
		err := pathx.MkdirIfNotExist(v.Filename)
		if err != nil {
			return nil, err
		}
	}
	serviceName := strings.TrimSuffix(proto.Name, filepath.Ext(proto.Name))
	return &cocosDirContext{
		ctx:         ctx,
		inner:       inner,
		serviceName: stringx.From(strings.ReplaceAll(serviceName, "-", "")),
	}, nil
}

// Cocos
func (d *cocosDirContext) GetCocos() Dir {
	return d.inner[cocos]
}

func (d *cocosDirContext) GetCCAsset() Dir {
	return d.inner[ccAsset]
}

func (d *cocosDirContext) GetCCResource() Dir {
	return d.inner[ccResource]
}

func (d *cocosDirContext) GetCCResourceConfig() Dir {
	return d.inner[ccResourceConfig]
}

func (d *cocosDirContext) GetCCScript() Dir {
	return d.inner[ccScript]
}

func (d *cocosDirContext) GetCCScriptLib() Dir {
	return d.inner[ccScriptLib]
}

func (d *cocosDirContext) GetCCScriptLibProtobuf() Dir {
	return d.inner[ccScriptLibProtobuf]
}

func (d *cocosDirContext) GetCCScriptLibBezierAnim() Dir {
	return d.inner[ccScriptLibBezierAnim]
}

func (d *cocosDirContext) GetCCScriptBootstrap() Dir {
	return d.inner[ccScriptBootstrap]
}

func (d *cocosDirContext) GetCCScriptController() Dir {
	return d.inner[ccScriptController]
}

func (d *cocosDirContext) GetCCScriptModel() Dir {
	return d.inner[ccScriptModel]
}

func (d *cocosDirContext) GetCCScriptService() Dir {
	return d.inner[ccScriptService]
}

func (d *cocosDirContext) GetCCScriptServiceAudio() Dir {
	return d.inner[ccScriptServiceAudio]
}

func (d *cocosDirContext) GetCCScriptServiceNet() Dir {
	return d.inner[ccScriptServiceNet]
}

func (d *cocosDirContext) GetCCScriptServiceProxy() Dir {
	return d.inner[ccScriptServiceProxy]
}

func (d *cocosDirContext) GetCCScriptUtil() Dir {
	return d.inner[ccScriptUtil]
}

func (d *cocosDirContext) GetCCScriptView() Dir {
	return d.inner[ccScriptView]
}

func (d *cocosDirContext) GetCCTool() Dir {
	return d.inner[ccTool]
}

func (d *cocosDirContext) GetCCToolProtoCompile() Dir {
	return d.inner[ccToolProtoCompile]
}
