from test_util import *


badhost = b'GET / HTTP/1.1\r\nHost: 127.0.0.1\r\n\r\n'
goodhost = b'GET /../ HTTP/1.1\r\nHost: jhtMac.local\r\n\r\n'

def testhost():
    print(send_msg(badhost))
    sleep(1)
    print(send_msg(goodhost))

if __name__ == "__main__":
    testhost()