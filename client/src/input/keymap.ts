import { Action } from "./action"

let keys = {
    "KeyW|false|false|false|false": (e: KeyboardEvent) => {
        console.log("Up")
    },
    "KeyA": (e: KeyboardEvent) => {
        console.log("Left")
    },
    "KeyS": (e: KeyboardEvent) => {
        console.log("Down")
    },
    "KeyD": (e: KeyboardEvent) => {
        console.log("Right")
    },
}

/**
 * KeyMap connects key presses to game actions.
 * 
 * @see {@link https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent|Keyboard Event}
 */
export class KeyMap {
    protected map: Map<string, Action>

    constructor() {
        this.map = new Map<string, Action>()
    }

    setKey(code: string, action: Action) {
        this.map.set(code, action)
    }

    setKeyFromEvent(e: KeyboardEvent, action: Action) {
        let code = KeyMap.keyCode(e)
        this.setKey(code, action)
    }

    getAction(code: string): Action | null {
        if (this.map.has(code)) {
            return this.map.get(code)
        }
        return null
    }

    getActionFromEvent(e: KeyboardEvent): Action | null {
        let code = KeyMap.keyCode(e)
        return this.getAction(code)
    }

    /**
     * keyCode generates a string based on what keys are being pressed
     * @param e 
     * @returns 
     */
    static keyCode(e: KeyboardEvent): string {
        return `${e.code}${(e.altKey) ? "+alt" : ""}${(e.ctrlKey) ? "+ctrl" : ""}${(e.metaKey) ? "+meta" : ""}${(e.shiftKey) ? "+shift" : ""}`
    }
}
