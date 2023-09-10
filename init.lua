vim.o.shiftwidth = 0
vim.o.softtabstop = -1

vim.api.nvim_create_autocmd('FileType', {
  pattern = '*',
  callback = function(args)
    vim.lsp.start({
      name = 'fakelsp',
      cmd = { 'go', 'run', '.' },
      root_dir = vim.fn.expand('~'),
      settings = {},
    })
  end
})
