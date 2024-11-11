import { KeyMap } from "./keymap";

export class KeyboardHandler {
    map: KeyMap

    constructor(map: KeyMap) {
        this.map = map
    }

    start() {
        document.addEventListener("keydown", this.keydown.bind(this), true)
        document.addEventListener("keyup", this.keyup.bind(this), true)
        document.addEventListener("keypress", this.keypress.bind(this), true)
    }
    stop() {
        document.removeEventListener("keydown", this.keydown)
        document.removeEventListener("keyup", this.keyup)
        document.removeEventListener("keypress", this.keypress)
    }

    protected keydown(e: KeyboardEvent) {
        this.log(e)
    }
    protected keyup(e: KeyboardEvent) {
        this.log(e)
    }
    protected keypress(e: KeyboardEvent) {
        this.log(e)
    }

    protected log(e: KeyboardEvent) {
        if (!e.repeat) {
            console.log({
                "alt": e.altKey,
                "code": e.code,
                "ctrl": e.ctrlKey,
                "key": e.key,
                "keyCode": e.keyCode,
                "repeat": e.repeat,
                "shift": e.shiftKey,
                "type": e.type,
                "which": e.which
            })
            return
        }

        console.log("repeat: ", e.key, e.type)
    }
}
