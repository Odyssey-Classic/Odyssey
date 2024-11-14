let keys = {
    "KeyW": (e: KeyboardEvent) => {
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
    keyDown(e: KeyboardEvent) {
        if (keys[e.code] != null) {
            keys[e.code](e)
        } else {
            console.log(e.code)
        }
    }

    keyUp(key: KeyboardEvent) {

    }

    keyPress(key: KeyboardEvent) {

    }
}
