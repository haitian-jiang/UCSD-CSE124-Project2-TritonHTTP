from test_util import *  # got socket

print(send_msg(send=2000))
time.sleep(4.99)
print(send_msg(send=800000, recv=800010, fulltxt=True))
