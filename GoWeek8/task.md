1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

C:\Program Files\Redis>redis-benchmark -d 10 -t get,set

====== GET ======
  10000 requests completed in 1.65 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

0.01% <= 1 milliseconds
0.85% <= 2 milliseconds
19.97% <= 3 milliseconds
55.31% <= 4 milliseconds
68.25% <= 5 milliseconds
73.86% <= 6 milliseconds
76.99% <= 7 milliseconds
80.36% <= 8 milliseconds
83.11% <= 9 milliseconds
84.89% <= 10 milliseconds
86.57% <= 11 milliseconds
88.23% <= 12 milliseconds
89.72% <= 13 milliseconds
90.95% <= 14 milliseconds
91.80% <= 15 milliseconds
92.43% <= 16 milliseconds
93.04% <= 17 milliseconds
93.30% <= 18 milliseconds
93.42% <= 19 milliseconds
93.72% <= 20 milliseconds
93.86% <= 21 milliseconds
93.91% <= 22 milliseconds
93.97% <= 23 milliseconds
94.30% <= 24 milliseconds
94.50% <= 25 milliseconds
94.69% <= 26 milliseconds
95.05% <= 27 milliseconds
95.31% <= 28 milliseconds
95.53% <= 29 milliseconds
96.04% <= 30 milliseconds
96.51% <= 31 milliseconds
96.82% <= 32 milliseconds
96.97% <= 33 milliseconds
97.14% <= 34 milliseconds
97.32% <= 35 milliseconds
97.56% <= 36 milliseconds
98.05% <= 37 milliseconds
98.26% <= 38 milliseconds
98.52% <= 39 milliseconds
98.83% <= 40 milliseconds
98.89% <= 41 milliseconds
98.92% <= 42 milliseconds
98.96% <= 43 milliseconds
98.98% <= 44 milliseconds
99.04% <= 45 milliseconds
99.05% <= 46 milliseconds
99.06% <= 52 milliseconds
99.07% <= 59 milliseconds
99.08% <= 60 milliseconds
99.09% <= 66 milliseconds
99.10% <= 67 milliseconds
99.14% <= 68 milliseconds
99.15% <= 70 milliseconds
99.16% <= 71 milliseconds
99.19% <= 73 milliseconds
99.20% <= 76 milliseconds
99.21% <= 77 milliseconds
99.23% <= 78 milliseconds
99.24% <= 79 milliseconds
99.27% <= 80 milliseconds
99.31% <= 81 milliseconds
99.34% <= 82 milliseconds
99.36% <= 83 milliseconds
99.37% <= 84 milliseconds
99.38% <= 86 milliseconds
99.39% <= 91 milliseconds
99.41% <= 92 milliseconds
99.42% <= 94 milliseconds
99.44% <= 97 milliseconds
99.47% <= 98 milliseconds
99.49% <= 101 milliseconds
99.50% <= 104 milliseconds
99.51% <= 109 milliseconds
99.52% <= 116 milliseconds
99.53% <= 119 milliseconds
99.54% <= 124 milliseconds
99.55% <= 128 milliseconds
99.56% <= 129 milliseconds
99.57% <= 136 milliseconds
99.58% <= 141 milliseconds
99.59% <= 144 milliseconds
99.62% <= 145 milliseconds
99.68% <= 149 milliseconds
99.70% <= 150 milliseconds
99.74% <= 157 milliseconds
99.83% <= 158 milliseconds
99.85% <= 167 milliseconds
99.86% <= 168 milliseconds
99.87% <= 172 milliseconds
99.88% <= 177 milliseconds
99.89% <= 181 milliseconds
99.90% <= 187 milliseconds
99.91% <= 199 milliseconds
99.92% <= 273 milliseconds
99.93% <= 344 milliseconds
99.94% <= 349 milliseconds
99.95% <= 359 milliseconds
99.96% <= 365 milliseconds
99.97% <= 371 milliseconds
99.98% <= 379 milliseconds
100.00% <= 384 milliseconds
6071.65 requests per second

====== SET ======
  10000 requests completed in 1.37 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

0.01% <= 1 milliseconds
1.34% <= 2 milliseconds
33.51% <= 3 milliseconds
67.46% <= 4 milliseconds
80.48% <= 5 milliseconds
85.33% <= 6 milliseconds
88.55% <= 7 milliseconds
90.30% <= 8 milliseconds
91.32% <= 9 milliseconds
92.27% <= 10 milliseconds
92.91% <= 11 milliseconds
93.66% <= 12 milliseconds
94.10% <= 13 milliseconds
94.54% <= 14 milliseconds
94.89% <= 15 milliseconds
95.02% <= 16 milliseconds
95.12% <= 17 milliseconds
95.24% <= 18 milliseconds
95.53% <= 19 milliseconds
95.83% <= 20 milliseconds
96.67% <= 21 milliseconds
97.66% <= 22 milliseconds
98.00% <= 23 milliseconds
98.20% <= 25 milliseconds
98.34% <= 26 milliseconds
98.46% <= 27 milliseconds
98.53% <= 28 milliseconds
98.83% <= 29 milliseconds
99.01% <= 30 milliseconds
99.18% <= 31 milliseconds
99.29% <= 32 milliseconds
99.34% <= 33 milliseconds
99.39% <= 34 milliseconds
99.43% <= 35 milliseconds
99.46% <= 37 milliseconds
99.47% <= 42 milliseconds
99.48% <= 44 milliseconds
99.49% <= 47 milliseconds
99.54% <= 48 milliseconds
99.57% <= 50 milliseconds
99.58% <= 53 milliseconds
99.59% <= 58 milliseconds
99.60% <= 61 milliseconds
99.61% <= 66 milliseconds
99.62% <= 67 milliseconds
99.63% <= 73 milliseconds
99.64% <= 78 milliseconds
99.65% <= 82 milliseconds
99.66% <= 84 milliseconds
99.67% <= 91 milliseconds
99.68% <= 94 milliseconds
99.69% <= 97 milliseconds
99.70% <= 101 milliseconds
99.71% <= 107 milliseconds
99.72% <= 114 milliseconds
99.73% <= 134 milliseconds
99.74% <= 144 milliseconds
99.75% <= 146 milliseconds
99.76% <= 149 milliseconds
99.77% <= 157 milliseconds
99.78% <= 161 milliseconds
99.79% <= 166 milliseconds
99.80% <= 171 milliseconds
99.81% <= 175 milliseconds
99.82% <= 177 milliseconds
99.83% <= 182 milliseconds
99.84% <= 189 milliseconds
99.85% <= 193 milliseconds
99.86% <= 195 milliseconds
99.90% <= 196 milliseconds
99.91% <= 206 milliseconds
99.92% <= 209 milliseconds
99.93% <= 215 milliseconds
99.94% <= 224 milliseconds
99.95% <= 226 milliseconds
99.96% <= 235 milliseconds
99.97% <= 258 milliseconds
99.98% <= 260 milliseconds
99.99% <= 266 milliseconds
100.00% <= 267 milliseconds
7293.95 requests per second

|get|||
|10|100000 requests completed in 1.65 seconds|6071.65 requests per second|
|20|100000 requests completed in 1.64 seconds|6043.02 requests per second|
|50|100000 requests completed in 1.81 seconds|5997.21 requests per second|
|100|100000 requests completed in 1.89 seconds|5933.48 requests per second|
|200|100000 requests completed in 1.88 seconds|6117.26 requests per second|
|1k|100000 requests completed in 1.89 seconds|5918.23 requests per second|
|5k|100000 requests completed in 2.73 seconds|5630.63 requests per second|
|set|||
|10|100000 requests completed in 1.37 seconds|7293.95 requests per second|
|20|100000 requests completed in 1.52 seconds|7037.13 requests per second|
|50|100000 requests completed in 1.47 seconds|7112.32 requests per second|
|100|100000 requests completed in 1.48 seconds|7107.25 requests per second|
|200|100000 requests completed in 1.46 seconds|7113.22 requests per second|
|1k|100000 requests completed in 1.74 seconds|6998.83 requests per second|
|5k|100000 requests completed in 2.13 seconds|6502.32 requests per second|

2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
写入不同数量长度的value，结合info memory，同等长度的value，写入数量越多，平均每个value占用内存越多
