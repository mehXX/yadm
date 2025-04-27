.PHONY: icloud restic sync

sync: karabiner sync_config

config:
	./bin/sync_config.sh

karabiner:
	./bin/sync_karabiner.sh

backup:
	./bin/backup.sh

plist:
	./plist/sync_plist.sh