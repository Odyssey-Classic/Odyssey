const metaCodes: string[] = [
    "AltLeft+alt",
    "AltRight+alt",
    "ControlLeft+ctrl",
    "ControlRight+ctrl",
    "MetaLeft+meta",
    "MetaRight+meta",
    "ShiftLeft+shift",
    "ShiftRight+shift",
] as const

const metaModifiers: Map<string, string> = new Map<string, string>([
    ["AltLeft", "+alt"],
    ["AltRight", "+alt"],
    ["ControlLeft", "+ctrl"],
    ["ControlRight", "+ctrl"],
    ["MetaLeft", "+meta"],
    ["MetaRight", "+meta"],
    ["ShiftLeft", "+shift"],
    ["ShiftRight", "+shift"],
])

export class KeyState {
    private keys: Map<string, state>

    constructor() {
        this.keys = new Map<string, state>()

        console.log("meta mod keys", metaModifiers.keys())
    }

    keyDown(code: string) {

        // Meta keys aren't usable individually.
        // Technical issue is that `keydown` will populate both the code and the
        // specific meta key field.
        // However, `keyup` will only have the code.
        // So, with our hash `keyCode()`, this means that meta keys will never
        // clear.
        if (metaCodes.includes(code)) {
            console.log("ignoring", code)
            return
        }

        let key = this.keys.get(code)
        if (!this.keys.has(code)) {
            key = new state(code)
            this.keys.set(code, key)
        }

        key.active = true
        key.clearable = false
    }

    keyUp(code: string) {
        // If a user releases the meta key before the other keys,
        // this needs to catch that and mark all combos including the meta key.
        if (metaModifiers.keys().find((v) => {
            return v.includes(code)
        })) {
            let metaCode = metaModifiers.get(code)
            console.log("clearing meta", code, metaCode)
            this.keys.forEach((key) => {
                key.clearable = key.code.includes(metaCode)
            })

            return
        }

        if (!this.keys.has(code)) {
            return
        }

        console.log("clearing", code)
        this.keys.get(code).clearable = true
    }

    /**
     * 
     * @param clear whether to clear the status of actions
     */
    getKeys(clear: boolean): string[] {
        let result: string[] = new Array<string>()

        this.keys.forEach((v, k, m) => {
            if (v.active) {
                result.push(v.code)
            }

            if (clear && v.clearable) {
                v.clear()
            }
        })

        return result
    }
}

class state {
    code: string // The combined key code for keys and modifiers.
    active: boolean = true
    clearable: boolean = false

    constructor(code: string) {
        this.code = code
    }

    clear() {
        this.active = false
        this.clearable = false
    }
}
