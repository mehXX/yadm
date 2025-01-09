import {KarabinerRules} from "../../types";

export const SlackChats: KarabinerRules =
    {
        description: "Map cmd + number to switch between slack chats",
        manipulators: [
            {
                description: "open chat with myself",
                conditions: [
                    {
                        bundle_identifiers: [
                            "^com.tinyspeck.slackmacgap$"
                        ],
                        type: "frontmost_application_if"
                    }
                ],
                from: {
                    key_code: "1",
                    modifiers: {
                        "mandatory": [
                            "left_command"
                        ],
                    }
                },
                to: [
                    {
                        shell_command: "open -a \"Slack\" \"slack://channel?team=T03CR7VBN0N&id=D06ME5C2RT5\"",
                    }
                ],
                type: "basic"
            },
            {
                description: "backend",
                conditions: [
                    {
                        bundle_identifiers: [
                            "^com.tinyspeck.slackmacgap$"
                        ],
                        type: "frontmost_application_if"
                    }
                ],
                from: {
                    key_code: "2",
                    modifiers: {
                        "mandatory": [
                            "left_command"
                        ],
                    }
                },
                to: [
                    {
                        shell_command: "open -a \"Slack\" \"slack://channel?team=T03CR7VBN0N&id=C064BU913GR\"",
                    }
                ],
                type: "basic"
            },
            {
                description: "dev",
                conditions: [
                    {
                        bundle_identifiers: [
                            "^com.tinyspeck.slackmacgap$"
                        ],
                        type: "frontmost_application_if"
                    }
                ],
                from: {
                    key_code: "3",
                    modifiers: {
                        "mandatory": [
                            "left_command"
                        ],
                    }
                },
                to: [
                    {
                        shell_command: "open -a \"Slack\" \"slack://channel?team=T03CR7VBN0N&id=C04QLGZ7492\"",
                    }
                ],
                type: "basic"
            },
            {
                description: "alerts",
                conditions: [
                    {
                        bundle_identifiers: [
                            "^com.tinyspeck.slackmacgap$"
                        ],
                        type: "frontmost_application_if"
                    }
                ],
                from: {
                    key_code: "4",
                    modifiers: {
                        "mandatory": [
                            "left_command"
                        ],
                    }
                },
                to: [
                    {
                        shell_command: "open -a \"Slack\" \"slack://channel?team=T03CR7VBN0N&id=C06KKQLE2N8\"",
                    }
                ],
                type: "basic"
            },
            {
                description: "Ivan",
                conditions: [
                    {
                        bundle_identifiers: [
                            "^com.tinyspeck.slackmacgap$"
                        ],
                        type: "frontmost_application_if"
                    }
                ],
                from: {
                    key_code: "5",
                    modifiers: {
                        "mandatory": [
                            "left_command"
                        ],
                    }
                },
                to: [
                    {
                        shell_command: "open -a \"Slack\" \"slack://channel?team=T03CR7VBN0N&id=D06NE86L6NL\"",
                    }
                ],
                type: "basic"
            },
            {
                description: "uzum",
                conditions: [
                    {
                        bundle_identifiers: [
                            "^com.tinyspeck.slackmacgap$"
                        ],
                        type: "frontmost_application_if"
                    }
                ],
                from: {
                    key_code: "1",
                    modifiers: {
                        "mandatory": [
                            "left_control"
                        ],
                    }
                },
                to: [
                    {
                        key_code: "1",
                        modifiers: [
                            "left_command"
                        ]
                    },
                ],
                type: "basic"
            },
            {
                description: "raycast",
                conditions: [
                    {
                        bundle_identifiers: [
                            "^com.tinyspeck.slackmacgap$"
                        ],
                        type: "frontmost_application_if"
                    }
                ],
                from: {
                    key_code: "2",
                    modifiers: {
                        "mandatory": [
                            "left_control"
                        ],
                    }
                },
                to: [
                    {
                        key_code: "2",
                        modifiers: [
                            "left_command"
                        ]
                    },
                ],
                type: "basic"
            },
        ]
    }

export const SlackRules: KarabinerRules[] = [
    SlackChats
]
