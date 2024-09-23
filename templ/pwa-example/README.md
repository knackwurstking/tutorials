# PWA (templ) Example

## Gettings Started

**Installing the templ cli**

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

**Neovim setup (lsp-zero)**

```lua
-- ~/.config/nvim/after/plugin/templ.lua
-- IDK where is the neovim configuration on Mac or Windows,
-- so you need to do some research :)

local function setup_templ()
    local lspconfig = require 'lspconfig'
    local configs = require 'lspconfig.configs'

    -- start the templ language server for go projects with .templ files
    configs.templ = {
        default_config = {
            cmd = { "templ", "lsp", "-http=localhost:7474", "-log=/tmp/templ.log" },
            filetypes = { "templ" },
            root_dir = lspconfig.util.root_pattern("go.mod", ".git"),
            settings = {},
        },
    }
    lspconfig.templ.setup{}

    -- register .templ as a filetype
    vim.filetype.add({ extension = { templ = "templ" } })
    lspconfig.html.setup({
        on_attach = lsp.on_attach,
        capabilities = lsp.capabilities,
        filetypes = { "html", "templ" },
    })

    -- htmx, the frontend library of peace
    lspconfig.htmx.setup({
        on_attach = lsp.on_attach,
        capabilities = lsp.capabilities,
        filetypes = { "html", "templ" },
    })

    -- needed tailwindcss classes auto complete
    lspconfig.tailwindcss.setup({
        on_attach = lsp.on_attach,
        capabilities = lsp.capabilities,
        filetypes = { "templ", "astro", "javascript", "typescript", "react" },
        init_options = { userLanguages = { templ = "html" } },
    })

    -- needed for auto tag insertion
    lspconfig.emmet_ls.setup({
        on_attach = lsp.on_attach,
        capabilities = lsp.capabilities,
        filetypes = { "templ", "astro", "javascript", "typescript", "react" },
        init_options = { userLanguages = { templ = "html" } },
    })

    -- format thingy
    vim.api.nvim_create_autocmd({ "BufWritePost" }, { -- IDK the docs said to do the format before saving the file, but it only makes the formatter freak out.
        pattern = { "*.templ" },
        callback = function()
            local file_name = vim.api.nvim_buf_get_name(0) -- Get file name of file in current buffer
            vim.cmd(":silent !templ fmt " .. file_name)

            local bufnr = vim.api.nvim_get_current_buf()
            if vim.api.nvim_get_current_buf() == bufnr then
                vim.cmd('e!')
            end
        end
    })
end

setup_templ()
```

**Project structure**

- `components/` - templ components.
- `db/` - Database access code used to access the spending logs.
- `handlers/` - HTTP handlers.
- `static/` - Files that are available to the public.
- `services/` - Services used by the handlers.
- `.gitignore` - Some stuff are not worthy of being committed.
- `Dockerfile` - Container configuration to run the application with the glorious Docker.
- `Makefile` - A runner and builder script to run the templ thing alongside tailwindcss, and it has the build commands.
- `main.go` - The entrypoint to our application.
