import { Controller } from "./controller"
import { Character } from "../characters/character"
import { Direction } from "../characters/direction"

export class AIController implements Controller {
    character: Character
    private timer: number = 0

    constructor(character: Character) {
        this.character = character
    }

    update(deltaMS: number) {
        this.timer += deltaMS
        if (this.timer > 1000) { // Change direction every second
            this.timer = 0
            const directions = [Direction.Up, Direction.Down, Direction.Left, Direction.Right]
            const randomDirection = directions[Math.floor(Math.random() * directions.length)]
            this.character.movement.startMove(randomDirection) // Ensure movement starts
        }

        // Update the character's movement
        this.character.update(deltaMS)
    }
}
