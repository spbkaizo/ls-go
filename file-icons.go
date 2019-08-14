package main

import (
	"strings"
)

func getIconForFile(name, ext string) string {
	key := strings.ToLower(ext)
	alias, hasAlias := aliases[key]
	if hasAlias {
		key = alias
	}
	icon := icons["file"]
	betterIcon, hasBetterIcon := icons[key]
	if hasBetterIcon {
		icon = betterIcon
	}
	fullName := name
	if ext != "" {
		fullName += "." + ext
	}
	bestIcon, hasBestIcon := icons[strings.ToLower(fullName)]
	if hasBestIcon {
		icon = bestIcon
	}
	return icon
}

func getIconForFolder(name string) string {
	icon := folders["folder"]
	betterIcon, hasBetterIcon := folders[name]
	if hasBetterIcon {
		icon = betterIcon
	}
	return icon
}

var icons = map[string]string{
	"ai":           "\ue7b4",
	"android":      "\ue70e",
	"apple":        "\uf179",
	"audio":        "\uf1c7",
	"avro":         "\ue60b",
	"binary":       "\uf471",
	"c":            "\ue61e",
	"cfg":          "\uf423",
	"cargo.lock":   "\uf487",
	"cargo.toml":   "\uf487",
	"clj":          "\ue768",
	"coffee":       "\ue751",
	"conf":         "\ue615",
	"cpp":          "\ue61d",
	"cr":           "\ue23e",
	"cs":           "\uf81a",
	"cson":         "\ue601",
	"css":          "\ue749",
	"d":            "\ue7af",
	"dart":         "\ue798",
	"db":           "\uf1c0",
	"deb":          "\uf306",
	"diff":         "\uf440",
	"doc":          "\uf1c2",
	"dockerfile":   "\ue7b0",
	"dpkg":         "\uf17c",
	"ebook":        "\uf02d",
	"elm":          "\ue62c",
	"env":          "\uf462",
	"erl":          "\ue7b1",
	"ex":           "\ue62d",
	"ics":          "\uf073",
	"key":          "\uf43d",
	"file":         "\uf15b",
	"font":         "\uf031",
	"gform":        "\uf298",
	"git":          "\ue702",
	"go":           "\ue724",
	"gruntfile.js": "\ue74c",
	"gulpfile.js":  "\ue610",
	"h":            "\uf0fd",
	"hs":           "\ue777",
	"html":         "\uf13b",
	"image":        "\uf1c5",
	"iml":          "\ue7b5",
	"ini":          "\uf669",
	"java":         "\ue738",
	"jenkinsfile":  "\ue767",
	"js":           "\ue781",
	"json":         "\ue60b",
	"jsx":          "\ue7ba",
	"mjs":          "\ue718",
	"less":         "\ue758",
	"lock":         "\uf720",
	"log":          "\uf18d",
	"lua":          "\ue620",
	"makefile":     "\ue20f",
	"md":           "\uf48a",
	"mustache":     "\ue60f",
	"npmignore":    "\ue71e",
	"patch":        "\uf440",
	"pdf":          "\uf1c1",
	"php":          "\ue73d",
	"pl":           "\ue769",
	"ppt":          "\uf1c4",
	"psd":          "\ue7b8",
	"py":           "\ue606",
	"passwd":       "\uf023",
	"r":            "\uf25d",
	"rs":           "\ue7a8",
	"rb":           "\ue21e",
	"rdb":          "\ue76d",
	"rss":          "\uf09e",
	"rpm":          "\uf17c",
	"rubydoc":      "\ue73b",
	"sass":         "\ue603",
	"scala":        "\ue737",
	"shell":        "\uf489",
	"sqlite3":      "\ue7c4",
	"sol":          "\ufcb9",
	"styl":         "\ue600",
	"swift":        "\ue755",
	"sql":          "\ue706",
	"tex":          "\ue600",
	"toml":         "\uf669",
	"ts":           "\ufbe4",
	"twig":         "\ue61c",
	"txt":          "\uf15c",
	"vagrantfile":  "\ue21e",
	"video":        "\uf03d",
	"vim":          "\ue62b",
	"windows":      "\uf17a",
	"xls":          "\uf1c3",
	"xml":          "\ue619",
	"vimrc":        "\ue62b",
	"yml":          "\ue601",
	"zip":          "\uf410",
}

var aliases = map[string]string{
	"apk":              "android",
	"gradle":           "android",
	"ds_store":         "apple",
	"localized":        "apple",
	"mp3":              "audio",
	"ogg":              "audio",
	"m4a":              "audio",
	"wav":              "audio",
	"flac":             "audio",
	"alac":             "audio",
	"aac":              "audio",
	"wma":              "audio",
	"mka":              "audio",
	"opus":             "audio",
	"pickle":           "binary",
	"pkl":              "binary",
	"pb":               "binary",
	"o":                "binary",
	"cc":               "cpp",
	"cxx":              "cpp",
	"c++":              "cpp",
	"conf":             "cfg",
	"config":           "cfg",
	"gpg":              "key",
	"pgp":              "key",
	"pem":              "key",
	"crt":              "key",
	"cer":              "key",
	"der":              "key",
	"pfx":              "key",
	"p7b":              "key",
	"editorconfig":     "conf",
	"rc":               "conf",
	"scss":             "css",
	"docx":             "doc",
	"gdoc":             "doc",
	"mobi":             "ebook",
	"epub":             "ebook",
	"eot":              "font",
	"otf":              "font",
	"ttf":              "font",
	"woff":             "font",
	"woff2":            "font",
	"gitattributes":    "git",
	"gitconfig":        "git",
	"gitignore":        "git",
	"gitmirrorall":     "git",
	"gitmodules":       "git",
	"gitignore_global": "git",
	"lhs":              "hs",
	"hpp":              "h",
	"hh":               "h",
	"h++":              "h",
	"hxx":              "h",
	"htm":              "html",
	"bmp":              "image",
	"gif":              "image",
	"ico":              "image",
	"jpeg":             "image",
	"jpg":              "image",
	"png":              "image",
	"svg":              "image",
	"webp":             "image",
	"tiff":             "image",
	"pxm":              "image",
	"ppm":              "image",
	"pgm":              "image",
	"pbm":              "image",
	"pnm":              "image",
	"stl":              "image",
	"eps":              "image",
	"dvi":              "image",
	"cbr":              "image",
	"cbz":              "image",
	"xpm":              "image",
	"orf":              "image",
	"nef":              "image",
	"jar":              "java",
	"properties":       "json",
	"webmanifest":      "json",
	"tsx":              "jsx",
	"license":          "md",
	"markdown":         "md",
	"mkd":              "md",
	"rdoc":             "md",
	"readme":           "md",
	"gslides":          "ppt",
	"pptx":             "ppt",
	"pyc":              "py",
	"whl":              "py",
	"ipynb":            "ebook",
	"rdata":            "r",
	"rds":              "r",
	"rmd":              "r",
	"gemfile":          "rb",
	"gemspec":          "rb",
	"guardfile":        "rb",
	"procfile":         "rb",
	"rakefile":         "rb",
	"rspec":            "rb",
	"rspec_parallel":   "rb",
	"rspec_status":     "rb",
	"ru":               "rb",
	"erb":              "rubydoc",
	"slim":             "rubydoc",
	"awk":              "shell",
	"bash":             "shell",
	"bash_history":     "shell",
	"bash_profile":     "shell",
	"bashrc":           "shell",
	"fish":             "shell",
	"sh":               "shell",
	"zsh":              "shell",
	"zsh-theme":        "shell",
	"zshrc":            "shell",
	"ksh":              "shell",
	"csh":              "shell",
	"stylus":           "styl",
	"tsql":             "sql",
	"psql":             "sql",
	"plsql":            "sql",
	"plpgsql":          "sql",
	"cls":              "tex",
	"avi":              "video",
	"mkv":              "video",
	"mp4":              "video",
	"ogv":              "video",
	"webm":             "video",
	"mov":              "video",
	"flv":              "video",
	"m2v":              "video",
	"mpeg":             "video",
	"mpg":              "video",
	"ogm":              "video",
	"vob":              "video",
	"bat":              "windows",
	"cmd":              "windows",
	"exe":              "windows",
	"csv":              "xls",
	"gsheet":           "xls",
	"xlsx":             "xls",
	"xul":              "xml",
	"yaml":             "yml",
	"gz":               "zip",
	"rar":              "zip",
	"tar":              "zip",
	"tgz":              "zip",
	"Z":                "zip",
	"z":                "zip",
	"bz2":              "zip",
	"7z":               "zip",
	"iso":              "zip",
	"dmg":              "zip",
	"tc":               "zip",
	"par":              "zip",
	"xz":               "zip",
	"txz":              "zip",
	"lzma":             "zip",
}

var folders = map[string]string{
	".atom":                 "\ue764",
	".aws":                  "\ue7ad",
	".docker":               "\ue7b0",
	".gem":                  "\ue21e",
	".git":                  "\ue5fb",
	".git-credential-cache": "\ue5fb",
	".github":               "\ue5fd",
	".npm":                  "\ue5fa",
	".nvm":                  "\ue718",
	".rvm":                  "\ue21e",
	".Trash":                "\uf1f8",
	".vscode":               "\ue70c",
	".vim":                  "\ue62b",
	"config":                "\ue5fc",
	"folder":                "\uf115",
	"hidden":                "\uf023",
	"node_modules":          "\ue5fa",
}

var otherIcons = map[string]string{
	"link":   "\uf0c1",
	"device": "\uf0a0",
}
