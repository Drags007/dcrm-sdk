/*
 *  Copyright (C) 2018-2019  Fusion Foundation Ltd. All rights reserved.
 *  Copyright (C) 2018-2019  caihaijun@fusion.org
 *
 *  This library is free software; you can redistribute it and/or
 *  modify it under the Apache License, Version 2.0.
 *
 *  This library is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  
 *
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package lib

import (
	"github.com/fsn-dev/dcrm-sdk/internal/common/math/random"
	"math/big"
	"time"
)

var (
    //SafePrime = make(chan *big.Int, 1000)
    SafePrime = make([]*big.Int, 4)
    RndInt = make(chan *big.Int, 1000)
)

func GenRandomSafePrime(length int) {
    p1 := random.GetSafeRandomPrimeInt(length/2)
    SafePrime[0] = p1
    p2 := random.GetSafeRandomPrimeInt(length/2)
    SafePrime[1] = p2
    p3 := random.GetSafeRandomPrimeInt(length/2)
    SafePrime[2] = p3
    p4 := random.GetSafeRandomPrimeInt(length/2)
    SafePrime[3] = p4
    return
    
    /*for {
	if len(SafePrime) < 1000 {
	    rndInt := <-RndInt
	    p := random.GetSafeRandomPrimeInt2(length/2,rndInt)
	    if p != nil {
		SafePrime <-p
		time.Sleep(time.Duration(10000)) //1000 000 000 == 1s
	    }
	}
    }*/
}

func GenRandomInt(length int) {
    return //tmp

    for {
	if len(RndInt) < 1000 {
	    p := random.GetSafeRandomInt(length/2)
	    RndInt <-p
	    
	    time.Sleep(time.Duration(10000)) //1000 000 000 == 1s
	}
    }
}

