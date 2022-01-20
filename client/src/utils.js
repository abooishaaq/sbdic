export const duplicate = (arr) => {
    const set = new Set();
    for (const i of arr) {
        if (set.has(i)) return i;
        set.add(i);
    }
    return null;
};
