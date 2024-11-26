import { KeyMap } from "./keymap";

export class KeyboardHandler {
    map: KeyMap

    constructor(map: KeyMap) {
        this.map = map
    }

    start() {
        document.addEventListener("keydown", this.keyDown.bind(this), true)
        document.addEventListener("keyup", this.keyUp.bind(this), true)
    }
    stop() {
        document.removeEventListener("keydown", this.keyDown)
        document.removeEventListener("keyup", this.keyUp)
    }

    protected keyDown(e: KeyboardEvent) {
        console.log(KeyMap.keyCode(e))
        let action = this.map.getActionFromEvent(e)
    }
    protected keyUp(e: KeyboardEvent) {
        console.log(KeyMap.keyCode(e))
        let action = this.map.getActionFromEvent(e)
    }
}
