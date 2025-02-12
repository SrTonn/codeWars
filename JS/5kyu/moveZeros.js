/**
 * link: https://www.codewars.com/kata/52597aa56021e91c93000cb0
 */

const moveZeros = function (arr) {
  return [...arr.filter(i => i !== 0), ...arr.filter(i => i === 0)]
}

console.log(moveZeros([false, 1, 0, 1, 2, 0, 1, 3, 'a']))
// returns[false,1,1,2,1,3,"a",0,0]
