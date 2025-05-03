import { describe, expect, it, beforeEach } from '@jest/globals'

import { Movement } from "./movement";
import { Direction } from "../characters/direction";
import { Character } from "../characters/character";
import { Config } from "../config/config";

describe("Movement.canMove", () => {
    let movement: Movement;
    let character: Character;

    beforeEach(() => {
        movement = new Movement();
        character = new Character();
        Config.columns = 10; // Set map dimensions
        Config.rows = 10;
    });

    it("should allow movement within map boundaries", () => {
        character.position.x = 1
        character.position.y = 1
        expect(movement.canMove(character, Direction.Up)).toBe(true);
        expect(movement.canMove(character, Direction.Down)).toBe(true);
        expect(movement.canMove(character, Direction.Left)).toBe(true);
        expect(movement.canMove(character, Direction.Right)).toBe(true);
    });

    it("should prevent movement beyond the top edge of the map", () => {
        character.position.y = 0; // Top edge
        expect(movement.canMove(character, Direction.Up)).toBe(false);
    });

    it("should prevent movement beyond the bottom edge of the map", () => {
        character.position.y = Config.rows - 1; // Bottom edge
        expect(movement.canMove(character, Direction.Down)).toBe(false);
    });

    it("should prevent movement beyond the left edge of the map", () => {
        character.position.x = 0; // Left edge
        expect(movement.canMove(character, Direction.Left)).toBe(false);
    });

    it("should prevent movement beyond the right edge of the map", () => {
        character.position.x = Config.columns - 1; // Right edge
        expect(movement.canMove(character, Direction.Right)).toBe(false);
    });
});
