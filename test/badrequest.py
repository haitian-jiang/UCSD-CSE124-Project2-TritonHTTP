from test_util import *  # got socket

# bad requests
post = b'POST / HTTP/1.1\r\nHost: localhost\r\n\r\n'
url = b'GET asdkf HTTP/1.1\r\nHost: localhost\r\n\r\n'
http = b'GET / HTTP/1.2\r\nHost: localhost\r\n\r\n'
maohao = b'GET / HTTP/1.1\r\nHost localhost\r\n\r\n'
crlf = b'GET / HTTP/1.1\r\nHost: localhost\r\n'
host = b'GET / HTTP/1.1\r\n\r\n'
keyvalue = b'GET / HTTP/1.1\r\nHost localhost\r\n\r\n'

# close request
close = b'GET / HTTP/1.1\r\nHost: localhost\r\nconnection:close\r\n\r\n'

def timeout_badrequest():
    s = get_socket()
    print(send_msg(GDRQ, s))
    sleep(4.995)
    response = send_msg(gdrq + b'tail:'+b'a'*900000+b'\r\n\r\n', s)
    print(response)
    sleep(1)

def other_badrequests():
    rqs = [post, url, http, maohao, crlf, host, keyvalue, close]
    for rq in rqs:
        print(send_msg(rq))

if __name__ == "__main__":
    # timeout_badrequest()
    # sleep(10)
    other_badrequests()