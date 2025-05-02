import { KeyboardHandler } from "./handler"
import { Movement } from "../movement/movement"
import { Direction } from "../characters/direction"

export class MovementInput {
    private keyboard: KeyboardHandler
    private movement: Movement

    constructor(movement: Movement) {
        this.keyboard = new KeyboardHandler()
        this.movement = movement
    }

    start() {
        this.keyboard.start()
    }

    stop() {
        this.keyboard.stop()
    }

    update() {
        const keys = this.keyboard.getKeys(true)
        let moving = false

        keys.forEach((key) => {
            switch (key) {
                case "KeyW":
                    this.movement.startMove(Direction.Up)
                    moving = true
                    break
                case "KeyS":
                    this.movement.startMove(Direction.Down)
                    moving = true
                    break
                case "KeyA":
                    this.movement.startMove(Direction.Left)
                    moving = true
                    break
                case "KeyD":
                    this.movement.startMove(Direction.Right)
                    moving = true
                    break
            }
        })

        if (!moving) {
            this.movement.stopMove()
        }
    }
}
