/* eslint-disable no-console */
//link: https://www.codewars.com/kata/5a3dd29055519e23ec000074/train/javascript
function checkExam(array1, array2) {
  let score = 0
  for(let i = 0; i < array1.length; i++) {
    if(array1[i] === array2[i]) {
      score += 4
    } else if(array2[i] !== '') {
      score -= 1
    }
  }
  return score < 0 ? 0 : score
}

console.log(checkExam(['a', 'a', 'b', 'b'], ['a', 'c', 'b', 'd']), 6)
console.log(checkExam(['a', 'a', 'c', 'b'], ['a', 'a', 'b', '']), 7)
console.log(checkExam(['a', 'a', 'b', 'c'], ['a', 'a', 'b', 'c']), 16)
console.log(checkExam(['b', 'c', 'b', 'a'], ['', 'a', 'a', 'c']), 0)

/**
 * describe("Basic tests",() => {
  it("Testing for fixed tests", () => {
    assert.strictEqual(checkExam(["a", "a", "b", "b"], ["a", "c", "b", "d"]), 6);
    assert.strictEqual(checkExam(["a", "a", "c", "b"], ["a", "a", "b",  ""]), 7);
    assert.strictEqual(checkExam(["a", "a", "b", "c"], ["a", "a", "b", "c"]), 16);
    assert.strictEqual(checkExam(["b", "c", "b", "a"], ["",  "a", "a", "c"]), 0);  
  });
});
 */