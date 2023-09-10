vim.o.expandtab = false
vim.o.shiftwidth = 0
vim.o.softtabstop = -1

vim.cmd.packadd("packer.nvim")

require("packer").startup(function(use)
  use("wbthomason/packer.nvim")
  use({
    "neovim/nvim-lspconfig",
    config = function()
      require("lspconfig").eslint.setup({})
    end
  })
end)


