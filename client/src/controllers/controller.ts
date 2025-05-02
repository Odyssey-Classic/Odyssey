import { Character } from "../characters/character"

export interface Controller {
    character: Character
    update(deltaMS: number): void
}
