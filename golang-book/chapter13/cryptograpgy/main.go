/*
	Hash function: Takes a set of data and reduces it to a smalled fixed size.
	In Go are broken into two categories: 
	CRYPTOGRAPHIC and NON-CRYPTOGRAPHIC

	The non-crypthographic can be found underneath the hash package and include 
	adler32, crc32, crc64 and fnv.

	The cryptographic hash functions are similar to their noncryptographic counterparts
	but they have the added property of being hard to reverse.
*/