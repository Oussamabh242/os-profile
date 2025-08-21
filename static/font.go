package static

import (
	"embed"
)

//go:embed IosevkaNerdFont-Regular.ttf
var IosevkaTTF []byte

//go:embed posts/*
var PostsFS embed.FS
