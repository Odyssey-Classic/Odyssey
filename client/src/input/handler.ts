import { KeyState } from "./state"

export class KeyboardHandler {
    state: KeyState

    constructor() {
        this.state = new KeyState()
    }

    start() {
        document.addEventListener("keydown", this.keyDown.bind(this), true)
        document.addEventListener("keyup", this.keyUp.bind(this), true)
    }
    stop() {
        document.removeEventListener("keydown", this.keyDown)
        document.removeEventListener("keyup", this.keyUp)
    }

    getKeys(clear: boolean): string[] {
        return this.state.getKeys(clear)
    }

    protected keyDown(e: KeyboardEvent) {
        if (e.repeat) {
            console.debug("repeated key", e.code)
            return
        }

        console.log("down", e.code, KeyboardHandler.keyCode(e))

        this.state.keyDown(KeyboardHandler.keyCode(e))
    }

    protected keyUp(e: KeyboardEvent) {
        console.log("up", e.code, KeyboardHandler.keyCode(e))

        this.state.keyUp(KeyboardHandler.keyCode(e))
    }

    /**
     * keyCode generates a string based on what keys are being pressed.
     * This allows us to turn each combination of keys into a unique string
     * that can be used in a map.
     * @param e 
     * @returns
     */
    static keyCode(e: KeyboardEvent): string {
        return `${e.code}${(e.altKey) ? "+alt" : ""}${(e.ctrlKey) ? "+ctrl" : ""}${(e.metaKey) ? "+meta" : ""}${(e.shiftKey) ? "+shift" : ""}`
    }
}
