#!/bin/bash

cd karabiner || return
yarn run build  && \
rm -f "$HOME/.config/karabiner/karabiner.json" && \
cp "$HOME/settings/karabiner/karabiner.json" "$HOME/.config/karabiner"