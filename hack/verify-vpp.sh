#!/bin/bash
# List interfaces
#   Name             Idx State Counter    	Count
# GigabitEthernet2/2/0 5 up   rx packets	391
#                             rx bytes		54248
#                             tx packets	34
#                             tx bytes		1666
#                             drops			375
#                             punts			7
#                             ip4			37
# GigabitEthernet2/7/0              6        down
# local0                            0        down
# pg/stream-0                       1        down
# pg/stream-1                       2        down
# pg/stream-2                       3        down
# pg/stream-3                       4        down
vppctl show int

# Setup a interface
vppctl set int ip address GigabitEthernet2/2/0 192.168.0.100/24
vppctl set int state GigabitEthernet2/2/0 up # now 192.168.0.100 should be ping by other hosts
vppctl sh ip arp

# Check the stats
vppctl show run

# Check the error counters
vppctl show error

# trace packets
vppctl trace add dpdk-input 5
vppctl show trace
