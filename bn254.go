package bn254

import "math/big"

func bigFromHex(hex string) *big.Int {
	if len(hex) > 1 && hex[:2] == "0x" {
		hex = hex[2:]
	}
	n, _ := new(big.Int).SetString(hex, 16)
	return n
}

var inp uint64 = 9786893198990664585

var modulus = fe{0x3c208c16d87cfd47, 0x97816a916871ca8d, 0xb85045b68181585d, 0x30644e72e131a029}

var zero = &fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000}

var r1 = &fe{0xd35d438dc58f0d9d, 0x0a78eb28f5c70b3d, 0x666ea36f7879462c, 0x0e0a77c19a07df2f}

var one = &fe{0xd35d438dc58f0d9d, 0x0a78eb28f5c70b3d, 0x666ea36f7879462c, 0x0e0a77c19a07df2f}

var r2 = &fe{0xf32cfc5b538afa89, 0xb5e71911d44501fb, 0x47ab1eff0a417ff6, 0x06d89f71cab8351f}

var _one = &fe{0x0000000000000001, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000}

var pbig, _ = new(big.Int).SetString("21888242871839275222246405745257275088696311157297823662689037894645226208583", 10)

var rbig, _ = new(big.Int).SetString("6350874878119819312338956282401532409788428879151445726012394534686998597021", 10)

var pMinus3Over4 = bigFromHex("0xc19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f51")

var pPlus1Over4 = bigFromHex("0xc19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52")

var pMinus1Over2 = bigFromHex("0x183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea3")

// -1
var nonResidue1 = &fe{0x68c3488912edefaa, 0x8d087f6872aabf4f, 0x51e1a24709081231, 0x2259d6b14729c0fa}

// -1 + 0 * u
var negativeOne2 = &fe2{
	fe{0x68c3488912edefaa, 0x8d087f6872aabf4f, 0x51e1a24709081231, 0x2259d6b14729c0fa},
	fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
}

// 9 + 1 * u
var nonResidue2 = &fe2{
	fe{0xf60647ce410d7ff7, 0x2f3d6f4dd31bd011, 0x2943337e3940c6d1, 0x1d9598e8a7e39857},
	fe{0xd35d438dc58f0d9d, 0x0a78eb28f5c70b3d, 0x666ea36f7879462c, 0x0e0a77c19a07df2f},
}

// 2 ^ (-1)
var twoInv = &fe{0x87bee7d24f060572, 0xd0fd2add2f1c6ae5, 0x8f5f7492fcfd4f44, 0x1f37631a3d9cbfac}

// Curve constants

// Group order
var q = bigFromHex("0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47")

// export curve order
var Order = q

// b coefficient for G1
var b = &fe{0x7a17caa950ad28d7, 0x1f6ac17ae15521b9, 0x334bea4e696bd284, 0x2a1f6744ce179d8e}

// G1 generator
var g1One = PointG1{
	fe{0xd35d438dc58f0d9d, 0x0a78eb28f5c70b3d, 0x666ea36f7879462c, 0x0e0a77c19a07df2f},
	fe{0xa6ba871b8b1e1b3a, 0x14f1d651eb8e167b, 0xccdd46def0f28c58, 0x1c14ef83340fbe5e},
	fe{0xd35d438dc58f0d9d, 0x0a78eb28f5c70b3d, 0x666ea36f7879462c, 0x0e0a77c19a07df2f},
}

// b coeffiecient for G2
var b2 = &fe2{
	fe{0x3bf938e377b802a8, 0x020b1b273633535d, 0x26b7edf049755260, 0x2514c6324384a86d},
	fe{0x38e7ecccd1dcff67, 0x65f0b37d93ce0d3e, 0xd749d0dd22ac00aa, 0x0141b9ce4a688d4d},
}

// G2 generator
var g2One = PointG2{
	fe2{
		fe{0x8e83b5d102bc2026, 0xdceb1935497b0172, 0xfbb8264797811adf, 0x19573841af96503b},
		fe{0xafb4737da84c6140, 0x6043dd5a5802d8c4, 0x09e950fc52a02f86, 0x14fef0833aea7b6b},
	},
	fe2{
		fe{0x619dfa9d886be9f6, 0xfe7fd297f59e9b78, 0xff9e1a62231b7dfe, 0x28fd7eebae9e4206},
		fe{0x64095b56c71856ee, 0xdc57f922327d3cbb, 0x55f935be33351076, 0x0da4a0e693fd6482},
	},
	fe2{
		fe{0xd35d438dc58f0d9d, 0x0a78eb28f5c70b3d, 0x666ea36f7879462c, 0x0e0a77c19a07df2f},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
}

var u = bigFromHex("0x44e992b44a6909f1")

var sixUPlus2 = bigFromHex("0x19d797039be763ba8")

var nonResidueInPMinusOver2 = fe2{
	fe{0xe4bbdd0c2936b629, 0xbb30f162e133bacb, 0x31a9d1b6f9645366, 0x253570bea500f8dd},
	fe{0xa1d77ce45ffe77c7, 0x07affd117826d1db, 0x6d16bd27bb7edc6b, 0x2c87200285defecc},
}

var frobeniusCoeffs61 = [6]fe2{
	fe2{
		fe{0xd35d438dc58f0d9d, 0x0a78eb28f5c70b3d, 0x666ea36f7879462c, 0x0e0a77c19a07df2f},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0xb5773b104563ab30, 0x347f91c8a9aa6454, 0x7a007127242e0991, 0x1956bcd8118214ec},
		fe{0x6e849f1ea0aa4757, 0xaa1c7b6d89f89141, 0xb6e713cdfae0ca3a, 0x26694fbb4e82ebc3},
	},
	fe2{
		fe{0x3350c88e13e80b9c, 0x7dce557cdb5e56b9, 0x6001b4b8b615564a, 0x2682e617020217e0},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0xc9af22f716ad6bad, 0xb311782a4aa662b2, 0x19eeaf64e248c7f4, 0x20273e77e3439f82},
		fe{0xacc02860f7ce93ac, 0x3933d5817ba76b4c, 0x69e6188b446c8467, 0x0a46036d4417cc55},
	},
	fe2{
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
}

var frobeniusCoeffs62 = [6]fe2{
	fe2{
		fe{0xd35d438dc58f0d9d, 0x0a78eb28f5c70b3d, 0x666ea36f7879462c, 0x0e0a77c19a07df2f},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x7361d77f843abe92, 0xa5bb2bd3273411fb, 0x9c941f314b3e2399, 0x15df9cddbb9fd3ec},
		fe{0x5dddfd154bd8c949, 0x62cb29a5a4445b60, 0x37bc870a0c7dd2b9, 0x24830a9d3171f0fd},
	},
	fe2{
		fe{0x71930c11d782e155, 0xa6bb947cffbe3323, 0xaa303344d4741444, 0x2c3b3f0d26594943},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x448a93a57b6762df, 0xbfd62df528fdeadf, 0xd858f5d00e9bd47a, 0x06b03d4d3476ec58},
		fe{0x2b19daf4bcc936d1, 0xa1a54e7a56f4299f, 0xb533eee05adeaef1, 0x170c812b84dda0b2},
	},
	fe2{
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
}

var frobeniusCoeffs12 = [12]fe2{
	fe2{
		fe{0xd35d438dc58f0d9d, 0x0a78eb28f5c70b3d, 0x666ea36f7879462c, 0x0e0a77c19a07df2f},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0xaf9ba69633144907, 0xca6b1d7387afb78a, 0x11bded5ef08a2087, 0x02f34d751a1f3a7c},
		fe{0xa222ae234c492d72, 0xd00f02a4565de15b, 0xdc2ff3a253dfc926, 0x10a75716b3899551},
	},
	fe2{
		fe{0xca8d800500fa1bf2, 0xf0c5d61468b39769, 0x0e201271ad0d4418, 0x04290f65bad856e6},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x365316184e46d97d, 0x0af7129ed4c96d9f, 0x659da72fca1009b5, 0x08116d8983a20d23},
		fe{0xb1df4af7c39c1939, 0x3d9f02878a73bf7f, 0x9b2220928caf0ae0, 0x26684515eff054a6},
	},
	fe2{
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x68c3488912edefaa, 0x8d087f6872aabf4f, 0x51e1a24709081231, 0x2259d6b14729c0fa},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
	fe2{
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		fe{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
	},
}
