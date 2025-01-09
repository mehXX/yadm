.PHONY: icloud restic sync

sync: karabiner sync_config
karabiner_config: karabiner sync_config

sync_config:
	./bin/sync_config.sh

icloud_dotfiles:
	./bin/icloud_dotfiles.sh

karabiner:
	./bin/sync_karabiner.sh

is_up_to_date:
	./bin/check_up_to_date.sh

