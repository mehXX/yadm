import {KarabinerRules} from "../../types";

export const CopyCodeBlockChatGPTOrion: KarabinerRules =
    {
        description: "Map Control+C to Command+Shift+; in Safari",
        manipulators: [
            {
                description: "Map Control+C to Command+Shift+; in Safari",
                conditions: [
                    {
                        bundle_identifiers: [
                            "^com.kagi.kagimacOS$"
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
