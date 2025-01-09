import { KarabinerRules } from "../../types";

export const MapControlSpaceToOptionCMD: KarabinerRules =
  {
      description: "control + spacebar -> option + cmd",
      manipulators: [
          {
              description: "control + spacebar -> option + cmd",
              type: "basic",
              from: {
                  key_code: "spacebar",
                  modifiers: {
                      "mandatory": [
                          "left_control"
                      ]
                  }
              },
              to: [
                  {
                      key_code: "left_option",
                      modifiers: [
                          "left_command"
                      ]
                  }
              ]
          }
      ]
  }
