import { Direction } from "../../characters/direction";
import { Character } from "../../characters/character";

export interface MovementState {
    update(character: Character, deltaMS: number): void;
    startMove(character: Character, direction: Direction): void;
    stopMove(character: Character): void;
}
