// 'abba' & 'baab' == true
// 'abba' & 'bbaa' == true
// 'abba' & 'abbba' == false
// 'abba' & 'abca' == false

function anagrams(word, words) {
  return words.filter((el) => el.split('').sort().join('') === word.split('').sort().join(''))
}

console.log(anagrams('abba', ['aabb', 'abcd', 'bbaa', 'dada'])) // => ['aabb', 'bbaa']

console.log(anagrams('racer', ['crazer', 'carer', 'racar', 'caers', 'racer'])) // => ['carer', 'racer']

console.log(anagrams('laser', ['lazing', 'lazy', 'lacer'])) // => []