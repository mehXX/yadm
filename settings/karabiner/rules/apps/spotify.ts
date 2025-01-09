import { KarabinerRules } from "../../types";

export const FnSpotifyAddToQueueLikedSongs: KarabinerRules =
    {
        description: "Fn as add to queue shortcut spotify in liked songs/etc",
        manipulators: [
            {
                type: "basic",
                from: {
                    key_code: "a",
                    modifiers: {
                        "mandatory": [
                            "left_command", "left_option", "left_shift",
                        ],
                    }
                },
                to: [
                    {
                        pointing_button: "button2"
                    },
                    {
                        key_code: "down_arrow"
                    },
                    {
                        key_code: "down_arrow"
                    },
                    {
                        key_code: "down_arrow"
                    },
                    {
                        key_code: "keypad_enter"
                    },
                    {
                        pointing_button: "button2"
                    },
                    {
                        pointing_button: "button1"
                    }
                ],
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.spotify.client"
                        ]
                    },
                ]
            }
        ]
    };
