package handlers

import (
	"embed"
	"io/fs"
	"net/http"
)

func HandleFS(publicFS embed.FS, root string) http.Handler {
	public, _ := fs.Sub(publicFS, root)
	return http.FileServer(http.FS(public))
}
