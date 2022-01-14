from test_util import *

goodhost = b'GET /subdir1 HTTP/1.1\r\nHost: jhtMac.local\r\n\r\n'

def testhost():
    s = get_socket()
    print(send_msg(goodhost, s, recv=409600))
    # sleep(7)
    print(send_msg(goodhost, s, recv=409600))


if __name__ == "__main__":
    testhost()