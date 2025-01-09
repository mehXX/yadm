import { KarabinerRules } from "../../types";

export const MapRightCmdOptionToChangeLanguage: KarabinerRules =
    {
        description: "Right and Left Command for Russian and English",
        manipulators: [
            {
                description: "Right Command for Russian",
                type: "basic",
                from: {
                    key_code: "right_option",
                    modifiers: {
                        optional: ["any"]
                    }
                },
                to: [
                    {
                        select_input_source: {
                            input_source_id: "com.apple.keylayout.Russian"
                        }
                    }
                ]
            },
            {
                type: "basic",
                from: {
                    key_code: "right_command",
                    modifiers: {
                        optional: ["any"]
                    }
                },
                to: [
                    {
                        select_input_source: {
                            input_source_id: "com.apple.keylayout.ABC"
                        }
                    }
                ]
            },
        ]
    };
