package configs

// Routes - contains routes list
// "name" - uniq template file name in ./templates/ without extension
// "path" = url
var Routes = map[string]map[string]string{
	"assets": {
		"name": "assets",
		"path": "/assets/",
	},
	"index": {
		"name": "index",
		"path": "/",
	},
}
