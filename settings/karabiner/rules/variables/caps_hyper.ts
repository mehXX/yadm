import { KarabinerRules } from "../../types";

export const CapsHyper: KarabinerRules =
  {
      description: "Caps -> Hyper key and hyper layer",
      manipulators: [
          {
              description: "Caps -> Hyper Key",
              from: {
                  key_code: "caps_lock",
                  modifiers: {
                      optional: ["any"]
                  }
              },
              to: [
                  {
                      set_variable: {
                          name: "hyper",
                          value: 1
                      }
                  },
                  {
                      key_code: "left_shift",
                      modifiers: ["left_command", "left_control", "left_option"],
                  }
              ],
              to_after_key_up: [
                  {
                      set_variable: {
                          name: "hyper",
                          value: 0
                      }
                  }
              ],
              type: "basic"
          },

      ]
  };