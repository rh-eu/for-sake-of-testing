set background=dark
set t_Co=256

let mapleader = ","

set number " show line numbers
set splitbelow " this will cause all splits to open below (including term)
set splitright " this will cause all vsplits to open right-hand

set nocompatible              " be iMproved, required
filetype off                  " required


call plug#begin('~/.vim/plugged')

" Make sure you use single quotes

" Shorthand notation; fetches https://github.com/junegunn/vim-easy-align
Plug 'junegunn/vim-easy-align'

" Any valid git URL is allowed
Plug 'https://github.com/junegunn/vim-github-dashboard.git'

" Multiple Plug commands can be written in a single line using | separators
Plug 'SirVer/ultisnips' | Plug 'honza/vim-snippets'


" Vim-go and friends
Plug 'fatih/vim-go', { 'do': ':GoInstallBinaries' }
Plug 'AndrewRadev/splitjoin.vim'
" Plug 'fatih/molokai'
Plug 'tomasr/molokai'
Plug 'ctrlpvim/ctrlp.vim'

" Python ...
"Plug 'nvie/vim-flake8'

"Git
Plug 'tpope/vim-fugitive'

" NERDTree
Plug 'preservim/nerdtree'

" Vim airline
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'


" Vim-Polyglot and Language packs
Plug 'sheerun/vim-polyglot'
Plug 'godlygeek/tabular'
Plug 'plasticboy/vim-markdown'
Plug 'chrisbra/csv.vim'
Plug 'ekalinin/Dockerfile.vim'
Plug 'pangloss/vim-javascript'
Plug 'elzr/vim-json'

" Emmet-vim
Plug 'mattn/emmet-vim'

" coc-nvim
"Plug 'neoclide/coc.nvim', {'branch': 'release'}

" fzf - the command line fuzzy finder 
Plug 'junegunn/fzf.vim'


" Initialize plugin system
call plug#end()


" Global Settings
:set tabstop=2
:set expandtab
:set shiftwidth=2
:set autoindent
:set smartindent

" Bind nohl
" Removes highlight of your last search
" ``<C>`` stands for ``CTRL`` and therefore ``<C-n>`` stands for ``CTRL+n``
noremap <C-n> :nohl<CR>
vnoremap <C-n> :nohl<CR>
inoremap <C-n> :nohl<CR>

" Quicksave command
noremap <C-Z> :update<CR>
vnoremap <C-Z> <C-C>:update<CR>
inoremap <C-Z> <C-O>:update<CR>

" Quicksave command
noremap <C-Z> :update<CR>
vnoremap <C-Z> <C-C>:update<CR>
inoremap <C-Z> <C-O>:update<CR>

" Quick quit command
noremap <Leader>e :quit<CR> " Quit current window
noremap <Leader>E :qa!<CR> " Quit all windows

" bind Ctrl+<movement> keys to move around the windows, instead of using
" Ctrl+w + <movement>
" Every unnecessary keystroke that can be saved is good for your health :)
map <c-j> <c-w>j
map <c-k> <c-w>k
map <c-l> <c-w>l
map <c-h> <c-w>h

" easier moving between tabs
map <Leader>n <esc>:tabprevious<CR>
map <Leader>m <esc>:tabnext<CR>

" map sort function to a key
vnoremap <Leader>s :sort<CR>

" easier moving of code blocks
" Try to go into visual mode (v), thenselect several lines of code here and
" then press ``>`` several times.
vnoremap < <gv " better indentation
vnoremap > >gv " better indentation

" pastetoogle
nnoremap <F2> :set invpaste paste?<CR>
set pastetoggle=<F2>
set showmode


" Begin vim-go
"
let g:go_fmt_command = "goimports"

let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1

" color settings
let g:molokai_original = 1
let g:rehash256 = 1
colorscheme molokai

"
" End vim-go
"

" Begin NERDTree
autocmd vimenter * NERDTree
map <Leader>t :NERDTreeToggle<CR>

" End NERDTree


" Begin vim-airline
"
let g:airline#extensions#tabline#enabled = 1
"
" End vim-airline
