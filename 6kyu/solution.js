/**
 * link: https://www.codewars.com/kata/51b6249c4612257ac0000005/train/javascript
 */

const table = {
  I: 1,
  V: 5,
  X: 10,
  L: 50,
  C: 100,
  D: 500,
  M: 1000
}

function solution(roman){
  roman = roman.toUpperCase().split('')
  let counterRepeated = 1
  return roman.reduce((acc, curr, index) => {
    if(counterRepeated > 3) return 'Error'
    let next = table[roman[index + 1]] || 0

    if(table[curr] >= next){
      counterRepeated = table[curr] === next ? counterRepeated+1 : 1
      return acc + table[curr]
    } else {
      return acc - table[curr]
    }
  }, 0)
}

// describe("Tests", () => {
//   it("test", () => {
console.log(solution('XXI'), 21)
console.log(solution('I'), 1)
console.log(solution('IV'), 4)
console.log(solution('MMVIII'), 2008)
console.log(solution('MDCLXVI'), 1666)
// });
// });
