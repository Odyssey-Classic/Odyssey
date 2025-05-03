import { Controller } from "./controller"
import { Character } from "../characters/character"
import { MovementInput } from "../input/movement_input"

export class PlayerController implements Controller {
    character: Character
    private movementInput: MovementInput

    constructor(character: Character) {
        this.character = character
        this.movementInput = new MovementInput(character.movement)
        this.movementInput.start()
    }

    update(deltaMS: number) {
        this.movementInput.update(this.character)
    }
}
