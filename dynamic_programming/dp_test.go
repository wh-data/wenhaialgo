package dynamic_programming

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFragGoStair(t *testing.T) {
	fmt.Println(FragGoStair(3))
	fmt.Println(FragGoStair(4))
	fmt.Println(FragGoStair(5))
	fmt.Println(FragGoStair(6))
	fmt.Println(FragGoStairPossibility(4))
}

func TestBackpack(Z *testing.T) {
	w := []int{1, 2, 5, 6, 7}
	v := []int{1, 6, 18, 22, 28}
	art := []Article{
		{
			Weight: 1,
			Value:  1,
		},
		{
			Weight: 2,
			Value:  6,
		},
		{
			Weight: 5,
			Value:  18,
		},
		{
			Weight: 6,
			Value:  22,
		},
		{
			Weight: 7,
			Value:  28,
		},
	}
	c := 11
	bp := &Backpack{}
	fmt.Println(BackpackIssueByGreedyAlgo(c, w, v))
	fmt.Println("before algo ", bp)
	BackpackIssueByBackTrackAlgo(0, c, art, bp)
	fmt.Println("after back tracking ", bp)
}

func TestBackpackDP(Z *testing.T) {
	wt := []int{100, 77, 22, 29, 50, 99}
	val := []int{5, 92, 22, 87, 46, 90}
	c := 100
	fmt.Println(BackpackIssueByDP(wt, val, c))
}

func TestBB(t *testing.T) {
	fmt.Println("td"[len("td")-1:])
	fmt.Println(string("test"[2]))
}

func TestIsPalindrome(t *testing.T) {
	strs := []string{"a", "aa", "abba"}
	for _, s := range strs {
		fmt.Println(isPalindrome(s))
	}
}

func TestGetLongestPalindrome(t *testing.T) {
	s := []string{"baab", "abcba", "aa", " | a", "a", "ttest", "asdfadsfasegewtragdasfgasdbghgaadsfasdfadsfasdfcas"}
	start := time.Now()
	for _, x := range s {
		_, p := GetPalindrome(x)
		fmt.Println(p)
	}
	fmt.Println(time.Now().Sub(start).Nanoseconds())
}

func TestNMPath(t *testing.T) {
	fmt.Println(n_m_path(4, 3))
}

// testing cases in 牛客
func TestBback_pack_v3(t *testing.T) {
	c := 2169 //8466
	w := make([]int, 0)
	v := make([]int, 0)
	//raw :=
	//	`1 335
	//225 170
	//359 479
	//465 463
	//146 206
	//328 282
	//492 462
	//443 496
	//437 328
	//105 392
	//154 403
	//383 293
	//217 422
	//396 219
	//227 448
	//39 272
	//413 370
	//300 168
	//395 36
	//312 204
	//334 323
	//165 174
	//212 142
	//369 254
	//145 48
	//258 163
	//360 38
	//242 224
	//279 30
	//36 317
	//343 191
	//107 289
	//443 41
	//149 265
	//306 447
	//230 391
	//351 371
	//102 7
	//49 394
	//124 130
	//455 85
	//341 257
	//377 467
	//309 432
	//440 445
	//324 127
	//39 38
	//83 119
	//42 430
	//116 334
	//159 140
	//431 205
	//307 478
	//387 174
	//246 22
	//73 425
	//330 271
	//74 278
	//13 98
	//291 487
	//137 162
	//268 356
	//75 156
	//53 32
	//151 351
	//225 442
	//431 467
	//192 108
	//338 8
	//288 458
	//384 254
	//410 446
	//259 210
	//89 222
	//447 423
	//31 7
	//169 414
	//92 401
	//156 263
	//360 411
	//38 125
	//484 49
	//42 96
	//351 103
	//337 292
	//21 375
	//22 97
	//200 349
	//485 169
	//235 282
	//500 54`
	raw :=
		`241 490
43 165
414 120
205 92
233 319
206 251
40 476
423 304
248 99
149 85
365 472
76 414
213 46
179 47
263 270
486 20
445 290
41 366
9 246
371 319
324 102
473 133
88 153
264 71
104 402
28 424
470 101
66 16
44 29
89 348
138 444
464 410
182 50
343 89
61 109
259 222
389 455
191 147
344 450
121 431
68 249
284 37
227 36
39 186
130 354
249 225
360 424
267 258
456 445
227 319
26 412
2 356
497 50
16 85
343 465
414 76
197 143
73 449
107 427
430 174
206 405
313 127
94 376
37 66
142 237
495 315
153 257
339 437
356 483
132 16
342 231
12 126
187 138
151 191
135 163
354 394
453 417
263 9
455 234
135 304
257 304
125 149
214 318
29 110`
	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		eles := strings.Split(line, " ")
		tem_w, _ := strconv.Atoi(strings.TrimSpace(eles[0]))
		tem_v, _ := strconv.Atoi(strings.TrimSpace(eles[1]))
		w = append(w, tem_w)
		v = append(v, tem_v)
	}
	fmt.Println(len(lines))
	fmt.Println(w)
	fmt.Println(v)
	fmt.Println(BackpackIssueByDP(w, v, c))
}
