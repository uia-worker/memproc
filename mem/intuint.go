package main

import (
	"testing"
	"fmt"
)

func AddToInt8(base int8, addi int8) int8 {
    return base + addi
}

func AddToUInt8(base uint8, addi uint8) uint8 {
    return base + addi
}

func main() {
    var addi int8 = 127

    fmt.Printf("%d", addi + 2)
}

func TestAddToInt8(t *testing.T) {
	type test struct {
                base int8
                addi int8
		want int8
	}

	tests := []test {
		{base: 126, addi: 1, want: 127},
	}

	for _, tc := range tests {
 	 got := AddToInt8(tc.base, tc.addi)
 	 if tc.want != got {
 		 t.Errorf("expected: %v, got: %v", tc.want, got)
 	 }
   }	

}

func TestAddToUInt8(t *testing.T) {
        type test struct {
                base uint8
                addi uint8
                want uint8
        }

        tests := []test {
                {base: 127, addi: 1, want: 128},
        }

        for _, tc := range tests {
         got := AddToUInt8(tc.base, tc.addi)
         if tc.want != got {
                 t.Errorf("expected: %v, got: %v", tc.want, got)
         }
   }   

}
