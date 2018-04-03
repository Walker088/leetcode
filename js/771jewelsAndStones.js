/**
 * @param {string} J
 * @param {string} S
 * @return {number}
 */
var numJewelsInStones = function(J, S) {
    const jewels = Array.from(J);
    const stones = Array.from(S);
    let count = 0;

    for (let j of jewels){
        const myJewels = stones.filter(s => s === j);
        
        count += myJewels.length;
    }
    
    return count;
};

var numJewelsInStones2 = function(J, S) {
    const jewelReg = new RegExp(`[${J}]`, 'g');
    const myJewels = S.match(jewelReg);
    return myJewels ? myJewels.length : 0; 
};