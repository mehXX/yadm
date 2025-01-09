import { KarabinerRules } from "../../types";

export const FnToSemiModifier: KarabinerRules =
    {
        description: "fn to brackets",
        manipulators: [
            {
                description: "control + spacebar -> option + cmd",
                type: "basic",
                from: {
                    key_code: "fn",
                },
                to: [
                    {
                        key_code: "left_option",
                        modifiers: [
                            "left_shift", "left_option", "left_command"
                        ]
                    }
                ]
            }
        ]
    }
