import { Direction } from "../../characters/direction";

export interface MovementState {
    update(deltaMS: number): void;
    startMove(direction: Direction): void;
    stopMove(): void;
}
