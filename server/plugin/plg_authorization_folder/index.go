package plg_authorization_folder

import (
	"strings"

	. "github.com/mickael-kerjean/filestash/server/common"
)

func init() {
	Hooks.Register.AuthorisationMiddleware(AuthM{})
}

type AuthM struct{}

var adminList = []string{"kp"}

const root = "/music-share/"

func checkNameInPath(path, login string) bool {
	if strings.Contains(strings.Join(adminList, ","), login) {
		return true
	}
	trimmedPath := strings.TrimPrefix(path, root)
	base_folder := trimmedPath[0:strings.Index(trimmedPath, "/")]
	if base_folder == login {
		return true
	} else {
		return false
	}
}

func (this AuthM) Ls(ctx *App, path string) error {
	Log.Stdout("LS %+v", ctx.Session)
	return nil
}

func (this AuthM) Cat(ctx *App, path string) error {
	Log.Stdout("CAT: %v", ctx.Session)
	return nil
}

func (this AuthM) Mkdir(ctx *App, path string) error {
	Log.Stdout("MKDIR %+v", ctx.Session)
	if checkNameInPath(path, ctx.Session["session_token"]) {
		return nil
	} else {
		return ErrNotAllowed
	}
}

func (this AuthM) Rm(ctx *App, path string) error {
	Log.Stdout("RM %+v", ctx.Session)
	if checkNameInPath(path, ctx.Session["session_token"]) {
		return nil
	} else {
		return ErrNotAllowed
	}
}

func (this AuthM) Mv(ctx *App, from string, to string) error {
	Log.Stdout("MV %+v", ctx.Session)
	if checkNameInPath(from, ctx.Session["session_token"]) && checkNameInPath(to, ctx.Session["session_token"]) {
		return nil
	} else {
		return ErrNotAllowed
	}
}

func (this AuthM) Save(ctx *App, path string) error {
	Log.Stdout("SAVE %+v", ctx.Session)
	if checkNameInPath(path, ctx.Session["session_token"]) {
		return nil
	} else {
		return ErrNotAllowed
	}
}

func (this AuthM) Touch(ctx *App, path string) error {
	Log.Stdout("TOUCH %+v", ctx.Session)
	if checkNameInPath(path, ctx.Session["session_token"]) {
		return nil
	} else {
		return ErrNotAllowed
	}
}
