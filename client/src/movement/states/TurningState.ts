import { Movement } from "../movement";
import { MovementState } from "./MovementState";
import { IdleState } from "./IdleState";
import { Direction } from "../../characters/direction";

export class TurningState implements MovementState {
    private movement: Movement;

    constructor(movement: Movement) {
        this.movement = movement;
        this.movement.turnTimeMS = 100; // Turn time in milliseconds
    }

    update(deltaMS: number): void {
        this.movement.turnTimeMS -= deltaMS;
        if (this.movement.turnTimeMS <= 0) {
            this.movement.setState(new IdleState(this.movement));
        }
    }

    startMove(direction: Direction): void {
        // No specific behavior for starting a move in TurningState
    }

    stopMove(): void {
        // No specific behavior for stopping in TurningState
    }
}
