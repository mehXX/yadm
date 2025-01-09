import { KarabinerRules } from "../../types";

export const CmdShiftVNoStyle: KarabinerRules =
    {
        description: "cmd shift v no style",
        manipulators: [
            {
                description: "cmd shift v no style",
                type: "basic",
                from: {
                    key_code: "v",
                    modifiers: {
                        "mandatory": [
                            "left_command", "left_shift",
                        ]
                    }
                },
                to: [
                    {
                        key_code: "slash",
                        modifiers: [
                            "left_option", "left_control", "left_shift",
                        ]
                    }
                ]
            }
        ]
    }
