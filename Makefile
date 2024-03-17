build:
	wails build .

setup:
	git submodule add -b gh-pages --force --depth 1 https://github.com/takecx/scratch-gui frontend
	go install github.com/wailsapp/wails/v2/cmd/wails@latest
