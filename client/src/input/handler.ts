import { KeyMap } from "./keymap";

export class KeyboardHandler {
    map: KeyMap

    constructor(map: KeyMap) {
        this.map = map
    }

    start() {
        document.addEventListener("keydown", this.keyDown.bind(this), true)
        document.addEventListener("keyup", this.keyUp.bind(this), true)
        document.addEventListener("keypress", this.keyPress.bind(this), true)
    }
    stop() {
        document.removeEventListener("keydown", this.keyDown)
        document.removeEventListener("keyup", this.keyUp)
        document.removeEventListener("keypress", this.keyPress)
    }

    protected keyDown(e: KeyboardEvent) {
        this.map.keyDown(e)
    }
    protected keyUp(e: KeyboardEvent) {
        this.map.keyUp(e)
    }
    protected keyPress(e: KeyboardEvent) {
        this.map.keyPress(e)
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
