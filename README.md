# Bloom Filters

This is the simple implementation of bloom filters in Go

## Installation

```
go get github.com/thisisnitish/bloom-filters
```

## Run the Project

```
go run main.go
```

## Behaviour

### Bloom Filter Size vs False Positive Rate

As you can see here, as the size increases of the bloom filters, the false positive rate
decreases. In short, lager the size of bloom filter, smaller the false positivity rate.

```
size:  100 fp_rate:  0.498 fp_rate_perc:  49.8
size:  1100 fp_rate:  0.183 fp_rate_perc:  18.3
size:  2100 fp_rate:  0.111 fp_rate_perc:  11.1
size:  3100 fp_rate:  0.089 fp_rate_perc:  8.9
size:  4100 fp_rate:  0.065 fp_rate_perc:  6.5
size:  5100 fp_rate:  0.038 fp_rate_perc:  3.8
size:  6100 fp_rate:  0.038 fp_rate_perc:  3.8
size:  7100 fp_rate:  0.035 fp_rate_perc:  3.5
size:  8100 fp_rate:  0.025 fp_rate_perc:  2.5
size:  9100 fp_rate:  0.024 fp_rate_perc:  2.4
size:  10100 fp_rate:  0.026 fp_rate_perc:  2.6
size:  11100 fp_rate:  0.027 fp_rate_perc:  2.7
size:  12100 fp_rate:  0.02 fp_rate_perc:  2
size:  13100 fp_rate:  0.02 fp_rate_perc:  2
size:  14100 fp_rate:  0.019 fp_rate_perc:  1.9
size:  15100 fp_rate:  0.015 fp_rate_perc:  1.5
size:  16100 fp_rate:  0.017 fp_rate_perc:  1.7
size:  17100 fp_rate:  0.02 fp_rate_perc:  2
size:  18100 fp_rate:  0.011 fp_rate_perc:  1.1
size:  19100 fp_rate:  0.013 fp_rate_perc:  1.3
size:  20100 fp_rate:  0.007 fp_rate_perc:  0.7
size:  21100 fp_rate:  0.011 fp_rate_perc:  1.1
size:  22100 fp_rate:  0.003 fp_rate_perc:  0.3
size:  23100 fp_rate:  0.011 fp_rate_perc:  1.1
size:  24100 fp_rate:  0.01 fp_rate_perc:  1
size:  25100 fp_rate:  0.01 fp_rate_perc:  1
size:  26100 fp_rate:  0.011 fp_rate_perc:  1.1
size:  27100 fp_rate:  0.012 fp_rate_perc:  1.2
size:  28100 fp_rate:  0.008 fp_rate_perc:  0.8
size:  29100 fp_rate:  0.005 fp_rate_perc:  0.5
```

![Bloom Filter Size vs False Positivity Rate ](./assets/bloom_filter_vs_fpr.png)
