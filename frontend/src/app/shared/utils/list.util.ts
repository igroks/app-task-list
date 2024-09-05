export class ListUtil {
  static toSorted<T>(array: T[], key: keyof T, reverse: boolean): T[] {
    let reverseFactor = reverse ? -1 : 1;

    return array.sort((a: T, b: T) => {
      if (a[key] > b[key]) {
        return 1 * reverseFactor;
      } else if (a[key] < b[key]) {
          return -1 * reverseFactor;
      }
      return 0;
    })
  }
}
