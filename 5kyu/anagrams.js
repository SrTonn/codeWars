// 'abba' & 'baab' == true
// 'abba' & 'bbaa' == true
// 'abba' & 'abbba' == false
// 'abba' & 'abca' == false

function anagrams(word, words) {
  return words.filter((el) => el.split('').sort().join('') === word.split('').sort().join(''))
}

console.log(anagrams('abba', ['aabab', 'abcd', 'bbaaa', 'dada']))

const str = 'word'

console.log(""+[...str].sort())
console.log([1,2,3]==[1,2,3])

console.log(str.split`${d}`[0])

                    
