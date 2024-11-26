export class Action {
    private name: string
    private clearable: boolean
    private active: boolean

    constructor(name: string) {
        this.name = name
    }

    down() {
        this.active = true
        this.clearable = false
    }

    up() {
        this.clearable = true
    }

    clear() {
        if (this.clearable) {
            this.active = false
        }
    }
}
