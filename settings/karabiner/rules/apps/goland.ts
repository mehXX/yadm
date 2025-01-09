import { KarabinerRules } from "../../types";

export const GolandChangeLanguage: KarabinerRules =
    {
        description: "Change language while openning goland/iterm",
        manipulators: [
            {
                description: "new tab terminal goland",
                from: {
                    key_code: "2",
                    modifiers: {
                        mandatory: ["left_command"]
                    }
                },
                to: [
                    {
                        key_code: "2",
                        modifiers: ["left_command"]
                    },
                    {
                        select_input_source: {
                            input_source_id: "com.apple.keylayout.ABC"
                        }
                    }
                ],
                type: "basic",
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.jetbrains.goland$"
                        ]
                    }
                ]
            },
            {
                description: "goland switch to english",
                from: {
                    key_code: "3",
                    modifiers: {
                        mandatory: ["left_option"]
                    }
                },
                to: [
                    {
                        key_code: "3",
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
            {
                description: "new tab iterm2",
                from: {
                    key_code: "k",
                    modifiers: {
                        mandatory: ["left_command", "left_shift"]
                    }
                },
                to: [
                    {
                        key_code: "k",
                        modifiers: ["left_option"]
                    },
                    {
                        key_code: "f",
                        modifiers: ["left_command", "left_shift"]
                    },
                ],
                type: "basic",
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.jetbrains.goland$"
                        ]
                    }
                ],
            },
            {
                description: "cmd + arrow right",
                from: {
                    key_code: "right_arrow",
                    modifiers: {
                        mandatory: ["left_command"]
                    }
                },
                to: [
                    {
                        key_code: "right_arrow",
                        modifiers: ["fn"]
                    },
                ],
                type: "basic",
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.jetbrains.goland$"
                        ]
                    }
                ],
            },
            {
                description: "cmd + arrow left",
                from: {
                    key_code: "left_arrow",
                    modifiers: {
                        mandatory: ["left_command"]
                    }
                },
                to: [
                    {
                        key_code: "left_arrow",
                        modifiers: ["fn"]
                    },
                ],
                type: "basic",
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.jetbrains.goland$"
                        ]
                    }
                ],
            },
        ]
    };
