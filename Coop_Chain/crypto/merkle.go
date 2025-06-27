package crypto

import (
	"crypto/sha256"
)

// MerkleRoot вычисляет Merkle-root для списка хэшей (каждый хэш — []byte)
func MerkleRoot(hashes [][]byte) []byte {
	n := len(hashes)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return hashes[0]
	}
	var nextLevel [][]byte
	for i := 0; i < n; i += 2 {
		if i+1 < n {
			combined := append(hashes[i], hashes[i+1]...)
			h := sha256.Sum256(combined)
			nextLevel = append(nextLevel, h[:])
		} else {
			h := sha256.Sum256(hashes[i])
			nextLevel = append(nextLevel, h[:])
		}
	}
	return MerkleRoot(nextLevel)
}

// MerkleProof возвращает proof для элемента с индексом idx
// (упрощённая версия, возвращает список соседних хэшей)
func MerkleProof(hashes [][]byte, idx int) [][]byte {
	var proof [][]byte
	n := len(hashes)
	if idx < 0 || idx >= n {
		return proof
	}
	for n > 1 {
		var nextLevel [][]byte
		for i := 0; i < n; i += 2 {
			var pair []byte
			if i+1 < n {
				pair = append(hashes[i], hashes[i+1]...)
			} else {
				pair = hashes[i]
			}
			h := sha256.Sum256(pair)
			nextLevel = append(nextLevel, h[:])
			if i == idx || i+1 == idx {
				if i == idx && i+1 < n {
					proof = append(proof, hashes[i+1])
				} else if i+1 == idx {
					proof = append(proof, hashes[i])
				}
			}
		}
		idx /= 2
		hashes = nextLevel
		n = len(hashes)
	}
	return proof
}
