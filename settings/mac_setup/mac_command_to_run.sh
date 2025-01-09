#!/bin/bash

# show language input indicator
defaults write kCFPreferencesAnyApplication TSMLanguageIndicatorEnabled 0
defaults write kCFPreferencesAnyApplication "com.apple.keyboard.fnState" -int 0
defaults write kCFPreferencesAnyApplication "com.apple.mouse.doubleClickThreshold" -string "1.7"
defaults write kCFPreferencesAnyApplication "com.apple.mouse.scaling" -string "2.5"
defaults write kCFPreferencesAnyApplication "com.apple.sound.beep.feedback" -int 0
defaults write kCFPreferencesAnyApplication "com.apple.sound.beep.flash" -int 0
defaults write kCFPreferencesAnyApplication "com.apple.sound.beep.volume" -int 0
defaults write kCFPreferencesAnyApplication "com.apple.sound.uiaudio.enabled" -int 0
defaults write kCFPreferencesAnyApplication "com.apple.springing.delay" -int 0
defaults write kCFPreferencesAnyApplication "com.apple.springing.enabled" -int 1
defaults write kCFPreferencesAnyApplication "com.apple.swipescrolldirection" -int 1
defaults write kCFPreferencesAnyApplication "com.apple.trackpad.forceClick" -int 1
defaults write kCFPreferencesAnyApplication "com.apple.trackpad.scaling" -string "1.5"
defaults write kCFPreferencesAnyApplication "com.apple.trackpad.scrolling" -string "0.3125"
defaults write kCFPreferencesAnyApplication AppleLanguages -array "en-RU" "ru-RU"


# Disable holding a key for the special character menu
defaults write -g ApplePressAndHoldEnabled -bool false

# Make keys repeat really fast
defaults write -g KeyRepeat -int 1
defaults write NSGlobalDomain InitialKeyRepeat -int 10

# enable ssh
systemsetup -setremotelogin on &>/dev/null

# Sets default save target to be a local disk, not iCloud
defaults write -g NSDocumentSaveNewDocumentsToCloud -bool false

# Show status bar
defaults write com.apple.Safari ShowStatusBar -bool false
defaults write com.apple.Safari ShowOverlayStatusBar -bool false
# Show the full URL in the address bar (note: this still hides the scheme)
defaults write com.apple.Safari ShowFullURLInSmartSearchField -bool false
# Safari opens with: last session
defaults write com.apple.Safari AlwaysRestoreSessionAtLaunch -bool true
# Enable the Develop menu and the Web Inspector in Safari
defaults write com.apple.Safari IncludeDevelopMenu -bool true
# Enable Safari’s debug menu
defaults write com.apple.Safari IncludeInternalDebugMenu -bool true
# Update extensions automatically
defaults write com.apple.Safari InstallExtensionUpdatesAutomatically -bool true
# Make Safari’s search banners default to Contains instead of Starts With
defaults write com.apple.Safari FindOnPageMatchesWordStartsOnly -bool false
# Website use of location services
# 0 = Deny without prompting
# 1 = Prompt for each website once each day
# 2 = Prompt for each website one time only
defaults write com.apple.Safari SafariGeolocationPermissionPolicy -int 2

defaults write com.apple.Safari SendDoNotTrackHTTPHeader -bool true


# Disable "removing from iCloud Drive" warnings
defaults write com.apple.finder FXEnableRemoveFromICloudDriveWarning -bool false

# Tweak the Dock
defaults write com.apple.dock autohide -bool true
defaults write com.apple.dock autohide-delay -bool false
defaults write com.apple.dock show-recents -bool false
defaults write com.apple.dock largesize -int 32
defaults write com.apple.dock orientation right
defaults write com.apple.Dock showhidden -bool yes
defaults write com.apple.dock tilesize -int 57

# Disable smart quotes and dashes as they're annoying when typing code:
defaults write NSGlobalDomain NSAutomaticQuoteSubstitutionEnabled -bool false
defaults write NSGlobalDomain NSAutomaticDashSubstitutionEnabled -bool false

# Disable the crash reporter dialog:
defaults write com.apple.CrashReporter DialogType -string "none"

# FINDER
# Disable the warning before emptying the Trash:
defaults write com.apple.finder WarnOnEmptyTrash -bool false
# Use list view in all Finder windows by default:
defaults write com.apple.finder FXPreferredViewStyle -string "Nlsv"
# Show path bar in Finder:
defaults write com.apple.finder ShowPathbar -bool true


# Disable the "Are you sure you want to open this application?" dialog:
defaults write com.apple.LaunchServices LSQuarantine -bool false


# Turn off automatic updates:
defaults write com.apple.SoftwareUpdate AutomaticCheckEnabled -bool false
# Turn off auto-download of updates:
defaults write com.apple.SoftwareUpdate AutomaticDownload -int 0
# Turn off auto-install of macOS updates:
defaults write com.apple.commerce AutoUpdate -bool false

# disable siri
defaults write com.apple.assistant.support "Assistant Enabled" -bool false

# Disable Game Center:
defaults write com.apple.gamed Disabled -bool true

# To show the Library folder in Finder
chflags nohidden ~/Library
# Avoid creating .DS_Store files on network or USB volumes:
defaults write com.apple.desktopservices DSDontWriteNetworkStores -bool true


#  Turn Off Spotlight Indexing
# sudo mdutil -a -i off

# Mail.app

# Disable send and reply animations in Mail.app
defaults write com.apple.mail DisableReplyAnimations -bool true
defaults write com.apple.mail DisableSendAnimations -bool true

# Copy email addresses as `foo@example.com` instead of `Foo Bar <foo@example.com>` in Mail.app
defaults write com.apple.mail AddressesIncludeNameOnPasteboard -bool false

# Add the keyboard shortcut ⌘ + Enter to send an email in Mail.app
defaults write com.apple.mail NSUserKeyEquivalents -dict-add "Send" "@\U21a9"

# Display emails in threaded mode, sorted by date (oldest at the top)
defaults write com.apple.mail DraftsViewerAttributes -dict-add "DisplayInThreadedMode" -string "yes"
defaults write com.apple.mail DraftsViewerAttributes -dict-add "SortedDescending" -string "yes"
defaults write com.apple.mail DraftsViewerAttributes -dict-add "SortOrder" -string "received-date"

# Disable inline attachments (just show the icons)
defaults write com.apple.mail DisableInlineAttachmentViewing -bool true