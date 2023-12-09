package main

const DIFFERENCE = 1e7
const XOR_CODE = 6139246
const ALPHABET = 67108863
const EXCEEDED = ALPHABET + 1

func koromo2maj(id int) int {
	data := id ^ XOR_CODE
	temp := data & ALPHABET
	majsoul_id := 0

	temp = (temp&524287)<<7 | temp>>19

	majsoul_id = temp + (data & EXCEEDED) + DIFFERENCE
	return majsoul_id
}
