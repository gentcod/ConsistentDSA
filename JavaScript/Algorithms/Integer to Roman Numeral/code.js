const integerToRoman = (n) => {
	let roman = ""
	const keys = [1000,500,100,50,10,5,1]
	const literals = {
		1000: "M",
		500: "D",
		100: "C",
		50: "L",
		10: "X",
		5: "V",
		1 : "I",
	}

	for (let i = 0; i < keys.length; i++) {		
		if (n % keys[i] != n) {
			const div = (n - (n % keys[i])) / keys[i]			
			if (div == 4) {
				if (roman.length > 0 && literals[keys[i-1]] == roman[roman.length-1]) {					
					roman = roman.split("").slice(0,-1).join("")
					roman += literals[keys[i]] + literals[keys[i-2]]
				} else {
					roman += literals[keys[i]] + literals[keys[i-1]]
				}
			} else {
				roman += evalLiteral(div, literals[keys[i]])
			}
		}
		n = n % keys[i]
	}

	return roman
}

const evalLiteral = (div, str) => {
	let s = ""
	for (let i = div; i > 0; i--) {
		s += str
	}

	return s
}

console.log(integerToRoman(4));

