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
    }

    keyDown(code: string) {
        let key = this.keys.get(code)
        if (!this.keys.has(code)) {
            key = new state(code)
            this.keys.set(code, key)
        }

        key.active = true
        key.clearable = false
    }

    keyUp(code: string) {
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
    code: string
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
