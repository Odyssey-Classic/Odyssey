import { KeyboardHandler } from "./handler"
import { Movement } from "../movement/movement"
import { Direction } from "../characters/direction"
import { Character } from "../characters/character" // Import Character

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

    update(character: Character) { // Changed type to Character
        const keys = this.keyboard.getKeys(true)
        let moving = false

        keys.forEach((key) => {
            switch (key) {
                case "KeyW":
                    this.movement.startMove(character, Direction.Up) // Correct argument order
                    moving = true
                    break
                case "KeyS":
                    this.movement.startMove(character, Direction.Down) // Correct argument order
                    moving = true
                    break
                case "KeyA":
                    this.movement.startMove(character, Direction.Left) // Correct argument order
                    moving = true
                    break
                case "KeyD":
                    this.movement.startMove(character, Direction.Right) // Correct argument order
                    moving = true
                    break
            }
        })

        if (!moving) {
            this.movement.stopMove(character) // Pass character to stopMove
        }
    }
}
