import {KarabinerRules} from "../../types";

export const CopyCodeBlockChatGPTSafari: KarabinerRules =
    {
        description: "Map Control+C to Command+Shift+; in Safari",
        manipulators: [
            {
                description: "Map Control+C to Command+Shift+; in Safari",
                conditions: [
                    {
                        bundle_identifiers: [
                            "^com.apple.Safari$"
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
                    }
                },
                to: [
                    {
                        key_code: "semicolon",
                        modifiers: [
                            "left_command",
                            "left_shift"
                        ]
                    }
                ],
                type: "basic"
            },
        ]

    }
