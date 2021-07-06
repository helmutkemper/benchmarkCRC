# benchmark for CRC algorithm

Baseado no código

[https://github.com/patrickmn/picugen](https://github.com/patrickmn/picugen)

```shell
goos: darwin
goarch: arm64
BenchmarkAdler32
BenchmarkAdler32-8           	23755016	        51.34 ns/op
BenchmarkCRC32IEEE
BenchmarkCRC32IEEE-8         	39377883	        30.44 ns/op
BenchmarkCRC31Castagnoli
BenchmarkCRC31Castagnoli-8   	39839700	        30.02 ns/op
BenchmarkCRC31Koopman
BenchmarkCRC31Koopman-8      	 4124380	       289.9 ns/op
BenchmarkCRC64ISO
BenchmarkCRC64ISO-8          	 9016206	       134.1 ns/op
BenchmarkCRC64ECMA
BenchmarkCRC64ECMA-8         	 8886650	       133.4 ns/op
BenchmarkFNV32
BenchmarkFNV32-8             	 8146906	       145.2 ns/op
BenchmarkFNV32a
BenchmarkFNV32a-8            	 8193144	       145.7 ns/op
BenchmarkFNV64
BenchmarkFNV64-8             	 8198060	       145.5 ns/op
BenchmarkFNV64a
BenchmarkFNV64a-8            	 8174168	       145.3 ns/op
BenchmarkHmacSha256
BenchmarkHmacSha256-8        	 5323273	       225.3 ns/op
BenchmarkHmacMd5
BenchmarkHmacMd5-8           	 2315974	       517.9 ns/op
BenchmarkHmacSha1
BenchmarkHmacSha1-8          	 5201324	       230.4 ns/op
BenchmarkHmacSha512
BenchmarkHmacSha512-8        	  620304	      1825 ns/op
BenchmarkMd5
BenchmarkMd5-8               	 3869907	       307.0 ns/op
BenchmarkMd52
BenchmarkMd52-8              	 1666381	       718.6 ns/op
BenchmarkSha1
BenchmarkSha1-8              	 9310050	       128.1 ns/op
BenchmarkSha254
BenchmarkSha254-8            	 9437179	       126.1 ns/op
BenchmarkSha256
BenchmarkSha256-8            	 9391630	       126.0 ns/op
BenchmarkSha512
BenchmarkSha512-8            	 1295797	       926.9 ns/op
```