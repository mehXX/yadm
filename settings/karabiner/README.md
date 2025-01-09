# Karabiner Elements configuration

Author: https://github.com/mxstbr/karabiner/tree/main

## Installation

1. Install & start [Karabiner Elements](https://karabiner-elements.pqrs.org/)
2. Clone this repository 
3. Delete the default `~/.config/karabiner` folder 
4. Create a symlink with `ln -s ~/github/meh/karabiner ~/.config` (where `~/github/meh/karabiner` is your local path to where you cloned the repository)
5. [Restart karabiner_console_user_server](https://karabiner-elements.pqrs.org/docs/manual/misc/configuration-file-path/) with

``` 
1. launchctl kickstart -k gui/`id -u`/org.pqrs.karabiner.karabiner_console_user_server 
```

## Development

```
yarn install
```

to install the dependencies. (one-time only)

```
yarn run build
```

builds the `karabiner.json` from the `karabiner_rules.ts`.

```
yarn run watch
```

watches the TypeScript files and rebuilds whenever they change.
