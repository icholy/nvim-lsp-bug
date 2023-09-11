# Steps

* Install gopls `go install golang.org/x/tools/gopls@latest`
* Open file `nvim --clean -u init.lua main.go`
* Insert tab character `i<TAB><ESC>`
* Format the file `:lua vim.lsp.buf.format()`

