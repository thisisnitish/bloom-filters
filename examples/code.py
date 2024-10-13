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


"""
This is the code of Number of Hash Functions vs False Positive Rate
which produces a chart which you can see in README.md
"""

import matplotlib.pyplot as plt
import math
x = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20]
y = [3.1,0.4,0.2,0.2,0,0,0,0,0,0,0,0,0,0,0,0,0.1,0.1,0.1,0.1]
plt.plot(
    x, y,
    marker='o'
)

new_list = range(math.floor(min(y)), math.ceil(max(y))+1)
# new_x = range(math.floor(min(x)), math.ceil(max(x))+1)
# plt.xticks(new_x)
plt.yticks(new_list)
plt.grid(True)
plt.title('fp rate vs number of hash functions')
plt.xlabel('number of hash functions')
plt.ylabel('fp rate')
plt.legend(['fp rate'])
plt.show()