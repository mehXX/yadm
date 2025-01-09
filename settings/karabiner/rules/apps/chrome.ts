import {KarabinerRules} from "../../types";

export const CopyCodeBlockChatGPTChrome: KarabinerRules =
    {
        description: "Map Control+C to Command+Shift+; in Chrome",
        manipulators: [
            {
                description: "Map Control+C to Command+Shift+; in Chrome",
                conditions: [
                    {
                        bundle_identifiers: [
                            "^com.google.Chrome$"
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
            {
                conditions: [
                    {
                        bundle_identifiers: [
                            "^com.google.Chrome$"
                        ],
                        type: "frontmost_application_if"
                    }
                ],
                    from: {
                    key_code: "c",
                    modifiers: {
                        mandatory: ["left_command", "left_control", "left_option", "left_shift"]
                    }
                },
                to: [
                    {
                        key_code: "l",
                        modifiers: ["left_command"]
                    },
                    {
                        key_code: "c",
                        modifiers: ["left_command"]
                    },
                    {
                        key_code: "escape",
                    },
                    {
                        key_code: "escape",
                    },
                ],

                type: "basic"
            }
        ]
    }

