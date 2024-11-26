import { Action } from "./action"

export class ActionState {
    private actions: Map<Action, boolean>
    constructor() {
        this.actions = new Map<Action, boolean>()
    }

    keyDown(action: Action) {
        this.actions.set(action, true)
    }

    keyUp(action: Action) {
        this.actions.set(action, false)
    }

    /**
     * 
     * @param clear whether to clear the status of actions
     */
    getActions(clear = true): MapIterator<Action> {
        let result = this.actions.keys()

        if (clear) {
            this.actions.forEach((value, key, map) => {
                if (!value) {
                    map.delete(key)
                }
            })
        }

        return result
    }
}
