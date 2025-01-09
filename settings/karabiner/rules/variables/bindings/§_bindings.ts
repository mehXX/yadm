import { KarabinerRules } from "../../../types";
import { createKeySubLayers, open } from "../../../utils";


export function NonUSBackSlashBindings(): KarabinerRules[] {
    return createKeySubLayers("non_us_backslash", {
        spacebar: {
            n: open("https://news.ycombinator.com"),
            j: open("https://jira.uzum.com/"),
            g: open("https://git.uzum.io/"),
            y: open("https://youtube.com/"),
            t: open("https://translate.yandex.com/"),
            l: open("https://linkedin.com"),
            c: open("https://chat.openai.com"),
            h: open("https://hn.algolia.com"),
            r: open("-a \"Google Chrome\" https://rezka.ag/"),
            p: open("https://perplexity.ai"),
            o: open("https://openrouter.ai/docs"),
            i: open("https://huggingface.co/spaces/Xenova/the-tokenizer-playground"),
            a: open("https://chat.lmsys.org"),
            v: open("https://vas3k.club"),
            1: open("raycast://extensions/benvp/audio-device/set-input-device"),
            2: open("raycast://extensions/benvp/audio-device/set-output-device")
        }
    });
}