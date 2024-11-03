import _default from "ts-jest";

export default function two_crystal_balls(breaks: boolean[]): number {
    const jumpAmount :number = Math.floor(Math.sqrt(breaks.length));

    let i :number = jumpAmount;
    for (; i < breaks.length; i += jumpAmount) {
        if (breaks[i]) {
            break;
        }
    }

    i -= jumpAmount;

    for (let j :number = 0; j < jumpAmount && i < breaks.length; ++j, ++i) {
        if (breaks[i]) {
            return i;
        }
    }

    return -1;
}