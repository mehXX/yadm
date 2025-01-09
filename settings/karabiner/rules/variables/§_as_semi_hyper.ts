import { KarabinerRules } from "../../types";

export const NonUSBackSlashAsSemiHyper: KarabinerRules =
  {
      description: "Make § act as Control when held and non_us_backslash layer",
      manipulators: [
          {
              description: "Make § act as Control when held",
              from: {
                  key_code: "non_us_backslash",
                  modifiers: {
                      "optional": [
                          "any"
                      ]
                  }
              },
              to: [
                  {
                      set_variable: {
                          name: "non_us_backslash",
                          value: 1
                      }
                  },
                  {
                      key_code: "left_control",
                      modifiers: [
                          "left_option",
                          "left_shift"
                      ]
                  }
              ],
              to_after_key_up: [
                  {
                      set_variable: {
                          name: "non_us_backslash",
                          value: 0
                      }
                  }
              ],
              conditions: [
                  {
                      type: "frontmost_application_unless",
                      bundle_identifiers: [
                          "^com.spotify.client"
                      ]
                  },
              ],
              type: "basic"
          }
      ]
  };
