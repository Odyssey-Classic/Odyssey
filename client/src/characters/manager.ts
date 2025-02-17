import { Character } from "./character";

export class Manager {
    characters: Character[]

    update(deltaMS: number) {
        this.characters.forEach((c) => {
            c.update(deltaMS)
        })
    }
}
