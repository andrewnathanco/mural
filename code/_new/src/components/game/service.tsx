export function get_digits(key: number): { row: number; col: number } {
  if (key < 0 || key > 99 || isNaN(key) || !Number.isInteger(key)) {
    return { row: 0, col: 0 };
  }

  const tens = Math.floor(key / 10);
  const ones = key % 10;

  return { row: tens, col: ones };
}

export function get_penalty_from_key(key: number) {
  const { row, col } = get_digits(key);
  return Math.min(row, col, 10 - row - 1, 10 - col - 1) + 1;
}
