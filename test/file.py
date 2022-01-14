from test_util import *  # got s
msg1 = b'GET /.. HTTP/1.1\r\nHost: localhost:888\r\n\r\n'
msg2 = b'GET /.. HTTP/1.1\r\nHost: jhtmac.local:888\r\n\r\n'
msg404 = b'GET /skdjf HTTP/1.1\r\nHost: jhtmac.local:888\r\n\r\n'
print(send_msg(msg1))
print(send_msg(msg2))
print(send_msg(msg404))
