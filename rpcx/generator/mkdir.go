package generator

import (
	"path/filepath"
	"strings"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"

	"github.com/gps-gaming-gs/goleaf/rpcx/parser"

	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
)

const (
	wd      = "wd"
	pb      = "pb"
	protoGo = "proto-go"

	wdconf            = "wdconf"
	server            = "server"
	base              = "base"
	config            = "conf"
	game              = "game"
	gameInternal      = "game-internal"
	gameInternalLogic = "game-internal-logic"
	gamedata          = "gamedata"
	gate              = "gate"
	gateInternal      = "gate-internal"
	login             = "login"
	loginInternal     = "login-internal"
	msg               = "msg"
)

type (
	// DirContext defines a rpc service directories context
	DirContext interface {
		GetConf() Dir
		GetBase() Dir
		GetServer() Dir
		GetConfig() Dir
		GetGame() Dir
		GetGameInternal() Dir
		GetGameInternalLogic() Dir
		GetGameData() Dir
		GetGate() Dir
		GetGateInternal() Dir
		GetLogin() Dir
		GetLoginInternal() Dir
		GetMsg() Dir
		GetPb() Dir
		GetProtoGo() Dir
		GetMain() Dir
		GetServiceName() stringx.String
		SetPbDir(pbDir, grpcDir string)
	}

	// Dir defines a directory
	Dir struct {
		Base            string
		Filename        string
		Package         string
		GetChildPackage func(childPath string) (string, error)
	}

	defaultDirContext struct {
		inner       map[string]Dir
		serviceName stringx.String
		ctx         *ctx.ProjectContext
	}
)

func mkdir(ctx *ctx.ProjectContext, proto parser.Proto, _ *conf.Config, c *ZRpcContext) (DirContext,
	error) {
	inner := make(map[string]Dir)
	confDir := filepath.Join(ctx.WorkDir, "conf")
	serverDir := filepath.Join(ctx.WorkDir, "server")
	baseDir := filepath.Join(serverDir, "base")
	configDir := filepath.Join(serverDir, "conf")
	gameDir := filepath.Join(serverDir, "game")
	gameInternalDir := filepath.Join(gameDir, "internal")
	gameInternalLogicDir := filepath.Join(gameInternalDir, "logic")
	gameDataDir := filepath.Join(serverDir, "gamedata")
	gateDir := filepath.Join(serverDir, "gate")
	gateInternalDir := filepath.Join(gateDir, "internal")
	loginDir := filepath.Join(serverDir, "login")
	loginInternalDir := filepath.Join(loginDir, "internal")
	msgDir := filepath.Join(serverDir, "msg")
	pbDir := filepath.Join(ctx.WorkDir, proto.GoPackage)
	protoGoDir := pbDir
	if c != nil {
		pbDir = msgDir
		protoGoDir = msgDir
	}

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

	inner[wd] = Dir{
		Filename: ctx.WorkDir,
		Package: filepath.ToSlash(filepath.Join(ctx.Path,
			strings.TrimPrefix(ctx.WorkDir, ctx.Dir))),
		Base: filepath.Base(ctx.WorkDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(ctx.WorkDir, childPath)
		},
	}

	inner[wdconf] = Dir{
		Filename: confDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(confDir, ctx.Dir))),
		Base:     filepath.Base(confDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(confDir, childPath)
		},
	}

	inner[server] = Dir{
		Filename: serverDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(serverDir, ctx.Dir))),
		Base:     filepath.Base(serverDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(serverDir, childPath)
		},
	}
	inner[base] = Dir{
		Filename: baseDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(baseDir, ctx.Dir))),
		Base:     filepath.Base(baseDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(baseDir, childPath)
		},
	}
	inner[config] = Dir{
		Filename: configDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(configDir, ctx.Dir))),
		Base:     filepath.Base(configDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(configDir, childPath)
		},
	}
	inner[game] = Dir{
		Filename: gameDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(gameDir, ctx.Dir))),
		Base:     filepath.Base(gameDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(gameDir, childPath)
		},
	}
	inner[gameInternal] = Dir{
		Filename: gameInternalDir,
		Package: filepath.ToSlash(filepath.Join(ctx.Path,
			strings.TrimPrefix(gameInternalDir, ctx.Dir))),
		Base: filepath.Base(gameInternalDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(gameInternalDir, childPath)
		},
	}
	inner[gameInternalLogic] = Dir{
		Filename: gameInternalLogicDir,
		Package: filepath.ToSlash(filepath.Join(ctx.Path,
			strings.TrimPrefix(gameInternalLogicDir, ctx.Dir))),
		Base: filepath.Base(gameInternalLogicDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(gameInternalLogicDir, childPath)
		},
	}
	inner[gamedata] = Dir{
		Filename: gameDataDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(gameDataDir, ctx.Dir))),
		Base:     filepath.Base(gameDataDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(gameDataDir, childPath)
		},
	}

	inner[gate] = Dir{
		Filename: gateDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(gateDir, ctx.Dir))),
		Base:     filepath.Base(gateDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(gateDir, childPath)
		},
	}

	inner[gateInternal] = Dir{
		Filename: gateInternalDir,
		Package: filepath.ToSlash(filepath.Join(ctx.Path,
			strings.TrimPrefix(gateInternalDir, ctx.Dir))),
		Base: filepath.Base(gateInternalDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(gateInternalDir, childPath)
		},
	}

	inner[login] = Dir{
		Filename: loginDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(loginDir, ctx.Dir))),
		Base:     filepath.Base(loginDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(loginDir, childPath)
		},
	}

	inner[loginInternal] = Dir{
		Filename: loginInternalDir,
		Package: filepath.ToSlash(filepath.Join(ctx.Path,
			strings.TrimPrefix(loginInternalDir, ctx.Dir))),
		Base: filepath.Base(loginInternalDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(loginInternalDir, childPath)
		},
	}

	inner[msg] = Dir{
		Filename: msgDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(msgDir, ctx.Dir))),
		Base:     filepath.Base(msgDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(msgDir, childPath)
		},
	}

	inner[pb] = Dir{
		Filename: pbDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(pbDir, ctx.Dir))),
		Base:     filepath.Base(pbDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(pbDir, childPath)
		},
	}

	inner[protoGo] = Dir{
		Filename: protoGoDir,
		Package: filepath.ToSlash(filepath.Join(ctx.Path,
			strings.TrimPrefix(protoGoDir, ctx.Dir))),
		Base: filepath.Base(protoGoDir),
		GetChildPackage: func(childPath string) (string, error) {
			return getChildPackage(protoGoDir, childPath)
		},
	}

	for _, v := range inner {
		err := pathx.MkdirIfNotExist(v.Filename)
		if err != nil {
			return nil, err
		}
	}
	serviceName := strings.TrimSuffix(proto.Name, filepath.Ext(proto.Name))
	return &defaultDirContext{
		ctx:         ctx,
		inner:       inner,
		serviceName: stringx.From(strings.ReplaceAll(serviceName, "-", "")),
	}, nil
}

func (d *defaultDirContext) SetPbDir(pbDir, grpcDir string) {
	d.inner[pb] = Dir{
		Filename: pbDir,
		Package:  filepath.ToSlash(filepath.Join(d.ctx.Path, strings.TrimPrefix(pbDir, d.ctx.Dir))),
		Base:     filepath.Base(pbDir),
	}

	d.inner[protoGo] = Dir{
		Filename: grpcDir,
		Package: filepath.ToSlash(filepath.Join(d.ctx.Path,
			strings.TrimPrefix(grpcDir, d.ctx.Dir))),
		Base: filepath.Base(grpcDir),
	}
}

func (d *defaultDirContext) GetConf() Dir {
	return d.inner[wdconf]
}

func (d *defaultDirContext) GetBase() Dir {
	return d.inner[base]
}

func (d *defaultDirContext) GetServer() Dir {
	return d.inner[server]
}

func (d *defaultDirContext) GetConfig() Dir {
	return d.inner[config]
}

func (d *defaultDirContext) GetGame() Dir {
	return d.inner[game]
}

func (d *defaultDirContext) GetGameInternal() Dir {
	return d.inner[gameInternal]
}

func (d *defaultDirContext) GetGameInternalLogic() Dir {
	return d.inner[gameInternalLogic]
}

func (d *defaultDirContext) GetGameData() Dir {
	return d.inner[gamedata]
}

func (d *defaultDirContext) GetGate() Dir {
	return d.inner[gate]
}

func (d *defaultDirContext) GetGateInternal() Dir {
	return d.inner[gateInternal]
}

func (d *defaultDirContext) GetLogin() Dir {
	return d.inner[login]
}

func (d *defaultDirContext) GetLoginInternal() Dir {
	return d.inner[loginInternal]
}

func (d *defaultDirContext) GetPb() Dir {
	return d.inner[pb]
}

func (d *defaultDirContext) GetProtoGo() Dir {
	return d.inner[protoGo]
}

func (d *defaultDirContext) GetMsg() Dir {
	return d.inner[msg]
}

func (d *defaultDirContext) GetMain() Dir {
	return d.inner[wd]
}

func (d *defaultDirContext) GetServiceName() stringx.String {
	return d.serviceName
}

// Valid returns true if the directory is valid
func (d *Dir) Valid() bool {
	return len(d.Filename) > 0 && len(d.Package) > 0
}
