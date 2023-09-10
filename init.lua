vim.o.expandtab = false
vim.o.shiftwidth = 0
vim.o.softtabstop = -1

vim.cmd.packadd("nvim-lspconfig")

require("lspconfig").eslint.setup({})
