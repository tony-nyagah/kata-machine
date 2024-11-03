export default function bs_list(haystack: number[], needle: number): boolean {
    let low :number = 0;
    let high :number = haystack.length;

    do {
        const middle :number = Math.floor(low + (high - low) / 2);
        const value :number = haystack[middle];

        if (value === needle) {
            return true;
        } else if (value > needle) {
            high = middle;
        } else {
            low = middle + 1;
        }
    } while (low < high);

    return false;
}
