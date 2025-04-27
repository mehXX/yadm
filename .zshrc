# ZSH
# ---------------------------------------------------------------------------------------------------------------------------------------------------------------- #
export ZSH="$HOME/.oh-my-zsh"

ZSH_THEME="af-magic"

HISTFILE=~/.zhistory
SAVEHIST=100000
HISTSIZE=100000
export REPORTTIME=1
export UPDATE_ZSH_DAYS=30

setopt INC_APPEND_HISTORY
setopt EXTENDED_HISTORY
setopt HIST_EXPIRE_DUPS_FIRST
setopt HIST_IGNORE_DUPS
setopt HIST_IGNORE_ALL_DUPS
setopt HIST_IGNORE_SPACE
setopt HIST_FIND_NO_DUPS
setopt HIST_SAVE_NO_DUPS
setopt HIST_BEEP
setopt AUTO_LIST
setopt AUTO_RESUME
setopt MULTIOS
setopt NO_CASE_GLOB
setopt CHASE_LINKS

plugins=(git zsh-autosuggestions zsh-syntax-highlighting fzf sudo)

source $ZSH/oh-my-zsh.sh

# partial completion suggestions
zstyle ':completion:*' list-suffixes
zstyle ':completion:*' expand prefix suffix

# case insensitive path-completion
zstyle ':completion:*' matcher-list 'm:{[:lower:][:upper:]}={[:upper:][:lower:]}' 'm:{[:lower:][:upper:]}={[:upper:][:lower:]} l:|=* r:|=*' 'm:{[:lower:][:upper:]}={[:upper:][:lower:]} l:|=* r:|=*' 'm:{[:lower:][:upper:]}={[:upper:][:lower:]} l:|=* r:|=*'

RPS1=""

#GO
# ---------------------------------------------------------------------------------------------------------------------------------------------------------------- #

#go env -w GOSUMDB="sum.golang.org"
#go env -w GOPROXY="https://goproxy.io,direct"
export GOPATH=$HOME/go

#ENV
# ---------------------------------------------------------------------------------------------------------------------------------------------------------------- #

eval "$(/opt/homebrew/bin/brew shellenv)"

export EDITOR=micro

export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:"$HOME/.DevUtils"
export PATH=$PATH:"$HOME/.local/bin"
export PATH=$PATH:"/usr/local/go/bin"
export PATH=$PATH:"$HOME/settings/dev_utils"
export PATH=$PATH:"/opt/homebrew/Cellar/libpq/17.4_1/bin"

export HOMEBREW_NO_AUTO_UPDATE=1
export PGDATABASE="postgres"
export PGUSER="postgres"
export PGPASSWORD='postgres'
export GIT_USER_CONFIG=$USER

export TIMEFMT="%E"

# FUNCTIONS
# ---------------------------------------------------------------------------------------------------------------------------------------------------------------- #
alias history='fc -il 1'
alias clear='clear && printf "\e[3J" clear && printf "\e[3J" && ls'
alias ls="eza --group-directories-first"
alias grep="rg"
alias cd="z"
alias ctags="`brew --prefix`/bin/ctags"
alias project="cd $(go list -m -e -json | jq -r .Dir)"
alias rm="trash"
# Force rebuild: The -B flag (also --always-make) tells make to unconditionally make all targets, regardless of whether they are up to date or not.
alias make="make -B"
alias resticprofile="resticprofile -c $HOME/settings/restic_profiles/profiles.conf"

function chpwd() {
    emulate -L zsh
    ls
}

# WORK RELATED
# ---------------------------------------------------------------------------------------------------------------------------------------------------------------- #


# START
# ---------------------------------------------------------------------------------------------------------------------------------------------------------------- #

eval "$(zoxide init zsh)"

ls
