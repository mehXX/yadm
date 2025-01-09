import * as fs from "fs";
import { KarabinerRules, SimpleKarabinerRules } from "./types";
import {CopyCodeBlockRayCast, RaycastSwitchToEnglish} from "./rules/apps/raycast";
import {MapControlSpaceToOptionCMD} from "./rules/remap/control_space_to_option_cmd";
import {
    MapRightCmdOptionToChangeLanguage
} from "./rules/languages/remap_map_language_change";
import {CapsHyper} from "./rules/variables/caps_hyper";
import {NonUSBackSlashAsSemiHyper} from "./rules/variables/§_as_semi_hyper";
import {CommandsNumberToChangeChatsTelegram, FnRightClickTelegram} from "./rules/apps/telegram";
import {FnSpotifyAddToQueueLikedSongs} from "./rules/apps/spotify";
import {GolandChangeLanguage} from "./rules/apps/goland";
import {NonUSBackSlashBindings} from "./rules/variables/bindings/§_bindings";
import {SlackRules} from "./rules/apps/slack";
import {CopyCodeBlockChatGPTSafari} from "./rules/apps/safari";
import {CapsBindings} from "./rules/variables/bindings/caps_bindings";
import {CopyCodeBlockChatGPTChrome} from "./rules/apps/chrome";
import {CmdShiftVNoStyle} from "./rules/remap/semi_hyper_key_slash";
import {FnToSemiModifier} from "./rules/remap/fn_to_brackets";
import {NotesBackForth} from "./rules/apps/notes";

const simple_modifications: SimpleKarabinerRules[] = [
    {
        from: {
            "apple_vendor_top_case_key_code": "keyboard_fn"
        },
        to: [
            {
                "key_code": "left_control"
            }
        ]
    },
    {
        from: {
            "key_code": "left_control"
        },
        to: [
            {
                "apple_vendor_top_case_key_code": "keyboard_fn"
            }
        ]
    }
];

const rules: KarabinerRules[] = [
    // Remap
    MapControlSpaceToOptionCMD,
    // use numbers as symbols, use f1-f12 as numbers
    // Variables
    // Caps
    CapsHyper,
    CapsBindings,
    // §
    NonUSBackSlashAsSemiHyper,
    ...NonUSBackSlashBindings(),
    CmdShiftVNoStyle,
    // fn
    FnToSemiModifier,
    // Telegram
    FnRightClickTelegram,
    CommandsNumberToChangeChatsTelegram,
    // Safari
    CopyCodeBlockChatGPTSafari,
    // Languages
    MapRightCmdOptionToChangeLanguage,
    // Spotify
    FnSpotifyAddToQueueLikedSongs,
    // Raycast
    CopyCodeBlockRayCast,
    RaycastSwitchToEnglish,
    // iTerm
    // Goland
    GolandChangeLanguage,
    NotesBackForth,
    // Slack
    ...SlackRules,
    // Chrome
    CopyCodeBlockChatGPTChrome,
];

fs.writeFileSync(
  "karabiner.json",
  JSON.stringify(
    {
        global: {
            "ask_for_confirmation_before_quitting": false,
            "check_for_updates_on_startup": false,
            "show_in_menu_bar": false,
            "show_profile_name_in_menu_bar": false,
            "unsafe_ui": false
        },
        profiles: [
            {
                name: "Default",
                simple_modifications,
                complex_modifications: {
                    rules
                }
            }
        ]
    },
    null,
    2
  )
);
