import { KarabinerRules } from "../../types";

export const NotesBackForth: KarabinerRules =
    {
        description: "Change language while openning goland/iterm",
        manipulators: [
            {
                description: "go back",
                from: {
                    key_code: "open_bracket",
                    modifiers: {
                        mandatory: ["left_command"]
                    }
                },
                to: [
                    {
                        key_code: "open_bracket",
                        modifiers: ["left_option", "left_command"]
                    },
                ],
                type: "basic",
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.apple.Notes$"
                        ]
                    }
                ]
            },
            {
                description: "go forward",
                from: {
                    key_code: "close_bracket",
                    modifiers: {
                        mandatory: ["left_command"]
                    }
                },
                to: [
                    {
                        key_code: "close_bracket",
                        modifiers: ["left_option", "left_command"]
                    },
                ],
                type: "basic",
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.apple.Notes$"
                        ]
                    }
                ]
            },
            {
                description: "replace shortcut",
                from: {
                    key_code: "r",
                    modifiers: {
                        mandatory: ["left_command", "left_shift"]
                    }
                },
                to: [
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
                            "^com.apple.Notes$"
                        ]
                    }
                ]
            },
            {
                description: "search shortcut",
                from: {
                    key_code: "f",
                    modifiers: {
                        mandatory: ["left_command", "left_shift"]
                    }
                },
                to: [
                    {
                        key_code: "f",
                        modifiers: ["left_command", "left_option"]
                    },
                ],
                type: "basic",
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.apple.Notes$"
                        ]
                    }
                ]
            },
        ]
    };
