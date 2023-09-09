vim.o.expandtab = true

vim.cmd.packadd("packer.nvim")

require("packer").startup(function(use)
  use("wbthomason/packer.nvim")
  use("tpope/vim-sleuth")
  use({
    "neovim/nvim-lspconfig",
    config = function()
      require("lspconfig").eslint.setup({})
    end
  })
end)


