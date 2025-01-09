import {KarabinerRules} from "../../types";

export const CopyCodeBlockRayCast: KarabinerRules =
    {
        description: "Map Control+C to Option+Command+C in Raycast",
        manipulators: [
            {
                description: "Map Control+C to Option+Command+C in Raycast",
                conditions: [
                    {
                        bundle_identifiers: [
                            "com.raycast.macos"
                        ],
                        type: "frontmost_application_if"
                    }
                ],
                from: {
                    key_code: "c",
                    modifiers: {
                        "mandatory": [
                            "control"
                        ],
                        "optional": [
                            "any"
                        ]
                    }
                },
                to: [
                    {
                        key_code: "c",
                        modifiers: [
                            "left_option",
                            "left_command"
                        ]
                    }
                ],
                type: "basic"
            }
        ]
    };


export const RaycastSwitchToEnglish: KarabinerRules =
    {
        description: "Cmd change Language to abc on some shortcuts",
        manipulators: [
            {
                description: "raycast english layout",
                from: {
                    key_code: "spacebar",
                    modifiers: {
                        mandatory: ["left_command"]
                    }
                },
                to: [
                    {
                        key_code: "spacebar",
                        modifiers: ["left_command"]
                    },
                    {
                        select_input_source: {
                            input_source_id: "com.apple.keylayout.ABC"
                        }
                    }
                ],
                type: "basic"
            },
            {
                description: "raycast ai switch to english",
                from: {
                    key_code: "a",
                    modifiers: {
                        mandatory: ["left_option"]
                    }
                },
                to: [
                    {
                        key_code: "a",
                        modifiers: ["left_option"]
                    },
                    {
                        select_input_source: {
                            input_source_id: "com.apple.keylayout.ABC"
                        }
                    }
                ],
                type: "basic",
            },
        ]
    };


