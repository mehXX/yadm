import { KarabinerRules } from "../../types";

export const FnRightClickTelegram: KarabinerRules =
    {
        description: "Fn as right click in telegram",
        manipulators: [
            {
                type: "basic",
                from: {
                    key_code: "fn",
                },
                to: [
                    {
                        pointing_button: "button2"
                    }
                ],
                to_after_key_up: [
                    {
                        pointing_button: "button1"
                    }
                ],
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.tdesktop.Telegram"
                        ]
                    },
                ]
            }
        ]
    };

export const CommandsNumberToChangeChatsTelegram: KarabinerRules =
    {
        description: "command numbers to change chats in telegram",
        manipulators: [
            {
                type: "basic",
                from: {
                    key_code: "3",
                    modifiers: {
                        mandatory: [
                            "left_command",
                        ]
                    }
                },
                to: [
                    {
                        key_code: "2",
                        modifiers: ["left_command"],
                    },
                    {
                        key_code: "tab",
                        modifiers: ["left_control"],
                    },
                ],
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.tdesktop.Telegram$"
                        ]
                    },
                ]
            },
            {
                type: "basic",
                from: {
                    key_code: "4",
                    modifiers: {
                        mandatory: [
                            "left_command",
                        ]
                    }
                },
                to: [
                    {
                        key_code: "2",
                        modifiers: ["left_command"],
                    },
                    {
                        key_code: "tab",
                        modifiers: ["left_control"],
                    },
                    {
                        key_code: "tab",
                        modifiers: ["left_control"],
                    },
                ],
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.tdesktop.Telegram$"
                        ]
                    },
                ]
            },
            {
                type: "basic",
                from: {
                    key_code: "5",
                    modifiers: {
                        mandatory: [
                            "left_command",
                        ]
                    }
                },
                to: [
                    {
                        key_code: "2",
                        modifiers: ["left_command"],
                    },
                    {
                        key_code: "tab",
                        modifiers: ["left_control"],
                    },
                    {
                        key_code: "tab",
                        modifiers: ["left_control"],
                    },
                    {
                        key_code: "tab",
                        modifiers: ["left_control"],
                    },
                ],
                conditions: [
                    {
                        type: "frontmost_application_if",
                        bundle_identifiers: [
                            "^com.tdesktop.Telegram$"
                        ]
                    },
                ]
            },
        ]
    };

