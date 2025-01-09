import { KarabinerRules } from "../../../types";

export const CapsBindings: KarabinerRules =
    {
        description: "caps + r as enter",
        manipulators: [
            {
                from: {
                    key_code: "e",
                    modifiers: {
                        mandatory: ["left_command", "left_control", "left_option", "left_shift"]
                    }
                },
                to: [
                    {
                        key_code: "return_or_enter",
                    },
                ],
                type: "basic"
            },
            {
                from: {
                    key_code: "d",
                    modifiers: {
                        mandatory: ["left_command", "left_control", "left_option", "left_shift"]
                    }
                },
                to: [
                    {
                        key_code: "delete_or_backspace",
                    },
                ],
                type: "basic"
            },

        ]
    }