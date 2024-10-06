"""
This is the code of Bloom Filter Size vs False Positive Rate
which produces a chart which you can see in README.md
"""

import matplotlib.pyplot as plt
plt.plot(
    [100, 1100, 2100, 3100, 4100, 5100, 6100, 7100, 8100, 9100, 10100, 11100, 12100, 13100, 14100, 15100, 16100, 17100, 18100, 19100, 20100, 21100, 22100, 23100, 24100, 25100, 26100, 27100, 28100, 29100], 
    [49.8, 18.3, 11.1, 8.9, 6.5, 3.8, 3.8, 3.5, 2.5, 2.4, 2.6, 2.7, 2, 2, 1.9, 1.5, 1.7, 2, 1.1, 1.3, 0.7, 1.1, 0.3, 1.1, 1, 1, 1.1, 1.2, 0.8, 0.5],
    marker='o'
)
plt.grid(True)
plt.title('fp rate vs size of bloom filters')
plt.xlabel('size of bloom filters')
plt.ylabel('fp rate')
plt.legend(['fp rate'])
plt.show()